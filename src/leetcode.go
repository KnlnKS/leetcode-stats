package src

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"
)

var ErrNotFound = errors.New("user not found")

const userAgent = "Mozilla/5.0 (compatible; leetcode-stats/2.0; +https://github.com/KnlnKS/leetcode-stats)"

type Site struct {
	GraphQL  string
	Calendar string
	Query    string
	Referer  string
	Origin   string
}

var Global = Site{
	GraphQL:  "https://leetcode.com/graphql/",
	Calendar: "https://leetcode.com/graphql/",
	Query:    `query userProfileCalendar($username:String!){matchedUser(username:$username){userCalendar{streak totalActiveDays submissionCalendar}}}`,
	Referer:  "https://leetcode.com/",
	Origin:   "https://leetcode.com",
}

var China = Site{
	GraphQL:  "https://leetcode.cn/graphql/",
	Calendar: "https://leetcode.cn/graphql/noj-go/",
	Query:    `query userProfileCalendar($userSlug:String!){userCalendar(userSlug:$userSlug){streak:recentStreak totalActiveDays submissionCalendar}}`,
	Referer:  "https://leetcode.cn/",
	Origin:   "https://leetcode.cn",
}

type Stats struct {
	Solved     [3]int
	Total      [3]int
	Attempting int
}

type Calendar struct {
	Days       map[int64]int
	Total      int
	ActiveDays int
	MaxStreak  int
}

var client = &http.Client{Timeout: 8 * time.Second}

func post(ctx context.Context, site Site, url, query string, vars map[string]any, out any) error {
	body, err := json.Marshal(map[string]any{"query": query, "variables": vars})
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", site.Referer)
	req.Header.Set("Origin", site.Origin)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("leetcode returned %d", resp.StatusCode)
	}
	var env struct {
		Data   json.RawMessage `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(&env); err != nil {
		return err
	}
	if len(env.Errors) > 0 {
		return errors.New(env.Errors[0].Message)
	}
	return json.Unmarshal(env.Data, out)
}

const progressQuery = `query userProfileUserQuestionProgressV2($userSlug:String!){userProfileUserQuestionProgressV2(userSlug:$userSlug){numAcceptedQuestions{count difficulty}numFailedQuestions{count difficulty}numUntouchedQuestions{count difficulty}}}`

type diffCount struct {
	Count      int    `json:"count"`
	Difficulty string `json:"difficulty"`
}

func diffIndex(d string) int {
	switch d {
	case "EASY":
		return 0
	case "MEDIUM":
		return 1
	case "HARD":
		return 2
	}
	return -1
}

func GetStats(ctx context.Context, site Site, user string) (Stats, error) {
	var data struct {
		Progress *struct {
			Accepted  []diffCount `json:"numAcceptedQuestions"`
			Failed    []diffCount `json:"numFailedQuestions"`
			Untouched []diffCount `json:"numUntouchedQuestions"`
		} `json:"userProfileUserQuestionProgressV2"`
	}
	if err := post(ctx, site, site.GraphQL, progressQuery, map[string]any{"userSlug": user}, &data); err != nil {
		return Stats{}, err
	}
	if data.Progress == nil || len(data.Progress.Accepted) == 0 {
		return Stats{}, ErrNotFound
	}
	var s Stats
	add := func(list []diffCount, dst *[3]int) {
		for _, c := range list {
			if i := diffIndex(c.Difficulty); i >= 0 {
				dst[i] += c.Count
			}
		}
	}
	add(data.Progress.Accepted, &s.Solved)
	add(data.Progress.Accepted, &s.Total)
	add(data.Progress.Failed, &s.Total)
	add(data.Progress.Untouched, &s.Total)
	for _, c := range data.Progress.Failed {
		s.Attempting += c.Count
	}
	return s, nil
}

func GetCalendar(ctx context.Context, site Site, user string) (Calendar, error) {
	raw := struct {
		Streak     int    `json:"streak"`
		ActiveDays int    `json:"totalActiveDays"`
		Submission string `json:"submissionCalendar"`
	}{}
	var err error
	if site.Calendar == Global.Calendar {
		var data struct {
			MatchedUser *struct {
				UserCalendar json.RawMessage `json:"userCalendar"`
			} `json:"matchedUser"`
		}
		if err = post(ctx, site, site.Calendar, site.Query, map[string]any{"username": user}, &data); err != nil {
			return Calendar{}, err
		}
		if data.MatchedUser == nil {
			return Calendar{}, ErrNotFound
		}
		err = json.Unmarshal(data.MatchedUser.UserCalendar, &raw)
	} else {
		var data struct {
			UserCalendar json.RawMessage `json:"userCalendar"`
		}
		if err = post(ctx, site, site.Calendar, site.Query, map[string]any{"userSlug": user}, &data); err != nil {
			return Calendar{}, err
		}
		err = json.Unmarshal(data.UserCalendar, &raw)
	}
	if err != nil {
		return Calendar{}, err
	}
	counts := map[string]int{}
	if raw.Submission != "" {
		if err := json.Unmarshal([]byte(raw.Submission), &counts); err != nil {
			return Calendar{}, err
		}
	}
	cal := Calendar{Days: make(map[int64]int, len(counts)), ActiveDays: raw.ActiveDays, MaxStreak: raw.Streak}
	for k, v := range counts {
		var sec int64
		if _, err := fmt.Sscan(k, &sec); err != nil {
			continue
		}
		cal.Days[sec] += v
		cal.Total += v
	}
	return cal, nil
}

func (c Calendar) levels() [4]int {
	active := make([]int, 0, len(c.Days))
	for _, v := range c.Days {
		if v > 0 {
			active = append(active, v)
		}
	}
	sort.Ints(active)
	var th [4]int
	for i, p := range [4]int{25, 50, 75, 100} {
		idx := len(active)*p/100 - 1
		if idx < 0 {
			idx = 0
		}
		if len(active) > 0 {
			th[i] = active[idx]
		}
	}
	return th
}

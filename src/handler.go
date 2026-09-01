package src

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var userPattern = regexp.MustCompile(`^[A-Za-z0-9._-]{1,64}$`)

func Serve(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	user, graph := route(r.URL.Path, q.Get("username"))
	theme := GetTheme(q.Get("theme"))
	font := GetFont(q.Get("font"))
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")

	fail := func(title, detail string) {
		w.Header().Set("Cache-Control", "no-store")
		var b bytes.Buffer
		if err := RenderError(&b, theme, systemFont, title, detail); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b.Bytes())
	}

	if user == "" {
		fail("Missing username", "Add ?username=your-handle to the URL")
		return
	}
	if !userPattern.MatchString(user) {
		fail("Invalid username", "Letters, digits, dot, dash and underscore only")
		return
	}

	site := Global
	if q.Has("cn") || q.Has("CN") || q.Get("site") == "cn" {
		site = China
	}

	ctx, cancel := context.WithTimeout(r.Context(), 6*time.Second)
	defer cancel()

	var b bytes.Buffer
	var err error
	if graph {
		var cal Calendar
		if cal, err = GetCalendar(ctx, site, user); err == nil {
			err = RenderGraph(&b, cal, theme, font, graphWidth(q.Get("width")), time.Now().UTC())
		}
	} else {
		var stats Stats
		if stats, err = GetStats(ctx, site, user); err == nil {
			err = RenderStats(&b, stats, theme, font)
		}
	}
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			fail("User not found", "No LeetCode profile for "+user)
		} else {
			fail("Something went wrong", "Could not reach LeetCode")
		}
		return
	}

	w.Header().Set("Cache-Control", "public, max-age=1800, s-maxage=1800, stale-while-revalidate=86400")
	w.Write(b.Bytes())
}

func graphWidth(raw string) float64 {
	width, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 836
	}
	return max(300, min(2000, width))
}

func route(path, fallback string) (string, bool) {
	segments := strings.FieldsFunc(path, func(c rune) bool { return c == '/' })
	if len(segments) > 0 && segments[0] == "api" {
		segments = segments[1:]
	}
	graph := len(segments) > 0 && segments[len(segments)-1] == "graph"
	if graph {
		segments = segments[:len(segments)-1]
	}
	if len(segments) > 0 {
		return segments[0], graph
	}
	return fallback, graph
}

package handler

import (
	"fmt"
	"leetcode-stats/src"
	"leetcode-stats/templates"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()
	w.Header().Add("Content-Type", "image/svg+xml")
	w.Header().Add("Cache-Control", "s-max-age=60, stale-while-revalidate")
	if q["username"] != nil && len(q["username"][0]) > 0 {
		submissionData, err := src.GetSubmissionStats(q["username"][0])
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		var theme src.Theme
		if q["theme"] != nil && len(q["theme"][0]) > 0 {
			theme = src.GenerateTheme(q["theme"][0])
		} else {
			theme = src.GenerateTheme("light")
		}

		tmpl := templates.GetSubmissionStatsTemplate()
		tmpl.Execute(w, struct {
			SubmissionData src.SubmissionData
			Theme          src.Theme
		}{
			SubmissionData: submissionData,
			Theme:          theme,
		})
	} else {
		fmt.Fprintf(w, "<svg>Something went wrong s</svg>")
	}
}

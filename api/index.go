package handler

import (
	"fmt"
	"leetcode-stats/templates"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()
	w.Header().Add("Content-Type", "image/svg+xml")
	w.Header().Add("Cache-Control", "s-max-age=60, stale-while-revalidate")
	if q["username"] != nil && len(q["username"][0]) > 0 {

		fmt.Fprintf(w, templates.GetPath())
	} else {
		fmt.Fprintf(w, "<svg>Something went wrong s</svg>")
	}
}

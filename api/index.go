package handler

import (
	"fmt"
	"leetcode-stats/src"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()
	if q["username"] != nil && len(q["username"][0]) > 0 {
		respData, err := src.GetSubmissionStats(q["username"][0])
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		log.Print(respData)
	}
	w.Header().Add("Content-Type", "image/svg+xml")
	w.Header().Add("Cache-Control", "s-max-age=60, stale-while-revalidate")
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()
	w.Header().Add("Content-Type", "image/svg+xml")
	w.Header().Add("Cache-Control", "s-max-age=60, stale-while-revalidate")
	if q["username"] != nil && len(q["username"][0]) > 0 {

		fmt.Fprintf(w, basepath)
	} else {
		fmt.Fprintf(w, "<svg>Something went wrong s</svg>")
	}
}

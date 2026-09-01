package handler

import (
	"net/http"

	"leetcode-stats/src"
)

func Handler(w http.ResponseWriter, r *http.Request) { src.Serve(w, r) }

package templates

import (
	"html/template"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func GetSubmissionStatsTemplate() *template.Template {
	tmpl, err := template.ParseFiles("/vercel/templates/submissionData.svg")
	if err != nil {
		panic(err)
	}
	return tmpl
}

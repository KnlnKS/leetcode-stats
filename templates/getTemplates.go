package templates

import (
	"html/template"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func GetSubmissionStatsTemplate() *template.Template {
	tmpl, err := template.ParseFiles(basepath + "/submissionData.svg")
	if err != nil {
		panic(err)
	}
	return tmpl
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		*files = append(*files, path)
		return nil
	}
}

func GetPath() []string {
	var files []string

	err := filepath.Walk(basepath, visit(&files))
	if err != nil {
		panic(err)
	}

	return files
}

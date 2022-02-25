package templates

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func GetSubmissionStatsTemplate() *template.Template {
	var files []string
	err := filepath.Walk(basepath, visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
	tmpl, err := template.ParseFiles(basepath + "/submissionData.svg")
	if err != nil {
		panic(err)
	}
	return tmpl
}

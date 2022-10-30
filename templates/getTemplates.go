package templates

import (
	_ "embed"
	"text/template"
)

//go:embed submissionStats.svg
var submissionStatsSVG string
var submissionStatsTmpl *template.Template = nil

func GetSubmissionStatsTemplate() *template.Template {
	if submissionStatsTmpl == nil {
		submissionStatsTmpl, _ = template.New("SubmissionStats").Parse(`{{define "T"}}` + submissionStatsSVG + `{{end}}`)
	}
	return submissionStatsTmpl
}

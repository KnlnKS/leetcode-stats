package src

import (
	"embed"
	"io"
	"strconv"
	"strings"
	"text/template"
)

//go:embed templates/*.svg
var templateFS embed.FS

var templates = template.Must(template.New("").Funcs(template.FuncMap{
	"n": num,
}).ParseFS(templateFS, "templates/*.svg"))

func num(v float64) string {
	s := strconv.FormatFloat(v, 'f', 4, 64)
	if strings.Contains(s, ".") {
		s = strings.TrimRight(s, "0")
		s = strings.TrimSuffix(s, ".")
	}
	return s
}

const (
	circumference = 264.0
	interGapDeg   = 8.1
	startAngle    = 225.0
)

type Font struct {
	FontFamily string
	FontCSS    string
}

type Seg struct {
	Rot       float64
	Track     float64
	TrackRest float64
	Prog      float64
	ProgRest  float64
	Color     string
}

type Pill struct {
	Y      float64
	LabelY float64
	ValueY float64
	Label  string
	Value  string
	Color  string
}

type StatsView struct {
	Theme
	Font
	Solved     int
	Total      int
	Attempting int
	Segs       []Seg
	Pills      []Pill
}

type ErrorView struct {
	Theme
	Font
	Title  string
	Detail string
}

var pillLabels = [3]string{"Easy", "Med.", "Hard"}

func RenderStats(w io.Writer, s Stats, th Theme, f Font) error {
	v := StatsView{Theme: th, Font: f, Attempting: s.Attempting}
	var totalAll int
	for i := 0; i < 3; i++ {
		v.Solved += s.Solved[i]
		totalAll += s.Total[i]
	}
	v.Total = totalAll
	gapLen := circumference * interGapDeg / 360
	budget := circumference*(360-90)/360 - 2*gapLen
	rot := startAngle
	for i := 0; i < 3; i++ {
		if totalAll > 0 && s.Total[i] > 0 {
			seg := budget * float64(s.Total[i]) / float64(totalAll)
			prog := seg * float64(s.Solved[i]) / float64(s.Total[i])
			v.Segs = append(v.Segs, Seg{
				Rot:       rot,
				Track:     seg,
				TrackRest: circumference - seg,
				Prog:      prog,
				ProgRest:  circumference - prog,
				Color:     th.Difficulty[i],
			})
			rot += (seg + gapLen) * 360 / circumference
		}
		y := 16 + float64(i)*52
		v.Pills = append(v.Pills, Pill{
			Y:      y,
			LabelY: y + 17.3,
			ValueY: y + 35.3,
			Label:  pillLabels[i],
			Value:  strconv.Itoa(s.Solved[i]) + "/" + strconv.Itoa(s.Total[i]),
			Color:  th.Difficulty[i],
		})
	}
	return templates.ExecuteTemplate(w, "stats.svg", v)
}

func RenderError(w io.Writer, th Theme, f Font, title, detail string) error {
	return templates.ExecuteTemplate(w, "error.svg", ErrorView{Theme: th, Font: f, Title: title, Detail: detail})
}

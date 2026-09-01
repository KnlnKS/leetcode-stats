package src

import (
	"embed"
	"io"
	"strconv"
	"strings"
	"text/template"
	"time"
)

//go:embed templates/*.svg
var templateFS embed.FS

var templates = template.Must(template.New("").Funcs(template.FuncMap{
	"n": num,
	"c": comma,
}).ParseFS(templateFS, "templates/*.svg"))

func num(v float64) string {
	s := strconv.FormatFloat(v, 'f', 4, 64)
	if strings.Contains(s, ".") {
		s = strings.TrimRight(s, "0")
		s = strings.TrimSuffix(s, ".")
	}
	return s
}

func comma(v int) string {
	s := strconv.Itoa(v)
	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	return s
}

const (
	circumference = 264.0
	interGapDeg   = 8.1
	startAngle    = 225.0
	blockWidth    = 8.86
	blockGap      = 11.52
	areaGap       = 15.95
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

type Cell struct {
	X    float64
	Y    float64
	Fill string
}

type MonthLabel struct {
	X    float64
	Text string
}

type GraphView struct {
	Theme
	Font
	W          float64
	H          float64
	Total      int
	ActiveDays int
	MaxStreak  int
	GridW      float64
	Cells      []Cell
	Labels     []MonthLabel
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

func RenderGraph(w io.Writer, cal Calendar, th Theme, f Font, width float64, now time.Time) error {
	v := GraphView{
		Theme:      th,
		Font:       f,
		W:          width,
		H:          width * 184 / 836,
		Total:      cal.Total,
		ActiveDays: cal.ActiveDays,
		MaxStreak:  cal.MaxStreak,
	}
	th4 := cal.levels()
	end := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	start := end.AddDate(-1, 0, 0)
	areaX := 0.0
	cols := 0
	for m := time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, time.UTC); !m.After(end); m = m.AddDate(0, 1, 0) {
		first := m
		if first.Before(start) {
			first = start
		}
		last := m.AddDate(0, 1, -1)
		if last.After(end) {
			last = end
		}
		pad := int(first.Weekday())
		days := int(last.Sub(first).Hours()/24) + 1
		cols = (pad + days + 6) / 7
		for i := 0; i < days; i++ {
			day := first.AddDate(0, 0, i)
			idx := pad + i
			v.Cells = append(v.Cells, Cell{
				X:    areaX + float64(idx/7)*blockGap,
				Y:    float64(idx%7) * blockGap,
				Fill: levelColor(cal.Days[day.Unix()], th4, th),
			})
		}
		if days >= 14 {
			v.Labels = append(v.Labels, MonthLabel{
				X:    areaX + (float64(cols)-2.5)/2*blockGap,
				Text: first.Format("Jan"),
			})
		}
		areaX += float64(cols-1)*blockGap + areaGap
	}
	v.GridW = areaX - areaGap + blockWidth
	return templates.ExecuteTemplate(w, "graph.svg", v)
}

func levelColor(count int, th4 [4]int, th Theme) string {
	if count <= 0 {
		return th.Empty
	}
	for i := 0; i < 4; i++ {
		if count <= th4[i] {
			return th.Levels[i]
		}
	}
	return th.Levels[3]
}

func RenderError(w io.Writer, th Theme, f Font, title, detail string) error {
	return templates.ExecuteTemplate(w, "error.svg", ErrorView{Theme: th, Font: f, Title: title, Detail: detail})
}

package src

type Theme struct {
	Card       string
	Fg         string
	Muted      string
	PillBg     string
	Success    string
	Label1     string
	Label2     string
	Label3     string
	Empty      string
	Levels     [4]string
	Difficulty [3]string
}

var difficulty = [3]string{"#1CBABA", "#FFB700", "#F63737"}

func GetTheme(name string) Theme {
	if name == "dark" {
		return Theme{
			Card:       "#282828",
			Fg:         "#F5F5F5",
			Muted:      "#A8A8A8",
			PillBg:     "rgba(255,255,255,0.06)",
			Success:    "#029727",
			Label1:     "#FFFFFF",
			Label2:     "rgba(239,241,246,0.75)",
			Label3:     "rgba(239,242,246,0.6)",
			Empty:      "rgba(255,255,255,0.08)",
			Levels:     [4]string{"#016620", "#109932", "#28C244", "#7FE18B"},
			Difficulty: difficulty,
		}
	}
	return Theme{
		Card:       "#FFFFFF",
		Fg:         "#1A1A1A",
		Muted:      "#737373",
		PillBg:     "rgba(0,0,0,0.02)",
		Success:    "#02B128",
		Label1:     "#262626",
		Label2:     "rgba(38,38,38,0.75)",
		Label3:     "rgba(60,60,67,0.6)",
		Empty:      "rgba(0,0,0,0.04)",
		Levels:     [4]string{"#AFF0B4", "#4DD161", "#01B328", "#008024"},
		Difficulty: difficulty,
	}
}

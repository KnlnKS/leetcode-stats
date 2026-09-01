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
			PillBg:     "#FFFFFF0F",
			Success:    "#029727",
			Label1:     "#FFFFFF",
			Label2:     "#EFF1F6BF",
			Label3:     "#EFF2F699",
			Empty:      "#FFFFFF14",
			Levels:     [4]string{"#016620", "#109932", "#28C244", "#7FE18B"},
			Difficulty: difficulty,
		}
	}
	return Theme{
		Card:       "#FFFFFF",
		Fg:         "#1A1A1A",
		Muted:      "#737373",
		PillBg:     "#00000005",
		Success:    "#02B128",
		Label1:     "#262626",
		Label2:     "#262626BF",
		Label3:     "#3C3C4399",
		Empty:      "#0000000A",
		Levels:     [4]string{"#AFF0B4", "#4DD161", "#01B328", "#008024"},
		Difficulty: difficulty,
	}
}

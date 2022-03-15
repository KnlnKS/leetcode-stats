package src

type Theme struct {
	Background       string
	Text             string
	SecondaryText    string
	EasyDifficulty   string
	MediumDifficulty string
	HardDifficulty   string
}

func GenerateTheme(theme string) Theme {
	switch theme {
	case "dark":
		return Theme{
			Background:       "#292A2B",
			Text:             "#CFCCC7",
			SecondaryText:    "#7A807C",
			EasyDifficulty:   "#43A047",
			MediumDifficulty: "#FB8C00",
			HardDifficulty:   "#E91E63",
		}
	case "midnight":
		return Theme{
			Background:       "#1A1B27",
			Text:             "#70A4FC",
			SecondaryText:    "#BE90F2",
			EasyDifficulty:   "#38BCAD",
			MediumDifficulty: "#38BCAD",
			HardDifficulty:   "#38BCAD",
		}
	default:
		return Theme{
			Background:       "#FFF",
			Text:             "#262626",
			SecondaryText:    "#8A8A8E",
			EasyDifficulty:   "#43A047",
			MediumDifficulty: "#FB8C00",
			HardDifficulty:   "#E91E63",
		}
	}
}

package src

type Theme struct {
	Background       string
	Text             string
	Text2            string
	Text3            string
	Text4            string
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
			Text2:            "rgba(239,241,246,.75)",
			Text3:            "rgba(239,242,246,.6)",
			Text4:            "rgba(235,235,245,.3)",
			EasyDifficulty:   "#43A047",
			MediumDifficulty: "#FB8C00",
			HardDifficulty:   "#E91E63",
		}
	default:
		return Theme{
			Background:       "#FFF",
			Text:             "#262626",
			Text2:            "rgba(38,38,38,.75)",
			Text3:            "rgba(60, 60, 67, .6)",
			Text4:            "rgba(60, 60, 67, .3)",
			EasyDifficulty:   "#43A047",
			MediumDifficulty: "#FB8C00",
			HardDifficulty:   "#E91E63",
		}
	}
}

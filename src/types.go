package src

type SubmissionData struct {
	AllQuestionsCount []struct {
		Count int `json:"count"`
	} `json:"allQuestionsCount"`
	MatchedUser struct {
		ProblemsSolvedBeatsStats []struct {
			Percentage float64 `json:"percentage"`
		} `json:"problemsSolvedBeatsStats"`
		SubmitStatsGlobal struct {
			AcSubmissionNum []struct {
				Count      int     `json:"count"`
				Percentage float64 `json:"percentage"`
			} `json:"acSubmissionNum"`
		} `json:"submitStatsGlobal"`
	} `json:"matchedUser"`
}

type UserQuestionCount struct {
	UserProfileUserQuestionProgress struct {
		NumAcceptedQuestions []struct {
			Count int `json:"count"`
		} `json:"numAcceptedQuestions"`
		NumFailedQuestions []struct {
			Count int `json:"count"`
		} `json:"numFailedQuestions"`
		NumUntouchedQuestions []struct {
			Count int `json:"count"`
		} `json:"numUntouchedQuestions"`
	} `json:"userProfileUserQuestionProgress"`
}

type UserProblemsSolved struct {
	ProblemsSolvedBeatsStats []struct {
		Percentage float64 `json:"percentage"`
	} `json:"problemsSolvedBeatsStats"`
}

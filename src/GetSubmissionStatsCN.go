package src

import (
	"context"
	"math"

	"github.com/machinebox/graphql"
)

func GetSubmissionStatsCN(username string) (SubmissionData, error) {
	client1 := graphql.NewClient("https://leetcode.cn/graphql")

	req1 := graphql.NewRequest(`
    query userQuestionProgress($userSlug: String!) {
		userProfileUserQuestionProgress(userSlug: $userSlug) {
		  numAcceptedQuestions {
			count
		  }
		  numFailedQuestions {
			count
		  }
		  numUntouchedQuestions {
			count
		  }
		}
	  }
  `)
	req1.Var("userSlug", username)

	ctx1 := context.Background()

	var respData1 UserQuestionCount
	if err := client1.Run(ctx1, req1, &respData1); err != nil {
		return SubmissionData{}, err
	}

	client2 := graphql.NewClient("https://leetcode.cn/graphql/noj-go/")

	req2 := graphql.NewRequest(`
    query problemSolvedBeatsStats($userSlug: String!) {
		problemsSolvedBeatsStats(userSlug: $userSlug) {
		  percentage
		}
	  }
  `)
	req2.Var("userSlug", username)

	ctx2 := context.Background()

	var respData2 UserProblemsSolved
	if err := client2.Run(ctx2, req2, &respData2); err != nil {
		return SubmissionData{}, err
	}

	var respData SubmissionData

	respData.AllQuestionsCount = []struct {
		Count int `json:"count"`
	}{
		{
			Count: 0,
		},
		{
			Count: 0,
		},
		{
			Count: 0,
		},
		{
			Count: 0,
		},
	}
	for i := 0; i < 3; i++ {
		respData.AllQuestionsCount[i+1].Count = respData1.UserProfileUserQuestionProgress.NumAcceptedQuestions[i].Count + respData1.UserProfileUserQuestionProgress.NumFailedQuestions[i].Count + respData1.UserProfileUserQuestionProgress.NumUntouchedQuestions[i].Count
	}
	respData.AllQuestionsCount[0].Count = respData.AllQuestionsCount[1].Count + respData.AllQuestionsCount[2].Count + respData.AllQuestionsCount[3].Count

	respData.MatchedUser.ProblemsSolvedBeatsStats = respData2.ProblemsSolvedBeatsStats
	for i := 0; i < len(respData.MatchedUser.ProblemsSolvedBeatsStats); i++ {
		respData.MatchedUser.ProblemsSolvedBeatsStats[i].Percentage = math.Round(respData.MatchedUser.ProblemsSolvedBeatsStats[i].Percentage*10) / 10
	}

	respData.MatchedUser.SubmitStatsGlobal.AcSubmissionNum = []struct {
		Count      int     `json:"count"`
		Percentage float64 `json:"percentage"`
	}{
		{
			Count:      0,
			Percentage: 0,
		},
		{
			Count:      0,
			Percentage: 0,
		},
		{
			Count:      0,
			Percentage: 0,
		},
		{
			Count:      0,
			Percentage: 0,
		},
	}
	for i := 0; i < 3; i++ {
		respData.MatchedUser.SubmitStatsGlobal.AcSubmissionNum[i+1].Count = respData1.UserProfileUserQuestionProgress.NumAcceptedQuestions[i].Count
		percentage := float64(respData1.UserProfileUserQuestionProgress.NumAcceptedQuestions[i].Count) / float64(respData.AllQuestionsCount[i+1].Count)
		respData.MatchedUser.SubmitStatsGlobal.AcSubmissionNum[i+1].Percentage = math.Round(percentage*1000000) / 10000
	}
	respData.MatchedUser.SubmitStatsGlobal.AcSubmissionNum[0].Count = respData.MatchedUser.SubmitStatsGlobal.AcSubmissionNum[1].Count + respData.MatchedUser.SubmitStatsGlobal.AcSubmissionNum[2].Count + respData.MatchedUser.SubmitStatsGlobal.AcSubmissionNum[3].Count

	return respData, nil
}

package src

import (
	"context"
	"math"

	"github.com/machinebox/graphql"
)

func GetSubmissionStats(username string) (SubmissionData, error) {
	client := graphql.NewClient("https://leetcode.com/graphql")

	req := graphql.NewRequest(`
    query userProblemsSolved($username: String!) {
		allQuestionsCount {
		  count
		}
		matchedUser(username: $username) {
		  problemsSolvedBeatsStats {
			percentage
		  }
		  submitStatsGlobal {
			acSubmissionNum {
			  count
			}
		  }
		}
	  }
  `)
	req.Var("username", username)

	ctx := context.Background()

	var respData SubmissionData
	if err := client.Run(ctx, req, &respData); err != nil {
		return respData, err
	}

	for i := 0; i < len(respData.MatchedUser.ProblemsSolvedBeatsStats); i++ {
		respData.MatchedUser.ProblemsSolvedBeatsStats[i].Percentage = math.Round(respData.MatchedUser.ProblemsSolvedBeatsStats[i].Percentage*10) / 10
	}

	for i := 0; i < len(respData.MatchedUser.SubmitStatsGlobal.AcSubmissionNum); i++ {
		percentage := float64(respData.MatchedUser.SubmitStatsGlobal.AcSubmissionNum[i].Count) / float64(respData.AllQuestionsCount[i].Count)
		respData.MatchedUser.SubmitStatsGlobal.AcSubmissionNum[i].Percentage = math.Round(percentage*1000000) / 10000
	}

	return respData, nil
}

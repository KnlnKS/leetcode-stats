package src

import (
	"context"

	"github.com/machinebox/graphql"
)

type SubmissionData struct {
	AllQuestionsCount []struct {
		Difficulty string `json:"difficulty"`
		Count      int    `json:"count"`
	} `json:"allQuestionsCount"`
	MatchedUser struct {
		SubmitStats struct {
			AcSubmissionNum []struct {
				Difficulty  string `json:"difficulty"`
				Count       int    `json:"count"`
				Submissions int    `json:"submissions"`
			} `json:"acSubmissionNum"`
		} `json:"submitStats"`
	} `json:"matchedUser"`
}

func GetSubmissionStats(username string) (SubmissionData, error) {
	client := graphql.NewClient("https://leetcode.com/graphql")

	req := graphql.NewRequest(`
	query getUserProfile($username: String!) {
	  allQuestionsCount {
		difficulty
		count
	  }
	  matchedUser(username: $username) {
		submitStats {
		  acSubmissionNum {
			difficulty
			count
			submissions
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
	return respData, nil
}

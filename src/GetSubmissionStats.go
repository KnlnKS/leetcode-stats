package src

import (
	"context"

	"github.com/machinebox/graphql"
)

func GetSubmissionStats(username string) (interface{}, error) {
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

	var respData interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return respData, nil
}

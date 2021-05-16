import { request, gql } from "graphql-request";

const query = gql`
  query getUserProfile($username: String!) {
    allQuestionsCount {
      difficulty
      count
      __typename
    }
    matchedUser(username: $username) {
      username
      socialAccounts
      githubUrl
      contributions {
        points
        questionCount
        testcaseCount
        __typename
      }
      profile {
        starRating
        ranking
        __typename
      }
      submitStats {
        acSubmissionNum {
          difficulty
          count
          submissions
          __typename
        }
        totalSubmissionNum {
          difficulty
          count
          submissions
          __typename
        }
        __typename
      }

      __typename
    }
  }
`;

const variables = { username: "knlnks" };

request("https://leetcode.com/graphql", query, variables).then((data) =>
  console.log(data?.matchedUser?.submitStats?.acSubmissionNum)
);

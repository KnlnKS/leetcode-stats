import { request, gql } from "graphql-request";

import { IGetSubmissionStats } from "../types";

const query = gql`
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
`;

const variables = { username: "knlnks" };

export const getSubmissionStats = async (): Promise<IGetSubmissionStats> =>
  await request("https://leetcode.com/graphql", query, variables);

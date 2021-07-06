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

export const getSubmissionStats = async (
  username
): Promise<IGetSubmissionStats> =>
  await request("https://leetcode.com/graphql", query, { username });

import { request, gql } from "graphql-request";

const query = gql`
  query getUserProfile($username: String!) {
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
        realName
        websites
        countryName
        skillTags
        company
        school
        starRating
        aboutMe
        userAvatar
        reputation
        ranking
        __typename
      }
      __typename
    }
  }
`;

const variables = {"username": "knlnks"}

request("https://leetcode.com/graphql", query, variables).then((data) =>
  console.log(data)
);

export interface AllQuestionsCountNode {
  difficulty: "All" | "Easy" | "Medium" | "Hard";
  count: number;
}

export interface SubmissionCountNode extends AllQuestionsCountNode {
  submissions: number;
}

export interface IGetSubmissionStats {
  allQuestionsCount: AllQuestionsCountNode[];
  matchedUser: {
    submitStats: {
      acSubmissionNum: SubmissionCountNode[];
    };
  };
}

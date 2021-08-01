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

export type Theme = "light" | "dark" | "midnight";

export interface IGenerateStatSVG {
  allQuestionsCount: AllQuestionsCountNode[];
  acSubmissionNum: SubmissionCountNode[];
  theme?: Theme;
}

export type ThemeColours = {
  bg: string;
  text: string;
  secondaryText: string;
  altDiffText?: string;
};

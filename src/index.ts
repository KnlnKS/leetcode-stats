import * as tools from "simple-svg-tools";

import { generateStatSVG, getSubmissionStats } from "./svg";

const main = async () => {
  const {
    allQuestionsCount,
    matchedUser: {
      submitStats: { acSubmissionNum },
    },
  } = await getSubmissionStats();
  const svgString = generateStatSVG(allQuestionsCount, acSubmissionNum);
  const svg = new tools.SVG(svgString);
};

main();

import * as fs from "fs";

import { generateStatSVG, getSubmissionStats } from "./svg";

const main = async () => {
  const {
    allQuestionsCount,
    matchedUser: {
      submitStats: { acSubmissionNum },
    },
  } = await getSubmissionStats(process.env.LEETCODE_USERNAME);
  const svgString = generateStatSVG(allQuestionsCount, acSubmissionNum);
  fs.writeFile("generated/stats.svg", svgString, (err) => {
    if (err) console.log(err);
  });
};

main();

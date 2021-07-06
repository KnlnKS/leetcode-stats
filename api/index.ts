import express from "express";
import { generateStatSVG, getSubmissionStats } from "./svg";

const app = express();

app.get("/api", async (req, res) => {
  const { username } = req?.query;

  if (username) {
    const {
      allQuestionsCount,
      matchedUser: {
        submitStats: { acSubmissionNum },
      },
    } = await getSubmissionStats(username);
    res.setHeader("Content-Type", "image/svg+xml");
    res.end(generateStatSVG(allQuestionsCount, acSubmissionNum));
  }
});

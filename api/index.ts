import express from "express";
import { generateStatSVG, getSubmissionStats } from "./svg";

const app = express();

app.get("/api", async (req, res) => {
  /*
  const { username } = req?.query;

  if (username) {
    const {
      allQuestionsCount,
      matchedUser: {
        submitStats: { acSubmissionNum },
      },
    } = await getSubmissionStats(username);


    res.setHeader("Content-Type", "image/svg+xml");
    res.setHeader('Cache-Control', 's-max-age=60, stale-while-revalidate');
    res.status(200).end(generateStatSVG(allQuestionsCount, acSubmissionNum));
  } else res.status(400).end("Missing Parameter");
  */
 res.end("yo")
});

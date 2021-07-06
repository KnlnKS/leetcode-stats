import express from "express";
import { generateStatSVG, getSubmissionStats } from "./svg";

const app = express();

app.get("/", (_, res) => res.send("<em>Nothing here...</em>"));

app.get("/api", async (req, res, next) => {
  const { username } = req?.query;

  if (username) {
    let data;

    try {
      data = await getSubmissionStats(username);
      const {
        allQuestionsCount,
        matchedUser: {
          submitStats: { acSubmissionNum },
        },
      } = data;

      res.setHeader("Content-Type", "image/svg+xml");
      res.setHeader("Cache-Control", "s-max-age=60, stale-while-revalidate");
      res.status(200).send(generateStatSVG(allQuestionsCount, acSubmissionNum));
    } catch (err) {
      res.status(400).send("Username does not exist");
    }
  } else res.status(400).send("Missing Parameter");
});

const port = process.env.PORT || 2001;
app.listen(port, () =>
  console.log(`Server running on ${port}, http://localhost:${port}`)
);

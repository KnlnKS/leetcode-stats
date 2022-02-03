import { generateStatSVG, getSubmissionStats } from "../../src";

export default async function handler(req, res) {
  const { username, theme = "light" } = req?.query;

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
      res
        .status(200)
        .send(generateStatSVG({ allQuestionsCount, acSubmissionNum, theme }));
    } catch (err) {
      res.status(400).send("Username does not exist");
    }
  } else res.status(400).send("Missing Parameter");
}

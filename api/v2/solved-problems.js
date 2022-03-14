import Chrome from "../../chrome";

async function handler(req, res) {
  const { username, theme = "light" } = req.query;
  if (!username) {
    res.status(400).send("Username is required");
    return;
  }

  const chrome = new Chrome();
  await chrome.getPage(`https://leetcode.com/${username}`, theme);
  const image = await chrome.takeScreenshot(".min-w-max");

  res.setHeader("Content-Type", "image/png");
  res.setHeader('Cache-Control', 's-maxage=3600');
  res.status(200).send(image);
}

export default handler;

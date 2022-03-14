import Chrome from "../../chrome";

async function handler(req, res) {
  const { username, theme = "light" } = req.query;
  if (!username) {
    res.status(400).send("Username is required");
    return;
  }
  if (theme !== "light" && theme !== "dark") {
    res.status(400).send("Theme must be either 'light' or 'dark'");
    return;
  }
  const chrome = new Chrome();
  await chrome.getPage(`https://leetcode.com/${username}`, theme);
  const image = await chrome.takeScreenshot(".min-w-max");
  chrome.close();

  res.setHeader("Content-Type", "image/png");
  res.status(200).send(image);
}

export default handler;

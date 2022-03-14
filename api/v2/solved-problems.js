import getBrowserInstance from "../../browser/getBrowserInstance";

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

  const browser = await getBrowserInstance();
  const page = await browser.newPage();
  await page.goto(`https://leetcode.com/${username}`);
  if (theme === "dark") {
    await page.emulateMediaFeatures([
      {
        name: "prefers-color-scheme",
        value: "dark",
      },
    ]);
  }
  const element = await page.$(".min-w-max");
  const image = await element.screenshot();
  browser.close();

  res.setHeader("Content-Type", "image/png");
  res.status(200).send(image);
}

export default handler;

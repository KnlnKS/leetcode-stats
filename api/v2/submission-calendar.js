import Chrome from "../../src/chrome";

async function handler(req, res) {
  const { username, theme = "light" } = req.query;
  if (!username) {
    res.status(400).send("Username is required");
    return;
  }

  const chrome = new Chrome();
  await chrome.getPage(`https://leetcode.com/${username}`, theme);
  await chrome.page.evaluate(() => {
    const listboxButton = document.querySelector(
      "#headlessui-listbox-button-5"
    );
    listboxButton.parentNode.removeChild(listboxButton);
    const infoButton = document.evaluate(
      '//*[@id="__next"]/body/div/div/div/div[2]/div[2]/div/div[1]/div[1]/div',
      document,
      null,
      XPathResult.FIRST_ORDERED_NODE_TYPE,
      null
    ).singleNodeValue;
    infoButton.parentNode.removeChild(infoButton);
  });
  const image = await chrome.takeScreenshot(
    '//*[@id="__next"]/body/div/div/div/div[2]/div[2]/div',
    true
  );

  res.setHeader("Content-Type", "image/png");
  res.setHeader('Cache-Control', 's-maxage=3600');
  res.status(200).send(image);
}

export default handler;

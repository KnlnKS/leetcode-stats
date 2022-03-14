import Chrome from "../../chrome";

async function handler(event) {
  const { username, theme = "light" } = JSON.parse(event.body);
  if (!username) {
    return {
      statusCode: 400,
      body: JSON.stringify({ message: `Error, no username provided!` }),
    };
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

  return {
    statusCode: 200,
    headers: {
      "Content-type": "image/png",
    },
    body: image.toString("base64"),
    isBase64Encoded: true,
  };
}

module.exports = { handler }

const Chrome = require("./vercel/src/chrome/index.js");

exports.handler = async function (event) {
  const { username = null, theme = "light" } = event.queryStringParameters;
  if (!username) {
    return {
      statusCode: 400,
      body: JSON.stringify({ message: `Error, no username provided!` }),
    };
  }

  const chrome = new Chrome();
  await chrome.getPage(`https://leetcode.com/${username}`, theme);
  const image = await chrome.takeScreenshot(".min-w-max");

  return {
    statusCode: 200,
    headers: {
      "Content-type": "image/png",
    },
    body: image.toString("base64"),
    isBase64Encoded: true,
  };
};

const chromium = require("chrome-aws-lambda");
const redis = require("redis");

class Chrome {
  constructor() {
    this.browser = null;
    this.page = null;
  }

  async getBrowserInstance() {
    const executablePath = await chromium.executablePath;

    this.browser = await chromium.puppeteer.launch({
      args: chromium.args,
      headless: true,
      defaultViewport: {
        width: 1920,
        height: 1080,
        deviceScaleFactor: 2,
      },
      ignoreHTTPSErrors: true,
      executablePath,
      headless: chromium.headless,
    });
  }

  async getPage(url, theme) {
    if (!this.browser) {
      await this.getBrowserInstance();
    }
    this.page = await this.browser.newPage();
    await this.page.goto(url);

    if (theme === "dark") {
      await this.page.emulateMediaFeatures([
        {
          name: "prefers-color-scheme",
          value: "dark",
        },
      ]);
    }
  }

  async takeScreenshot(selector, isXPath) {
    let element = await (isXPath
      ? this.page.$x(selector)
      : this.page.$(selector));
    if (Array.isArray(element)) {
      element = element[0];
    }
    const image = await element.screenshot();
    return image;
  }

  async close() {
    await this.browser.close();
  }

  async evaluate(fn) {
    return await this.page.evaluate(fn);
  }
}

exports.handler = async function (event) {
  const { username, theme = "light" } = event.queryStringParameters;
  if (!username) {
    return {
      statusCode: 400,
      body: JSON.stringify({ message: `Error, no username provided!` }),
    };
  }
  if (!["light", "dark"].includes(theme)) {
    return {
      statusCode: 400,
      body: JSON.stringify({ message: `Error, invalid theme!` }),
    };
  }

  const client = redis.createClient({
    host: process.env.REDIS_HOST,
    port: process.env.REDIS_PORT,
    password: process.env.REDIS_PASSWORD,
  });
  await client.connect();

  const value = await client
    .get(`${username}:${theme}`)
    .then((resp) => JSON.parse(resp));
  if (value.headers["Cache-Control"] < Date.now()) {
    return value;
  }

  const chrome = new Chrome();
  await chrome.getPage(`https://leetcode.com/${username}`, theme);
  const image = await chrome.takeScreenshot(".min-w-max");
  const resp = {
    statusCode: 200,
    headers: {
      "Content-type": "image/png",
      "Cache-Control": Date.now() + 1000 * 60 * 60,
    },
    body: image.toString("base64"),
    isBase64Encoded: true,
  };

  await client.set(`${username}:${theme}`, JSON.stringify(resp));
  return resp;
};

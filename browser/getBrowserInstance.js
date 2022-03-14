import chromium from "chrome-aws-lambda";

export default async function getBrowserInstance() {
  const executablePath = await chromium.executablePath;
  const options = {
    args: chromium.args,
    headless: true,
    defaultViewport: {
      width: 1920,
      height: 1080,
    },
    ignoreHTTPSErrors: true,
  };

  return !executablePath
    ? require("puppeteer").launch(options)
    : chromium.puppeteer.launch({
        ...options,
        executablePath,
        headless: chromium.headless,
      });
}

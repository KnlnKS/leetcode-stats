import { minify } from "html-minifier";
import { IGenerateStatSVG, Theme, ThemeColours } from "../types";

const getThemeColours = (theme?: Theme): ThemeColours => {
  if (theme === "dark") {
    return {
      bg: "#292A2B",
      text: "#cfccc7",
      secondaryText: "#7A807C",
    };
  } else if (theme === "midnight") {
    return {
      bg: "#1A1B27",
      text: "#70A4FC",
      secondaryText: "#BE90F2",
      altDiffText: "#38BCAD",
    };
  }
  return {
    bg: "#fff",
    text: "#262626",
    secondaryText: "rgba(60, 60, 67, 0.6)",
  };
};

export const generateStatSVG = ({
  allQuestionsCount,
  acSubmissionNum,
  theme,
}: IGenerateStatSVG) => {
  const { bg, text, secondaryText, altDiffText } = getThemeColours(theme);
  return minify(
    `
    <svg width="330" height="180" xmlns="http://www.w3.org/2000/svg">
    <style>
      svg {
        font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial, sans-serif;
        font-size: 14px;
        line-height: 1.5;
      }
      #background {
        width: calc(100% - 10px);
        height: calc(100% - 10px);
        fill: ${bg};
        rx: 8px;
        ry: 8px;
      }
      .total-solved-container {
        display: -webkit-box;
        display: -ms-flexbox;
        display: flex;
        -webkit-box-pack: justify;
        -ms-flex-pack: justify;
        justify-content: space-between;
        color: #9e9e9e;
        height: 33px;
      }
      .stat-wrapper {
        display: flex;
        flex-direction: column;
        text-align: start;
      }
      .problems-solved {
        color: ${secondaryText};
        font-size: 12px;
        line-height: 17px;
        margin-bottom: 2px;
        white-space: nowrap;
        font-weight: 500;
      }
      .total-solved-count {
        font-size: 22px;
        font-weight: 600;
        line-height: 100%;
        white-space: nowrap;
        color: ${text};
        font-size: 22px;
      }
      .difficulty-label {
        font-size: 12px;
        font-weight: normal;
        line-height: 17px;
        margin-bottom: 2px;
        white-space: nowrap;
      }
      .easy {
        color: ${altDiffText ?? "rgb(67, 160, 71)"};
      }
      .medium {
        color: ${altDiffText ?? "rgb(251, 140, 0)"};
      }
      .hard {
        color: ${altDiffText ?? "rgb(233, 30, 99)"};
      }
      .total-solved-container .total-count::before {
        content: '/';
        margin: 0 1px;
      }
      .total-solved-container .total-count {
        color: ${secondaryText};
        font-size: 12px;
        font-weight: 500;
        line-height: 14px;
      }
      .solved {
        color: ${text};
        font-size: 14px;
        font-weight: 600;
        line-height: 100%;
        white-space: nowrap;
      }
      .top {
        height: 108px;
      }
      foreignObject {
        width: calc(100% - 10px - 32px);
        height: calc(100% - 10px - 24px);
      }
    </style>
    <g>
      <rect x="5" y="5" id="background" />
      <g>
        <foreignObject x="21" y="17" width="318" height="176">
          <div xmlns="http://www.w3.org/1999/xhtml">

            <div class="stat-wrapper top" size="108">
              <div class="problems-solved">Problems Solved</div>
              <div class="total-solved-count">${acSubmissionNum[0].count}</div>
            </div>

            <div class="total-solved-container">
              <div class="stat-wrapper" data-difficulty="Easy">
                <div class="difficulty-label easy">Easy</div>
                <div class="solved">${
                  acSubmissionNum[1].count
                }<span class="total-count">${allQuestionsCount[1].count}</span>
                </div>
              </div>
              <div class="stat-wrapper" data-difficulty="Medium">
                <div class="difficulty-label medium">Medium</div>
                <div class="solved">${
                  acSubmissionNum[2].count
                }<span class="total-count">${allQuestionsCount[2].count}</span>
                </div>
              </div>
              <div class="stat-wrapper" data-difficulty="Hard">
                <div class="difficulty-label hard">Hard</div>
                <div class="solved">${
                  acSubmissionNum[3].count
                }<span class="total-count">${allQuestionsCount[3].count}</span>
                </div>
              </div>
            </div>

          </div>
        </foreignObject>
      </g>
    </g>
  </svg>`,
    { minifyCSS: true }
  );
};

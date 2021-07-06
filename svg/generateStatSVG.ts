import { AllQuestionsCountNode, SubmissionCountNode } from "../types";

export const generateStatSVG = (
  allQuestionsCount: AllQuestionsCountNode[],
  acSubmissionNum: SubmissionCountNode[]
) =>
  `
    <svg width="330" height="180" xmlns="http://www.w3.org/2000/svg">
    <style>
      svg {
        font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial, sans-serif, Apple Color Emoji, Segoe UI Emoji;
        font-size: 14px;
        line-height: 1.5;
      }

      #background {
        width: calc(100% - 10px);
        height: calc(100% - 10px);
        fill: white;
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
        color: rgba(60, 60, 67, 0.6);
        font-size: 12px;
        line-height: 17px;
        margin-bottom: 2px;
        white-space: nowrap;
        font-weight: 500;
      }

      .total-solved-count {
        font-size: 22px;
        color: #212121;
        font-weight: 600;
        line-height: 100%;
        white-space: nowrap;
        color: rgb(38, 38, 38);
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
        color: rgb(67, 160, 71);
      }

      .medium {
        color: rgb(251, 140, 0);
      }

      .hard {
        color: rgb(233, 30, 99);
      }

      .total-solved-container .total-count::before {
        content: '/';
        margin: 0 1px;
      }

      .total-solved-container .total-count {
        color: rgba(60, 60, 67, 0.3);
        font-size: 12px;
        font-weight: 500;
        line-height: 14px;
      }

      .solved {
        color: rgba(38, 38, 38, 0.75);
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
                <div class="solved">${acSubmissionNum[1].count}<span class="total-count">${allQuestionsCount[1].count}</span>
                </div>
              </div>
              <div class="stat-wrapper" data-difficulty="Medium">
                <div class="difficulty-label medium">Medium</div>
                <div class="solved">${acSubmissionNum[2].count}<span class="total-count">${allQuestionsCount[2].count}</span>
                </div>
              </div>
              <div class="stat-wrapper" data-difficulty="Hard">
                <div class="difficulty-label hard">Hard</div>
                <div class="solved">${acSubmissionNum[3].count}<span class="total-count">${allQuestionsCount[3].count}</span>
                </div>
              </div>
            </div>


          </div>

        </foreignObject>
      </g>
    </g>
  </svg>`;

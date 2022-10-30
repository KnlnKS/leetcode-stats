package templates

import (
	"text/template"
)

var (
	submissionStatsTmpl *template.Template = nil
)

func GetSubmissionStatsTemplate() *template.Template {
	if submissionStatsTmpl == nil {
		submissionStatsTmpl, _ = template.New("SubmissionStats").Parse(`
			{{define "T"}}
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

				            fill: {{.Theme.Background}};

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
				            color: {{.Theme.SecondaryText}};
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

				            color: {{.Theme.Text}};
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
				            color: {{.Theme.EasyDifficulty}};
				        }

				        .medium {
				            color: {{.Theme.MediumDifficulty}};
				        }

				        .hard {
				            color: {{.Theme.HardDifficulty}};
				        }

				        .total-solved-container .total-count::before {
				            content: '/';
				            margin: 0 1px;
				        }

				        .total-solved-container .total-count {
				            color: {{.Theme.SecondaryText}};
				            font-size: 12px;
				            font-weight: 500;
				            line-height: 14px;
				        }

				        .solved {
				            color: {{.Theme.Text}};
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
				                        <div class="total-solved-count">{{(index .SubmissionData.MatchedUser.SubmitStats.AcSubmissionNum 0).Count}}</div>
				                    </div>
				                    <div class="total-solved-container">
				                        <div class="stat-wrapper" data-difficulty="Easy">
				                            <div class="difficulty-label easy">Easy</div>
				                            <div class="solved">
				                                {{(index .SubmissionData.MatchedUser.SubmitStats.AcSubmissionNum 1).Count}}<span class="total-count">{{(index .SubmissionData.AllQuestionsCount 1).Count}}</span>
				                            </div>
				                        </div>
				                        <div class="stat-wrapper" data-difficulty="Medium">
				                            <div class="difficulty-label medium">Medium</div>
				                            <div class="solved">
				                                {{(index .SubmissionData.MatchedUser.SubmitStats.AcSubmissionNum 2).Count}}<span class="total-count">{{(index .SubmissionData.AllQuestionsCount 2).Count}}</span>
				                            </div>
				                        </div>
				                        <div class="stat-wrapper" data-difficulty="Hard">
				                            <div class="difficulty-label hard">Hard</div>
				                            <div class="solved">
				                                {{(index .SubmissionData.MatchedUser.SubmitStats.AcSubmissionNum 3).Count}}<span class="total-count">{{(index .SubmissionData.AllQuestionsCount 3).Count}}</span>
				                            </div>
				                        </div>
				                    </div>
				                </div>
				            </foreignObject>
				        </g>
				    </g>
				</svg>
			{{end}}
	`)
	}
	return submissionStatsTmpl
}

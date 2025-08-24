package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/dongjune8931/SAA_BOT/internal/ai"
	"github.com/dongjune8931/SAA_BOT/internal/notion"
	"github.com/spf13/cobra"
)

// 사용자 답변을 저장할 구조체
var qs = []*survey.Question{
	{
		Name:     "title",
		Prompt:   &survey.Input{Message: "문제 번호(제목)를 입력하세요:"},
		Validate: survey.Required,
	},
	{
		Name: "question",
		Prompt: &survey.Multiline{ // Editor에서 Multiline으로 변경
			Message: "문제 내용을 모두 붙여넣은 후, Ctrl+D를 눌러 입력을 완료하세요.",
		},
		Validate: survey.Required,
	},
}

var addCmd = &cobra.Command{
	// ... 이하 코드는 이전과 동일합니다 ...
	Use:   "add",
	Short: "대화형으로 새로운 SAA 문제를 분석하고 Notion에 추가합니다.",
	Long:  `add 명령어를 실행하면, 문제 번호와 내용을 순서대로 질문합니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			Title    string `survey:"title"`
			Question string `survey:"question"`
		}{}

		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println("\n작업이 취소되었습니다.")
			return
		}

		fmt.Println("\n🧠 AI가 문제를 분석하고 있습니다. 잠시만 기다려주세요...")
		analysisResult, err := ai.AnalyzeQuestion(answers.Question)
		if err != nil {
			fmt.Printf("❌ AI 분석 중 오류 발생: %v\n", err)
			return
		}
		fmt.Println("✨ AI 분석 완료!")

		fullContent := fmt.Sprintf("## 📝 문제 원문\n%s\n---\n%s", answers.Question, analysisResult)

		fmt.Println("📝 Notion에 페이지를 생성하고 있습니다...")
		err = notion.CreatePage(answers.Title, fullContent)
		if err != nil {
			fmt.Printf("❌ Notion 페이지 생성 중 오류 발생: %v\n", err)
			return
		}
		fmt.Println("✅ 성공! Notion에서 새로운 문제 노트를 확인하세요.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const asciiArt = `

  ██████╗  █████╗  █████╗    ██████╗  ██████╗ ████████╗
 ██╔════╝ ██╔══██╗██╔══██╗   ██╔══██╗██╔═══██╗╚══██╔══╝
 ███████╗ ███████║███████║   ██████╔╝██║   ██║   ██║   
 ╚════██║ ██╔══██║██╔══██║   ██╔══██╗██║   ██║   ██║   
 ███████╔╝ ██║  ██║██║  ██║   ██████╔╝╚██████╔╝   ██║   
 ╚══════╝  ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═════╝  ╚═════╝    ╚═╝                                                

`
const mainHelpText = "AI를 이용해 SAA 문제 노트를 Notion에 자동으로 정리하는 CLI 툴입니다.\n'add'를 입력해 새 노트를 추가하거나 'help'로 명령어 목록을 확인하세요. 종료는 'exit'입니다."

var rootCmd = &cobra.Command{
	Use:   "saa-note",
	Short: "AI 기반 SAA 문제 노트 자동 정리 툴",
	Long:  mainHelpText,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "오류가 발생했습니다: '%s'", err)
		os.Exit(1)
	}
}

// PrintBanner는 프로그램 시작 시 환영 메시지를 출력하는 공개 함수입니다.
func PrintBanner() {
	fmt.Println(asciiArt)
	fmt.Println(mainHelpText)
	fmt.Println("--------------------------------------------------------------------")
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}

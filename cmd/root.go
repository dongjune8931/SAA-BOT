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

var rootCmd = &cobra.Command{
	Use:   "saa-note",
	Short: "AI를 이용해 SAA 문제 노트를 Notion에 자동으로 정리하는 CLI 툴입니다.",
	Long:  asciiArt + "\n\nAI를 이용해 SAA(AWS Certified Solutions Architect - Associate) 문제와 해설, 오답 노트를 Notion에 자동으로 정리해주는 CLI 툴입니다.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "오류가 발생했습니다: '%s'", err)
		os.Exit(1)
	}
}

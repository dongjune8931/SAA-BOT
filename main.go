package main

import (
	"bufio"
	"fmt"
	"github.com/dongjune8931/SAA_BOT/cmd"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func main() {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	// 1. 프로그램 시작 시 배너와 기본 설명 출력
	cmd.PrintBanner()

	// 2. 사용자 입력을 계속 받기 위한 스캐너 생성
	scanner := bufio.NewScanner(os.Stdin)

	// 3. 전용 셸 루프 시작
	for {
		// 4. 전용 프롬프트 출력
		fmt.Print("SAA-BOT > ")

		// 5. 사용자 입력 대기
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		line := scanner.Text()

		// 6. 'exit' 또는 'quit' 입력 시 루프 종료
		if line == "exit" || line == "quit" {
			break
		}

		// 7. 입력받은 명령어를 공백 기준으로 분리
		args := strings.Fields(line)
		if len(args) == 0 {
			continue // 아무것도 입력 안 하면 다시 프롬프트 표시
		}

		// 8. Cobra가 처리할 수 있도록 인자 설정 후 실행
		rootCmd := cmd.GetRootCmd() // rootCmd를 가져오는 함수 필요
		rootCmd.SetArgs(args)
		rootCmd.Execute()
	}

	fmt.Println("SAA-BOT을 종료합니다.")
}

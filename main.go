package main

import (
	"github.com/dongjune8931/SAA_BOT/cmd"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// .env 파일에서 환경변수 로드
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	// Cobra 애플리케이션 실행
	cmd.Execute()
}

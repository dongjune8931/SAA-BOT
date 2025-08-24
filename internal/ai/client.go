package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const geminiAPIURL = "https://generativelanguage.googleapis.com/v1/models/gemini-1.5-flash:generateContent?key="

// Gemini API 요청/응답 구조체 정의
type GeminiRequest struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content Content `json:"content"`
}

// AI에게 문제 분석을 요청하는 함수
func AnalyzeQuestion(question string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY is not set")
	}

	prompt := buildPrompt(question)

	reqBody := GeminiRequest{
		Contents: []Content{
			{Parts: []Part{{Text: prompt}}},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	client := &http.Client{Timeout: 60 * time.Second}
	req, err := http.NewRequest("POST", geminiAPIURL+apiKey, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status: %s, body: %s", resp.Status, string(respBody))
	}

	var geminiResp GeminiResponse
	if err := json.Unmarshal(respBody, &geminiResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		return geminiResp.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("no content received from AI")
}

// 프롬프트를 만드는 헬퍼 함수
func buildPrompt(question string) string {
	return fmt.Sprintf(`
당신은 AWS SAA 시험을 준비하는 학생을 돕는 최고의 AWS 전문가 튜터입니다.
주어진 AWS SAA 문제와 선택지를 분석하여 다음 형식에 맞춰 Markdown으로 정리해주세요.

### 1. 문제 분석
이 문제가 어떤 AWS 지식을 테스트하는지, 핵심 키워드는 무엇인지 설명해줘.

### 2. 정답 및 해설
정답이 무엇이며, 왜 그 선택지가 정답인지 상세히 설명해줘.

### 3. 오답 노트
나머지 오답 선택지들이 각각 왜 틀렸는지 명확하게 설명해줘.

### 4. 핵심 AWS 서비스 개념
이 문제와 직접적으로 관련된 AWS 서비스들의 핵심 개념과 주요 특징을 정리해줘.

---
[문제 시작]
%s
[문제 끝]
`, question)
}

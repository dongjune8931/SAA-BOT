package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const notionAPIURL = "https://api.notion.com/v1/pages"

// Notion API 요청 구조체 정의
type NotionRequest struct {
	Parent     Parent     `json:"parent"`
	Properties Properties `json:"properties"`
	Children   []Block    `json:"children"`
}

type Parent struct {
	DatabaseID string `json:"database_id"`
}

type Properties struct {
	Title TitleProperty `json:"문제"` // 데이터베이스의 Title 속성 이름과 일치해야 함
}

type TitleProperty struct {
	Title []Text `json:"title"`
}

type Text struct {
	Text Content `json:"text"`
}

type Content struct {
	Content string `json:"content"`
}

type Block struct {
	Object    string    `json:"object"`
	Type      string    `json:"type"`
	Paragraph Paragraph `json:"paragraph,omitempty"`
}

type Paragraph struct {
	RichText []Text `json:"rich_text"`
}

// Notion에 새 페이지를 생성하는 함수
func CreatePage(title, content string) error {
	apiKey := os.Getenv("NOTION_API_KEY")
	databaseID := os.Getenv("NOTION_DATABASE_ID")
	if apiKey == "" || databaseID == "" {
		return fmt.Errorf("NOTION_API_KEY or NOTION_DATABASE_ID is not set")
	}

	// AI가 생성한 Markdown 텍스트를 Notion Block 구조로 변환
	var blocks []Block
	for _, line := range strings.Split(content, "\n") {
		blocks = append(blocks, Block{
			Object: "block",
			Type:   "paragraph",
			Paragraph: Paragraph{
				RichText: []Text{
					{Text: Content{Content: line}},
				},
			},
		})
	}

	reqBody := NotionRequest{
		Parent: Parent{DatabaseID: databaseID},
		Properties: Properties{
			Title: TitleProperty{
				Title: []Text{
					{Text: Content{Content: title}},
				},
			},
		},
		Children: blocks,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("POST", notionAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2022-06-28")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// 에러 응답 본문을 읽어 더 자세한 정보 제공
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status: %s, body: %s", resp.Status, string(bodyBytes))
	}

	return nil
}

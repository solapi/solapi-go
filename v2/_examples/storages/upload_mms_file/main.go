package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/solapi/solapi-go/v2/client"
	"github.com/solapi/solapi-go/v2/storages"
)

func mustReadAndEncode(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("failed to read file:", err)
		os.Exit(1)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func main() {
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	if apiKey == "" || apiSecret == "" {
		fmt.Println("API KEY 또는 API SECRET이 설정되지 않았습니다.")
		os.Exit(1)
	}

	// 예제 파일 경로
	exampleDir, _ := os.Getwd()
	filePath := filepath.Join(exampleDir, "test.jpg")

	encoded := mustReadAndEncode(filePath)

	c := client.NewClient(apiKey, apiSecret)
	req := storages.UploadFileRequest{
		File: encoded,
		Name: "test.jpg",
		Type: "MMS",
	}
	resp, err := c.Storages.Upload(context.Background(), req)
	if err != nil {
		fmt.Println("upload error:", err)
		os.Exit(1)
	}

	fmt.Printf("uploaded fileId=%s name=%s type=%s url=%s\n", resp.FileID, resp.Name, resp.Type, resp.URL)
}

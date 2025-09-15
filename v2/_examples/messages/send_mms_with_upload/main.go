package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/solapi/solapi-go/v2/client"
	"github.com/solapi/solapi-go/v2/messages"
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
	to := "수신번호 입력"
	from := "계정에 등록한 발신번호 입력"
	if apiKey == "" || apiSecret == "" || to == "" || from == "" {
		fmt.Println("환경변수(API_KEY, API_SECRET, TO, FROM)가 필요합니다.")
		os.Exit(1)
	}

	// 예제 이미지 경로: 현재 디렉터리 기준 test.jpg
	// 또는 임의 경로를 FILE_PATH 환경변수로 지정 가능
	filePath := os.Getenv("FILE_PATH")
	if filePath == "" {
		wd, _ := os.Getwd()
		filePath = filepath.Join(wd, "test.jpg")
	}

	encoded := mustReadAndEncode(filePath)

	c := client.NewClient(apiKey, apiSecret)

	// 1) 스토리지 업로드: 파일을 올리고 fileId를 받습니다.
	upReq := storages.UploadFileRequest{
		File: encoded,
		Name: filepath.Base(filePath),
		Type: "MMS",
	}
	upRes, err := c.Storages.Upload(context.Background(), upReq)
	if err != nil {
		fmt.Println("upload error:", err)
		os.Exit(1)
	}

	// 2) MMS 전송: messages.Message의 ImageID에 업로드된 fileId 설정
	// Subject(제목)은 필요에 따라 설정하세요.
	// Type은 필요에 따라 설정하세요.
	msg := messages.Message{
		To:      to,
		From:    from,
		Type:    "MMS",
		Subject: "MMS 발송 테스트",
		Text:    "이미지 한 장이 포함된 MMS",
		ImageID: upRes.FileID,
	}

	res, err := c.Send(msg)
	if err != nil {
		fmt.Println("send error:", err)
		os.Exit(1)
	}
	fmt.Printf("groupId=%s total=%d registeredFailed=%d\n", res.GroupID, res.GroupInfo.Count.Total, res.GroupInfo.Count.RegisteredFailed)
}

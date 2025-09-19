package main

import (
	"context"
	"fmt"
	"os"

	"github.com/solapi/solapi-go/v2/client"
	"github.com/solapi/solapi-go/v2/groups"
	"github.com/solapi/solapi-go/v2/messages"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	groupId := os.Getenv("GROUP_ID")
	if apiKey == "" || apiSecret == "" || groupId == "" {
		fmt.Println("환경변수(API_KEY, API_SECRET, GROUP_ID)가 필요합니다.")
		os.Exit(1)
	}

	c := client.NewClient(apiKey, apiSecret)

	req := groups.AddGroupMessagesRequest{
		Messages: []messages.Message{
			{To: "수신번호", From: "발신번호", Text: "그룹에 추가되는 첫 번째 메시지"},
			{To: "수신번호", From: "발신번호", Text: "그룹에 추가되는 두 번째 메시지"},
		},
	}

	res, err := c.Groups.AddMessages(context.Background(), groupId, req)
	if err != nil {
		fmt.Println("add messages error:", err)
		os.Exit(1)
	}

	fmt.Printf("groupId=%s total=%d registeredFailed=%d\n", groupId, res.GroupInfo.Count.Total, res.GroupInfo.Count.RegisteredFailed)
}



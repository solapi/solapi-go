package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/solapi/solapi-go/v2/client"
	"github.com/solapi/solapi-go/v2/groups"
	"github.com/solapi/solapi-go/v2/messages"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	if apiKey == "" || apiSecret == "" {
		fmt.Println("환경변수(API_KEY, API_SECRET)가 필요합니다.")
		os.Exit(1)
	}

	c := client.NewClient(apiKey, apiSecret)

	// 항상 새 그룹 생성 (사전 조회 없음)
	created, err := c.Groups.Create(context.Background(), groups.CreateGroupOptions{AllowDuplicates: true})
	if err != nil {
		fmt.Println("create group error:", err)
		os.Exit(1)
	}
	groupId := created.GroupID
	fmt.Printf("created groupId=%s\n", groupId)

	// 전송 전에 그룹에 메시지 추가 (예제 통합)
	// 반드시 발신번호/수신번호는 01000000000 형식으로 입력해야 합니다.
	req := groups.AddGroupMessagesRequest{
		Messages: []messages.Message{
			{To: "수신번호", From: "계정에 등록한 발신번호", Text: "그룹에 추가되는 첫 번째 메시지"},
			{To: "수신번호", From: "계정에 등록한 발신번호", Text: "그룹에 추가되는 두 번째 메시지"},
		},
	}
	addRes, err := c.Groups.AddMessages(context.Background(), groupId, req)
	if err != nil {
		fmt.Println("add messages error:", err)
		os.Exit(1)
	}
	fmt.Printf("groupId=%s total=%d registeredFailed=%d\n", groupId, addRes.GroupInfo.Count.Total, addRes.GroupInfo.Count.RegisteredFailed)

	res, err := c.Groups.Send(context.Background(), groupId)
	if err != nil {
		fmt.Println("send group error:", err)
		os.Exit(1)
	}

	fmt.Printf("groupId=%s total=%d status=%s scheduled=%s\n", res.GroupID, res.GroupInfo.Count.Total, res.Status, res.ScheduledDate)
	// 실패 목록이 있다면 몇 건인지 출력
	if len(res.FailedMessageList) > 0 {
		fmt.Printf("failed=%d firstStatus=%s\n", len(res.FailedMessageList), res.FailedMessageList[0].StatusMessage)
	}

	// 발송 후 그룹 메시지 목록 조회 (list_group_messages 예제 통합)
	listRes, err := c.Groups.ListMessages(context.Background(), groupId, groups.ListMessagesQuery{Limit: 10})
	if err != nil {
		fmt.Println("list group messages error:", err)
		os.Exit(1)
	}
	fmt.Printf("limit=%d startKey=%s nextKey=%s\n", listRes.Limit, listRes.StartKey, listRes.NextKey)
	for id, m := range listRes.MessageList {
		b, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			fmt.Printf("messageId=%s marshal error: %v\n", id, err)
			continue
		}
		fmt.Printf("messageId=%s\n%s\n", id, string(b))
	}
}

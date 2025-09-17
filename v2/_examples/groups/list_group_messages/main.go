package main

import (
	"context"
	"fmt"
	"os"

	"github.com/solapi/solapi-go/v2/client"
	"github.com/solapi/solapi-go/v2/groups"
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

	resp, err := c.Groups.ListMessages(context.Background(), groupId, groups.ListMessagesQuery{Limit: 10})
	if err != nil {
		fmt.Println("list group messages error:", err)
		os.Exit(1)
	}

	fmt.Printf("limit=%d startKey=%s nextKey=%s\n", resp.Limit, resp.StartKey, resp.NextKey)
	for id, m := range resp.MessageList {
		fmt.Printf("%s to=%s from=%s type=%s status=%s text=%s\n", id, m.To, m.From, m.Type, m.Status, m.Text)
	}
}

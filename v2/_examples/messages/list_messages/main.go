package main

import (
	"context"
	"fmt"
	"os"

	"github.com/solapi/solapi-go/v2/client"
	"github.com/solapi/solapi-go/v2/messages"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	if apiKey == "" || apiSecret == "" {
		fmt.Println("API KEY 또는 API SECRET이 설정되지 않았습니다.")
		os.Exit(1)
	}

	c := client.NewClient(apiKey, apiSecret)

	// 간단 조회: 가장 최근 데이터부터 10건
	resp, err := c.Messages.List(
		context.Background(),
		messages.ListQuery{Limit: 10},
	)
	if err != nil {
		fmt.Println("list error:", err)
		os.Exit(1)
	}

	fmt.Printf("limit=%d startKey=%s nextKey=%s\n", resp.Limit, resp.StartKey, resp.NextKey)
	for id, m := range resp.MessageList {
		fmt.Printf("%s to=%s from=%s type=%s status=%s\n", id, m.To, m.From, m.Type, m.Status)
	}
}

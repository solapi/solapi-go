package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/solapi/solapi-go/v2/client"
	"github.com/solapi/solapi-go/v2/groups"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	if apiKey == "" || apiSecret == "" {
		fmt.Println("환경변수(API_KEY, API_SECRET)가 필요합니다.")
		os.Exit(1)
	}

	c := client.NewClient(apiKey, apiSecret)

	resp, err := c.Groups.ListGroups(context.Background(), groups.ListGroupsQuery{Limit: 10})
	if err != nil {
		fmt.Println("list groups error:", err)
		os.Exit(1)
	}

	startKey := "null"
	if resp.StartKey != nil {
		startKey = *resp.StartKey
	}
	fmt.Printf("limit=%d startKey=%s nextKey=%s\n", resp.Limit, startKey, resp.NextKey)

	for id, g := range resp.GroupList {
		b, err := json.MarshalIndent(g, "", "  ")
		if err != nil {
			fmt.Printf("groupId=%s marshal error: %v\n", id, err)
			continue
		}
		fmt.Printf("groupId=%s\n%s\n", id, string(b))
	}
}

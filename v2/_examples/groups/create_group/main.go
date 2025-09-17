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
	appId := os.Getenv("APP_ID")
	if apiKey == "" || apiSecret == "" {
		fmt.Println("API KEY 또는 API SECRET이 설정되지 않았습니다.")
		os.Exit(1)
	}

	c := client.NewClient(apiKey, apiSecret)

	opts := groups.CreateGroupOptions{
		AllowDuplicates: true,
		AppId:           appId,
	}

	res, err := c.Groups.Create(context.Background(), opts)
	if err != nil {
		fmt.Println("create group error:", err)
		os.Exit(1)
	}

	fmt.Printf("created groupId=%s\n", res.GroupID)
}



package main

import (
	"fmt"

	"github.com/solapi/solapi-go"
)

func main() {
	client := solapi.NewClient()

	//  파라미터들
	params := make(map[string]string)

	// SetCustomConfig
	/*
		client.Messages.Config = map[string]string{
			"APIKey": "Another API KEY",
		}
	*/

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Messages.CreateGroup(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}

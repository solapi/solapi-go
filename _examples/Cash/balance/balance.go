package main

import (
	"fmt"

	"github.com/solapi/solapi-go"
)

func main() {
	client := solapi.NewClient()

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Cash.Balance()
	if err != nil {
		fmt.Println(err)
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}

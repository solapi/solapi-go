package main

import (
	"fmt"

	"github.com/solapi/solapi-go"
)

func main() {
	client := solapi.NewClient()
	client.Messages.Config = map[string]string{
		"APIKey":    "", // solapi apikey
		"APISecret": "", // solapi secretkey
		"Protocol":  "https",
		"Domain":    "api.solapi.com",
		"Prefix":    "",
		"AppId":     "", // 이곳에 앱 아이디 입력 시 그룹 생성, 메시지 발송 시 추가로 입력할 필요 없습니다.
	}
	// 발송요청을 할 그룹 아이디
	groupId := "G4V20200824233846LAI57B7SKJACJ40"

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Messages.SendGroup(groupId)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}

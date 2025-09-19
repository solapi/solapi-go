package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/solapi/solapi-go/v2/client"
	"github.com/solapi/solapi-go/v2/messages"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	// 발신번호 및 수신번호 입력 형식은 01000000000 형식으로 입력하세요.
	to := "수신번호 입력"
	from := "계정에 등록한 발신번호 입력"
	if apiKey == "" || apiSecret == "" {
		fmt.Println("API KEY 또는 API SECRET이 설정되지 않았습니다.")
		os.Exit(1)
	}

	c := client.NewClient(apiKey, apiSecret)

	// 수신번호가 여러 개라면 To 항목을 삭제하고 ToList를 사용하세요.
	// 예시) ToList := []string{"수신번호1", "수신번호2", "수신번호3"}
	msg := messages.Message{
		To:   to,
		From: from,
		Text: "SOLAPI GO SDK V2 문자 발송 테스트",
	}

	// 여러 건의 메시지를 보내고 싶다면 아래 코드로 메시지 구조체를 만들어서 보내세요!
	// 수신번호가 중복된다면 allowDuplicates를 true로 설정하세요.
	// msgs := []messages.Message{
	// 	{
	// 		To:   to,
	// 		From: from,
	// 		Text: "SOLAPI GO SDK V2 문자 발송 테스트1",
	// 	},
	// 	{
	// 		To:   to,
	// 		From: from,
	// 		Text: "SOLAPI GO SDK V2 문자 발송 테스트2",
	// 	},
	// 	{
	// 		To:   to,
	// 		From: from,
	// 		Text: "SOLAPI GO SDK V2 문자 발송 테스트3",
	// 	},
	// }
	// allowDuplicates := true

	// res, err := c.Messages.Send(context.Background(), msgs, messages.SendOptions{
	// 	AllowDuplicates: &allowDuplicates,
	// })

	res, err := c.Messages.Send(context.Background(), msg)
	if err != nil {
		fmt.Println("send error:", err)
		os.Exit(1)
	}
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println("marshal error:", err)
		os.Exit(1)
	}
	fmt.Println(string(b))
}

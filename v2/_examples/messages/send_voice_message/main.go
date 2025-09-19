package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/solapi/solapi-go/v2/client"
	"github.com/solapi/solapi-go/v2/messages"
)

func buildVoiceMessage(to, from string) messages.Message {
	return messages.Message{
		To:   to,
		From: from,
		Text: "음성 메시지 테스트입니다, 실제 수신자에게 들리는 내용입니다.",
		Type: "VOICE",
		VoiceOptions: &messages.VoiceOptions{
			VoiceType: "FEMALE",
		},
	}
}

func main() {
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	to := "수신번호 입력"
	from := "계정에 등록한 발신번호 입력"
	if apiKey == "" || apiSecret == "" {
		fmt.Println("API KEY 또는 API SECRET이 설정되지 않았습니다.")
		os.Exit(1)
	}

	c := client.NewClient(apiKey, apiSecret)

	msg := buildVoiceMessage(to, from)

	// 필요하다면 아래 옵션을 설정하세요.
	// msg.VoiceOptions.HeaderMessage = "보이스 메시지 테스트" // 메시지 시작에 나오는 머릿말, 최대 135자
	// msg.VoiceOptions.TailMessage = "보이스 메시지 테스트"   // 통화 종료 시 나오는 꼬릿말, 최대 135자
	// msg.VoiceOptions.ReplyRange = 1                    // 수신자가 누를 수 있는 다이얼 범위(1~9), counselorNumber와 함께 사용 불가
	// msg.VoiceOptions.CounselorNumber = "상담번호"        // 수신자가 0번을 누르면 연결되는 번호, replyRange와 함께 사용 불가

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

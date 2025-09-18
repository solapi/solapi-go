package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/solapi/solapi-go/v2/client"
	"github.com/solapi/solapi-go/v2/messages"
)

// sendReservedSMS 는 delay 뒤에 예약 발송을 수행하고 응답 구조체를 리턴합니다.
func sendReservedSMS(ctx context.Context, apiKey, apiSecret, to, from string, delay time.Duration) (messages.DetailGroupMessageResponse, error) {
	c := client.NewClient(apiKey, apiSecret)

	// delay 뒤 예약발송 시간 (RFC3339, UTC)
	scheduledAt := time.Now().UTC().Add(delay).Format(time.RFC3339)

	msg := messages.Message{
		To:   to,
		From: from,
		Text: "SOLAPI GO SDK V2 예약 문자 발송 테스트",
	}

	showMessageList := true

	return c.Messages.Send(ctx, msg, messages.SendOptions{ScheduledDate: scheduledAt, ShowMessageList: &showMessageList})
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

	res, err := sendReservedSMS(context.Background(), apiKey, apiSecret, to, from, 10*time.Minute)
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

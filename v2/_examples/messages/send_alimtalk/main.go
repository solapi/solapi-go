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

	// 단일 알림톡 발송 예제
	fmt.Println("=== 단일 알림톡 발송 예제 ===")
	msg := messages.Message{
		To:   to,
		From: from,
		KakaoOptions: &messages.KakaoOptions{
			PfID:       "연동한 비즈니스 채널의 pfId",
			TemplateID: "등록한 알림톡 템플릿의 ID",
			Variables:  map[string]string{
				// 치환문구가 있는 경우 추가, 반드시 key, value 모두 string으로 기입해야 합니다.
				// "#{변수명}": "임의의 값",
			},
			// disableSms 값을 true로 줄 경우 문자로의 대체발송이 비활성화 됩니다.
			// DisableSms: &[]bool{true}[0],
		},
	}

	res, err := c.Messages.Send(context.Background(), msg)
	if err != nil {
		fmt.Println("알림톡 발송 실패:", err)
		os.Exit(1)
	}

	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println("JSON 변환 실패:", err)
		os.Exit(1)
	}
	fmt.Println("발송 결과:")
	fmt.Println(string(b))

	// 여러 메시지 알림톡 발송 예제
	fmt.Println("\n=== 여러 메시지 알림톡 발송 예제 ===")
	msgs := []messages.Message{
		{
			To:   to,
			From: from,
			KakaoOptions: &messages.KakaoOptions{
				PfID:       "연동한 비즈니스 채널의 pfId",
				TemplateID: "등록한 알림톡 템플릿의 ID",
				Variables:  map[string]string{
					// 치환문구가 있는 경우 추가, 반드시 key, value 모두 string으로 기입해야 합니다.
					// "#{변수명}": "임의의 값",
				},
				// disableSms 값을 true로 줄 경우 문자로의 대체발송이 비활성화 됩니다.
				// DisableSms: &[]bool{true}[0],
			},
		},
		{
			To:   "수신번호2 입력",
			From: from,
			KakaoOptions: &messages.KakaoOptions{
				PfID:       "연동한 비즈니스 채널의 pfId",
				TemplateID: "등록한 알림톡 템플릿의 ID",
				Variables:  map[string]string{
					// 치환문구가 있는 경우 추가, 반드시 key, value 모두 string으로 기입해야 합니다.
					// "#{변수명}": "임의의 값",
				},
				// disableSms 값을 true로 줄 경우 문자로의 대체발송이 비활성화 됩니다.
				// DisableSms: &[]bool{true}[0],
			},
		},
	}

	// 여러 메시지 발송 시 중복 수신번호를 허용하려면 AllowDuplicates를 true로 설정하세요.
	multiRes, err := c.Messages.Send(context.Background(), msgs, messages.SendOptions{
		// AllowDuplicates: &[]bool{true}[0],
	})
	if err != nil {
		fmt.Println("여러 메시지 알림톡 발송 실패:", err)
		os.Exit(1)
	}

	multiB, err := json.MarshalIndent(multiRes, "", "  ")
	if err != nil {
		fmt.Println("JSON 변환 실패:", err)
		os.Exit(1)
	}
	fmt.Println("여러 메시지 발송 결과:")
	fmt.Println(string(multiB))
}

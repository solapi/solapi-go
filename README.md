# SOLAPI SDK for Go

[Site](https://www.solapi.com/) |
[Docs](https://developers.solapi.com/) |
[Examples](https://github.com/solapi/solapi-go/tree/master/v2/_examples) |

SOLAPI 서비스를 이용하실 때, Golang에서 문자 메시지 발송 및 조회 관련 기능들을 쉽게 사용하실 수 있도록 만들어진 Golang 전용의 SOLAPI SDK 입니다.

## 발송 예제 코드

```go
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
	to := "수신번호 입력"
	from := "계정에 등록한 발신번호 입력"
	if apiKey == "" || apiSecret == "" {
		fmt.Println("API KEY 또는 API SECRET이 설정되지 않았습니다.")
		os.Exit(1)
	}

	c := client.NewClient(apiKey, apiSecret)

	msg := messages.Message{
		To:   to,
		From: from,
		Text: "SOLAPI GO SDK V2 문자 발송 테스트",
	}

	res, err := c.Messages.Send(context.Background(), msg)
	if err != nil {
		fmt.Println("send error:", err)
		os.Exit(1)
	}
	fmt.Printf("총 발송 접수 요청건수=%d 접수 실패 건 수=%d\n", res.GroupInfo.Count.Total, res.GroupInfo.Count.RegisteredFailed)
}

```

자세한 각각의 케이스 별 예제는 [예제 폴더](https://github.com/solapi/solapi-go/tree/master/v2/_examples)에서 확인 해 보실 수 있습니다!

## 설치방법

실제 SOLAPI Golang SDK를 적용 할 프로젝트 폴더에서 터미널을 실행 한 다음, 아래와 같은 명령어를 입력 해 주세요!

```
go get github.com/solapi/solapi-go/v2
```

## 주의사항

해당 레포지터리 내 루트 디렉터리 상 바로 확인할 수 있는 [_examples](https://github.com/solapi/solapi-go/tree/master/_examples) 폴더는 구 버전의 예제 코드를 안내하고 있습니다.  
실제 최신 예제는 [v2/_examples](https://github.com/solapi/solapi-go/tree/master/v2/_examples) 폴더를 확인 해 주세요!
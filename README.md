# solapi-go

[Site](https://www.solapi.com/) |
[Docs](https://docs.solapi.com/) |
[Examples](https://github.com/solapi/solapi-go/tree/master/_examples) |

문자 메시지 발송 및 조회 관련 기능들을 쉽게 사용하실 수 있도록 만들어진 SDK 입니다.

## Example

```go
require (
        github.com/solapi/solapi-go
)

func main() {
	client := solapi.NewClient()
	client.Messages.Config = map[string]string{
	"APIKey":    "", // solapi apikey
	"APISecret": "",  // solapi secretkey
	"Protocol":  "https",
	"Domain":    "api.solapi.com",
	"Prefix":    "",
	"AppId":     "", // 이곳에 앱 아이디 입력 시 그룹 생성, 메시지 발송 시 추가로 입력할 필요 없습니다.
}

	// Message Data
	// 관련 파라미터들은 https://docs.solapi.com에서 확인 가능합니다.
	message := make(map[string]interface{})
	message["to"] = "01000000000"
	message["from"] = "029302266"
	message["text"] = "Test Message"
	message["type"] = "SMS"

	params := make(map[string]interface{})
	params["message"] = message

	// Call API Resource
	result, err := client.Message.SendSimpleMessage(params)
	if err != nil {
		fmt.Println(err)
	}
}
```

[examples folder](https://github.com/solapi/solapi-go/tree/master/_examples)에서 자세한 예제파일들을 확인하세요.

## Installation

```
go get github.com/solapi/solapi-go
```

## Configs

```
{
  "APIKey": "NCSVYGF1IK5PUKDA",
  "APISecret": "FSD4ER2WYPZQVDBPKMLOZVAWTGYBDTRW",
  "Protocol": "https",
  "Domain": "api.solapi.com",
  "Prefix": "",
  "AppId": "" // 이곳에 앱 아이디 입력 시 그룹 생성, 메시지 발송 시 추가로 입력할 필요 없습니다.
}
```

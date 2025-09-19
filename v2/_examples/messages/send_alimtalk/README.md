# 알림톡 발송 예제 (Kakao Alimtalk)

카카오 알림톡(이미지 알림톡 포함) 발송 예제입니다.

## 사전 준비

1. **API 키 설정**: 환경변수에 SOLAPI API 키를 설정하세요.
   ```bash
   export API_KEY="your_api_key"
   export API_SECRET="your_api_secret"
   ```

2. **비즈니스 채널 연동**: 카카오 비즈니스 채널을 연동하고 PfID를 확인하세요.
3. **알림톡 템플릿 등록**: 사용할 알림톡 템플릿을 등록하고 템플릿 ID를 확인하세요.

## 실행 방법

```bash
# 예제 실행
go run main.go
```

## 코드 설명

### 단일 알림톡 발송

```go
msg := messages.Message{
    To:   "수신번호",
    From: "발신번호",
    KakaoOptions: &messages.KakaoOptions{
        PfID:       "비즈니스 채널 PfID",
        TemplateID: "알림톡 템플릿 ID",
        Variables: map[string]string{
            "#{변수명}": "치환할 값",
        },
    },
}

res, err := c.Messages.Send(context.Background(), msg)
```

### 여러 메시지 발송

```go
msgs := []messages.Message{...}
res, err := c.Messages.Send(context.Background(), msgs)
```

### 옵션 설명

- **PfID**: 연동한 비즈니스 채널의 고유 ID (필수)
- **TemplateID**: 등록한 알림톡 템플릿의 ID (필수)
- **Variables**: 템플릿의 치환문구를 대체할 값들 (선택)
- **DisableSms**: true로 설정하면 문자 대체발송이 비활성화됩니다 (선택)
- **AllowDuplicates**: 중복 수신번호 허용 여부 (선택)

## 주의사항

1. 발신번호와 수신번호는 반드시 `-`, `*` 등 특수문자를 제거한 형식으로 입력하세요.
   - 예: `01012345678` (올바름)
   - 예: `010-1234-5678` (잘못됨)

2. 한 번 호출 당 최대 10,000건까지 발송 가능합니다.

3. 치환문구(Variables)를 사용할 경우 반드시 key와 value 모두 string 타입이어야 합니다.

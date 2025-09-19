# SMS 예약 발송 예제

SMS 문자 메시지를 예약 발송하는 예제입니다.

## 사전 준비

1. **API 키 설정**: 환경변수에 SOLAPI API 키를 설정하세요.
   ```bash
   export API_KEY="your_api_key"
   export API_SECRET="your_api_secret"
   ```

2. **발신번호 등록**: SOLAPI 콘솔에서 발신번호를 등록하세요.

## 실행 방법

```bash
# 예제 실행 (10분 후 예약 발송)
go run main.go
```

## 코드 설명

### 예약 발송 함수

```go
func sendReservedSMS(ctx context.Context, apiKey, apiSecret, to, from string, delay time.Duration) (messages.DetailGroupMessageResponse, error) {
    c := client.NewClient(apiKey, apiSecret)

    // delay 뒤 예약발송 시간 (RFC3339, UTC)
    scheduledAt := time.Now().UTC().Add(delay).Format(time.RFC3339)

    msg := messages.Message{
        To:   to,
        From: from,
        Text: "예약 문자 메시지 내용",
    }

    showMessageList := true

    return c.Messages.Send(ctx, msg, messages.SendOptions{
        ScheduledDate:   scheduledAt,
        ShowMessageList: &showMessageList,
    })
}
```

### 예약 발송 호출

```go
// 10분 후 예약 발송
res, err := sendReservedSMS(context.Background(), apiKey, apiSecret, to, from, 10*time.Minute)
```

### 옵션 설명

- **To**: 수신번호 (필수, 숫자만 입력)
- **From**: 발신번호 (필수, 등록된 발신번호만 사용 가능)
- **Text**: 메시지 내용 (필수)
- **ScheduledDate**: 예약 발송 시간 (RFC3339 형식, UTC, 필수)
- **ShowMessageList**: 메시지 목록 표시 여부 (선택)

## 주의사항

1. 발신번호와 수신번호는 반드시 `-`, `*` 등 특수문자를 제거한 형식으로 입력하세요.
   - 예: `01012345678` (올바름)
   - 예: `010-1234-5678` (잘못됨)

2. 예약 발송 시간은 현재 시간보다 미래여야 합니다.

3. 예약 발송 시 현재 시간보다 과거의 시간을 입력할 경우 즉시 발송됩니다.

4. 메시지 내용은 최대 90바이트까지 지원됩니다. (한글 기준 약 45자)

5. 예약 발송은 최대 6개월 까지 설정 가능합니다.

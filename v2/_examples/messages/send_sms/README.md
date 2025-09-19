# SMS 문자 발송 예제

SMS 문자 메시지 발송 예제입니다.

## 사전 준비

1. **API 키 설정**: 환경변수에 SOLAPI API 키를 설정하세요.
   ```bash
   export API_KEY="your_api_key"
   export API_SECRET="your_api_secret"
   ```

2. **발신번호 등록**: SOLAPI 콘솔에서 발신번호를 등록하세요.

## 실행 방법

```bash
# 예제 실행
go run main.go
```

## 코드 설명

### 단일 SMS 발송

```go
msg := messages.Message{
    To:   "수신번호",
    From: "발신번호",
    Text: "문자 메시지 내용",
}

res, err := c.Messages.Send(context.Background(), msg)
```

### 여러 메시지 발송

```go
msgs := []messages.Message{
    {
        To:   "수신번호1",
        From: "발신번호",
        Text: "첫 번째 메시지",
    },
    {
        To:   "수신번호2",
        From: "발신번호",
        Text: "두 번째 메시지",
    },
}

// 중복 수신번호 허용 시
allowDuplicates := true
res, err := c.Messages.Send(context.Background(), msgs, messages.SendOptions{
    AllowDuplicates: &allowDuplicates,
})
```

### 옵션 설명

- **To**: 수신번호 (필수, 숫자만 입력)
- **From**: 발신번호 (필수, 등록된 발신번호만 사용 가능)
- **Text**: 메시지 내용 (필수)
- **AllowDuplicates**: 중복 수신번호 허용 여부 (선택)

## 주의사항

1. 발신번호와 수신번호는 반드시 `-`, `*` 등 특수문자를 제거한 형식으로 입력하세요.
   - 예: `01012345678` (올바름)
   - 예: `010-1234-5678` (잘못됨)

2. 한 번 호출 당 최대 10,000건까지 발송 가능합니다.

3. 메시지 내용은 한글 기준 최대 1,000자 까지 입력할 수 있으며, 한글 기준 메시지 내용이 45자가 넘어가면 자동으로 LMS으로 전환됩니다.

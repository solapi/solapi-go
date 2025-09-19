# 음성 메시지 발송 예제

음성 메시지(TTS)를 발송하는 예제입니다.

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

### 음성 메시지 생성 함수

```go
func buildVoiceMessage(to, from string) messages.Message {
    return messages.Message{
        To:   to,
        From: from,
        Text: "음성 메시지 내용입니다.",
        Type: "VOICE",
        VoiceOptions: &messages.VoiceOptions{
            VoiceType: "FEMALE", // 또는 "MALE"
        },
    }
}
```

### 음성 옵션 설정

```go
msg.VoiceOptions = &messages.VoiceOptions{
    VoiceType:       "FEMALE",        // 음성 타입 (MALE/FEMALE)
    HeaderMessage:   "머릿말 메시지",  // 통화 시작 시 나오는 메시지 (최대 135자)
    TailMessage:     "꼬릿말 메시지",  // 통화 종료 시 나오는 메시지 (최대 135자)
    ReplyRange:      1,              // 수신자가 누를 수 있는 다이얼 범위 (1-9)
    CounselorNumber: "상담번호",      // 0번을 누르면 연결되는 번호, 01000000000 형식 입력
}
```

### 옵션 설명

- **To**: 수신번호 (필수, 숫자만 입력)
- **From**: 발신번호 (필수, 등록된 발신번호만 사용 가능)
- **Text**: 음성으로 변환될 텍스트 내용 (필수)
- **Type**: 메시지 타입, "VOICE"로 설정 (필수)
- **VoiceOptions.VoiceType**: 음성 타입 ("MALE" 또는 "FEMALE")
- **VoiceOptions.HeaderMessage**: 통화 시작 시 나오는 메시지 (최대 135자)
- **VoiceOptions.TailMessage**: 통화 종료 시 나오는 메시지 (최대 135자)
- **VoiceOptions.ReplyRange**: 수신자가 누를 수 있는 다이얼 범위 (1-9, CounselorNumber와 함께 사용 불가)
- **VoiceOptions.CounselorNumber**: 0번을 누르면 연결되는 번호 (ReplyRange와 함께 사용 불가)

## 주의사항

1. 발신번호와 수신번호는 반드시 `-`, `*` 등 특수문자를 제거한 형식으로 입력하세요.
   - 예: `01012345678` (올바름)
   - 예: `010-1234-5678` (잘못됨)

2. ReplyRange와 CounselorNumber는 함께 사용할 수 없습니다.

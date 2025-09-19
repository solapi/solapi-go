# 그룹 메시지 추가 예제

기존 그룹에 메시지를 추가하는 예제입니다.

## 사전 준비

1. **API 키 설정**: 환경변수에 SOLAPI API 키를 설정하세요.
   ```bash
   export API_KEY="your_api_key"
   export API_SECRET="your_api_secret"
   ```

2. **그룹 생성**: 메시지를 추가할 그룹이 있어야 합니다. (create_group 예제 참조)

3. **그룹 ID 설정**: 생성된 그룹의 ID를 환경변수로 설정하세요.
   ```bash
   export GROUP_ID="your_group_id"
   ```

## 실행 방법

```bash
# 예제 실행
go run main.go
```

## 코드 설명

### 그룹 메시지 추가 요청

```go
req := groups.AddGroupMessagesRequest{
    Messages: []messages.Message{
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
    },
}
```

### 메시지 추가 실행

```go
res, err := c.Groups.AddMessages(context.Background(), groupId, req)
if err != nil {
    fmt.Println("add messages error:", err)
    return
}
```

### 결과 확인

```go
fmt.Printf("그룹 ID: %s\n", groupId)
fmt.Printf("총 메시지 수: %d\n", res.GroupInfo.Count.Total)
fmt.Printf("등록 실패 수: %d\n", res.GroupInfo.Count.RegisteredFailed)
fmt.Printf("등록 성공 수: %d\n", res.GroupInfo.Count.RegisteredSuccess)
```

### 메시지 옵션들

```go
message := messages.Message{
    To:   "수신번호",
    From: "발신번호",
    Text: "메시지 내용",

    // 선택적 옵션들
    Type:    "SMS",           // 메시지 타입 (SMS, LMS, MMS 등)
    Subject: "제목",          // LMS/MMS 제목
    ImageID: "이미지ID",      // MMS용 이미지 ID

    // 카카오톡 옵션
    KakaoOptions: &messages.KakaoOptions{
        PfID:       "비즈니스 채널 ID",
        TemplateID: "템플릿 ID",
        Variables: map[string]string{
            "#{이름}": "홍길동",
        },
    },
}
```

## 주의사항

1. 그룹이 존재하지 않으면 에러가 발생합니다.

2. 그룹에 추가된 메시지는 즉시 발송되지 않습니다. 그룹 발송을 별도로 실행해야 합니다.

3. 한 번에 최대 10,000건까지 메시지를 추가할 수 있습니다.

4. 발신번호와 수신번호는 반드시 `-`, `*` 등 특수문자를 제거한 형식으로 입력하세요.
   - 예: `01012345678` (올바름)
   - 예: `010-1234-5678` (잘못됨)

5. 메시지 타입에 따라 필수 파라미터가 다를 수 있습니다.

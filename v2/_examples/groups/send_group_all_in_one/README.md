# 그룹 전체 워크플로우 예제

그룹 생성부터 메시지 추가, 발송, 결과 조회까지의 전체 과정을 하나의 예제로 통합한 예제입니다.

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

### 1. 그룹 생성

```go
created, err := c.Groups.Create(context.Background(), groups.CreateGroupOptions{
    AllowDuplicates: true, // 중복 수신번호 허용
})
groupId := created.GroupID
```

### 2. 메시지 추가

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

addRes, err := c.Groups.AddMessages(context.Background(), groupId, req)
```

### 3. 그룹 발송

```go
res, err := c.Groups.Send(context.Background(), groupId)
```

### 4. 결과 확인

```go
fmt.Printf("그룹 ID: %s\n", res.GroupID)
fmt.Printf("총 메시지 수: %d\n", res.GroupInfo.Count.Total)
fmt.Printf("그룹 상태: %s\n", res.Status)
fmt.Printf("예약 시간: %s\n", res.ScheduledDate)

// 실패한 메시지 확인
if len(res.FailedMessageList) > 0 {
    fmt.Printf("실패 건수: %d\n", len(res.FailedMessageList))
    fmt.Printf("첫 번째 실패 사유: %s\n", res.FailedMessageList[0].StatusMessage)
}
```

### 5. 메시지 목록 조회

```go
listRes, err := c.Groups.ListMessages(context.Background(), groupId, groups.ListMessagesQuery{
    Limit: 10,
})

for messageId, message := range listRes.MessageList {
    fmt.Printf("메시지 ID: %s\n", messageId)
    fmt.Printf("수신번호: %s\n", message.To)
    fmt.Printf("상태: %s\n", message.Status)
}
```

## 워크플로우 단계

### 단계 1: 그룹 생성
- 새로운 메시지 그룹을 생성합니다.
- 중복 수신번호 허용 옵션을 설정할 수 있습니다.

### 단계 2: 메시지 추가
- 생성된 그룹에 메시지를 추가합니다.
- 여러 개의 메시지를 한 번에 추가할 수 있습니다.
- 각 메시지는 개별적으로 설정할 수 있습니다.

### 단계 3: 그룹 발송
- 그룹의 모든 메시지를 발송합니다.
- 즉시 발송되거나 예약 발송될 수 있습니다.

### 단계 4: 결과 조회
- 발송 결과를 확인합니다.
- 성공/실패 건수를 확인할 수 있습니다.
- 실패한 메시지의 상세 정보를 확인할 수 있습니다.

### 단계 5: 메시지 목록 조회
- 그룹에 포함된 모든 메시지의 상태를 확인합니다.
- 각 메시지의 발송 결과를 개별적으로 확인할 수 있습니다.

## 그룹 발송 옵션

```go
// 예약 발송
res, err := c.Groups.Send(context.Background(), groupId, groups.SendOptions{
    ScheduledDate: "2023-12-31T23:59:59Z",
})
```

## 주의사항

1. 발신번호와 수신번호는 반드시 `-`, `*` 등 특수문자를 제거한 형식으로 입력하세요.
   - 예: `01012345678` (올바름)
   - 예: `010-1234-5678` (잘못됨)

2. 그룹 생성 후 24시간 이내에 발송하지 않으면 그룹이 만료될 수 있습니다.

3. 한 그룹에 최대 10,000건까지 메시지를 추가할 수 있습니다.

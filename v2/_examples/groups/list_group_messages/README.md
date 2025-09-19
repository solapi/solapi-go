# 그룹 메시지 목록 조회 예제

특정 그룹에 포함된 메시지 목록을 조회하는 예제입니다.

## 사전 준비

1. **API 키 설정**: 환경변수에 SOLAPI API 키를 설정하세요.
   ```bash
   export API_KEY="your_api_key"
   export API_SECRET="your_api_secret"
   ```

2. **그룹 ID 설정**: 조회할 그룹의 ID를 환경변수로 설정하세요.
   ```bash
   export GROUP_ID="your_group_id"
   ```

## 실행 방법

```bash
# 예제 실행
go run main.go
```

## 코드 설명

### 그룹 메시지 목록 조회

```go
resp, err := c.Groups.ListMessages(
    context.Background(),
    groupId,
    groups.ListMessagesQuery{Limit: 10},
)
```

### 고급 조회 옵션들

```go
// 페이지네이션
query := groups.ListMessagesQuery{
    StartKey: "이전 조회의 NextKey 값",
    Limit:    10,
}

// 날짜 범위로 필터링
query := groups.ListMessagesQuery{
    StartDate: "2023-01-01 00:00:00",
    EndDate:   "2023-01-31 23:59:59",
    Limit:     10,
}
```

### 조회 결과 처리

```go
fmt.Printf("조회 건수: %d\n", resp.Limit)
fmt.Printf("시작 키: %s\n", resp.StartKey)
fmt.Printf("다음 키: %s\n", resp.NextKey)
fmt.Printf("총 메시지 수: %d\n", len(resp.MessageList))

for messageId, message := range resp.MessageList {
    fmt.Printf("메시지 ID: %s\n", messageId)
    fmt.Printf("수신번호: %s\n", message.To)
    fmt.Printf("발신번호: %s\n", message.From)
    fmt.Printf("메시지 타입: %s\n", message.Type)
    fmt.Printf("상태: %s\n", message.Status)
    fmt.Printf("메시지 내용: %s\n", message.Text)
    fmt.Printf("생성 시간: %s\n", message.DateCreated)
    fmt.Printf("발송 시간: %s\n", message.DateSent)
}
```

### 옵션 설명

- **StartKey**: 페이지네이션 시작 키
- **Limit**: 조회 건수 (기본값: 20, 최대값: 500)
- **StartDate**: 시작 날짜 (YYYY-MM-DD HH:mm:ss 형식)
- **EndDate**: 종료 날짜 (YYYY-MM-DD HH:mm:ss 형식)

## 응답 필드

- **MessageList**: 메시지 맵 (키: 메시지ID, 값: 메시지 객체)
- **Limit**: 요청한 건수
- **StartKey**: 현재 페이지의 시작 키
- **NextKey**: 다음 페이지의 시작 키 (없으면 빈 문자열)

## 주의사항

1. 한 번에 최대 500건까지 조회할 수 있습니다.

2. 페이지네이션이 필요한 경우 NextKey를 사용하여 다음 페이지를 조회하세요.

3. 날짜 필터를 사용할 때는 StartDate와 EndDate를 함께 지정해야 합니다.

4. 메시지 상태에 따라 조회 결과가 다를 수 있습니다.

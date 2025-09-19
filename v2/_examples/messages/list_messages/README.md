# 메시지 목록 조회 예제

발송된 메시지 목록을 조회하는 예제입니다.

## 사전 준비

1. **API 키 설정**: 환경변수에 SOLAPI API 키를 설정하세요.
   ```bash
   export API_KEY="your_api_key"
   export API_SECRET="your_api_secret"
   ```

2. **발송 이력**: 조회할 메시지가 있어야 합니다.

## 실행 방법

```bash
# 예제 실행
go run main.go
```

## 코드 설명

### 기본 메시지 목록 조회

```go
// 가장 최근 데이터부터 10건 조회
resp, err := c.Messages.List(
    context.Background(),
    messages.ListQuery{Limit: 10},
)
```

### 고급 조회 옵션들

```go
// 특정 메시지 ID로 조회
query := messages.ListQuery{
    MessageID: "메시지ID",
    Limit:     10,
}

// 그룹 ID로 조회
query := messages.ListQuery{
    GroupID: "그룹ID",
    Limit:   10,
}

// 수신번호로 조회
query := messages.ListQuery{
    To:    "01012345678",
    Limit: 10,
}

// 발신번호로 조회
query := messages.ListQuery{
    From:  "01012345678",
    Limit: 10,
}

// 메시지 타입으로 조회
query := messages.ListQuery{
    TypeIn: []string{"SMS", "LMS", "MMS"},
    Limit:  10,
}

// 날짜 범위로 조회
query := messages.ListQuery{
    DateType:  "CREATED", // CREATED, UPDATED
    StartDate: "2023-01-01 00:00:00",
    EndDate:   "2023-01-31 23:59:59",
    Limit:     10,
}

// 페이지네이션
query := messages.ListQuery{
    StartKey: "이전 조회의 NextKey 값",
    Limit:    10,
}

resp, err := c.Messages.List(context.Background(), query)
```

### 조회 결과 처리

```go
fmt.Printf("총 건수: %d\n", len(resp.MessageList))
fmt.Printf("다음 페이지 키: %s\n", resp.NextKey)

for id, message := range resp.MessageList {
    fmt.Printf("ID: %s\n", id)
    fmt.Printf("수신번호: %s\n", message.To)
    fmt.Printf("발신번호: %s\n", message.From)
    fmt.Printf("메시지 타입: %s\n", message.Type)
    fmt.Printf("상태: %s\n", message.Status)
    fmt.Printf("발송 시간: %s\n", message.DateSent)
    fmt.Printf("생성 시간: %s\n", message.DateCreated)
}
```

### 옵션 설명

- **MessageID**: 특정 메시지 ID로 조회
- **GroupID**: 그룹 ID로 조회
- **To**: 수신번호로 조회
- **From**: 발신번호로 조회
- **TypeIn**: 메시지 타입 배열 (SMS, LMS, MMS, ATA, CTA 등)
- **DateType**: 날짜 필터 타입 (CREATED, UPDATED, SENT)
- **StartDate/EndDate**: 날짜 범위 (YYYY-MM-DD HH:mm:ss 형식)
- **StartKey**: 페이지네이션 시작 키
- **Limit**: 조회 건수 (기본값: 20, 최대값: 500)

## 응답 필드

- **MessageList**: 메시지 맵 (키: 메시지ID, 값: 메시지 객체)
- **Limit**: 요청한 건수
- **StartKey**: 현재 페이지의 시작 키
- **NextKey**: 다음 페이지의 시작 키

## 주의사항

1. 한 번에 최대 500건까지 조회할 수 있습니다.(기본 값: 20건)

2. 페이지네이션이 필요한 경우 NextKey를 사용하여 다음 페이지를 조회하세요.

3. 날짜 범위 조회 시 DateType을 반드시 지정해야 합니다.

4. 메시지 타입은 대문자로 입력해야 합니다.

# 그룹 목록 조회 예제

생성된 메시지 그룹 목록을 조회하는 예제입니다.

## 사전 준비

1. **API 키 설정**: 환경변수에 SOLAPI API 키를 설정하세요.
   ```bash
   export API_KEY="your_api_key"
   export API_SECRET="your_api_secret"
   ```

2. **그룹 데이터**: 조회할 그룹이 있어야 합니다.

## 실행 방법

```bash
# 예제 실행
go run main.go
```

## 코드 설명

### 기본 그룹 목록 조회

```go
// 가장 최근 그룹부터 10건 조회
resp, err := c.Groups.ListGroups(
    context.Background(),
    groups.ListGroupsQuery{Limit: 10},
)
```

### 고급 조회 옵션들

```go
// 페이지네이션
query := groups.ListGroupsQuery{
    StartKey: "이전 조회의 NextKey 값",
    Limit:    10,
}

// 날짜 범위로 조회
query := groups.ListGroupsQuery{
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
fmt.Printf("총 그룹 수: %d\n", len(resp.GroupList))

for groupId, group := range resp.GroupList {
    fmt.Printf("그룹 ID: %s\n", groupId)
    fmt.Printf("그룹 상태: %s\n", group.Status)
    fmt.Printf("총 메시지 수: %d\n", group.Count.Total)
    fmt.Printf("발송 성공 수: %d\n", group.Count.SentSuccess)
    fmt.Printf("발송 실패 수: %d\n", group.Count.SentFailed)
    fmt.Printf("생성 시간: %s\n", group.DateCreated)
    fmt.Printf("발송 시간: %s\n", group.DateSent)
}
```

### 옵션 설명

- **StartKey**: 페이지네이션 시작 키
- **Limit**: 조회 건수 (기본값: 20, 최대값: 100)
- **StartDate**: 시작 날짜 (UTC 형식)
- **EndDate**: 종료 날짜 (UTC 형식)
- **Status**: 그룹 상태 필터

## 그룹 상태 값

- **PENDING**: 대기 중 (메시지 등록됨, 발송 전)
- **SENDING**: 발송 중
- **COMPLETE**: 발송 완료
- **FAILED**: 발송 실패

## 그룹 정보 필드

- **Count.Total**: 총 메시지 수
- **Count.SentTotal**: 발송된 메시지 수
- **Count.SentSuccess**: 발송 성공 수
- **Count.SentFailed**: 발송 실패 수
- **Count.SentPending**: 발송 대기 수
- **DateCreated**: 그룹 생성 시간
- **DateSent**: 그룹 발송 시간
- **DateCompleted**: 그룹 완료 시간

## 응답 필드

- **GroupList**: 그룹 맵 (키: 그룹ID, 값: 그룹 정보 객체)
- **Limit**: 요청한 건수
- **StartKey**: 현재 페이지의 시작 키
- **NextKey**: 다음 페이지의 시작 키

## 주의사항

1. 한 번에 최대 500건까지 조회할 수 있습니다.

2. 페이지네이션이 필요한 경우 NextKey를 사용하여 다음 페이지를 조회하세요.

3. 날짜 필터를 사용할 때는 StartDate와 EndDate를 함께 지정해야 합니다.

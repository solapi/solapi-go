# 그룹 생성 예제

새로운 메시지 그룹을 생성하는 예제입니다.

## 사전 준비

1. **API 키 설정**: 환경변수에 SOLAPI API 키를 설정하세요.
   ```bash
   export API_KEY="your_api_key"
   export API_SECRET="your_api_secret"
   ```

2. **앱 ID 설정 (선택)**: 애플리케이션 ID가 있는 경우 설정하세요.
   ```bash
   export APP_ID="your_app_id"
   ```

## 실행 방법

```bash
# 기본 옵션으로 그룹 생성
go run main.go

# 앱 ID 지정하여 그룹 생성
export APP_ID="your_app_id"
go run main.go
```

## 코드 설명

### 그룹 생성 옵션 설정

```go
opts := groups.CreateGroupOptions{
    AllowDuplicates: true,     // 중복 수신번호 허용 여부
    AppId:           appId,    // 애플리케이션 ID (선택)
}
```

### 그룹 생성 실행

```go
res, err := c.Groups.Create(context.Background(), opts)
if err != nil {
    fmt.Println("create group error:", err)
    return
}
```

### 생성된 그룹 ID 확인

```go
fmt.Printf("생성된 그룹 ID: %s\n", res.GroupID)
```

### 그룹 생성 옵션들

```go
opts := groups.CreateGroupOptions{
    AllowDuplicates: true,        // 중복 수신번호 허용 (기본값: false)
    AppId:           "앱ID",      // 애플리케이션 식별자 (선택)
    Strict:          false,       // 엄격 모드 (기본값: false)
}
```

## 그룹 생성 옵션 설명

- **AllowDuplicates**: 동일한 수신번호로 여러 메시지를 보낼 수 있는지 여부
  - `true`: 중복 허용
  - `false`: 중복 불허 (기본값)

- **Strict**: 엄격 모드
  - `true`: 유효하지 않은 메시지는 등록되지 않음
  - `false`: 유효하지 않은 메시지도 등록 (기본값)

## 그룹 사용 방법

생성된 그룹에 메시지를 추가하려면:

```bash
export GROUP_ID="생성된_그룹_ID"
# add_group_messages 예제 실행
```

그룹의 모든 메시지를 발송하려면:

```bash
export GROUP_ID="생성된_그룹_ID"
# send_group_all_in_one 예제 실행
```

## 주의사항

1. 생성된 그룹은 메시지가 발송되기 전까지 24시간 동안 유지됩니다.

2. 그룹에 메시지를 추가하지 않고 발송을 시도하면 에러가 발생합니다.

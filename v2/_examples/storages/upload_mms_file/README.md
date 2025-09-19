# MMS 파일 업로드 예제

MMS 메시지에 사용할 이미지 파일을 스토리지에 업로드하는 예제입니다.

## 사전 준비

1. **API 키 설정**: 환경변수에 SOLAPI API 키를 설정하세요.
   ```bash
   export API_KEY="your_api_key"
   export API_SECRET="your_api_secret"
   ```

2. **이미지 파일 준비**: `test.jpg` 파일을 예제 디렉터리에 준비하세요.

## 실행 방법

```bash
# 기본 파일로 실행
go run main.go

# 다른 파일로 실행
# 파일을 test.jpg로 복사하거나 코드에서 경로를 수정하세요
```

## 코드 설명

### 파일 읽기 및 인코딩 함수

```go
func mustReadAndEncode(path string) string {
    b, err := os.ReadFile(path)
    if err != nil {
        fmt.Println("failed to read file:", err)
        os.Exit(1)
    }
    return base64.StdEncoding.EncodeToString(b)
}
```

### 파일 업로드 요청

```go
req := storages.UploadFileRequest{
    File: encoded,     // Base64로 인코딩된 파일 데이터
    Name: "test.jpg",  // 파일명
    Type: "MMS",       // 업로드 타입
}

resp, err := c.Storages.Upload(context.Background(), req)
```

### 업로드 결과 확인

```go
fmt.Printf("파일 ID: %s\n", resp.FileID)
fmt.Printf("파일명: %s\n", resp.Name)
fmt.Printf("파일 타입: %s\n", resp.Type)
fmt.Printf("파일 URL: %s\n", resp.URL)
```

### MMS 메시지에서 사용

```go
msg := messages.Message{
    To:      "수신번호",
    From:    "발신번호",
    Type:    "MMS",
    Subject: "MMS 제목",
    Text:    "MMS 메시지 내용",
    ImageID: resp.FileID, // 업로드된 파일 ID
}

res, err := c.Messages.Send(context.Background(), msg)
```

## 지원 파일 형식

- **이미지**: JPG, JPEG, PNG, GIF
- **문서**: PDF (특정 경우에 한함)
- **최대 크기**: 300KB
- **권장 해상도**: 640x480 이하

## 업로드 타입

- **MMS**: MMS 메시지용 파일
- **RCS**: RCS 메시지용 파일
- **FAX**: 팩스용 파일

## 파일 업로드 옵션

```go
req := storages.UploadFileRequest{
    File: encoded,
    Name: "custom_name.jpg",  // 사용자 정의 파일명
    Type: "MMS",
    // Link: "https://..."    // 외부 URL에서 파일 다운로드 (선택)
}
```

## 주의사항

1. 파일은 반드시 Base64로 인코딩하여 전송해야 합니다.

2. 업로드된 파일은 MMS 메시지에서만 사용할 수 있습니다.

3. 파일 업로드 후 반환된 FileID를 MMS 메시지의 ImageID 필드에 설정해야 합니다.

4. 업로드된 파일은 SOLAPI 스토리지에 저장되며, 별도의 비용이 발생할 수 있습니다.

5. 파일 크기가 300KB를 초과하면 업로드가 실패할 수 있습니다.

6. 동일한 파일을 여러 번 업로드하면 각각 다른 FileID가 생성됩니다.

7. 업로드된 파일은 영구적으로 저장되지 않을 수 있으니, 사용 후 FileID를 기록해두세요.

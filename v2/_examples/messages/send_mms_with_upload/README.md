# MMS 발송 예제 (파일 업로드 포함)

이미지 파일을 업로드하고 MMS 메시지를 발송하는 예제입니다.

## 사전 준비

1. **API 키 설정**: 환경변수에 SOLAPI API 키를 설정하세요.
   ```bash
   export API_KEY="your_api_key"
   export API_SECRET="your_api_secret"
   ```

2. **발신번호 등록**: SOLAPI 콘솔에서 발신번호를 등록하세요.

3. **이미지 파일 준비**: `test.jpg` 파일을 예제 디렉터리에 준비하거나 `FILE_PATH` 환경변수로 파일 경로를 지정하세요.

## 실행 방법

```bash
# 기본 이미지 파일(test.jpg)로 실행
go run main.go

# 다른 이미지 파일로 실행
export FILE_PATH="/path/to/your/image.jpg"
go run main.go
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

### 1. 파일 업로드

```go
upReq := storages.UploadFileRequest{
    File: encoded,           // Base64로 인코딩된 파일 데이터
    Name: filepath.Base(filePath), // 파일명
    Type: "MMS",             // 업로드 타입
}

upRes, err := c.Storages.Upload(context.Background(), upReq)
```

### 2. MMS 메시지 발송

```go
msg := messages.Message{
    To:      to,
    From:    from,
    Type:    "MMS",
    Subject: "MMS 제목",
    Text:    "MMS 메시지 내용",
    ImageID: upRes.FileID, // 업로드된 파일의 ID
}

res, err := c.Messages.Send(context.Background(), msg)
```

### 옵션 설명

- **To**: 수신번호 (필수, 숫자만 입력)
- **From**: 발신번호 (필수, 등록된 발신번호만 사용 가능)
- **Type**: 메시지 타입, "MMS"로 설정 (필수)
- **Subject**: MMS 제목 (선택)
- **Text**: 메시지 내용 (필수)
- **ImageID**: 업로드된 이미지 파일의 ID (필수)

## 지원 파일 형식

- **이미지**: JPG, JPEG, PNG
- **파일 크기**: 최대 200KB

## 주의사항

1. 발신번호와 수신번호는 반드시 `-`, `*` 등 특수문자를 제거한 형식으로 입력하세요.
   - 예: `01012345678` (올바름)
   - 예: `010-1234-5678` (잘못됨)

2. 파일 업로드 후 반환된 FileID를 사용하여 MMS를 발송해야 합니다.

3. MMS 메시지 텍스트는 최대 2,000바이트까지 지원됩니다. (한글 기준 약 1,000자)

4. 발송 간 이미지 파일 ID가 없는 경우 SMS/LMS로 자동 전환됩니다.

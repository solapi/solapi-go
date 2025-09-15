package storages

// UploadFileRequest represents POST /storage/v1/files body.
// file: base64 encoded content; optional name, type, link.
type UploadFileRequest struct {
	File string `json:"file"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Link string `json:"link,omitempty"`
}

// KakaoInfo is a placeholder for kakao field if needed by consumers.
// We don't expand structure as API does not mandate it for basic usage.
type KakaoInfo struct {
	Daou    string `json:"daou,omitempty"`
	Biztalk string `json:"biztalk,omitempty"`
}

// UploadFileResponse represents response from storage upload.
type UploadFileResponse struct {
	Kakao        KakaoInfo `json:"kakao,omitempty"`
	Type         string    `json:"type"`
	OriginalName string    `json:"originalName"`
	Link         string    `json:"link,omitempty"`
	FileID       string    `json:"fileId"`
	Name         string    `json:"name"`
	URL          string    `json:"url"`
	AccountID    string    `json:"accountId"`
	References   []any     `json:"references"`
	DateCreated  string    `json:"dateCreated"`
	DateUpdated  string    `json:"dateUpdated"`
}

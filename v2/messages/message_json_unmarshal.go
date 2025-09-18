package messages

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON ensures "to" supports both string and []string
func (m *Message) UnmarshalJSON(data []byte) error {
	type msgAlias struct {
		To           json.RawMessage   `json:"to"`
		From         string            `json:"from,omitempty"`
		Text         string            `json:"text,omitempty"`
		Subject      string            `json:"subject,omitempty"`
		ImageID      string            `json:"imageId,omitempty"`
		KakaoOptions *KakaoOptions     `json:"kakaoOptions,omitempty"`
		VoiceOptions *VoiceOptions     `json:"voiceOptions,omitempty"`
		FaxOptions   *FaxOptions       `json:"faxOptions,omitempty"`
		Country      string            `json:"country,omitempty"`
		CustomFields map[string]string `json:"customFields,omitempty"`
		Type         string            `json:"type,omitempty"`
		AutoType     *bool             `json:"autoTypeDetect,omitempty"`

		MessageID     string `json:"messageId,omitempty"`
		GroupID       string `json:"groupId,omitempty"`
		AccountID     string `json:"accountId,omitempty"`
		Status        string `json:"status,omitempty"`
		StatusCode    string `json:"statusCode,omitempty"`
		Replacement   *bool  `json:"replacement,omitempty"`
		DateCreated   string `json:"dateCreated,omitempty"`
		DateUpdated   string `json:"dateUpdated,omitempty"`
		DateProcessed string `json:"dateProcessed,omitempty"`
		DateReported  string `json:"dateReported,omitempty"`
		DateReceived  string `json:"dateReceived,omitempty"`
	}

	var v msgAlias
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	m.To = ""
	m.ToList = nil
	if len(v.To) != 0 && string(v.To) != "null" {
		var s string
		if err := json.Unmarshal(v.To, &s); err == nil {
			m.To = s
		} else {
			var arr []string
			if err2 := json.Unmarshal(v.To, &arr); err2 == nil {
				m.ToList = arr
			} else {
				return fmt.Errorf("messages: invalid to field, not string or []string: %s", string(v.To))
			}
		}
	}

	m.From = v.From
	m.Text = v.Text
	m.Subject = v.Subject
	m.ImageID = v.ImageID
	m.KakaoOptions = v.KakaoOptions
	m.VoiceOptions = v.VoiceOptions
	m.FaxOptions = v.FaxOptions
	m.Country = v.Country
	m.CustomFields = v.CustomFields
	m.Type = v.Type
	m.AutoType = v.AutoType

	m.MessageID = v.MessageID
	m.GroupID = v.GroupID
	m.AccountID = v.AccountID
	m.Status = v.Status
	m.StatusCode = v.StatusCode
	m.Replacement = v.Replacement
	m.DateCreated = v.DateCreated
	m.DateUpdated = v.DateUpdated
	m.DateProcessed = v.DateProcessed
	m.DateReported = v.DateReported
	m.DateReceived = v.DateReceived

	return nil
}

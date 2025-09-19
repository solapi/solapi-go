package messages

import "encoding/json"

// MarshalJSON ensures "to" supports both string and []string
func (m Message) MarshalJSON() ([]byte, error) {
	toVal := any(m.To)
	if len(m.ToList) > 0 {
		toVal = m.ToList
	}
	type msgAlias struct {
		To           any               `json:"to"`
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
	}
	return json.Marshal(msgAlias{
		To:           toVal,
		From:         m.From,
		Text:         m.Text,
		Subject:      m.Subject,
		ImageID:      m.ImageID,
		KakaoOptions: m.KakaoOptions,
		VoiceOptions: m.VoiceOptions,
		FaxOptions:   m.FaxOptions,
		Country:      m.Country,
		CustomFields: m.CustomFields,
		Type:         m.Type,
		AutoType:     m.AutoType,
	})
}

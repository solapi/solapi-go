package messages

import "encoding/json"

// Request types

// internal API-only types (not exported)
type apiAgent struct {
	SDKVersion string `json:"sdkVersion"`
	OSPlatform string `json:"osPlatform"`
	AppID      string `json:"appId,omitempty"`
}

type apiSendRequest struct {
	Messages        []Message `json:"messages"`
	AllowDuplicates *bool     `json:"allowDuplicates,omitempty"`
	Agent           *apiAgent `json:"agent,omitempty"`
	ScheduledDate   string    `json:"scheduledDate,omitempty"`
	ShowMessageList *bool     `json:"showMessageList,omitempty"`
}

// SendOptions is a user-facing configuration for SendMessages helper.
// It mirrors a subset of SendRequest minus Messages.

type SendOptions struct {
	AppId           string
	AllowDuplicates *bool
	ScheduledDate   string
	ShowMessageList *bool
}

type KakaoButton struct {
	Type          string `json:"type,omitempty"`
	Name          string `json:"name,omitempty"`
	LinkMobile    string `json:"linkMobile,omitempty"`
	LinkPc        string `json:"linkPc,omitempty"`
	SchemeIos     string `json:"schemeIos,omitempty"`
	SchemeAndroid string `json:"schemeAndroid,omitempty"`
	ChatExtra     string `json:"chatExtra,omitempty"`
	ChatEvent     string `json:"chatEvent,omitempty"`
	RelayId       string `json:"relayId,omitempty"`
	Keyword       string `json:"keyword,omitempty"`
}

// KakaoBMSOptions represents BMS-related options for Kakao messages.
// targeting should be one of "I", "M", or "N".
type KakaoBMSOptions struct {
	Targeting string `json:"targeting,omitempty"`
}

type KakaoOptions struct {
	PfID         string            `json:"pfId,omitempty"`
	TemplateID   string            `json:"templateId,omitempty"`
	DisableSms   *bool             `json:"disableSms,omitempty"`
	Title        string            `json:"title,omitempty"`
	Buttons      []KakaoButton     `json:"buttons,omitempty"`
	Variables    map[string]string `json:"variables,omitempty"`
	Replacements map[string]any    `json:"replacements,omitempty"`
	AdFlag       *bool             `json:"adFlag,omitempty"`
	ImageID      string            `json:"imageId,omitempty"`
	BMS          *KakaoBMSOptions  `json:"bms,omitempty"`
}

// VoiceOptions corresponds to voiceOptions payload for VOICE type
type VoiceOptions struct {
	VoiceType       string `json:"voiceType,omitempty"`
	HeaderMessage   string `json:"headerMessage,omitempty"`
	TailMessage     string `json:"tailMessage,omitempty"`
	ReplyRange      int    `json:"replyRange,omitempty"`
	CounselorNumber string `json:"counselorNumber,omitempty"`
}

// FaxOptions corresponds to faxOptions payload for FAX type
type FaxOptions struct {
	FileIDs []string `json:"fileIds"`
}

type Message struct {
	To           string            `json:"to"`
	ToList       []string          `json:"-"`
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

	// Common response-side fields (ignored on send if empty)
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

// MarshalJSON moved to message_json_marshal.go

type SendRequest struct {
	Messages        []Message `json:"messages"`
	AllowDuplicates *bool     `json:"allowDuplicates,omitempty"`
	ScheduledDate   string    `json:"scheduledDate,omitempty"`
	ShowMessageList *bool     `json:"showMessageList,omitempty"`

	// AppId is provided by SDK users; Service builds agent internally.
	AppId string `json:"-"`
}

// Response types

type GroupCount struct {
	Total             int `json:"total"`
	SentTotal         int `json:"sentTotal"`
	SentFailed        int `json:"sentFailed"`
	SentSuccess       int `json:"sentSuccess"`
	SentPending       int `json:"sentPending"`
	SentReplacement   int `json:"sentReplacement"`
	Refund            int `json:"refund"`
	RegisteredFailed  int `json:"registeredFailed"`
	RegisteredSuccess int `json:"registeredSuccess"`
}

type CountryCount map[string]int

type CoundForCharge struct {
	SMS             CountryCount `json:"sms,omitempty"`
	LMS             CountryCount `json:"lms,omitempty"`
	MMS             CountryCount `json:"mms,omitempty"`
	ATA             CountryCount `json:"ata,omitempty"`
	CTA             CountryCount `json:"cta,omitempty"`
	CTI             CountryCount `json:"cti,omitempty"`
	NSA             CountryCount `json:"nsa,omitempty"`
	RCSSMS          CountryCount `json:"rcs_sms,omitempty"`
	RCSLMS          CountryCount `json:"rcs_lms,omitempty"`
	RCSMMS          CountryCount `json:"rcs_mms,omitempty"`
	RCSTPL          CountryCount `json:"rcs_tpl,omitempty"`
	RCSITPL         CountryCount `json:"rcs_itpl,omitempty"`
	RCSLTPL         CountryCount `json:"rcs_ltpl,omitempty"`
	Fax             CountryCount `json:"fax,omitempty"`
	Voice           CountryCount `json:"voice,omitempty"`
	BMSText         CountryCount `json:"bms_text,omitempty"`
	BMSImage        CountryCount `json:"bms_image,omitempty"`
	BMSWide         CountryCount `json:"bms_wide,omitempty"`
	BMSWideItemList CountryCount `json:"bms_wide_item_list,omitempty"`
	BMSCarouselFeed CountryCount `json:"bms_carousel_feed,omitempty"`
	BMSPremiumVideo CountryCount `json:"bms_premium_video,omitempty"`
	BMSCommerce     CountryCount `json:"bms_commerce,omitempty"`
	BMSCarouselComm CountryCount `json:"bms_carousel_commerce,omitempty"`
}

type Amount struct {
	Requested   float64 `json:"requested"`
	Replacement float64 `json:"replacement"`
	Additional  float64 `json:"additional"`
	Refund      float64 `json:"refund"`
	Sum         float64 `json:"sum"`
}

type ProfitPerType struct {
	SMS             float64 `json:"sms"`
	LMS             float64 `json:"lms"`
	MMS             float64 `json:"mms"`
	ATA             float64 `json:"ata"`
	CTA             float64 `json:"cta"`
	CTI             float64 `json:"cti"`
	NSA             float64 `json:"nsa"`
	RCSSMS          float64 `json:"rcs_sms"`
	RCSLMS          float64 `json:"rcs_lms"`
	RCSMMS          float64 `json:"rcs_mms"`
	RCSTPL          float64 `json:"rcs_tpl"`
	RCSITPL         float64 `json:"rcs_itpl"`
	RCSLTPL         float64 `json:"rcs_ltpl"`
	Fax             float64 `json:"fax"`
	Voice           float64 `json:"voice"`
	BMSText         float64 `json:"bms_text"`
	BMSImage        float64 `json:"bms_image"`
	BMSWide         float64 `json:"bms_wide"`
	BMSWideItemList float64 `json:"bms_wide_item_list"`
	BMSCarouselFeed float64 `json:"bms_carousel_feed"`
	BMSPremiumVideo float64 `json:"bms_premium_video"`
	BMSCommerce     float64 `json:"bms_commerce"`
	BMSCarouselComm float64 `json:"bms_carousel_commerce"`
}

type AppInfo struct {
	Profit  ProfitPerType `json:"profit"`
	App     string        `json:"app"`
	Version string        `json:"version"`
}

type GroupInfo struct {
	Count          GroupCount     `json:"count"`
	CountForCharge CoundForCharge `json:"countForCharge,omitempty"`
	Balance        Amount         `json:"balance,omitempty"`
	Point          Amount         `json:"point,omitempty"`
	App            AppInfo        `json:"app,omitempty"`

	// Additional fields found in actual API response
	ID              string            `json:"_id"`
	ServiceMethod   string            `json:"serviceMethod"`
	APIVersion      string            `json:"apiVersion"`
	SDKVersion      string            `json:"sdkVersion"`
	OSPlatform      string            `json:"osPlatform"`
	Log             []LogEntry        `json:"log,omitempty"`
	Status          string            `json:"status"`
	DateSent        string            `json:"dateSent"`
	ScheduledDate   string            `json:"scheduledDate"`
	DateCompleted   string            `json:"dateCompleted"`
	IsRefunded      bool              `json:"isRefunded"`
	FlagUpdated     bool              `json:"flagUpdated"`
	CustomFields    map[string]string `json:"customFields"`
	Prepaid         bool              `json:"prepaid"`
	Strict          bool              `json:"strict"`
	Hint            json.RawMessage   `json:"hint"`
	MasterAccountID string            `json:"masterAccountId"`
	AllowDuplicates bool              `json:"allowDuplicates"`
	ClusterKey      *string           `json:"clusterKey"`
	AccountID       string            `json:"accountId"`
	GroupID         string            `json:"groupId"`
	Price           Price             `json:"price"` // Price field inside groupInfo
	DateCreated     string            `json:"dateCreated"`
	DateUpdated     string            `json:"dateUpdated"`
}

type LogEntry struct {
	CreateAt   string  `json:"createAt"`
	Message    string  `json:"message"`
	OldBalance float64 `json:"oldBalance"`
	NewBalance float64 `json:"newBalance"`
	OldPoint   float64 `json:"oldPoint"`
	NewPoint   float64 `json:"newPoint"`
	TotalPrice float64 `json:"totalPrice"`
}

type FailedMessage struct {
	To            string `json:"to"`
	From          string `json:"from"`
	Type          string `json:"type"`
	Country       string `json:"country"`
	MessageID     string `json:"messageId"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	AccountID     string `json:"accountId"`
}

// MessageListItem represents per-message result summary for send-many/detail
type MessageListItem struct {
	MessageID     string            `json:"messageId"`
	StatusCode    string            `json:"statusCode"`
	CustomFields  map[string]string `json:"customFields,omitempty"`
	StatusMessage string            `json:"statusMessage"`
}

type UnitPrice struct {
	SMS             float64 `json:"sms,omitempty"`
	LMS             float64 `json:"lms,omitempty"`
	MMS             float64 `json:"mms,omitempty"`
	ATA             float64 `json:"ata,omitempty"`
	CTA             float64 `json:"cta,omitempty"`
	CTI             float64 `json:"cti,omitempty"`
	NSA             float64 `json:"nsa,omitempty"`
	RCSSMS          float64 `json:"rcs_sms,omitempty"`
	RCSLMS          float64 `json:"rcs_lms,omitempty"`
	RCSMMS          float64 `json:"rcs_mms,omitempty"`
	RCSTPL          float64 `json:"rcs_tpl,omitempty"`
	RCSITPL         float64 `json:"rcs_itpl,omitempty"`
	RCSLTPL         float64 `json:"rcs_ltpl,omitempty"`
	Fax             float64 `json:"fax,omitempty"`
	Voice           float64 `json:"voice,omitempty"`
	BMSText         float64 `json:"bms_text,omitempty"`
	BMSImage        float64 `json:"bms_image,omitempty"`
	BMSWide         float64 `json:"bms_wide,omitempty"`
	BMSWideItemList float64 `json:"bms_wide_item_list,omitempty"`
	BMSCarouselFeed float64 `json:"bms_carousel_feed,omitempty"`
	BMSPremiumVideo float64 `json:"bms_premium_video,omitempty"`
	BMSCommerce     float64 `json:"bms_commerce,omitempty"`
	BMSCarouselComm float64 `json:"bms_carousel_commerce,omitempty"`
}

type Price map[string]UnitPrice

type DetailGroupMessageResponse struct {
	GroupInfo         GroupInfo         `json:"groupInfo"`
	FailedMessageList []FailedMessage   `json:"failedMessageList"`
	MessageList       []MessageListItem `json:"messageList,omitempty"`

	Status          string            `json:"status"`
	DateSent        string            `json:"dateSent"`
	ScheduledDate   string            `json:"scheduledDate"`
	DateCompleted   string            `json:"dateCompleted"`
	IsRefunded      bool              `json:"isRefunded"`
	FlagUpdated     bool              `json:"flagUpdated"`
	Prepaid         bool              `json:"prepaid"`
	Strict          bool              `json:"strict"`
	MasterAccountID string            `json:"masterAccountId"`
	AllowDuplicates bool              `json:"allowDuplicates"`
	ID              string            `json:"_id,omitempty"` // Support both _id and id
	AccountID       string            `json:"accountId"`
	APIVersion      string            `json:"apiVersion"`
	CustomFields    map[string]string `json:"customFields"`
	Hint            string            `json:"hint"`
	GroupID         string            `json:"groupId"`
	Price           Price             `json:"price"`

	ServiceMethod string     `json:"serviceMethod,omitempty"`
	SDKVersion    string     `json:"sdkVersion,omitempty"`
	OSPlatform    string     `json:"osPlatform,omitempty"`
	Log           []LogEntry `json:"log,omitempty"`

	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

// UnmarshalJSON implements custom JSON unmarshaling for DetailGroupMessageResponse
// to handle different API response formats
func (d *DetailGroupMessageResponse) UnmarshalJSON(data []byte) error {
	type detailGroupMessageResponseAlias struct {
		GroupInfo         GroupInfo         `json:"groupInfo"`
		FailedMessageList []FailedMessage   `json:"failedMessageList"`
		MessageList       []MessageListItem `json:"messageList,omitempty"`

		Status          string            `json:"status"`
		DateSent        string            `json:"dateSent"`
		ScheduledDate   *string           `json:"scheduledDate"` // Can be null
		DateCompleted   *string           `json:"dateCompleted"` // Can be null
		IsRefunded      bool              `json:"isRefunded"`
		FlagUpdated     bool              `json:"flagUpdated"`
		Prepaid         bool              `json:"prepaid"`
		Strict          bool              `json:"strict"`
		MasterAccountID *string           `json:"masterAccountId"` // Can be null
		AllowDuplicates bool              `json:"allowDuplicates"`
		ID              string            `json:"_id,omitempty"`
		IDAlt           string            `json:"id,omitempty"` // Alternative ID field
		AccountID       string            `json:"accountId"`
		APIVersion      string            `json:"apiVersion"`
		CustomFields    map[string]string `json:"customFields"`
		Hint            json.RawMessage   `json:"hint"` // Can be object or other format
		GroupID         string            `json:"groupId"`
		// Price is inside groupInfo, not at top level

		ServiceMethod string     `json:"serviceMethod,omitempty"`
		SDKVersion    string     `json:"sdkVersion,omitempty"`
		OSPlatform    string     `json:"osPlatform,omitempty"`
		Log           []LogEntry `json:"log,omitempty"`

		DateCreated string `json:"dateCreated"`
		DateUpdated string `json:"dateUpdated"`
	}

	var v detailGroupMessageResponseAlias
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	// Copy all fields - prefer top-level fields, fallback to groupInfo fields
	d.GroupInfo = v.GroupInfo
	d.FailedMessageList = v.FailedMessageList
	d.MessageList = v.MessageList

	// Use top-level fields if available, otherwise use groupInfo fields
	if v.Status != "" {
		d.Status = v.Status
	} else if v.GroupInfo.Status != "" {
		d.Status = v.GroupInfo.Status
	}

	if v.DateSent != "" {
		d.DateSent = v.DateSent
	} else if v.GroupInfo.DateSent != "" {
		d.DateSent = v.GroupInfo.DateSent
	}

	// Handle nullable fields
	if v.ScheduledDate != nil && *v.ScheduledDate != "" {
		d.ScheduledDate = *v.ScheduledDate
	} else if v.GroupInfo.ScheduledDate != "" {
		d.ScheduledDate = v.GroupInfo.ScheduledDate
	} else {
		d.ScheduledDate = ""
	}

	if v.DateCompleted != nil && *v.DateCompleted != "" {
		d.DateCompleted = *v.DateCompleted
	} else if v.GroupInfo.DateCompleted != "" {
		d.DateCompleted = v.GroupInfo.DateCompleted
	} else {
		d.DateCompleted = ""
	}

	if v.MasterAccountID != nil && *v.MasterAccountID != "" {
		d.MasterAccountID = *v.MasterAccountID
	} else if v.GroupInfo.MasterAccountID != "" {
		d.MasterAccountID = v.GroupInfo.MasterAccountID
	} else {
		d.MasterAccountID = ""
	}

	d.IsRefunded = v.IsRefunded || v.GroupInfo.IsRefunded
	d.FlagUpdated = v.FlagUpdated || v.GroupInfo.FlagUpdated
	d.Prepaid = v.Prepaid || v.GroupInfo.Prepaid
	d.Strict = v.Strict || v.GroupInfo.Strict
	d.AllowDuplicates = v.AllowDuplicates || v.GroupInfo.AllowDuplicates

	if v.AccountID != "" {
		d.AccountID = v.AccountID
	} else if v.GroupInfo.AccountID != "" {
		d.AccountID = v.GroupInfo.AccountID
	}

	if v.APIVersion != "" {
		d.APIVersion = v.APIVersion
	} else if v.GroupInfo.APIVersion != "" {
		d.APIVersion = v.GroupInfo.APIVersion
	}

	d.CustomFields = v.CustomFields
	if d.CustomFields == nil && v.GroupInfo.CustomFields != nil {
		d.CustomFields = v.GroupInfo.CustomFields
	}

	// Handle hint field - try to unmarshal as map, fallback to empty string
	if len(v.Hint) > 0 {
		var hintMap map[string]any
		if err := json.Unmarshal(v.Hint, &hintMap); err == nil {
			// Successfully unmarshaled as map, but we need to convert back to string for now
			// TODO: Update DetailGroupMessageResponse.Hint type to support map
			d.Hint = ""
		} else {
			// If unmarshaling fails, just set empty string
			d.Hint = ""
		}
	} else {
		d.Hint = ""
	}

	if v.GroupID != "" {
		d.GroupID = v.GroupID
	} else if v.GroupInfo.GroupID != "" {
		d.GroupID = v.GroupInfo.GroupID
	}

	d.ServiceMethod = v.ServiceMethod
	d.SDKVersion = v.SDKVersion
	d.OSPlatform = v.OSPlatform
	d.Log = v.Log
	if d.Log == nil && v.GroupInfo.Log != nil {
		d.Log = v.GroupInfo.Log
	}

	if v.DateCreated != "" {
		d.DateCreated = v.DateCreated
	} else if v.GroupInfo.DateCreated != "" {
		d.DateCreated = v.GroupInfo.DateCreated
	}

	if v.DateUpdated != "" {
		d.DateUpdated = v.DateUpdated
	} else if v.GroupInfo.DateUpdated != "" {
		d.DateUpdated = v.GroupInfo.DateUpdated
	}

	// Handle ID field - prefer _id, fallback to id
	if v.ID != "" {
		d.ID = v.ID
	} else if v.IDAlt != "" {
		d.ID = v.IDAlt
	}

	// Handle Price field - use from groupInfo since that's where it actually is
	d.Price = v.GroupInfo.Price

	return nil
}

// ListQuery describes filters for GET /messages/v4/list
type ListQuery struct {
	MessageID string
	GroupID   string
	To        string
	From      string
	TypeIn    []string

	DateType  string
	StartDate string
	EndDate   string

	StartKey string
	Limit    int
}

// MessageListResponse matches Kotlin SDK: messageList map + pagination fields
// kotlin: data class MessageListResponse(messageList: Map<String, Message>) : CommonListResponse
// CommonListResponse: limit, startKey, nextKey
type MessageListResponse struct {
	MessageList map[string]Message `json:"messageList"`
	Limit       int                `json:"limit"`
	StartKey    string             `json:"startKey"`
	NextKey     string             `json:"nextKey"`
}

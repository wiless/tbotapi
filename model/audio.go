package model

type Audio struct {
	FileBase
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
}

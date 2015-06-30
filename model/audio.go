package model

type Audio struct {
	File
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
}

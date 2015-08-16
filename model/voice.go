package model

type Voice struct {
	File
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
}

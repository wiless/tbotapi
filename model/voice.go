package model

// Voice represents a voice note
type Voice struct {
	FileBase
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
}

package model

// Audio represents an audio file to be treated as music
type Audio struct {
	FileBase
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
}

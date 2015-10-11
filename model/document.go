package model

// Document represents a general file
type Document struct {
	FileBase
	Thumbnail PhotoSize `json:"thumb"`
	Name      string    `json:"file_name"`
	MimeType  string    `json:"mime_type"`
}

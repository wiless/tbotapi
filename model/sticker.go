package model

// Sticker represents a sticker
type Sticker struct {
	FileBase
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Thumbnail PhotoSize `json:"thumb"`
}

package model

type Sticker struct {
	File
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Thumbnail PhotoSize `json:"thumb"`
}

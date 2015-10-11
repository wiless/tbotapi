package model

// PhotoSize represents one size of a photo or a thumbnail
type PhotoSize struct {
	FileBase
	Width  int `json:"width"`
	Height int `json:"height"`
}

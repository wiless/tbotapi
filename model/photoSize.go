package model

type PhotoSize struct {
	File
	Width  int `json:"width"`
	Height int `json:"height"`
}

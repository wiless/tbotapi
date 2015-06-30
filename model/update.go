package model

type Update struct {
	Id      int     `json:"id"`
	Message Message `json:"message"`
}

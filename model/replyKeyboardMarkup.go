package model

type ReplyKeyboardMarkup struct {
	Keyboard        [][]string `json:"keyboard"` // slice of keyboard lines
	ResizeKeyboard  bool       `json:"resize_keyboard"`
	OneTimeKeyboard bool       `json:"one_time_keyboard"`
	Selective       bool       `json:"selective"`
}

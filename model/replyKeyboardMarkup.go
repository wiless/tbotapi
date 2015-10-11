package model

// ReplyKeyboardMarkup represents a custom keyboard with reply options to be presented to clients
type ReplyKeyboardMarkup struct {
	Keyboard        [][]string `json:"keyboard"` // slice of keyboard lines
	ResizeKeyboard  bool       `json:"resize_keyboard"`
	OneTimeKeyboard bool       `json:"one_time_keyboard"`
	Selective       bool       `json:"selective"`
}

package model

// ReplyKeyboardHide contains the fields necessary to hide a custom keyboard
type ReplyKeyboardHide struct {
	HideKeyboard bool `json:"hide_keyboard"`
	Selective    bool `json:"selective"`
}

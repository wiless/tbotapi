package model

type ReplyMarkup interface {
	replyMarkup()
}

func (ReplyKeyboardHide) replyMarkup()   {}
func (ReplyKeyboardMarkup) replyMarkup() {}
func (ForceReply) replyMarkup()          {}

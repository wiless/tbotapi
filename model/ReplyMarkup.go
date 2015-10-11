package model

// ReplyMarkup is s marker interface for ReplyMarkups
type ReplyMarkup interface {
	replyMarkup()
}

func (ReplyKeyboardHide) replyMarkup()   {}
func (ReplyKeyboardMarkup) replyMarkup() {}
func (ForceReply) replyMarkup()          {}

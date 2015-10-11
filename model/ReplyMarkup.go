package model

type ReplyMarkup interface {
	replyMarkup()
}

func (ReplyKeyboardHide) replyMarkup()   {}
func (ReplyKeyboardMarkup) replyMarkup() {}
func (ForceReply) replyMarkup()          {}

/*func (r ReplyMarkup) MarshalJSON() ([]byte, error) {
	t := interface{}(r)

	switch t := t.(type) {
	case ReplyKeyboardMarkup:
		return json.Marshal(t)
	case ReplyKeyboardHide:
		return json.Marshal(t)
	case ForceReply:
		return json.Marshal(t)
	default:
		panic("tbotapi: ReplyMarkup must be of type ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply!")
	}
}*/

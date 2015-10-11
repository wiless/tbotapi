package model

type OutgoingMessage struct {
	OutgoingBase
	Text                  string    `json:"text"`
	DisableWebPagePreview bool      `json:"disable_web_page_preview,omitempty"`
	ParseMode             ParseMode `json:"parse_mode,omitempty"`
}

type Querystring map[string]string

func NewOutgoingMessage(recipient Recipient, text string) *OutgoingMessage {
	return &OutgoingMessage{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
		Text:      text,
		ParseMode: ModeDefault,
	}
}

func (om *OutgoingMessage) SetMarkdown(to bool) *OutgoingMessage {
	if to {
		om.ParseMode = ModeMarkdown
	} else {
		om.ParseMode = ModeDefault
	}
	return om
}

func (om *OutgoingMessage) SetDisableWebPagePreview(to bool) *OutgoingMessage {
	om.DisableWebPagePreview = to
	return om
}

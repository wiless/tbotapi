package model

// OutgoingMessage represents an outgoing message
type OutgoingMessage struct {
	OutgoingBase
	Text                  string    `json:"text"`
	DisableWebPagePreview bool      `json:"disable_web_page_preview,omitempty"`
	ParseMode             ParseMode `json:"parse_mode,omitempty"`
}

// Querystring is a type to represent querystring-applicable data
type Querystring map[string]string

// NewOutgoingMessage creates a new outgoing message
func NewOutgoingMessage(recipient Recipient, text string) *OutgoingMessage {
	return &OutgoingMessage{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
		Text:      text,
		ParseMode: ModeDefault,
	}
}

// SetMarkdown sets or resets whether the message should be parsed as markdown (optional)
func (om *OutgoingMessage) SetMarkdown(to bool) *OutgoingMessage {
	if to {
		om.ParseMode = ModeMarkdown
	} else {
		om.ParseMode = ModeDefault
	}
	return om
}

// SetDisableWebPagePreview disables web page previews for the message (optional)
func (om *OutgoingMessage) SetDisableWebPagePreview(to bool) *OutgoingMessage {
	om.DisableWebPagePreview = to
	return om
}

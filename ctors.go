package tbotapi

// NewOutgoingMessage creates a new outgoing message
func (api *TelegramBotAPI) NewOutgoingMessage(recipient Recipient, text string) *OutgoingMessage {
	return &OutgoingMessage{
		OutgoingBase: OutgoingBase{
			api:       api,
			Recipient: recipient,
		},
		Text:      text,
		ParseMode: ModeDefault,
	}
}

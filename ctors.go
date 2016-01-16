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

// NewOutgoingLocation creates a new outgoing location
func (api *TelegramBotAPI) NewOutgoingLocation(recipient Recipient, latitude, longitude float32) *OutgoingLocation {
	return &OutgoingLocation{
		OutgoingBase: OutgoingBase{
			api:       api,
			Recipient: recipient,
		},
		Latitude:  latitude,
		Longitude: longitude,
	}
}

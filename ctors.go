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

// NewOutgoingVideo creates a new outgoing video file
func (api *TelegramBotAPI) NewOutgoingVideo(recipient Recipient, filePath string) *OutgoingVideo {
	return &OutgoingVideo{
		OutgoingBase: OutgoingBase{
			api:       api,
			Recipient: recipient,
		},
		filePath: filePath,
	}
}

// NewOutgoingVideo creates a new outgoing video file
func (api *TelegramBotAPI) NewOutgoingVideoResend(recipient Recipient, fileID string) *OutgoingVideo {
	return &OutgoingVideo{
		OutgoingBase: OutgoingBase{
			api:       api,
			Recipient: recipient,
		},
		fileID: fileID,
	}
}

// NewOutgoingPhoto creates a new outgoing photo
func (api *TelegramBotAPI) NewOutgoingPhoto(recipient Recipient, filePath string) *OutgoingPhoto {
	return &OutgoingPhoto{
		OutgoingBase: OutgoingBase{
			api:       api,
			Recipient: recipient,
		},
		filePath: filePath,
	}
}

// NewOutgoingPhoto creates a new outgoing photo
func (api *TelegramBotAPI) NewOutgoingPhotoResend(recipient Recipient, fileID string) *OutgoingPhoto {
	return &OutgoingPhoto{
		OutgoingBase: OutgoingBase{
			api:       api,
			Recipient: recipient,
		},
		fileID: fileID,
	}
}

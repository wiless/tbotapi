package tbotapi

// NewOutgoingMessage creates a new outgoing message
func (api *TelegramBotAPI) NewOutgoingMessage(recipient Recipient, text string) *OutgoingMessage {
	return &OutgoingMessage{
		outgoingBase: outgoingBase{
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
		outgoingBase: outgoingBase{
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
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		filePath: filePath,
	}
}

// NewOutgoingVideo creates a new outgoing video file
func (api *TelegramBotAPI) NewOutgoingVideoResend(recipient Recipient, fileID string) *OutgoingVideo {
	return &OutgoingVideo{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		fileID: fileID,
	}
}

// NewOutgoingPhoto creates a new outgoing photo
func (api *TelegramBotAPI) NewOutgoingPhoto(recipient Recipient, filePath string) *OutgoingPhoto {
	return &OutgoingPhoto{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		filePath: filePath,
	}
}

// NewOutgoingPhoto creates a new outgoing photo
func (api *TelegramBotAPI) NewOutgoingPhotoResend(recipient Recipient, fileID string) *OutgoingPhoto {
	return &OutgoingPhoto{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		fileID: fileID,
	}
}

// NewOutgoingSticker creates a new outgoing sticker message
func (api *TelegramBotAPI) NewOutgoingSticker(recipient Recipient, filePath string) *OutgoingSticker {
	return &OutgoingSticker{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		filePath: filePath,
	}
}

// NewOutgoingSticker creates a new outgoing sticker message
func (api *TelegramBotAPI) NewOutgoingStickerResend(recipient Recipient, fileID string) *OutgoingSticker {
	return &OutgoingSticker{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		fileID: fileID,
	}
}

// NewOutgoingVoice creates a new outgoing voice note
func (api *TelegramBotAPI) NewOutgoingVoice(recipient Recipient, filePath string) *OutgoingVoice {
	return &OutgoingVoice{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		filePath: filePath,
	}
}

// NewOutgoingVoice creates a new outgoing voice note
func (api *TelegramBotAPI) NewOutgoingVoiceResend(recipient Recipient, fileID string) *OutgoingVoice {
	return &OutgoingVoice{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		fileID: fileID,
	}
}

// NewOutgoingAudio creates a new outgoing audio file
func (api *TelegramBotAPI) NewOutgoingAudio(recipient Recipient, filePath string) *OutgoingAudio {
	return &OutgoingAudio{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		filePath: filePath,
	}
}

// NewOutgoingAudio creates a new outgoing audio file
func (api *TelegramBotAPI) NewOutgoingAudioResend(recipient Recipient, fileID string) *OutgoingAudio {
	return &OutgoingAudio{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		fileID: fileID,
	}
}

// NewOutgoingDocument creates a new outgoing file
func (api *TelegramBotAPI) NewOutgoingDocument(recipient Recipient, filePath string) *OutgoingDocument {
	return &OutgoingDocument{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		filePath: filePath,
	}
}

// NewOutgoingDocument creates a new outgoing file
func (api *TelegramBotAPI) NewOutgoingDocumentResend(recipient Recipient, fileID string) *OutgoingDocument {
	return &OutgoingDocument{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		fileID: fileID,
	}
}

// NewOutgoingForward creates a new outgoing, forwarded message
func (api *TelegramBotAPI) NewOutgoingForward(recipient Recipient, origin Chat, messageID int) *OutgoingForward {
	return &OutgoingForward{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		FromChatID: NewRecipientFromChat(origin),
		MessageID:  messageID,
	}
}

func (api *TelegramBotAPI) NewOutgoingChatAction(recipient Recipient, action ChatAction) *OutgoingChatAction {
	return &OutgoingChatAction{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		Action: action,
	}
}

// NewOutgoingUserProfilePhotosRequest creates a new request for a users profile photos
func (api *TelegramBotAPI) NewOutgoingUserProfilePhotosRequest(userID int) *OutgoingUserProfilePhotosRequest {
	return &OutgoingUserProfilePhotosRequest{
		api:    api,
		UserID: userID,
	}
}

func (api *TelegramBotAPI) NewInlineQueryAnswer(queryID string, results []InlineQueryResult) *InlineQueryAnswer {
	return &InlineQueryAnswer{
		api:     api,
		QueryID: queryID,
		Results: results,
	}
}

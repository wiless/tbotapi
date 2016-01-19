package tbotapi

import "io"

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
func (api *TelegramBotAPI) NewOutgoingVideo(recipient Recipient, fileName string, reader io.Reader) *OutgoingVideo {
	return &OutgoingVideo{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		outgoingFileBase: outgoingFileBase{
			fileName: fileName,
			r:        reader,
		},
	}
}

// NewOutgoingVideoResend creates a new outgoing video file for re-sending
func (api *TelegramBotAPI) NewOutgoingVideoResend(recipient Recipient, fileID string) *OutgoingVideo {
	return &OutgoingVideo{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		outgoingFileBase: outgoingFileBase{
			fileID: fileID,
		},
	}
}

// NewOutgoingPhoto creates a new outgoing photo
func (api *TelegramBotAPI) NewOutgoingPhoto(recipient Recipient, fileName string, reader io.Reader) *OutgoingPhoto {
	return &OutgoingPhoto{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		outgoingFileBase: outgoingFileBase{
			fileName: fileName,
			r:        reader,
		},
	}
}

// NewOutgoingPhotoResend creates a new outgoing photo for re-sending
func (api *TelegramBotAPI) NewOutgoingPhotoResend(recipient Recipient, fileID string) *OutgoingPhoto {
	return &OutgoingPhoto{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		outgoingFileBase: outgoingFileBase{
			fileID: fileID,
		},
	}
}

// NewOutgoingSticker creates a new outgoing sticker message
func (api *TelegramBotAPI) NewOutgoingSticker(recipient Recipient, fileName string, reader io.Reader) *OutgoingSticker {
	return &OutgoingSticker{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		outgoingFileBase: outgoingFileBase{
			fileName: fileName,
			r:        reader,
		},
	}
}

// NewOutgoingStickerResend creates a new outgoing sticker message for re-sending
func (api *TelegramBotAPI) NewOutgoingStickerResend(recipient Recipient, fileID string) *OutgoingSticker {
	return &OutgoingSticker{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		outgoingFileBase: outgoingFileBase{
			fileID: fileID,
		},
	}
}

// NewOutgoingVoice creates a new outgoing voice note
func (api *TelegramBotAPI) NewOutgoingVoice(recipient Recipient, fileName string, reader io.Reader) *OutgoingVoice {
	return &OutgoingVoice{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		outgoingFileBase: outgoingFileBase{
			fileName: fileName,
			r:        reader,
		},
	}
}

// NewOutgoingVoiceResend creates a new outgoing voice note for re-sending
func (api *TelegramBotAPI) NewOutgoingVoiceResend(recipient Recipient, fileID string) *OutgoingVoice {
	return &OutgoingVoice{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		outgoingFileBase: outgoingFileBase{
			fileID: fileID,
		},
	}
}

// NewOutgoingAudio creates a new outgoing audio file
func (api *TelegramBotAPI) NewOutgoingAudio(recipient Recipient, fileName string, reader io.Reader) *OutgoingAudio {
	return &OutgoingAudio{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		outgoingFileBase: outgoingFileBase{
			fileName: fileName,
			r:        reader,
		},
	}
}

// NewOutgoingAudioResend creates a new outgoing audio file for re-sending
func (api *TelegramBotAPI) NewOutgoingAudioResend(recipient Recipient, fileID string) *OutgoingAudio {
	return &OutgoingAudio{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		outgoingFileBase: outgoingFileBase{
			fileID: fileID,
		},
	}
}

// NewOutgoingDocument creates a new outgoing file
func (api *TelegramBotAPI) NewOutgoingDocument(recipient Recipient, fileName string, reader io.Reader) *OutgoingDocument {
	return &OutgoingDocument{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		outgoingFileBase: outgoingFileBase{
			fileName: fileName,
			r:        reader,
		},
	}
}

// NewOutgoingDocumentResend creates a new outgoing file for re-sending
func (api *TelegramBotAPI) NewOutgoingDocumentResend(recipient Recipient, fileID string) *OutgoingDocument {
	return &OutgoingDocument{
		outgoingBase: outgoingBase{
			api:       api,
			Recipient: recipient,
		},
		outgoingFileBase: outgoingFileBase{
			fileID: fileID,
		},
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

// NewOutgoingChatAction creates a new outgoing chat action
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

// NewInlineQueryAnswer creates a new inline query answer
func (api *TelegramBotAPI) NewInlineQueryAnswer(queryID string, results []InlineQueryResult) *InlineQueryAnswer {
	return &InlineQueryAnswer{
		api:     api,
		QueryID: queryID,
		Results: results,
	}
}

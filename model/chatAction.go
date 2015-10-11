package model

// ChatAction represents an action to be shown to clients, indicating activity of the bot
type ChatAction string

// Represents all the possible ChatActions to be sent, see https://core.telegram.org/bots/api#sendchataction
const (
	ChatActionTyping         ChatAction = "typing"
	ChatActionUploadPhoto               = "upload_photo"
	ChatActionRecordVideo               = "record_video"
	ChatActionUploadVideo               = "upload_video"
	ChatActionRecordAudio               = "record_audio"
	ChatActionUploadAudio               = "upload_audio"
	ChatActionUploadDocument            = "upload_document"
	ChatActionFindLocation              = "find_location"
)

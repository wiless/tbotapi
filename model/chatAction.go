package model

type ChatAction string

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

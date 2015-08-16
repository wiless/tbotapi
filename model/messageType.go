package model

type MessageType int

const (
	TEXT MessageType = iota
	AUDIO
	DOCUMENT
	PHOTO
	STICKER
	VIDEO
	VOICE
	CONTACT
	LOCATION

	chatAction_beg
	NEW_CHAT_PARTICIPANT
	LEFT_CHAT_PARTICIPANT
	NEW_CHAT_TITLE
	NEW_CHAT_PHOTO
	DELETE_CHAT_PHOTO
	GROUP_CHAT_CREATED
	chatAction_end

	UNKNOWN
)

var types = map[MessageType]string{
	TEXT:     "TEXT",
	AUDIO:    "AUDIO",
	DOCUMENT: "DOCUMENT",
	PHOTO:    "PHOTO",
	STICKER:  "STICKER",
	VIDEO:    "VIDEO",
	VOICE:    "VOICE",
	CONTACT:  "CONTACT",
	LOCATION: "LOCATION",

	NEW_CHAT_PARTICIPANT:  "NEW_CHAT_PARTICIPANT",
	LEFT_CHAT_PARTICIPANT: "LEFT_CHAT_PARTICIPANT",
	NEW_CHAT_TITLE:        "NEW_CHAT_TITLE",
	NEW_CHAT_PHOTO:        "NEW_CHAT_PHOTO",
	DELETE_CHAT_PHOTO:     "DELETE_CHAT_PHOTO",
	GROUP_CHAT_CREATED:    "GROUP_CHAT_CREATED",

	UNKNOWN: "UNKNOWN",
}

func (mt MessageType) IsChatAction() bool {
	return mt > chatAction_beg && mt < chatAction_end
}

func (mt MessageType) String() string {
	val, ok := types[mt]
	if !ok {
		return "unknown"
	}
	return val
}

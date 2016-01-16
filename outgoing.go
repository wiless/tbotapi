package tbotapi

import "fmt"
import "encoding/json"

// OutgoingBase contains fields shared by most of the outgoing messages
type outgoingBase struct {
	api                 *TelegramBotAPI
	Recipient           Recipient   `json:"chat_id"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
	replyToMessageIDSet bool
	replyMarkupSet      bool
}

// SetReplyToMessageID sets the ID for the message to reply to (optional)
func (op *outgoingBase) SetReplyToMessageID(to int) {
	op.ReplyToMessageID = to
	op.replyToMessageIDSet = true
}

// SetReplyKeyboardMarkup sets the ReplyKeyboardMarkup (optional)
// Note that only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set.
// Attempting to set any of the other two or re-setting this will cause a panic.
func (op *outgoingBase) SetReplyKeyboardMarkup(to ReplyKeyboardMarkup) {
	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.ReplyMarkup = ReplyMarkup(to)
}

// SetReplyKeyboardHide sets the ReplyKeyboardHide (optional)
// Note that only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set.
// Attempting to set any of the other two or re-setting this will cause a panic.
func (op *outgoingBase) SetReplyKeyboardHide(to ReplyKeyboardHide) {
	if !to.HideKeyboard {
		return
	}

	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.ReplyMarkup = ReplyMarkup(to)
}

// SetForceReply sets ForceReply for this message (optional)
// Note that only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set.
// Attempting to set any of the other two or re-setting this will cause a panic.
func (op *outgoingBase) SetForceReply(to ForceReply) {
	if !to.ForceReply {
		return
	}

	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.ReplyMarkup = ReplyMarkup(to)
}

// GetBaseQueryString gets a Querystring representing this message
func (op *outgoingBase) getBaseQueryString() querystring {
	toReturn := map[string]string{}
	if op.Recipient.isChannel() {
		//Channel
		toReturn["chat_id"] = fmt.Sprint(*op.Recipient.ChannelID)
	} else {
		toReturn["chat_id"] = fmt.Sprint(*op.Recipient.ChatID)
	}

	if op.replyToMessageIDSet {
		toReturn["reply_to_message_id"] = fmt.Sprint(op.ReplyToMessageID)
	}

	if op.replyMarkupSet {
		b, err := json.Marshal(op.ReplyMarkup)
		if err != nil {
			panic(err)
		}
		toReturn["reply_markup"] = string(b)
	}

	return querystring(toReturn)
}

// OutgoingAudio represents an outgoing audio file
type OutgoingAudio struct {
	outgoingBase
	filePath  string
	fileID    string
	Duration  int    `json:"duration,omitempty"`
	Title     string `json:"title,omitempty"`
	Performer string `json:"performer,omitempty"`
}

// SetDuration sets a duration for the audio file (optional)
func (oa *OutgoingAudio) SetDuration(to int) *OutgoingAudio {
	oa.Duration = to
	return oa
}

// SetPerformer sets a performer for the audio file (optional)
func (oa *OutgoingAudio) SetPerformer(to string) *OutgoingAudio {
	oa.Performer = to
	return oa
}

// SetTitle sets a title for the audio file (optional)
func (oa *OutgoingAudio) SetTitle(to string) *OutgoingAudio {
	oa.Title = to
	return oa
}

// querystring implements querystringer to represent the audio file
func (oa *OutgoingAudio) querystring() querystring {
	toReturn := map[string]string(oa.getBaseQueryString())

	if oa.Duration != 0 {
		toReturn["duration"] = fmt.Sprint(oa.Duration)
	}

	if oa.Performer != "" {
		toReturn["performer"] = oa.Performer
	}

	if oa.Title != "" {
		toReturn["title"] = oa.Title
	}

	return querystring(toReturn)
}

// OutgoingDocument represents an outgoing file
type OutgoingDocument struct {
	outgoingBase
	filePath string
	fileID   string
}

// querystring implements querystringer to represent the outgoing file
func (od *OutgoingDocument) querystring() querystring {
	return od.getBaseQueryString()
}

// OutgoingForward represents an outgoing, forwarded message
type OutgoingForward struct {
	outgoingBase
	FromChatID Recipient `json:"from_chat_id"`
	MessageID  int       `json:"message_id"`
}

// OutgoingLocation represents an outgoing location on a map
type OutgoingLocation struct {
	outgoingBase
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// OutgoingMessage represents an outgoing message
type OutgoingMessage struct {
	outgoingBase
	Text                  string    `json:"text"`
	DisableWebPagePreview bool      `json:"disable_web_page_preview,omitempty"`
	ParseMode             ParseMode `json:"parse_mode,omitempty"`
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

// OutgoingPhoto represents an outgoing photo
type OutgoingPhoto struct {
	outgoingBase
	filePath string
	fileID   string
	Caption  string `json:"caption,omitempty"`
}

// SetCaption sets a caption for the photo (optional)
func (op *OutgoingPhoto) SetCaption(to string) *OutgoingPhoto {
	op.Caption = to
	return op
}

// querystring implements querystringer to represent the photo
func (op *OutgoingPhoto) querystring() querystring {
	toReturn := map[string]string(op.getBaseQueryString())

	if op.Caption != "" {
		toReturn["caption"] = op.Caption
	}

	return querystring(toReturn)
}

// OutgoingSticker represents an outgoing sticker message
type OutgoingSticker struct {
	outgoingBase
	filePath string
	fileID   string
}

// querystring implements querystringer to represent the sticker message
func (os *OutgoingSticker) querystring() querystring {
	return os.getBaseQueryString()
}

// OutgoingUserProfilePhotosRequest represents a request for a users profile photos
type OutgoingUserProfilePhotosRequest struct {
	api    *TelegramBotAPI
	UserID int `json:"user_id"`
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

// SetOffset sets an offset for the request (optional)
func (op *OutgoingUserProfilePhotosRequest) SetOffset(to int) *OutgoingUserProfilePhotosRequest {
	op.Offset = to
	return op
}

// SetLimit sets a limit for the request (optional)
func (op *OutgoingUserProfilePhotosRequest) SetLimit(to int) *OutgoingUserProfilePhotosRequest {
	op.Limit = to
	return op
}

// querystring implements querystringer to represent the request
func (op *OutgoingUserProfilePhotosRequest) querystring() querystring {
	toReturn := map[string]string{}
	toReturn["user_id"] = fmt.Sprint(op.UserID)

	if op.Offset != 0 {
		toReturn["offset"] = fmt.Sprint(op.Offset)
	}

	if op.Limit != 0 {
		toReturn["limit"] = fmt.Sprint(op.Limit)
	}

	return querystring(toReturn)
}

// OutgoingVideo represents an outgoing video file
type OutgoingVideo struct {
	outgoingBase
	fileID   string
	filePath string
	Duration int    `json:"duration,omitempty"`
	Caption  string `json:"caption,omitempty"`
}

// SetCaption sets a caption for the video file (optional)
func (ov *OutgoingVideo) SetCaption(to string) *OutgoingVideo {
	ov.Caption = to
	return ov
}

// SetDuration sets a duration for the video file (optional)
func (ov *OutgoingVideo) SetDuration(to int) *OutgoingVideo {
	ov.Duration = to
	return ov
}

// querystring implements querystringer to represent the outgoing video file
func (ov *OutgoingVideo) querystring() querystring {
	toReturn := map[string]string(ov.getBaseQueryString())

	if ov.Caption != "" {
		toReturn["caption"] = ov.Caption
	}

	if ov.Duration != 0 {
		toReturn["duration"] = fmt.Sprint(ov.Duration)
	}

	return querystring(toReturn)
}

// OutgoingVoice represents an outgoing voice note
type OutgoingVoice struct {
	outgoingBase
	filePath string
	fileID   string
	Duration int `json:"duration,omitempty"`
}

// SetDuration sets a duration of the voice note (optional)
func (ov *OutgoingVoice) SetDuration(to int) *OutgoingVoice {
	ov.Duration = to
	return ov
}

// querystring implements querystringer to represent the outgoing voice note
func (ov *OutgoingVoice) querystring() querystring {
	toReturn := map[string]string(ov.getBaseQueryString())

	if ov.Duration != 0 {
		toReturn["duration"] = fmt.Sprint(ov.Duration)
	}

	return querystring(toReturn)
}

// ReplyMarkup is s marker interface for ReplyMarkups
type ReplyMarkup interface {
	replyMarkup()
}

func (ReplyKeyboardHide) replyMarkup()   {}
func (ReplyKeyboardMarkup) replyMarkup() {}
func (ForceReply) replyMarkup()          {}

// ForceReply represents the values sent by a bot so that clients will be presented with a forced reply, see https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options to be presented to clients
type ReplyKeyboardMarkup struct {
	Keyboard        [][]string `json:"keyboard"` // slice of keyboard lines
	ResizeKeyboard  bool       `json:"resize_keyboard"`
	OneTimeKeyboard bool       `json:"one_time_keyboard"`
	Selective       bool       `json:"selective"`
}

// ReplyKeyboardHide contains the fields necessary to hide a custom keyboard
type ReplyKeyboardHide struct {
	HideKeyboard bool `json:"hide_keyboard"`
	Selective    bool `json:"selective"`
}

// ParseMode describes how a message should be parsed client-side
type ParseMode string

//ParseModes
const (
	ModeMarkdown = ParseMode("Markdown") // Parse as Markdown
	ModeDefault  = ParseMode("")         //Parse as text
)

type OutgoingChatAction struct {
	outgoingBase
	Action ChatAction `json:"action"`
}

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

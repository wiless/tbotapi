package tbotapi

import "fmt"
import "encoding/json"

// OutgoingBase contains fields shared by most of the outgoing messages
type OutgoingBase struct {
	api                 *TelegramBotAPI
	Recipient           Recipient   `json:"chat_id"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
	replyToMessageIDSet bool
	replyMarkupSet      bool
}

// SetReplyToMessageID sets the ID for the message to reply to (optional)
func (op *OutgoingBase) SetReplyToMessageID(to int) {
	op.ReplyToMessageID = to
	op.replyToMessageIDSet = true
}

// SetReplyKeyboardMarkup sets the ReplyKeyboardMarkup (optional)
// Note that only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set.
// Attempting to set any of the other two or re-setting this will cause a panic.
func (op *OutgoingBase) SetReplyKeyboardMarkup(to ReplyKeyboardMarkup) {
	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.ReplyMarkup = ReplyMarkup(to)
}

// SetReplyKeyboardHide sets the ReplyKeyboardHide (optional)
// Note that only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set.
// Attempting to set any of the other two or re-setting this will cause a panic.
func (op *OutgoingBase) SetReplyKeyboardHide(to ReplyKeyboardHide) {
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
func (op *OutgoingBase) SetForceReply(to ForceReply) {
	if !to.ForceReply {
		return
	}

	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.ReplyMarkup = ReplyMarkup(to)
}

// GetBaseQueryString gets a Querystring representing this message
func (op *OutgoingBase) GetBaseQueryString() Querystring {
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

	return Querystring(toReturn)
}

// OutgoingAudio represents an outgoing audio file
type OutgoingAudio struct {
	OutgoingBase
	Duration  int    `json:"duration,omitempty"`
	Title     string `json:"title,omitempty"`
	Performer string `json:"performer,omitempty"`
}

// NewOutgoingAudio creates a new outgoing audio file
func NewOutgoingAudio(recipient Recipient) *OutgoingAudio {
	return &OutgoingAudio{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
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

// GetQueryString returns a Querystring representing the audio file
func (oa *OutgoingAudio) GetQueryString() Querystring {
	toReturn := map[string]string(oa.GetBaseQueryString())

	if oa.Duration != 0 {
		toReturn["duration"] = fmt.Sprint(oa.Duration)
	}

	if oa.Performer != "" {
		toReturn["performer"] = oa.Performer
	}

	if oa.Title != "" {
		toReturn["title"] = oa.Title
	}

	return Querystring(toReturn)
}

// OutgoingDocument represents an outgoing file
type OutgoingDocument struct {
	OutgoingBase
}

// NewOutgoingDocument creates a new outgoing file
func NewOutgoingDocument(recipient Recipient) *OutgoingDocument {
	return &OutgoingDocument{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

// GetQueryString returns a Querystring representing the outgoing file
func (od *OutgoingDocument) GetQueryString() Querystring {
	return od.GetBaseQueryString()
}

// OutgoingForward represents an outgoing, forwarded message
type OutgoingForward struct {
	OutgoingBase
	FromChatID Recipient `json:"from_chat_id"`
	MessageID  int       `json:"message_id"`
}

// NewOutgoingForward creates a new outgoing, forwarded message
func NewOutgoingForward(recipient Recipient, origin Chat, messageID int) *OutgoingForward {
	return &OutgoingForward{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
		FromChatID: NewRecipientFromChat(origin),
		MessageID:  messageID,
	}
}

// OutgoingLocation represents an outgoing location on a map
type OutgoingLocation struct {
	OutgoingBase
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// GetQueryString returns a Querystring representing the location
func (ol *OutgoingLocation) GetQueryString() Querystring {
	toReturn := map[string]string(ol.GetBaseQueryString())

	toReturn["latitude"] = fmt.Sprint(ol.Latitude)
	toReturn["longitude"] = fmt.Sprint(ol.Longitude)

	return Querystring(toReturn)
}

// OutgoingMessage represents an outgoing message
type OutgoingMessage struct {
	OutgoingBase
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
	OutgoingBase
	Caption string `json:"caption,omitempty"`
}

// NewOutgoingPhoto creates a new outgoing photo
func NewOutgoingPhoto(recipient Recipient) *OutgoingPhoto {
	return &OutgoingPhoto{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

// SetCaption sets a caption for the photo (optional)
func (op *OutgoingPhoto) SetCaption(to string) *OutgoingPhoto {
	op.Caption = to
	return op
}

// GetQueryString returns a Querystring representing the photo
func (op *OutgoingPhoto) GetQueryString() Querystring {
	toReturn := map[string]string(op.GetBaseQueryString())

	if op.Caption != "" {
		toReturn["caption"] = op.Caption
	}

	return Querystring(toReturn)
}

// OutgoingSticker represents an outgoing sticker message
type OutgoingSticker struct {
	OutgoingBase
}

// NewOutgoingSticker creates a new outgoing sticker message
func NewOutgoingSticker(recipient Recipient) *OutgoingSticker {
	return &OutgoingSticker{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

// GetQueryString returns a Querystring representing the sticker message
func (os *OutgoingSticker) GetQueryString() Querystring {
	return os.GetBaseQueryString()
}

// OutgoingUserProfilePhotosRequest represents a request for a users profile photos
type OutgoingUserProfilePhotosRequest struct {
	UserID int `json:"user_id"`
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

// NewOutgoingUserProfilePhotosRequest creates a new request for a users profile photos
func NewOutgoingUserProfilePhotosRequest(userID int) *OutgoingUserProfilePhotosRequest {
	return &OutgoingUserProfilePhotosRequest{
		UserID: userID,
	}
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

// GetQueryString returns a Querystring representing the request
func (op *OutgoingUserProfilePhotosRequest) GetQueryString() Querystring {
	toReturn := map[string]string{}
	toReturn["user_id"] = fmt.Sprint(op.UserID)

	if op.Offset != 0 {
		toReturn["offset"] = fmt.Sprint(op.Offset)
	}

	if op.Limit != 0 {
		toReturn["limit"] = fmt.Sprint(op.Limit)
	}

	return Querystring(toReturn)
}

// OutgoingVideo represents an outgoing video file
type OutgoingVideo struct {
	OutgoingBase
	Duration int    `json:"duration,omitempty"`
	Caption  string `json:"caption,omitempty"`
}

// NewOutgoingVideo creates a new outgoing video file
func NewOutgoingVideo(recipient Recipient) *OutgoingVideo {
	return &OutgoingVideo{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
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

// GetQueryString returns a Querystring representing the outgoing video file
func (ov *OutgoingVideo) GetQueryString() Querystring {
	toReturn := map[string]string(ov.GetBaseQueryString())

	if ov.Caption != "" {
		toReturn["caption"] = ov.Caption
	}

	if ov.Duration != 0 {
		toReturn["duration"] = fmt.Sprint(ov.Duration)
	}

	return Querystring(toReturn)
}

// OutgoingVoice represents an outgoing voice note
type OutgoingVoice struct {
	OutgoingBase
	Duration int `json:"duration,omitempty"`
}

// NewOutgoingVoice creates a new outgoing voice note
func NewOutgoingVoice(recipient Recipient) *OutgoingVoice {
	return &OutgoingVoice{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

// SetDuration sets a duration of the voice note (optional)
func (ov *OutgoingVoice) SetDuration(to int) *OutgoingVoice {
	ov.Duration = to
	return ov
}

// GetQueryString returns a Querystring representing the outgoing voice note
func (ov *OutgoingVoice) GetQueryString() Querystring {
	toReturn := map[string]string(ov.GetBaseQueryString())

	if ov.Duration != 0 {
		toReturn["duration"] = fmt.Sprint(ov.Duration)
	}

	return Querystring(toReturn)
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

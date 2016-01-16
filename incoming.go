package tbotapi

import "fmt"
import "sort"

// BaseResponse contains the basic fields contained in every API response
type baseResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
	ErrorCode   int    `json:"error_code"`
}

// Audio represents an audio file to be treated as music
type Audio struct {
	FileBase
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
}

// Chat contains information about the chat a message originated from
type Chat struct {
	ID        int     `json:"id"`         // Unique identifier for this chat
	Type      string  `json:"type"`       // Type of chat, can be either "private", "group" or "channel". Check Is(PrivateChat|GroupChat|Channel)() methods
	Title     *string `json:"title"`      // Title for channels and group chats
	Username  *string `json:"username"`   // Username for private chats and channels if available
	FirstName *string `json:"first_name"` // First name of the other party in a private chat
	LastName  *string `json:"last_name"`  // Last name of the other party in a private chat
}

// IsPrivateChat checks if the chat is a private chat
func (c Chat) IsPrivateChat() bool {
	return c.Type == "private"
}

// IsGroupChat checks if the chat is a group chat
func (c Chat) IsGroupChat() bool {
	return c.Type == "group"
}

// IsChannel checks if the chat is a channel
func (c Chat) IsChannel() bool {
	return c.Type == "channel"
}

func (c Chat) String() string {
	toReturn := fmt.Sprint(c.ID)
	if c.IsPrivateChat() {
		toReturn += " (P) "
	} else if c.IsGroupChat() {
		toReturn += " (G) "
	} else {
		toReturn += " (C) "
	}

	if c.Title != nil {
		toReturn += "\"" + *c.Title + "\" "
	}

	if c.FirstName != nil {
		toReturn += *c.FirstName + " "
	}

	if c.LastName != nil {
		toReturn += *c.LastName + " "
	}

	if c.Username != nil {
		toReturn += "(@" + *c.Username + ")"
	}

	return toReturn
}

// Contact represents a phone contact
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	ID          int    `json:"user_id"`
}

// Document represents a general file
type Document struct {
	FileBase
	Thumbnail PhotoSize `json:"thumb"`
	Name      string    `json:"file_name"`
	MimeType  string    `json:"mime_type"`
}

// FileBase contains all the fields present in every file-like API response
type FileBase struct {
	ID   string `json:"file_id"`
	Size int    `json:"file_size"`
}

// File represents a file ready to be downloaded
type File struct {
	FileBase
	Path string `json:"file_path"`
}

// FileResponse represents the response sent by the API when requesting a file for download
type FileResponse struct {
	baseResponse
	File File `json:"result"`
}

// Location represents a point on the map
type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

// MessageResponse represents the response sent by the API on successful messages sent
type MessageResponse struct {
	baseResponse
	Message Message `json:"result"`
}

// Message represents a message
type Message struct {
	noReplyMessage
	ReplyToMessage *noReplyMessage `json:"reply_to_message"`
}

// IsForwarded checks if the message was forwarded
func (m *Message) IsForwarded() bool {
	return m.ForwardFrom != nil
}

// IsReply checks if the message is a reply
func (m *Message) IsReply() bool {
	return m.ReplyToMessage != nil
}

// Type determines the type of the message.
// Note that, for all these types, messages can still be replies or forwarded.
func (m *Message) Type() MessageType {
	if m.Text != nil {
		return TextType
	} else if m.Audio != nil {
		return AudioType
	} else if m.Document != nil {
		return DocumentType
	} else if m.Photo != nil {
		return PhotoType
	} else if m.Sticker != nil {
		return StickerType
	} else if m.Video != nil {
		return VideoType
	} else if m.Voice != nil {
		return VoiceType
	} else if m.Contact != nil {
		return ContactType
	} else if m.Location != nil {
		return LocationType
	} else if m.NewChatParticipant != nil {
		return NewChatParticipant
	} else if m.LeftChatParticipant != nil {
		return LeftChatParticipant
	} else if m.NewChatTitle != nil {
		return NewChatTitle
	} else if m.NewChatPhoto != nil {
		return NewChatPhoto
	} else if m.DeleteChatPhoto != nil {
		return DeletedChatPohoto
	} else if m.GroupChatCreated != nil {
		return GroupChatCreated
	}

	return Unknown
}

type noReplyMessage struct {
	Chat                Chat         `json:"chat"`                  // information about the chat
	ID                  int          `json:"message_id"`            // message id
	From                User         `json:"from"`                  // sender
	Date                int          `json:"date"`                  // timestamp
	ForwardFrom         *User        `json:"forward_from"`          // forwarded from who
	ForwardDate         *int         `json:"forward_date"`          // forwarded from when
	Text                *string      `json:"text"`                  // the actual text content
	Caption             *string      `json:"caption"`               // caption for photo or video messages
	Audio               *Audio       `json:"audio"`                 // information about audio contents
	Document            *Document    `json:"document"`              // information about file contents
	Photo               *[]PhotoSize `json:"photo"`                 // information about photo contents
	Sticker             *Sticker     `json:"sticker"`               // information about sticker contents
	Video               *Video       `json:"video"`                 // information about video contents
	Voice               *Voice       `json:"voice"`                 // information about voice message contents
	Contact             *Contact     `json:"contact"`               // information about contact contents
	Location            *Location    `json:"location"`              // information about location contents
	NewChatParticipant  *User        `json:"new_chat_participant"`  // information about a new chat participant
	LeftChatParticipant *User        `json:"left_chat_participant"` // information about a chat participant who left
	NewChatTitle        *string      `json:"new_chat_title"`        // information about changes in the group name
	NewChatPhoto        *[]PhotoSize `json:"new_chat_photo"`        // information about a new chat photo
	DeleteChatPhoto     *bool        `json:"delete_chat_photo"`     // information about a deleted chat photo
	GroupChatCreated    *bool        `json:"group_chat_created"`    // information about a created group chat
}

// MessageType is the type of a message
type MessageType int

// Message types
const (
	TextType     MessageType = iota // text messages
	AudioType                       // audio messages
	DocumentType                    // files
	PhotoType                       // photos
	StickerType                     // stickers
	VideoType                       // videos
	VoiceType                       // voice messages
	ContactType                     // contact information
	LocationType                    // locations

	chatActionsBegin
	NewChatParticipant  // joined chat participants
	LeftChatParticipant // left chat participants
	NewChatTitle        // chat title changes
	NewChatPhoto        // new chat photos
	DeletedChatPohoto   // deleted chat photos
	GroupChatCreated    // creation of a group chat
	chatActionsEnd

	Unknown // unknown (probably new due to API changes)
)

var types = map[MessageType]string{
	TextType:     "Text",
	AudioType:    "Audio",
	DocumentType: "Document",
	PhotoType:    "Photo",
	StickerType:  "Sticker",
	VideoType:    "Video",
	VoiceType:    "Voice",
	ContactType:  "Contact",
	LocationType: "Location",

	NewChatParticipant:  "NewChatParticipant",
	LeftChatParticipant: "LeftChatParticipant",
	NewChatTitle:        "NewChatTitle",
	NewChatPhoto:        "NewChatPhoto",
	DeletedChatPohoto:   "DeletedChatPhoto",
	GroupChatCreated:    "GroupChatCreated",

	Unknown: "UNKNOWN",
}

// IsChatAction checks if the MessageType is about changes in group chats
func (mt MessageType) IsChatAction() bool {
	return mt > chatActionsBegin && mt < chatActionsEnd
}

func (mt MessageType) String() string {
	val, ok := types[mt]
	if !ok {
		return types[Unknown]
	}
	return val
}

// PhotoSize represents one size of a photo or a thumbnail
type PhotoSize struct {
	FileBase
	Width  int `json:"width"`
	Height int `json:"height"`
}

// Sticker represents a sticker
type Sticker struct {
	FileBase
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Thumbnail PhotoSize `json:"thumb"`
}

// UpdateResponse represents the response sent by the API for a GetUpdates request
type updateResponse struct {
	baseResponse
	Update []Update `json:"result"`
}

// ByID is a wrapper to sort an []Update by ID
type byId []Update

func (a byId) Len() int           { return len(a) }
func (a byId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byId) Less(i, j int) bool { return a[i].ID < a[j].ID }

// Sort sorts all the updates contained in an UpdateResponse by their ID
func (resp *updateResponse) sort() {
	sort.Sort(byId(resp.Update))
}

// Update represents an incoming update
type Update struct {
	ID      int     `json:"update_id"`
	Message Message `json:"message"`
}

// UserResponse represents the response sent by the API on a GetMe request
type UserResponse struct {
	baseResponse
	User User `json:"result"`
}

// User represents a Telegram user or bot
type User struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
	Username  *string `json:"username"`
}

func (u User) String() string {
	if u.LastName != nil && u.Username != nil {
		return fmt.Sprintf("%d/%s %s (@%s)", u.ID, u.FirstName, u.LastName, u.Username)
	} else if u.LastName != nil {
		return fmt.Sprintf("%d/%s %s", u.ID, u.FirstName, u.LastName)
	} else if u.Username != nil {
		return fmt.Sprintf("%d/%s (@%s)", u.ID, u.FirstName, u.Username)
	}
	return fmt.Sprintf("%d/%s", u.ID, u.FirstName)
}

// UserProfilePhotosResponse represents the response sent by the API on a GetUserProfilePhotos request
type UserProfilePhotosResponse struct {
	baseResponse
	UserProfilePhotos UserProfilePhotos `json:"result"`
}

// UserProfilePhotos represents a users profile pictures
type UserProfilePhotos struct {
	TotalCount int         `json:"total_count"`
	Photos     []PhotoSize `json:"photos"`
}

// Video represents a video file
type Video struct {
	FileBase
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Duration  int       `json:"duration"`
	Thumbnail PhotoSize `json:"thumb"`
	MimeType  string    `json:"mime_type"`
	Caption   string    `json:"caption"`
}

// Voice represents a voice note
type Voice struct {
	FileBase
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
}

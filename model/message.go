package model

type MessageResponse struct {
	BaseResponse
	Message Message `json:"result"`
}

type Message struct {
	noReplyMessage
	ReplyToMessage noReplyMessage `json:"reply_to_message"`
}

type noReplyMessage struct {
	Id                  int         `json:"message_id"`
	From                User        `json:"from"`
	Date                int         `json:"date"`
	ChatUser            User        `json:"chat"`
	ChatGroup           GroupChat   `json:"chat"`
	ForwardFrom         User        `json:"forward_from"`
	ForwardDate         int         `json:"forward_date"`
	Text                string      `json:"text"`
	Audio               Audio       `json:"audio"`
	Document            Document    `json:"document"`
	Photo               []PhotoSize `json:"photo"`
	Sticker             Sticker     `json:"sticker"`
	Video               Video       `json:"video"`
	Contact             Contact     `json:"contact"`
	Location            Location    `json:"location"`
	NewChatParticipant  User        `json:"new_chat_participant"`
	LeftChatParticipant User        `json:"left_chat_participant"`
	NewChatTitle        string      `json:"new_chat_title"`
	NewChatPhoto        []PhotoSize `json:"new_chat_photo"`
	DeleteChatPhoto     bool        `json:"delete_chat_photo"`
	GroupChatCreated    bool        `json:"group_chat_created"`
}

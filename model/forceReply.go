package model

// ForceReply represents the values sent by a bot so that clients will be presented with a forced reply, see https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

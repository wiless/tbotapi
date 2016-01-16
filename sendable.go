package tbotapi

type sendable interface {
	Send() (*MessageResponse, error)
}

// Send sends the message.
// On success, the sent message is returned as a MessageResponse.
func (om *OutgoingMessage) Send() (*MessageResponse, error) {
	return om.api.send(om)
}

// Send sends the location.
// On success, the sent message is returned as a MessageResponse.
func (ol *OutgoingLocation) Send() (*MessageResponse, error) {
	return ol.api.send(ol)
}

// Send sends the video.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what bots can send, please check the API documentation.
// On success, the sent message is returned as a MessageResponse.
func (ov *OutgoingVideo) Send() (*MessageResponse, error) {
	return ov.api.send(ov)
}

// Send sends the photo.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what bots can send, please check the API documentation.
// On success, the sent message is returned as a MessageResponse.
func (op *OutgoingPhoto) Send() (*MessageResponse, error) {
	return op.api.send(op)
}

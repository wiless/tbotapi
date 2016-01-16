package tbotapi

type sendable interface {
	Send() (*MessageResponse, error)
}

func (om *OutgoingMessage) Send() (*MessageResponse, error) {
	return om.api.send(om)
}

func (ol *OutgoingLocation) Send() (*MessageResponse, error) {
	return ol.api.send(ol)
}

// Send sends a video.
// Use NewOutgoingVideo(Resend) to construct the video message and specify the file.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what bots can send, please check the API documentation.
// On success, the sent message is returned as a MessageResponse.
func (ov *OutgoingVideo) Send() (*MessageResponse, error) {
	return ov.api.send(ov)
}

// Send sends a photo.
// Use NewOutgoingPhoto(Resend) to construct the photo message and specify the file.
// Note that the Telegram servers may check the fileName for its extension.
// For current limitations on what bots can send, please check the API documentation.
// On success, the sent message is returned as a MessageResponse.
func (op *OutgoingPhoto) Send() (*MessageResponse, error) {
	return op.api.send(op)
}

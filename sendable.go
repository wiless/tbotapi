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

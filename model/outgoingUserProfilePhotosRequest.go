package model

import (
	"fmt"
)

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

package model

import (
	"fmt"
)

type OutgoingUserProfilePhotosRequest struct {
	UserID    int `json:"user_id"`
	Offset    int `json:"offset,omitempty"`
	Limit     int `json:"limit,omitempty"`
	offsetSet bool
	limitSet  bool
}

func NewOutgoingUserProfilePhotosRequest(userId int) *OutgoingUserProfilePhotosRequest {
	return &OutgoingUserProfilePhotosRequest{
		UserID: userId,
	}
}

func (op *OutgoingUserProfilePhotosRequest) SetOffset(to int) {
	op.Offset = to
	op.offsetSet = true
}

func (op *OutgoingUserProfilePhotosRequest) SetLimit(to int) {
	op.Limit = to
	op.limitSet = true
}

func (op *OutgoingUserProfilePhotosRequest) GetQueryString() Querystring {
	toReturn := map[string]string{}
	toReturn["user_id"] = fmt.Sprint(op.UserID)

	if op.offsetSet {
		toReturn["offset"] = fmt.Sprint(op.Offset)
	}

	if op.limitSet {
		toReturn["limit"] = fmt.Sprint(op.Limit)
	}

	return Querystring(toReturn)
}

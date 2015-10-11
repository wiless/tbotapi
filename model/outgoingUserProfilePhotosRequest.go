package model

import (
	"fmt"
)

type OutgoingUserProfilePhotosRequest struct {
	UserID int `json:"user_id"`
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

func NewOutgoingUserProfilePhotosRequest(userId int) *OutgoingUserProfilePhotosRequest {
	return &OutgoingUserProfilePhotosRequest{
		UserID: userId,
	}
}

func (op *OutgoingUserProfilePhotosRequest) SetOffset(to int) *OutgoingUserProfilePhotosRequest {
	op.Offset = to
	return op
}

func (op *OutgoingUserProfilePhotosRequest) SetLimit(to int) *OutgoingUserProfilePhotosRequest {
	op.Limit = to
	return op
}

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

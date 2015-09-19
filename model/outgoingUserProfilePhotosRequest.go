package model

import (
	"fmt"
	"net/url"
)

type OutgoingUserProfilePhotosRequestPub struct {
	UserId int `json:"user_id"`
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type OutgoingUserProfilePhotosRequest struct {
	userId    int
	offset    int
	limit     int
	offsetSet bool
	limitSet  bool
}

func NewOutgoingUserProfilePhotosRequest(userId int) *OutgoingUserProfilePhotosRequest {
	return &OutgoingUserProfilePhotosRequest{
		userId: userId,
	}
}

func (op *OutgoingUserProfilePhotosRequest) SetOffset(to int) {
	op.offset = to
	op.offsetSet = true
}

func (op *OutgoingUserProfilePhotosRequest) SetLimit(to int) {
	op.limit = to
	op.limitSet = true
}

func (op *OutgoingUserProfilePhotosRequest) GetQueryString() Querystring {
	toReturn := url.Values{}
	toReturn.Set("user_id", fmt.Sprint(op.userId))

	if op.offsetSet {
		toReturn.Set("offset", fmt.Sprint(op.offset))
	}

	if op.limitSet {
		toReturn.Set("limit", fmt.Sprint(op.limit))
	}

	return Querystring(toReturn)
}

func (op *OutgoingUserProfilePhotosRequest) GetPub() OutgoingUserProfilePhotosRequestPub {
	return OutgoingUserProfilePhotosRequestPub{
		UserId: op.userId,
		Offset: op.offset,
		Limit:  op.limit,
	}
}

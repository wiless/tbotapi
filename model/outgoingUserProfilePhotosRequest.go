package model

import (
	"fmt"
	"net/url"
)

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

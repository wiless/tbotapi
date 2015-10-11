package model

// UserProfilePhotosResponse represents the response sent by the API on a GetUserProfilePhotos request
type UserProfilePhotosResponse struct {
	BaseResponse
	UserProfilePhotos UserProfilePhotos `json:"result"`
}

// UserProfilePhotos represents a users profile pictures
type UserProfilePhotos struct {
	TotalCount int         `json:"total_count"`
	Photos     []PhotoSize `json:"photos"`
}

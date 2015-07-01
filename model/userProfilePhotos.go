package model

type UserProfilePhotosResponse struct {
	BaseResponse
	UserProfilePhotos UserProfilePhotos `json:"result"`
}

type UserProfilePhotos struct {
	TotalCount int         `json:"total_count"`
	Photos     []PhotoSize `json:"photos"`
}

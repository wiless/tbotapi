package model

// BaseResponse contains the basic fields contained in every API response
type BaseResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
	ErrorCode   int    `json:"error_code"`
}

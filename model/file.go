package model

// FileBase contains all the fields present in every file-like API response
type FileBase struct {
	ID   string `json:"file_id"`
	Size int    `json:"file_size"`
}

// File represents a file ready to be downloaded
type File struct {
	FileBase
	Path string `json:"file_path"`
}

// FileResponse represents the response sent by the API when requesting a file for download
type FileResponse struct {
	BaseResponse
	File File `json:"result"`
}

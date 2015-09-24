package model

type FileBase struct {
	Id   string `json:"file_id"`
	Size int    `json:"file_size"`
}

type File struct {
	FileBase
	Path string `json:"file_path"`
}

type FileResponse struct {
	BaseResponse
	File File `json:"result"`
}

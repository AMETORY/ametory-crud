package requests

// FileUploadRequest represents the structure for file upload input
type FileUploadRequest struct {
	File    []byte `form:"file" binding:"required"`
	Flipped string `form:"flipped"`
}

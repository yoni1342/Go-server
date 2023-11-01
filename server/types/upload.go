package types

type FileToUpload struct {
	Base64Data string `json:"base64"`
	Name       string `json:"name"`
}

type UploadInput struct {
	Input struct {
		Args []FileToUpload `json:"arg1"`
	} `json:"input"`
}

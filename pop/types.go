package pop

type VerifyPrompt struct {
	SystemPrompt string `json:"systemPrompt"`
}

type Response[T any] struct {
	// error code
	ErrCode int `json:"errCode"`
	// error Message
	ErrMsg string `json:"errMsg"`
	// return data
	Data T `json:"data,omitempty"`
}

package pop

type VerifyPrompt struct {
	SystemPrompt string `json:"systemPrompt"`
}

type Response[T any] struct {
	// 错误码
	ErrCode int `json:"errCode"`
	// 错误消息
	ErrMsg string `json:"errMsg"`
	// 返回数据
	Data T `json:"data,omitempty"`
}

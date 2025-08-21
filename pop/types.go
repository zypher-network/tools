package pop

type VerifyPrompt struct {
	SystemPrompt string `json:"systemPrompt"`
}

type Empty struct{}

type Response[T any] struct {
	// error code
	ErrCode int `json:"errCode"`
	// error Message
	ErrMsg string `json:"errMsg"`
	// return data
	Data T `json:"data,omitempty"`
}

type Proof struct {
	PiA      []string   `json:"pi_a"`
	PiB      [][]string `json:"pi_b"`
	PiC      []string   `json:"pi_c"`
	Protocol string     `json:"protocol"`
	Curve    string     `json:"curve"`
}

type VerifyResponse struct {
	Proof Proof  `json:"proof"`
	Hash  string `json:"hash"`
}

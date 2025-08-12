package pop

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Verify(ctx context.Context, systemPrompt []byte, apiKey string) (bool, error) {
	postUrl := fmt.Sprintf("https://pop-backend.zypher.game/verify/%s", apiKey)
	resp, err := http.Post(postUrl, "application/json", bytes.NewBuffer(systemPrompt))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var result Response[string]
	err = json.Unmarshal(body, &result)
	if err != nil {
		return false, err
	}
	return result.Data == "ok", nil
}

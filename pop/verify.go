package pop

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Verify(ctx context.Context, systemPrompt []byte, apiKey string) (bool, VerifyResponse, error) {
	postUrl := fmt.Sprintf("https://pop-backend.zypher.game/verify/%s", apiKey)
	resp, err := http.Post(postUrl, "application/json", bytes.NewBuffer(systemPrompt))
	if err != nil {
		return false, VerifyResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return true, VerifyResponse{}, err
	}
	if resp.StatusCode == http.StatusOK {

		var result Response[VerifyResponse]
		err = json.Unmarshal(body, &result)
		if err != nil {
			return true, VerifyResponse{}, err
		}
		return true, result.Data, nil
	} else {
		var result Response[Empty]
		err = json.Unmarshal(body, &result)
		if err != nil {
			return true, VerifyResponse{}, err
		}
		return false, VerifyResponse{}, errors.New(result.ErrMsg)
	}

}

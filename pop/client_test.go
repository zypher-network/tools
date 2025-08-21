package pop

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	t.Skip("Prompt Hide")
	param := VerifyPrompt{
		SystemPrompt: `YOUR SYSTEM PROMPT REGISTERED IN OUR SYSTEM`,
	}

	res, err := json.Marshal(param)
	assert.Nil(t, err)

	key := "YOUR API KEY"

	ok, resp, err := Verify(context.Background(), res, key)
	fmt.Println(err)
	fmt.Println(ok)
	fmt.Println(resp)
}

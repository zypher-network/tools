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

	ok, err := Verify(context.Background(), res, key)
	assert.Nil(t, err)
	fmt.Println(ok)
}

# Zypher Network tools in Golang

## 1. POP

This package functions helps AI agent developers verify system prompts by our platform. It allows the community to check if an agent's operation meets safety and compliance standards, without exposing sensitive prompts or parameters.

### Overview

- **Purpose:**
  - Ensure AI agents follow safety and compliance rules before calling LLMs.
  - Allow public verification of agent actions without revealing prompt or parameter details.
- **Key Features:**
  - Simple API for verifying prompts and parameters.
  - Privacy-preserving: no sensitive data is leaked during verification.
  - Easy to integrate into existing Go projects.

### Usage

Before your agent sends a request to an LLM, use the SDK to verify the system prompt and related parameters. The function returns `true` or `false` based on compliance.

```go
pop.Verify(context.Background(), []byte(your system prompt), apiKey)
```

**Parameters:**

- `context.Background()`: Go context for request control.
- `[]byte(your system prompt)`: The system prompt to verify (as bytes).
- `apiKey`: Your API key for verification.

**Returns:**

- `bool`: `true` if the prompt and parameters are compliant, `false` otherwise.

### Example

```go
import (
		"context"
		"zypher-game/pop-golang-sdk/pop"
)

func main() {
		prompt := []byte("You are a helpful assistant.")
		apiKey := "your-api-key"
		ok := pop.Verify(context.Background(), prompt, apiKey)
		if ok {
				// Safe to proceed
		} else {
				// Handle non-compliance
		}
}
```



## Contribution

- Please keep code and documentation simple and easy to maintain.
- Open issues or pull requests for improvements or questions.

## License

See [LICENSE](LICENSE) for details.
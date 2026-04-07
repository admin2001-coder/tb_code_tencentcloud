package lib

import (
	"context"
	"encoding/json"
	"log"

	"github.com/taubyte/go-sdk/event"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/tinfoilsh/tinfoil-go"
)

// Response structure
type ChatResponse struct {
	Model   string `json:"model"`
	Content string `json:"content"`
}

//export chat
func chat(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	// Create client
	client, err := tinfoil.NewClient(
		option.WithAPIKey("tk_XXXX"), // <-- replace with env later
	)
	if err != nil {
		log.Printf("client error: %v", err)
		return 1
	}

	// Call model
	resp, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Say this is a test"),
		},
		Model: "llama3-3-70b",
	})
	if err != nil {
		log.Printf("chat error: %v", err)
		return 1
	}

	// Build JSON response
	out := ChatResponse{
		Model:   "llama3-3-70b",
		Content: resp.Choices[0].Message.Content,
	}

	jsonBytes, err := json.Marshal(out)
	if err != nil {
		log.Printf("json error: %v", err)
		return 1
	}

	// Write HTTP response
	h.Headers().Set("Content-Type", "application/json")
	h.Write(jsonBytes)

	return 0
}
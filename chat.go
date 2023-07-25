package sense

import (
	"context"
	"errors"
	"net/http"
)

// Chat message role defined by the Sensa API.

type ModelName string

const (
	ChatMessageRoleUser                = "user"
	ChatMessageRoleAssistant           = "assistant"
	NovaPtcXsV1              ModelName = "nova-ptc-xs-v1"
	NovaPtcXlV1              ModelName = "nova-ptc-xl-v1"
)

var (
	ErrChatCompletionStreamNotSupported = errors.New("streaming is not supported with this method, please use CreateChatCompletionStream") //nolint:lll
)

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionRequest represents a request structure for chat completion API.
type ChatCompletionRequest struct {
	Model             string                  `json:"model"`
	KnowIDS           []string                `json:"know_ids"`
	MaxNewTokens      int                     `json:"max_new_tokens,omitempty"`
	Messages          []ChatCompletionMessage `json:"messages"`
	RepetitionPenalty float32                 `json:"repetition_penalty,omitempty"`
	Stream            bool                    `json:"stream,omitempty"`
	Temperature       *float32                `json:"temperature,omitempty"`
	TopP              *float32                `json:"top_p,omitempty"`
	User              string                  `json:"user,omitempty"`
}

type ChatCompletionChoice struct {
	Message      string `json:"message"`
	FinishReason string `json:"finish_reason"`
	Delta        string `json:"delta"`
}

// ChatCompletionResponse represents a response structure for chat completion API.
type ChatCompletionResponse struct {
	ID      string                 `json:"id"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   Usage                  `json:"usage"`
}

// CreateChatCompletion — API call to Create a completion for the chat message.
func (c *Client) CreateChatCompletion(
	ctx context.Context,
	request ChatCompletionRequest,
) (response ApiResponse[*ChatCompletionResponse], err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	urlSuffix := "/chat-completions"
	req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullURL(urlSuffix), request)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

package sense

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"
)

func TestClient_CreateChatCompletion(t *testing.T) {
	ak := ""
	sk := ""
	client, _ := NewClient(ak, sk)

	req := ChatCompletionRequest{
		Model: "nova-ptc-xs-v1",
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "  -d '{\"messages\": [{\"content\": \"如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度如何提高玻璃硬度\", \"role\": \"user\" }],\"model\": \"nova-ptc-xl-v1\"}' \n",
			},
		},
		Stream:      true,
		Temperature: nil,
		TopP:        nil,
	}
	r, err := client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
	for {
		fmt.Println(1)
		r, err := r.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println(1)
			}
			t.Error(err)
			break
		}
		t.Log(r.Data.Choices)
	}
}

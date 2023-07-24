package sense

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	utils "github.com/Lok-Lu/go-sense/internal"
)

var (
	ErrTooManyEmptyStreamMessages = errors.New("stream has sent too many empty messages")
)

type streamReader[T apiResponse] struct {
	emptyMessagesLimit uint
	isFinished         bool

	reader         *bufio.Reader
	response       *http.Response
	errAccumulator utils.ErrorAccumulator
	unmarshaler    utils.Unmarshaler
}

func (stream *streamReader[T]) Recv() (response ApiResponse[T], err error) {
	if stream.isFinished {
		err = io.EOF
		return
	}

	response, err = stream.processLines()
	return
}

func (stream *streamReader[T]) processLines() (ApiResponse[T], error) {
	var emptyMessagesCount uint

	for {
		rawLine, readErr := stream.reader.ReadBytes('\n')
		if readErr != nil {
			respErr := stream.unmarshalError()
			if respErr != nil {
				return *new(ApiResponse[T]), fmt.Errorf("error, %w", respErr.Error)
			}
			return *new(ApiResponse[T]), readErr
		}

		var headerData = []byte("data:")
		noSpaceLine := bytes.TrimSpace(rawLine)
		if !bytes.HasPrefix(noSpaceLine, headerData) {
			writeErr := stream.errAccumulator.Write(noSpaceLine)
			if writeErr != nil {
				return *new(ApiResponse[T]), writeErr
			}
			emptyMessagesCount++
			if emptyMessagesCount > stream.emptyMessagesLimit {
				return *new(ApiResponse[T]), ErrTooManyEmptyStreamMessages
			}

			continue
		}

		noPrefixLine := bytes.TrimPrefix(noSpaceLine, headerData)
		if string(noPrefixLine) == "[DONE]" {
			stream.isFinished = true
			return *new(ApiResponse[T]), io.EOF
		}

		var response ApiResponse[T]
		unmarshalErr := stream.unmarshaler.Unmarshal(noPrefixLine, &response)
		if unmarshalErr != nil {
			return *new(ApiResponse[T]), unmarshalErr
		}

		if response.IsNil() {
			var errRes ErrorResponse
			err := stream.unmarshaler.Unmarshal(noPrefixLine, &errRes)
			if err != nil {
				return *new(ApiResponse[T]), err
			}
			return *new(ApiResponse[T]), errRes.Error
		}

		return response, nil
	}
}

func (stream *streamReader[T]) unmarshalError() (errResp *ErrorResponse) {
	errBytes := stream.errAccumulator.Bytes()
	if len(errBytes) == 0 {
		return
	}

	err := stream.unmarshaler.Unmarshal(errBytes, &errResp)
	if err != nil {
		errResp = nil
	}

	return
}

func (stream *streamReader[T]) Close() {
	stream.response.Body.Close()
}

package sense

import (
	"context"
	"encoding/json"
	"fmt"
	utils "github.com/Lok-Lu/go-sense/internal"
	"io"
	"net/http"
)

type Client struct {
	config ClientConfig

	requestBuilder utils.RequestBuilder
}

func NewClient(ak, sk string) (*Client, error) {
	config, err := DefaultConfig(ak, sk)
	if err != nil {
		return nil, err
	}
	return NewClientWithConfig(config), nil
}

func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config:         config,
		requestBuilder: utils.NewRequestBuilder(),
	}
}

func (c *Client) sendRequest(req *http.Request, v any) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")

	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	c.setCommonHeaders(req)

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if isFailureStatusCode(res) {
		return c.handleErrorResp(res)
	}

	return decodeResponse(res.Body, v)
}

func (c *Client) setCommonHeaders(req *http.Request) {
	valid, _ := CheckJwtToken(c.config.authToken, c.config.secretKey)
	if !valid {
		token, _ := EncodeJwtToken(c.config.accessKey, c.config.secretKey)
		c.config.authToken = token
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.authToken))
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

func decodeResponse(body io.Reader, v any) error {
	if v == nil {
		return nil
	}

	if result, ok := v.(*string); ok {
		fmt.Println(11111)
		return decodeString(body, result)
	}
	return json.NewDecoder(body).Decode(v)
}

func decodeString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	*output = string(b)
	return nil
}

// fullURL returns full URL for request.
// args[0] is model name, if API type is Azure, model name is required to get deployment name.
func (c *Client) fullURL(suffix string) string {
	return fmt.Sprintf("%s%s", c.config.BaseURL, suffix)
}

func (c *Client) newStreamRequest(
	ctx context.Context,
	method string,
	urlSuffix string,
	body any) (*http.Request, error) {
	req, err := c.requestBuilder.Build(ctx, method, c.fullURL(urlSuffix), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	c.setCommonHeaders(req)
	return req, nil
}

func (c *Client) handleErrorResp(resp *http.Response) error {
	var errRes ErrorResponse

	err := json.NewDecoder(resp.Body).Decode(&errRes)
	if err != nil || errRes.Error == nil {
		reqErr := &RequestError{
			HTTPStatusCode: resp.StatusCode,
			Err:            err,
		}
		if errRes.Error != nil {
			reqErr.Err = errRes.Error
		}
		return reqErr
	}

	errRes.Error.HTTPStatusCode = resp.StatusCode
	return errRes.Error
}
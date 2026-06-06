// Package uploadpost provides a Go client for the Upload-Post API.
// It supports uploading videos, photos, text posts, and documents to
// TikTok, Instagram, YouTube, LinkedIn, Facebook, Pinterest, Threads,
// Reddit, Bluesky, and X (Twitter).
package uploadpost

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	defaultBaseURL = "https://api.upload-post.com/api"
	sdkSource      = "go"
)

// Client is the Upload-Post API client.
type Client struct {
	apiKey  string
	baseURL string
	http    *http.Client
}

// ClientOptions allows customizing the client.
type ClientOptions struct {
	// BaseURL overrides the default API base URL.
	BaseURL string
	// HTTPClient overrides the default HTTP client.
	HTTPClient *http.Client
}

// New creates a new Upload-Post API client with the given API key.
func New(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: defaultBaseURL,
		http:    &http.Client{},
	}
}

// NewWithOptions creates a new client with custom options.
func NewWithOptions(apiKey string, opts ClientOptions) *Client {
	c := New(apiKey)
	if opts.BaseURL != "" {
		c.baseURL = opts.BaseURL
	}
	if opts.HTTPClient != nil {
		c.http = opts.HTTPClient
	}
	return c
}

// APIError is returned when the API responds with an error HTTP status code.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("upload-post API error (HTTP %d): %s", e.StatusCode, e.Message)
}

// Bool returns a pointer to v. Useful for optional *bool fields.
func Bool(v bool) *bool { return &v }

// Int returns a pointer to v. Useful for optional *int fields.
func Int(v int) *int { return &v }

// Int64 returns a pointer to v. Useful for optional *int64 fields.
func Int64(v int64) *int64 { return &v }

// isURL returns true if s starts with http:// or https://.
func isURL(s string) bool {
	lower := strings.ToLower(s)
	return strings.HasPrefix(lower, "http://") || strings.HasPrefix(lower, "https://")
}

// extractErrorMessage pulls the human-readable error out of an API response body.
func extractErrorMessage(m map[string]json.RawMessage) string {
	if raw, ok := m["message"]; ok {
		var s string
		if json.Unmarshal(raw, &s) == nil && s != "" {
			return s
		}
	}
	if raw, ok := m["detail"]; ok {
		var s string
		if json.Unmarshal(raw, &s) == nil && s != "" {
			return s
		}
	}
	return "unknown error"
}

// doRaw executes an HTTP request and returns the raw response body.
func (c *Client) doRaw(ctx context.Context, method, endpoint string, body io.Reader, contentType string) ([]byte, int, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+endpoint, body)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Authorization", "Apikey "+c.apiKey)
	req.Header.Set("X-Upload-Post-Source", sdkSource)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	return data, resp.StatusCode, err
}

// doJSON executes a request and unmarshals the JSON response into dest.
func (c *Client) doJSON(ctx context.Context, method, endpoint string, body io.Reader, contentType string, dest interface{}) error {
	data, statusCode, err := c.doRaw(ctx, method, endpoint, body, contentType)
	if err != nil {
		return err
	}

	if statusCode >= 400 {
		var m map[string]json.RawMessage
		if json.Unmarshal(data, &m) == nil {
			return &APIError{StatusCode: statusCode, Message: extractErrorMessage(m)}
		}
		return &APIError{StatusCode: statusCode, Message: string(data)}
	}

	if dest != nil {
		return json.Unmarshal(data, dest)
	}
	return nil
}

// getJSON performs a GET request with query params and unmarshals into dest.
func (c *Client) getJSON(ctx context.Context, endpoint string, params map[string]string, dest interface{}) error {
	u, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return err
	}
	if len(params) > 0 {
		q := u.Query()
		for k, v := range params {
			if v != "" {
				q.Set(k, v)
			}
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Apikey "+c.apiKey)
	req.Header.Set("X-Upload-Post-Source", sdkSource)

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		var m map[string]json.RawMessage
		if json.Unmarshal(data, &m) == nil {
			return &APIError{StatusCode: resp.StatusCode, Message: extractErrorMessage(m)}
		}
		return &APIError{StatusCode: resp.StatusCode, Message: string(data)}
	}

	if dest != nil {
		return json.Unmarshal(data, dest)
	}
	return nil
}

// jsonBody serializes payload to JSON and returns a reader + content-type.
func jsonBody(payload interface{}) (io.Reader, string, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, "", err
	}
	return bytes.NewReader(data), "application/json", nil
}

// postJSON performs a POST with a JSON body and unmarshals the response into dest.
func (c *Client) postJSON(ctx context.Context, endpoint string, payload interface{}, dest interface{}) error {
	body, ct, err := jsonBody(payload)
	if err != nil {
		return err
	}
	return c.doJSON(ctx, http.MethodPost, endpoint, body, ct, dest)
}

// patchJSON performs a PATCH with a JSON body.
func (c *Client) patchJSON(ctx context.Context, endpoint string, payload interface{}, dest interface{}) error {
	body, ct, err := jsonBody(payload)
	if err != nil {
		return err
	}
	return c.doJSON(ctx, http.MethodPatch, endpoint, body, ct, dest)
}

// deleteJSON performs a DELETE with an optional JSON body.
func (c *Client) deleteJSON(ctx context.Context, endpoint string, payload interface{}, dest interface{}) error {
	var body io.Reader
	ct := ""
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		body = bytes.NewReader(data)
		ct = "application/json"
	}
	return c.doJSON(ctx, http.MethodDelete, endpoint, body, ct, dest)
}

// formBuilder builds a multipart/form-data request body.
type formBuilder struct {
	buf *bytes.Buffer
	w   *multipart.Writer
}

func newFormBuilder() *formBuilder {
	buf := &bytes.Buffer{}
	return &formBuilder{buf: buf, w: multipart.NewWriter(buf)}
}

func (f *formBuilder) set(key, value string) {
	_ = f.w.WriteField(key, value)
}

func (f *formBuilder) setIfNotEmpty(key, value string) {
	if value != "" {
		f.set(key, value)
	}
}

func (f *formBuilder) setBool(key string, v *bool) {
	if v != nil {
		f.set(key, strconv.FormatBool(*v))
	}
}

func (f *formBuilder) setInt(key string, v *int) {
	if v != nil {
		f.set(key, strconv.Itoa(*v))
	}
}

func (f *formBuilder) setInt64(key string, v *int64) {
	if v != nil {
		f.set(key, strconv.FormatInt(*v, 10))
	}
}

func (f *formBuilder) setArray(key string, values []string) {
	for _, v := range values {
		_ = f.w.WriteField(key, v)
	}
}

// setFile appends a file field: uploads the file if pathOrURL is a local path,
// or appends it as a plain string value if it is a URL.
func (f *formBuilder) setFile(key, pathOrURL string) error {
	if isURL(pathOrURL) {
		f.set(key, pathOrURL)
		return nil
	}
	file, err := os.Open(pathOrURL)
	if err != nil {
		return fmt.Errorf("failed to open %q: %w", pathOrURL, err)
	}
	defer file.Close()
	part, err := f.w.CreateFormFile(key, filepath.Base(pathOrURL))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	return err
}

// setFiles appends multiple file fields (e.g. photos[]).
func (f *formBuilder) setFiles(key string, pathsOrURLs []string) error {
	for _, p := range pathsOrURLs {
		if err := f.setFile(key, p); err != nil {
			return err
		}
	}
	return nil
}

func (f *formBuilder) close() { _ = f.w.Close() }

// postForm closes the builder and sends a multipart POST request.
func (c *Client) postForm(ctx context.Context, endpoint string, fb *formBuilder, dest interface{}) error {
	fb.close()
	return c.doJSON(ctx, http.MethodPost, endpoint, fb.buf, fb.w.FormDataContentType(), dest)
}

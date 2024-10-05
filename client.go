package kabustation

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	query "github.com/google/go-querystring/query"
)

type Client struct {
	APIPassword string
	BaseURL     string
	HTTPClient  *http.Client

	state state
}

type state struct {
	token string
}

func NewClient(
	apiPassword string,
	ops ...Option,
) *Client {
	c := &Client{
		APIPassword: apiPassword,
		BaseURL:     "https://api.kabus.gr.jp",
		HTTPClient:  http.DefaultClient,
	}
	for _, op := range ops {
		op(c)
	}
	return c
}

type Option func(*Client)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = httpClient
	}
}

func WithRateLimiter() Option {
	return func(c *Client) {
		c.HTTPClient.Transport = NewThrottledTransport(&ThrottledTransportConfig{
			Transport: http.DefaultTransport,
			Paths: map[string]RateLimitConfig{
				rateLimitPathOrder: {
					LimitPeriod:  1 * time.Second,
					RequestCount: 5,
				},
				rateLimitPathRegister: {
					LimitPeriod:  1 * time.Second,
					RequestCount: 10,
				},
				rateLimitPathInfo: {
					LimitPeriod:  1 * time.Second,
					RequestCount: 10,
				},
				rateLimitPathWallet: {
					LimitPeriod:  1 * time.Second,
					RequestCount: 10,
				},
			},
		})
	}
}

func (c *Client) newRequest(
	method string,
	path string,
	opts ...newRequestOption,
) (*http.Request, error) {
	req, err := http.NewRequest(method, c.BaseURL+path, nil)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(req)
	}

	return req, nil
}

type newRequestOption func(*http.Request)

func withQueryParams(d any) newRequestOption {
	return func(req *http.Request) {
		v, _ := query.Values(d)
		req.URL.RawQuery = v.Encode()
	}
}

func withBody(d any) newRequestOption {
	return func(req *http.Request) {
		b, _ := json.Marshal(d)
		req.Body = io.NopCloser(bytes.NewReader(b))
	}
}

func (c *Client) do(
	ctx context.Context,
	req *http.Request,
	v any,
) (*http.Response, error) {
	req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", c.state.token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if !isSuccess(res) {
		var e ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		}
		return res, &e
	}
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, res.Body)
		} else {
			if err := json.NewDecoder(res.Body).Decode(v); err != nil {
				return res, err
			}
		}
	}
	return res, nil
}

func isSuccess(res *http.Response) bool {
	return res.StatusCode/100 == 2
}

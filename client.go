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
	aPIPassword string
	baseURL     string
	httpClient  *http.Client

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
		aPIPassword: apiPassword,
		baseURL:     "http://localhost:8080",
		httpClient:  http.DefaultClient,
	}
	for _, op := range ops {
		op(c)
	}
	return c
}

type Option func(*Client)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

func WithRateLimiter() Option {
	return func(c *Client) {
		c.httpClient.Transport = NewThrottledTransport(&ThrottledTransportConfig{
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
	req, err := http.NewRequest(method, c.baseURL+path, nil)
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

	res, err := c.httpClient.Do(req)
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

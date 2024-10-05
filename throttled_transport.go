package kabustation

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// ThrottledTransport Rate Limited HTTP Client
type ThrottledTransport struct {
	roundTripperWrap http.RoundTripper
	ratelimiter      map[string]*rate.Limiter
}

func (c *ThrottledTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	limiter := c.fromContext(r.Context())
	if limiter != nil {
		if err := limiter.Wait(r.Context()); err != nil {
			return nil, err
		}
	}
	return c.roundTripperWrap.RoundTrip(r)
}

type ThrottledTransportConfig struct {
	Transport http.RoundTripper
	Paths     map[string]RateLimitConfig
}

type RateLimitConfig struct {
	LimitPeriod  time.Duration
	RequestCount int
}

// NewThrottledTransport wraps transportWrap with a rate limitter
// client.Transport = NewThrottledTransport(10*time.Seconds, 60, http.DefaultTransport) allows 60 requests every 10 seconds
func NewThrottledTransport(config *ThrottledTransportConfig) http.RoundTripper {
	limiters := make(map[string]*rate.Limiter)
	for k, v := range config.Paths {
		limiters[k] = rate.NewLimiter(rate.Every(v.LimitPeriod), v.RequestCount)
	}

	return &ThrottledTransport{
		roundTripperWrap: config.Transport,
		ratelimiter:      limiters,
	}
}

var ctxkey struct{}

const (
	rateLimitPathOrder    = "order"
	rateLimitPathRegister = "register"
	rateLimitPathInfo     = "info"
	rateLimitPathWallet   = "wallet"
)

func withRateLimiterContext(ctx context.Context, val string) context.Context {
	return context.WithValue(ctx, ctxkey, val)
}

func (c *ThrottledTransport) fromContext(ctx context.Context) *rate.Limiter {
	val, ok := ctx.Value(ctxkey).(string)
	if !ok {
		return nil
	}
	return c.ratelimiter[val]
}

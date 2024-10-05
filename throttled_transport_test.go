package kabustation

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestThrottledTransport_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	transport := &http.Transport{}
	config := &ThrottledTransportConfig{
		Transport: transport,
		Paths: map[string]RateLimitConfig{
			"test": {
				LimitPeriod:  time.Second,
				RequestCount: 1,
			},
		},
	}

	throttledTransport := NewThrottledTransport(config)

	client := &http.Client{
		Transport: throttledTransport,
	}

	req, _ := http.NewRequest("GET", server.URL, nil)
	ctx := withRateLimiterContext(req.Context(), "test")
	req = req.WithContext(ctx)

	_, err := client.Do(req)
	if err != nil {
		t.Errorf("ThrottledTransport was incorrect, got: %v", err)
	}
}

func TestThrottledTransport_WaitsAfterRateLimitExceeded(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	transport := &http.Transport{}
	config := &ThrottledTransportConfig{
		Transport: transport,
		Paths: map[string]RateLimitConfig{
			"test": {
				LimitPeriod:  time.Second,
				RequestCount: 1,
			},
		},
	}

	throttledTransport := NewThrottledTransport(config)

	client := &http.Client{
		Transport: throttledTransport,
	}

	req, _ := http.NewRequest("GET", server.URL, nil)
	ctx := withRateLimiterContext(req.Context(), "test")
	req = req.WithContext(ctx)

	start := time.Now()
	// First request should pass
	_, err := client.Do(req)
	if err != nil {
		t.Errorf("ThrottledTransport was incorrect, got: %v", err)
	}

	// Second request should be rate limited
	_, err = client.Do(req)
	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("ThrottledTransport was corrected, got: %v", err)
	}

	// Check if it waited for at least the limit period before returning
	if elapsed < config.Paths["test"].LimitPeriod {
		t.Errorf("ThrottledTransport did not wait after rate limit exceeded, waited: %v, expected: %v", elapsed, config.Paths["test"].LimitPeriod)
	}
}

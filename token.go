package kabustation

import (
	"context"
	"net/http"
)

type TokenResponse struct {
	ResultCode APIResult `json:"ResultCode"`
	Token      string    `json:"Token"`
}

func (c *Client) GetToken(
	ctx context.Context,
	isForced bool,
	isCashed bool,
) (*TokenResponse, error) {
	if c.state.token != "" && !isForced {
		return &TokenResponse{
			ResultCode: 0,
			Token:      c.state.token,
		}, nil
	}

	req, err := c.newRequest(http.MethodGet, "/token")
	if err != nil {
		return nil, err
	}

	var resp TokenResponse
	if _, err := c.do(ctx, req, &resp); err != nil {
		return nil, err
	}

	if resp.ResultCode.IsSuccess() && isCashed {
		c.state.token = resp.Token
	}

	return &resp, nil
}

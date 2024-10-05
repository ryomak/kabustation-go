package kabustation

import "fmt"

type ErrorResponse struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
}

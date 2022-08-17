package db

import "fmt"

type ErrHttp struct {
	StatusCode int
	Message    string `json:"error,omitempty"`
	Reason     string `json:"reason,omitempty"`
}

func (e ErrHttp) Error() string {
    return fmt.Sprintf("[%d]%s: %s", e.StatusCode, e.Message, e.Reason)
}

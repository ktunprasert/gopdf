package api

import (
	"errors"
	"fmt"

	"github.com/ktunprasert/gopdf/db"
	"github.com/ktunprasert/gopdf/domains"
)

func handleRepoError(err error) *domains.ErrorResponse {
	if err != nil {
		var errHttp *db.ErrHttp
		switch {
		case errors.As(err, &errHttp):
			return &domains.ErrorResponse{
				Success: false,
				Message: errHttp.Reason,
			}
		default:
			fmt.Println("unknown err", err)
		}
	}
	return nil
}

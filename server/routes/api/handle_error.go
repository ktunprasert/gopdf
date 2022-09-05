package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ktunprasert/gopdf/db"
	"github.com/ktunprasert/gopdf/domains"
)

func handleRepoError(err error, w http.ResponseWriter) *domains.ErrorResponse {
	if err != nil {
		fmt.Println("handleRepoError", err)
		var errHttp *db.ErrHttp
		switch {
		case errors.As(err, &errHttp):
            w.WriteHeader(errHttp.StatusCode)
			json.NewEncoder(w).Encode(&domains.ErrorResponse{
				Success:    false,
				Message:    errHttp.Reason,
			})
		default:
			fmt.Println("unknown err", err)
		}
	}
	return nil
}

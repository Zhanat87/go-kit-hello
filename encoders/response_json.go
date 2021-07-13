package encoders

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Zhanat87/go-kit-hello/contracts"
)

func EncodeResponseJSON(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(contracts.Errorer); ok && e.Error() != nil {
		EncodeErrorJSON(ctx, e.Error(), w)

		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

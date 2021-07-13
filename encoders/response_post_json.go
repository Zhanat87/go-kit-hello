package encoders

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Zhanat87/go-kit-hello/contracts"
)

func EncodeResponsePostJSON(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(contracts.Errorer); ok && e.Error() != nil {
		EncodeErrorJSON(ctx, e.Error(), w)

		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, _ := json.Marshal(response)
	var dat map[string]interface{}
	_ = json.Unmarshal(b, &dat)
	if _, ok := dat["location"]; ok {
		w.Header().Set("Location", dat["location"].(string))
	}
	w.WriteHeader(http.StatusCreated)

	return json.NewEncoder(w).Encode(response)
}

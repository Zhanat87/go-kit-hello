package encoders

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Zhanat87/go-kit-hello/errors"
)

func EncodeErrorJSON(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	case errors.OK:
		w.WriteHeader(http.StatusOK)
	case errors.Conflict, errors.CsvError, errors.CassandraReadError:
		w.WriteHeader(http.StatusConflict)
	case errors.NotFound:
		w.WriteHeader(http.StatusNotFound)
	case errors.AccessDenied:
		w.WriteHeader(http.StatusForbidden)
	case errors.ElasticConnectError, errors.S3ConnectError, errors.CassandraConnectError, errors.RabbitMQConnectError:
		w.WriteHeader(http.StatusServiceUnavailable)
	case errors.ContentNotFound:
		w.WriteHeader(http.StatusNoContent)
	case errors.DeserializeBug:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	case errors.InvalidCharacter:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(err)
}

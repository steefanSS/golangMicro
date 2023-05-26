package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/steefanSS/watermark-service/internal/util"
	endpoints "github.com/steefanSS/watermark-service/pkg/watermark/endpoint"
)

type endpoint endpoints.Set

func NewHTTPHandler(ep endpoint) http.Handler {
	m := http.NewServeMux()

	m.Handle("/healtz", httptransport.NewServer(
		ep.ServiceStatusEndpoint,
		decodeHTTPServiceStatusRequest,
		encodeResponse,
	))

	m.Handle("/status", httptransport.NewServer(
		ep.StatusEndpoint,
		decodeHTTPStatusRequest,
		encodeResponse,
	))

	m.Handle("/addDocument", httptransport.NewServer(
		ep.AddDocumentEndPoint,
		decodeHTTPAddDocumentRequest,
		encodeResponse,
	))

	m.Handle("/get", httptransport.NewServer(
		ep.GetEndpoint,
		decodeHTTPGetRequest,
		encodeResponse,
	))

	m.Handle("/watermark", httptransport.NewServer(
		ep.WatermarkEndpoint,
		decodeHTTPWatermarkRequest,
		encodeResponse,
	))

	return m
}

func decodeHTTPGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetRequest

	if r.ContentLength == 0 {
		logger.Log("Get request with no body")
		return req, nil
	}

	// Repeats too often
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeHTTPAddDocumentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.AddDocumentRequest

	//This will decode and store the value, but in case there is a problem, Decode returns "error" (func (*json.Decoder).Decode(v any) error)
	//in case someone wonders why only value is defined and instanced and err at that.
	err := json.NewDecoder(r.Body).Decode(&req)

	// Repeats too often
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeHTTPWatermarkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.WatermarkRequest

	// Repeats too often
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeHTTPStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.StatusRequest

	// Repeats too often
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeHTTPServiceStatusRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req endpoints.ServiceStatusRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, r interface{}) error {
	if e, ok := r.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(r)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json; character=utf-8")

	switch err {
	case util.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case util.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})

}

// Repeats too often
var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

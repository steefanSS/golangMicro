package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/steefanSS/watermark-service/pkg/watermark"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	AddDocumentEndPoint   endpoint.Endpoint
	StatusEndpoint        endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	WatermarkEndpoint     endpoint.Endpoint
}

func NewSetEndpoint(svc watermark.Service) Set {
	return Set{
		GetEndpoint:           MakeGetEndpoint(svc),
		AddDocumentEndPoint:   MakeAddDocumentEndpoint(svc),
		StatusEndpoint:        MakeStatusEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
		WatermarkEndpoint:     MakeWatermarkEndpoint(svc),
	}
}

// Examples of how the actual integration of enpoint struct is done(Request/Response) - factory functions

func MakeGetEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRequest)

		docs, err := svc.Get(ctx, req.Filters...)

		if err != nil {
			return GetResponse{docs, err.Error()}, nil
		}

		return GetResponse{docs, ""}, nil
	}
}

func MakeStatusEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(StatusRequest)

		status, err := svc.Status(ctx, req.TicketId)

		if err != nil {
			return StatusResponse{status, err.Error()}, nil
		}

		return StatusResponse{status, ""}, nil
	}
}

func MakeAddDocumentEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddDocumentRequest)

		tickedId, err := svc.AddDocument(ctx, req.Document)

		if err != nil {
			return AddDocumentResponse{tickedId, err.Error()}, nil
		}

		return AddDocumentResponse{tickedId, ""}, nil
	}
}

func MakeWatermarkEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WatermarkRequest)

		code, err := svc.Watermark(ctx, req.TickedId, req.Mark)

		if err != nil {
			return WatermarkResponse{code, err.Error()}, nil
		}

		return WatermarkResponse{code, ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		status, err := svc.ServiceStatus(ctx)

		if err != nil {
			return ServiceStatusResponse{status, err.Error()}, nil
		}

		return ServiceStatusResponse{status, ""}, nil
	}
}

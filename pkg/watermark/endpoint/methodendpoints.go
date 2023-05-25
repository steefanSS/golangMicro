package endpoints

import (
	"context"
	"errors"

	"github.com/steefanSS/watermark-service/internal"
)

type InterfaceSet Set

//Separated logic a bit more, this are the methods that are going to be used. The flow goes like this. Service defines concrete models, MakeEndpoint factory functions, wrap the types and form and point
//Finally the new type InterfaceSet of type Set is created and with it, concrete methods to be used to get the job done
//(didn't make sense to me to mix factory functions and methods in the same file, it's easily confusing for anyone who doesn't have experience with golang and in general structured programming lang like go)

func (s *InterfaceSet) Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest{Filters: filters})

	if err != nil {
		return []internal.Document{}, err
	}

	getResp := resp.(GetResponse)

	if getResp.Err != "" {
		return []internal.Document{}, errors.New(getResp.Err)
	}

	return getResp.Document, nil
}

func (s *InterfaceSet) Status(ctx context.Context, ticketID string) (internal.Status, error) {
	resp, err := s.StatusEndpoint(ctx, StatusRequest{TicketId: ticketID})

	if err != nil {
		return internal.Failed, err
	}

	stsResp := resp.(StatusResponse)

	if stsResp.Err != "" {
		return internal.Failed, errors.New(stsResp.Err)
	}

	return stsResp.Status, nil

}

func (s *InterfaceSet) AddDocument(ctx context.Context, doc *internal.Document) (string, error) {
	resp, err := s.AddDocumentEndPoint(ctx, AddDocumentRequest{Document: doc})

	if err != nil {
		return "", err
	}

	addDocResponse := resp.(AddDocumentResponse)

	if addDocResponse.Err != "" {
		return "", errors.New(addDocResponse.Err)
	}

	return addDocResponse.TicketId, nil
}

func (s *InterfaceSet) Watermark(ctx context.Context, ticketID, mark string) (int, error) {
	resp, err := s.WatermarkEndpoint(ctx, WatermarkRequest{TickedId: ticketID, Mark: mark})

	wmResp := resp.(WatermarkResponse)

	if err != nil {
		return wmResp.Code, err
	}

	if wmResp.Err != "" {
		return wmResp.Code, errors.New(wmResp.Err)
	}

	return wmResp.Code, nil
}

func (s *InterfaceSet) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})

	serviceStatus := resp.(ServiceStatusResponse)

	if err != nil {
		return serviceStatus.Code, err
	}

	if serviceStatus.Err != "" {
		return serviceStatus.Code, errors.New(serviceStatus.Err)
	}

	return serviceStatus.Code, nil
}

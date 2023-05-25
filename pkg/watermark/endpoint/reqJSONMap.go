package endpoints

//Package contains implentation of requests/response types that will be used in the actual implementation of endpoints

import "github.com/steefanSS/watermark-service/internal"

type GetRequest struct {
	Filters []internal.Filter `json:"filters,omitempty"`
}

type GetResponse struct {
	Document []internal.Document `json:"documents"`
	Err      string              `json:"err,omitempty"`
}

type StatusRequest struct {
	TicketId string `json:"ticketId"`
}

type StatusResponse struct {
	Status internal.Status `json:"status"`
	Err    string          `json:"err,omitempty"`
}

type WatermarkRequest struct {
	TickedId string `json:"ticketId"`
	Mark     string `json:"mark"`
}

type WatermarkResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err,omitempty"`
}

type AddDocumentRequest struct {
	Document *internal.Document `json:"document"`
}

type AddDocumentResponse struct {
	TicketId string `json:"ticketId"`
	Err      string `json:"err,omitempty"`
}

type ServiceStatusRequest struct{} //request only expects context

type ServiceStatusResponse struct {
	Code int    `json:"serviceStatus"`
	Err  string `json:"err,omitempty"`
}

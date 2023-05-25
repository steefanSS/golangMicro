package watermark

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/lithammer/shortuuid/v4"
	"github.com/steefanSS/watermark-service/internal"
)

//Service should always be represented as interface in Go Kit framework

type Service interface {
	Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error)
	Status(ctx context.Context, ticketID string) (internal.Status, error)
	Watermark(ctx context.Context, ticketID, mark string) (int, error)
	AddDocument(ctx context.Context, doc *internal.Document) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}

type watermarkService struct{}

// Just a constructor. When NewService constructor is used, it returns result from methods of empty new type watermarkService
func NewService() Service { return &watermarkService{} }

func (w *watermarkService) Get(_ context.Context, filter ...internal.Filter) ([]internal.Document, error) {

	doc := internal.Document{
		Content: "book",
		Title:   "Harry Potter and Half Blood Prince",
		Author:  "J.K. Rowling",
		Topic:   "Fiction and Magic",
	}

	return []internal.Document{doc}, nil
}

func (w *watermarkService) Status(_ context.Context, tickedID string) (internal.Status, error) {
	return internal.InProgress, nil
}

func (w *watermarkService) Watermark(_ context.Context, ticketID, mark string) (int, error) {
	return http.StatusOK, nil
}

func (w *watermarkService) AddDocument(_ context.Context, doc *internal.Document) (string, error) {
	newTicketId := shortuuid.New()

	return newTicketId, nil
}

func (w *watermarkService) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking the service status health")

	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestamp)
}

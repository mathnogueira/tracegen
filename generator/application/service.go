package application

import (
	"context"
	"fmt"

	"github.com/mathnogueira/tracegen/telemetry"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	Name       string
	Domain     string
	Operations []*Operation
	Tracer     trace.Tracer
}

// ChatGPT FTW
var domains = []string{
	"Account",
	"Advertisement",
	"Address",
	"Alert",
	"Answer",
	"Asset",
	"Attachment",
	"Article",
	"Badge",
	"Batch",
	"Category",
	"Cart",
	"Chat",
	"Code",
	"Comment",
	"Company",
	"Connection",
	"Container",
	"Conversation",
	"Customer",
	"Destination",
	"Detail",
	"Destination",
	"Document",
	"Employee",
	"Error",
	"Event",
	"Feedback",
	"File",
	"Form",
	"Frequency",
	"Group",
	"Inventory",
	"Invoice",
	"Link",
	"Location",
	"Log",
	"Match",
	"Menu",
	"Message",
	"Method",
	"Notification",
	"Note",
	"Order",
	"Partner",
	"Payment",
	"PaymentMethod",
	"Platform",
	"Playlist",
	"Profile",
	"Project",
	"Question",
	"Rating",
	"Record",
	"Region",
	"Report",
	"Reservation",
	"Result",
	"Review",
	"Role",
	"Rule",
	"Schedule",
	"Section",
	"Service",
	"Session",
	"Setting",
	"Score",
	"State",
	"Statistic",
	"Subscription",
	"Tag",
	"Task",
	"Template",
	"Ticket",
	"Token",
	"Transaction",
	"User",
	"Version",
	"Vote",
	"Widget",
}

var types = []string{
	"API",
	"Cache",
	"Logger",
	"Scheduler",
	"Index",
	"Worker",
}

type config struct {
	collectorEndpoint string
	insecure          bool
}

func NewService(options ...ServiceOption) *Service {
	config := config{}
	for _, option := range options {
		option(&config)
	}

	domain := random(domains)
	name := fmt.Sprintf("%s %s", domain, random(types))

	tracer, err := telemetry.GetTracer(context.Background(), telemetry.Config{
		ServiceName:       name,
		TracerName:        name,
		CollectorEndpoint: config.collectorEndpoint,
		Insecure:          config.insecure,
	})
	if err != nil {
		panic(err)
	}

	return &Service{
		Name:       name,
		Domain:     domain,
		Operations: make([]*Operation, 0),
		Tracer:     tracer,
	}
}

type ServiceOption func(*config)

func WithCollectorEndpoint(endpoint string) ServiceOption {
	return func(c *config) {
		c.collectorEndpoint = endpoint
	}
}

func WithInsecure(insecure bool) ServiceOption {
	return func(c *config) {
		c.insecure = insecure
	}
}

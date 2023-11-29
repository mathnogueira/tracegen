package telemetry

import (
	"errors"
)

type Config struct {
	ServiceName       string `json:"service_name"`
	TracerName        string `json:"tracer_name"`
	CollectorEndpoint string `json:"collector_endpoint"`
}

func (c Config) Validate() error {
	if c.CollectorEndpoint == "" {
		return errors.New("empty collector endpoint")
	}

	if c.TracerName == "" {
		return errors.New("empty tracer name")
	}

	if c.ServiceName == "" {
		return errors.New("empty service name")
	}

	return nil
}

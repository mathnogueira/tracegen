package generator

import "time"

type Config struct {
	NumberServices uint `json:"number_services"`
	NumberSpans    uint `json:"number_spans"`

	Collector CollectorConfig `json:"collector"`
	Delay     time.Duration   `json:"delay"`
}

type CollectorConfig struct {
	Endpoint string `json:"endpoint"`
}

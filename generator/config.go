package generator

type Config struct {
	NumberServices uint `json:"number_services"`
	NumberSpans    uint `json:"number_spans"`

	Collector CollectorConfig `json:"collector"`
}

type CollectorConfig struct {
	Endpoint string `json:"endpoint"`
}

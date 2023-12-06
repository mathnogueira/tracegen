package operation

import (
	"fmt"
	"math/rand"

	"go.opentelemetry.io/otel/attribute"
)

var namespaces = []string{
	"app",
	"service",
	"telemetry",
}

var keys = []string{
	"version",
	"type",
	"name",
	"library",
	"scope",
	"operation",
}

func GenerateRandomAttributes(n int) []attribute.KeyValue {
	attributes := make([]attribute.KeyValue, n)
	for i := 0; i < n; i++ {
		attributeName := fmt.Sprintf("%s.%s", random(namespaces), random(keys))
		attributes[i] = attribute.String(attributeName, randomString(16))
	}

	return attributes
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

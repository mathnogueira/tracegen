package application

import (
	"math/rand"
)

func random(input []string) string {
	return input[rand.Intn(len(input))]
}

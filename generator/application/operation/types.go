package operation

import (
	"fmt"
	"math/rand"
)

type Type string

const (
	TypeEntrypoint = "entrypoint"
	TypeStorage    = "storage"
	TypeProcessing = "processing"
	TypeMessaging  = "messaging"
)

var types = []Type{
	TypeEntrypoint,
	TypeStorage,
	TypeProcessing,
	TypeMessaging,
}

func GetRandomType() Type {
	return types[rand.Intn(len(types))]
}

func GetRandomName(t Type, entity string) string {
	var input []string
	if t == TypeEntrypoint {
		input = EntrypointOperations
	}

	if t == TypeStorage {
		input = StorageOperations
	}

	if t == TypeMessaging {
		input = MessagingOperations
	}

	if t == TypeProcessing {
		input = ProcessingOperations
	}

	randomName := random(input)
	return fmt.Sprintf(randomName, entity)
}

func random(options []string) string {
	return options[rand.Intn(len(options))]
}

var EntrypointOperations = []string{
	"Get %s",
	"Update %s",
	"Delete %s",
	"List %s",
	"Create %s",
}

var MessagingOperations = []string{
	"Get %s message",
	"Publish %s message",
}

var ProcessingOperations = []string{
	"Validate Request",
	"Authenticate",
	"Authorizate",
	"Transform %s into model",
	"Log",
	"Create metric",
}

var StorageOperations = []string{
	"Cache %s",
	"Insert %s into database",
	"Select %s from database",
	"Update %s in database",
	"Delete %s from database",
	"Insert %s into cache",
	"Select %s from cache",
	"Update %s in cache",
	"Delete %s from cache",
	"Create %s file",
	"Get %s file",
	"Update %s file",
	"Delete %s file",
}

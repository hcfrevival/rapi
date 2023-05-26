package controllers

import "go.mongodb.org/mongo-driver/mongo"

type Controller struct {
	Mongo          *mongo.Client
	DatabaseName   string
	CollectionName string
}

// ErrorWrapper is a json wrapper for strings used
// in Error Responses
type ErrorWrapper map[string]any

// CreateErrorResponse wraps an error in the
// ErrorWrapper decorator to create consistent
// error responses
func CreateErrorResponse(error string) ErrorWrapper {
	return ErrorWrapper{
		"message": error,
	}
}

// Package handlers Question API
//
// Documentation for Question API
//
//  Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
// swagger:meta
package handlers

import "github.com/NOOBDY/questions_server/data"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// A list of questions
// swagger:response questionsResponse
type questionsResponseWrapper struct {
	// All current questions
	// in: body
	Body []data.Question
}

// Data structure representing a single question
// swagger:response questionResponse
type questionResponseWrapper struct {
	// Newly created question
	// in: body
	Body data.Question
}

// swagger:parameters listQuestion
type productIDParamsWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

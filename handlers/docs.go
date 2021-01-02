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

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

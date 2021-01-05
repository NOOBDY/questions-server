// Code generated by go-swagger; DO NOT EDIT.

package questions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListQuestionsParams creates a new ListQuestionsParams object
// with the default values initialized.
func NewListQuestionsParams() *ListQuestionsParams {

	return &ListQuestionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListQuestionsParamsWithTimeout creates a new ListQuestionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListQuestionsParamsWithTimeout(timeout time.Duration) *ListQuestionsParams {

	return &ListQuestionsParams{

		timeout: timeout,
	}
}

// NewListQuestionsParamsWithContext creates a new ListQuestionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListQuestionsParamsWithContext(ctx context.Context) *ListQuestionsParams {

	return &ListQuestionsParams{

		Context: ctx,
	}
}

// NewListQuestionsParamsWithHTTPClient creates a new ListQuestionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListQuestionsParamsWithHTTPClient(client *http.Client) *ListQuestionsParams {

	return &ListQuestionsParams{
		HTTPClient: client,
	}
}

/*ListQuestionsParams contains all the parameters to send to the API endpoint
for the list questions operation typically these are written to a http.Request
*/
type ListQuestionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list questions params
func (o *ListQuestionsParams) WithTimeout(timeout time.Duration) *ListQuestionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list questions params
func (o *ListQuestionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list questions params
func (o *ListQuestionsParams) WithContext(ctx context.Context) *ListQuestionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list questions params
func (o *ListQuestionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list questions params
func (o *ListQuestionsParams) WithHTTPClient(client *http.Client) *ListQuestionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list questions params
func (o *ListQuestionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListQuestionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

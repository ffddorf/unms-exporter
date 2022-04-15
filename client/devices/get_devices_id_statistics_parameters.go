// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewGetDevicesIDStatisticsParams creates a new GetDevicesIDStatisticsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicesIDStatisticsParams() *GetDevicesIDStatisticsParams {
	return &GetDevicesIDStatisticsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesIDStatisticsParamsWithTimeout creates a new GetDevicesIDStatisticsParams object
// with the ability to set a timeout on a request.
func NewGetDevicesIDStatisticsParamsWithTimeout(timeout time.Duration) *GetDevicesIDStatisticsParams {
	return &GetDevicesIDStatisticsParams{
		timeout: timeout,
	}
}

// NewGetDevicesIDStatisticsParamsWithContext creates a new GetDevicesIDStatisticsParams object
// with the ability to set a context for a request.
func NewGetDevicesIDStatisticsParamsWithContext(ctx context.Context) *GetDevicesIDStatisticsParams {
	return &GetDevicesIDStatisticsParams{
		Context: ctx,
	}
}

// NewGetDevicesIDStatisticsParamsWithHTTPClient creates a new GetDevicesIDStatisticsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicesIDStatisticsParamsWithHTTPClient(client *http.Client) *GetDevicesIDStatisticsParams {
	return &GetDevicesIDStatisticsParams{
		HTTPClient: client,
	}
}

/* GetDevicesIDStatisticsParams contains all the parameters to send to the API endpoint
   for the get devices Id statistics operation.

   Typically these are written to a http.Request.
*/
type GetDevicesIDStatisticsParams struct {

	// ID.
	ID string

	/* Interval.

	   Interval
	*/
	Interval string

	// Period.
	Period string

	// Start.
	Start string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get devices Id statistics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDStatisticsParams) WithDefaults() *GetDevicesIDStatisticsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get devices Id statistics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicesIDStatisticsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) WithTimeout(timeout time.Duration) *GetDevicesIDStatisticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) WithContext(ctx context.Context) *GetDevicesIDStatisticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) WithHTTPClient(client *http.Client) *GetDevicesIDStatisticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) WithID(id string) *GetDevicesIDStatisticsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) SetID(id string) {
	o.ID = id
}

// WithInterval adds the interval to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) WithInterval(interval string) *GetDevicesIDStatisticsParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) SetInterval(interval string) {
	o.Interval = interval
}

// WithPeriod adds the period to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) WithPeriod(period string) *GetDevicesIDStatisticsParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) SetPeriod(period string) {
	o.Period = period
}

// WithStart adds the start to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) WithStart(start string) *GetDevicesIDStatisticsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get devices Id statistics params
func (o *GetDevicesIDStatisticsParams) SetStart(start string) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesIDStatisticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param interval
	qrInterval := o.Interval
	qInterval := qrInterval
	if qInterval != "" {

		if err := r.SetQueryParam("interval", qInterval); err != nil {
			return err
		}
	}

	// query param period
	qrPeriod := o.Period
	qPeriod := qrPeriod
	if qPeriod != "" {

		if err := r.SetQueryParam("period", qPeriod); err != nil {
			return err
		}
	}

	// query param start
	qrStart := o.Start
	qStart := qrStart
	if qStart != "" {

		if err := r.SetQueryParam("start", qStart); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
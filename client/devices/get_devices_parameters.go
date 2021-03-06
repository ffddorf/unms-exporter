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
	"github.com/go-openapi/swag"
)

// NewGetDevicesParams creates a new GetDevicesParams object
// with the default values initialized.
func NewGetDevicesParams() *GetDevicesParams {
	var ()
	return &GetDevicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicesParamsWithTimeout creates a new GetDevicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDevicesParamsWithTimeout(timeout time.Duration) *GetDevicesParams {
	var ()
	return &GetDevicesParams{

		timeout: timeout,
	}
}

// NewGetDevicesParamsWithContext creates a new GetDevicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDevicesParamsWithContext(ctx context.Context) *GetDevicesParams {
	var ()
	return &GetDevicesParams{

		Context: ctx,
	}
}

// NewGetDevicesParamsWithHTTPClient creates a new GetDevicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDevicesParamsWithHTTPClient(client *http.Client) *GetDevicesParams {
	var ()
	return &GetDevicesParams{
		HTTPClient: client,
	}
}

/*GetDevicesParams contains all the parameters to send to the API endpoint
for the get devices operation typically these are written to a http.Request
*/
type GetDevicesParams struct {

	/*Authorized*/
	Authorized *bool
	/*Role*/
	Role []string
	/*SiteID*/
	SiteID *string
	/*Type*/
	Type []string
	/*WithInterfaces*/
	WithInterfaces *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get devices params
func (o *GetDevicesParams) WithTimeout(timeout time.Duration) *GetDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get devices params
func (o *GetDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get devices params
func (o *GetDevicesParams) WithContext(ctx context.Context) *GetDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get devices params
func (o *GetDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get devices params
func (o *GetDevicesParams) WithHTTPClient(client *http.Client) *GetDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get devices params
func (o *GetDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorized adds the authorized to the get devices params
func (o *GetDevicesParams) WithAuthorized(authorized *bool) *GetDevicesParams {
	o.SetAuthorized(authorized)
	return o
}

// SetAuthorized adds the authorized to the get devices params
func (o *GetDevicesParams) SetAuthorized(authorized *bool) {
	o.Authorized = authorized
}

// WithRole adds the role to the get devices params
func (o *GetDevicesParams) WithRole(role []string) *GetDevicesParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the get devices params
func (o *GetDevicesParams) SetRole(role []string) {
	o.Role = role
}

// WithSiteID adds the siteID to the get devices params
func (o *GetDevicesParams) WithSiteID(siteID *string) *GetDevicesParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get devices params
func (o *GetDevicesParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithType adds the typeVar to the get devices params
func (o *GetDevicesParams) WithType(typeVar []string) *GetDevicesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get devices params
func (o *GetDevicesParams) SetType(typeVar []string) {
	o.Type = typeVar
}

// WithWithInterfaces adds the withInterfaces to the get devices params
func (o *GetDevicesParams) WithWithInterfaces(withInterfaces *bool) *GetDevicesParams {
	o.SetWithInterfaces(withInterfaces)
	return o
}

// SetWithInterfaces adds the withInterfaces to the get devices params
func (o *GetDevicesParams) SetWithInterfaces(withInterfaces *bool) {
	o.WithInterfaces = withInterfaces
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorized != nil {

		// query param authorized
		var qrAuthorized bool
		if o.Authorized != nil {
			qrAuthorized = *o.Authorized
		}
		qAuthorized := swag.FormatBool(qrAuthorized)
		if qAuthorized != "" {
			if err := r.SetQueryParam("authorized", qAuthorized); err != nil {
				return err
			}
		}

	}

	valuesRole := o.Role

	joinedRole := swag.JoinByFormat(valuesRole, "multi")
	// query array param role
	if err := r.SetQueryParam("role", joinedRole...); err != nil {
		return err
	}

	if o.SiteID != nil {

		// query param siteId
		var qrSiteID string
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {
			if err := r.SetQueryParam("siteId", qSiteID); err != nil {
				return err
			}
		}

	}

	valuesType := o.Type

	joinedType := swag.JoinByFormat(valuesType, "multi")
	// query array param type
	if err := r.SetQueryParam("type", joinedType...); err != nil {
		return err
	}

	if o.WithInterfaces != nil {

		// query param withInterfaces
		var qrWithInterfaces bool
		if o.WithInterfaces != nil {
			qrWithInterfaces = *o.WithInterfaces
		}
		qWithInterfaces := swag.FormatBool(qrWithInterfaces)
		if qWithInterfaces != "" {
			if err := r.SetQueryParam("withInterfaces", qWithInterfaces); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

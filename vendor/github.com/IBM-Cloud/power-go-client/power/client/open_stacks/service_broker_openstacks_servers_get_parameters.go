// Code generated by go-swagger; DO NOT EDIT.

package open_stacks

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

// NewServiceBrokerOpenstacksServersGetParams creates a new ServiceBrokerOpenstacksServersGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceBrokerOpenstacksServersGetParams() *ServiceBrokerOpenstacksServersGetParams {
	return &ServiceBrokerOpenstacksServersGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBrokerOpenstacksServersGetParamsWithTimeout creates a new ServiceBrokerOpenstacksServersGetParams object
// with the ability to set a timeout on a request.
func NewServiceBrokerOpenstacksServersGetParamsWithTimeout(timeout time.Duration) *ServiceBrokerOpenstacksServersGetParams {
	return &ServiceBrokerOpenstacksServersGetParams{
		timeout: timeout,
	}
}

// NewServiceBrokerOpenstacksServersGetParamsWithContext creates a new ServiceBrokerOpenstacksServersGetParams object
// with the ability to set a context for a request.
func NewServiceBrokerOpenstacksServersGetParamsWithContext(ctx context.Context) *ServiceBrokerOpenstacksServersGetParams {
	return &ServiceBrokerOpenstacksServersGetParams{
		Context: ctx,
	}
}

// NewServiceBrokerOpenstacksServersGetParamsWithHTTPClient creates a new ServiceBrokerOpenstacksServersGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceBrokerOpenstacksServersGetParamsWithHTTPClient(client *http.Client) *ServiceBrokerOpenstacksServersGetParams {
	return &ServiceBrokerOpenstacksServersGetParams{
		HTTPClient: client,
	}
}

/* ServiceBrokerOpenstacksServersGetParams contains all the parameters to send to the API endpoint
   for the service broker openstacks servers get operation.

   Typically these are written to a http.Request.
*/
type ServiceBrokerOpenstacksServersGetParams struct {

	/* OpenstackID.

	   Openstack ID
	*/
	OpenstackID string

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service broker openstacks servers get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceBrokerOpenstacksServersGetParams) WithDefaults() *ServiceBrokerOpenstacksServersGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service broker openstacks servers get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceBrokerOpenstacksServersGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service broker openstacks servers get params
func (o *ServiceBrokerOpenstacksServersGetParams) WithTimeout(timeout time.Duration) *ServiceBrokerOpenstacksServersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service broker openstacks servers get params
func (o *ServiceBrokerOpenstacksServersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service broker openstacks servers get params
func (o *ServiceBrokerOpenstacksServersGetParams) WithContext(ctx context.Context) *ServiceBrokerOpenstacksServersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service broker openstacks servers get params
func (o *ServiceBrokerOpenstacksServersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service broker openstacks servers get params
func (o *ServiceBrokerOpenstacksServersGetParams) WithHTTPClient(client *http.Client) *ServiceBrokerOpenstacksServersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service broker openstacks servers get params
func (o *ServiceBrokerOpenstacksServersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOpenstackID adds the openstackID to the service broker openstacks servers get params
func (o *ServiceBrokerOpenstacksServersGetParams) WithOpenstackID(openstackID string) *ServiceBrokerOpenstacksServersGetParams {
	o.SetOpenstackID(openstackID)
	return o
}

// SetOpenstackID adds the openstackId to the service broker openstacks servers get params
func (o *ServiceBrokerOpenstacksServersGetParams) SetOpenstackID(openstackID string) {
	o.OpenstackID = openstackID
}

// WithPvmInstanceID adds the pvmInstanceID to the service broker openstacks servers get params
func (o *ServiceBrokerOpenstacksServersGetParams) WithPvmInstanceID(pvmInstanceID string) *ServiceBrokerOpenstacksServersGetParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the service broker openstacks servers get params
func (o *ServiceBrokerOpenstacksServersGetParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBrokerOpenstacksServersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param openstack_id
	if err := r.SetPathParam("openstack_id", o.OpenstackID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
package scopes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteScopeParams creates a new DeleteScopeParams object
// with the default values initialized.
func NewDeleteScopeParams() *DeleteScopeParams {
	var ()
	return &DeleteScopeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScopeParamsWithTimeout creates a new DeleteScopeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteScopeParamsWithTimeout(timeout time.Duration) *DeleteScopeParams {
	var ()
	return &DeleteScopeParams{

		timeout: timeout,
	}
}

// NewDeleteScopeParamsWithContext creates a new DeleteScopeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteScopeParamsWithContext(ctx context.Context) *DeleteScopeParams {
	var ()
	return &DeleteScopeParams{

		Context: ctx,
	}
}

// NewDeleteScopeParamsWithHTTPClient creates a new DeleteScopeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteScopeParamsWithHTTPClient(client *http.Client) *DeleteScopeParams {
	var ()
	return &DeleteScopeParams{
		HTTPClient: client,
	}
}

/*DeleteScopeParams contains all the parameters to send to the API endpoint
for the delete scope operation typically these are written to a http.Request
*/
type DeleteScopeParams struct {

	/*OpID*/
	OpID *string
	/*IDName*/
	IDName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete scope params
func (o *DeleteScopeParams) WithTimeout(timeout time.Duration) *DeleteScopeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete scope params
func (o *DeleteScopeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete scope params
func (o *DeleteScopeParams) WithContext(ctx context.Context) *DeleteScopeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete scope params
func (o *DeleteScopeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete scope params
func (o *DeleteScopeParams) WithHTTPClient(client *http.Client) *DeleteScopeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete scope params
func (o *DeleteScopeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOpID adds the opID to the delete scope params
func (o *DeleteScopeParams) WithOpID(opID *string) *DeleteScopeParams {
	o.SetOpID(opID)
	return o
}

// SetOpID adds the opId to the delete scope params
func (o *DeleteScopeParams) SetOpID(opID *string) {
	o.OpID = opID
}

// WithIDName adds the iDName to the delete scope params
func (o *DeleteScopeParams) WithIDName(iDName string) *DeleteScopeParams {
	o.SetIDName(iDName)
	return o
}

// SetIDName adds the idName to the delete scope params
func (o *DeleteScopeParams) SetIDName(iDName string) {
	o.IDName = iDName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScopeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.OpID != nil {

		// header param Op-ID
		if err := r.SetHeaderParam("Op-ID", *o.OpID); err != nil {
			return err
		}

	}

	// path param idName
	if err := r.SetPathParam("idName", o.IDName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/sendinblue/APIv3-go-library/models"
)

// NewCreateFolderParams creates a new CreateFolderParams object
// with the default values initialized.
func NewCreateFolderParams() *CreateFolderParams {
	var ()
	return &CreateFolderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFolderParamsWithTimeout creates a new CreateFolderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateFolderParamsWithTimeout(timeout time.Duration) *CreateFolderParams {
	var ()
	return &CreateFolderParams{

		timeout: timeout,
	}
}

// NewCreateFolderParamsWithContext creates a new CreateFolderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateFolderParamsWithContext(ctx context.Context) *CreateFolderParams {
	var ()
	return &CreateFolderParams{

		Context: ctx,
	}
}

// NewCreateFolderParamsWithHTTPClient creates a new CreateFolderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateFolderParamsWithHTTPClient(client *http.Client) *CreateFolderParams {
	var ()
	return &CreateFolderParams{
		HTTPClient: client,
	}
}

/*CreateFolderParams contains all the parameters to send to the API endpoint
for the create folder operation typically these are written to a http.Request
*/
type CreateFolderParams struct {

	/*CreateFolder
	  Name of the folder

	*/
	CreateFolder *models.CreateUpdateFolder

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create folder params
func (o *CreateFolderParams) WithTimeout(timeout time.Duration) *CreateFolderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create folder params
func (o *CreateFolderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create folder params
func (o *CreateFolderParams) WithContext(ctx context.Context) *CreateFolderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create folder params
func (o *CreateFolderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create folder params
func (o *CreateFolderParams) WithHTTPClient(client *http.Client) *CreateFolderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create folder params
func (o *CreateFolderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateFolder adds the createFolder to the create folder params
func (o *CreateFolderParams) WithCreateFolder(createFolder *models.CreateUpdateFolder) *CreateFolderParams {
	o.SetCreateFolder(createFolder)
	return o
}

// SetCreateFolder adds the createFolder to the create folder params
func (o *CreateFolderParams) SetCreateFolder(createFolder *models.CreateUpdateFolder) {
	o.CreateFolder = createFolder
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFolderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateFolder == nil {
		o.CreateFolder = new(models.CreateUpdateFolder)
	}

	if err := r.SetBodyParam(o.CreateFolder); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

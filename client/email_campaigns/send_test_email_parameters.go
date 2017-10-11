// Code generated by go-swagger; DO NOT EDIT.

package email_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/models"
)

// NewSendTestEmailParams creates a new SendTestEmailParams object
// with the default values initialized.
func NewSendTestEmailParams() *SendTestEmailParams {
	var ()
	return &SendTestEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendTestEmailParamsWithTimeout creates a new SendTestEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendTestEmailParamsWithTimeout(timeout time.Duration) *SendTestEmailParams {
	var ()
	return &SendTestEmailParams{

		timeout: timeout,
	}
}

// NewSendTestEmailParamsWithContext creates a new SendTestEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendTestEmailParamsWithContext(ctx context.Context) *SendTestEmailParams {
	var ()
	return &SendTestEmailParams{

		Context: ctx,
	}
}

// NewSendTestEmailParamsWithHTTPClient creates a new SendTestEmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendTestEmailParamsWithHTTPClient(client *http.Client) *SendTestEmailParams {
	var ()
	return &SendTestEmailParams{
		HTTPClient: client,
	}
}

/*SendTestEmailParams contains all the parameters to send to the API endpoint
for the send test email operation typically these are written to a http.Request
*/
type SendTestEmailParams struct {

	/*CampaignID
	  Id of the campaign

	*/
	CampaignID int64
	/*EmailTo*/
	EmailTo *models.SendTestEmail

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send test email params
func (o *SendTestEmailParams) WithTimeout(timeout time.Duration) *SendTestEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send test email params
func (o *SendTestEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send test email params
func (o *SendTestEmailParams) WithContext(ctx context.Context) *SendTestEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send test email params
func (o *SendTestEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send test email params
func (o *SendTestEmailParams) WithHTTPClient(client *http.Client) *SendTestEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send test email params
func (o *SendTestEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the send test email params
func (o *SendTestEmailParams) WithCampaignID(campaignID int64) *SendTestEmailParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the send test email params
func (o *SendTestEmailParams) SetCampaignID(campaignID int64) {
	o.CampaignID = campaignID
}

// WithEmailTo adds the emailTo to the send test email params
func (o *SendTestEmailParams) WithEmailTo(emailTo *models.SendTestEmail) *SendTestEmailParams {
	o.SetEmailTo(emailTo)
	return o
}

// SetEmailTo adds the emailTo to the send test email params
func (o *SendTestEmailParams) SetEmailTo(emailTo *models.SendTestEmail) {
	o.EmailTo = emailTo
}

// WriteToRequest writes these params to a swagger request
func (o *SendTestEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignId
	if err := r.SetPathParam("campaignId", swag.FormatInt64(o.CampaignID)); err != nil {
		return err
	}

	if o.EmailTo == nil {
		o.EmailTo = new(models.SendTestEmail)
	}

	if err := r.SetBodyParam(o.EmailTo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

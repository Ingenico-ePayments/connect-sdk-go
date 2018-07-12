// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package merchant

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/internal/apiresource"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/captures"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/hostedcheckouts"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/hostedmandatemanagements"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/mandates"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/payments"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/payouts"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/productgroups"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/products"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/refunds"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/riskassessments"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/services"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/sessions"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/tokens"
)

// Client represents a merchant client. Thread-safe.
type Client struct {
	apiResource *apiresource.APIResource
}

// Hostedcheckouts represents the resource /{merchantId}/hostedcheckouts
// Create new hosted checkout
func (c *Client) Hostedcheckouts() *hostedcheckouts.Client {
	return hostedcheckouts.NewClient(c.apiResource, nil)
}

// Hostedmandatemanagements represents the resource /{merchantId}/hostedmandatemanagements
// Create new hosted mandate management
func (c *Client) Hostedmandatemanagements() *hostedmandatemanagements.Client {
	return hostedmandatemanagements.NewClient(c.apiResource, nil)
}

// Payments represents the resource /{merchantId}/payments
// Create, cancel and approve payments
func (c *Client) Payments() *payments.Client {
	return payments.NewClient(c.apiResource, nil)
}

// Captures represents the resource /{merchantId}/captures
// Get capture
func (c *Client) Captures() *captures.Client {
	return captures.NewClient(c.apiResource, nil)
}

// Refunds represents the resource /{merchantId}/refunds
// Create, cancel and approve refunds
func (c *Client) Refunds() *refunds.Client {
	return refunds.NewClient(c.apiResource, nil)
}

// Payouts represents the resource /{merchantId}/payouts
// Create, cancel and approve payouts
func (c *Client) Payouts() *payouts.Client {
	return payouts.NewClient(c.apiResource, nil)
}

// Productgroups represents the resource /{merchantId}/productgroups
// Get information about payment product groups
func (c *Client) Productgroups() *productgroups.Client {
	return productgroups.NewClient(c.apiResource, nil)
}

// Products represents the resource /{merchantId}/products
// Get information about payment products
func (c *Client) Products() *products.Client {
	return products.NewClient(c.apiResource, nil)
}

// Riskassessments represents the resource /{merchantId}/riskassessments
// Perform risk assessments on your customer data
func (c *Client) Riskassessments() *riskassessments.Client {
	return riskassessments.NewClient(c.apiResource, nil)
}

// Services represents the resource /{merchantId}/services
// Several services to help you
func (c *Client) Services() *services.Client {
	return services.NewClient(c.apiResource, nil)
}

// Tokens represents the resource /{merchantId}/tokens
// Create, delete and update tokens
func (c *Client) Tokens() *tokens.Client {
	return tokens.NewClient(c.apiResource, nil)
}

// Mandates represents the resource /{merchantId}/mandates
// Create, get and update mandates
func (c *Client) Mandates() *mandates.Client {
	return mandates.NewClient(c.apiResource, nil)
}

// Sessions represents the resource /{merchantId}/sessions
// Create new Session for Client2Server API calls
func (c *Client) Sessions() *sessions.Client {
	return sessions.NewClient(c.apiResource, nil)
}

// NewClient constructs a Merchant Client
//
// parent is the *apiresource.APIResource on top of which we want to build the new Merchant Client
func NewClient(parent *apiresource.APIResource, pathContext map[string]string) *Client {
	apiResource := apiresource.NewAPIResourceWithParent(parent, pathContext)

	return &Client{apiResource}
}

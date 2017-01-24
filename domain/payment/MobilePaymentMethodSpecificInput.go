// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// MobilePaymentMethodSpecificInput represents class MobilePaymentMethodSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_MobilePaymentMethodSpecificInput
type MobilePaymentMethodSpecificInput struct {
	AuthorizationMode              *string                               `json:"authorizationMode,omitempty"`
	DecryptedPaymentData           *DecryptedPaymentData                 `json:"decryptedPaymentData,omitempty"`
	EncryptedPaymentData           *string                               `json:"encryptedPaymentData,omitempty"`
	PaymentProduct320SpecificInput *MobilePaymentProduct320SpecificInput `json:"paymentProduct320SpecificInput,omitempty"`
	PaymentProductID               *int32                                `json:"paymentProductId,omitempty"`
	RequiresApproval               *bool                                 `json:"requiresApproval,omitempty"`
	SkipFraudService               *bool                                 `json:"skipFraudService,omitempty"`
	TransactionID                  *string                               `json:"transactionId,omitempty"`
}

// NewMobilePaymentMethodSpecificInput constructs a new MobilePaymentMethodSpecificInput
func NewMobilePaymentMethodSpecificInput() *MobilePaymentMethodSpecificInput {
	return &MobilePaymentMethodSpecificInput{}
}

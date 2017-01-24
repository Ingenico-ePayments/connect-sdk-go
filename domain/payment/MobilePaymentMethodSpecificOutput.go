// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// MobilePaymentMethodSpecificOutput represents class MobilePaymentMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_MobilePaymentMethodSpecificOutput
type MobilePaymentMethodSpecificOutput struct {
	AuthorisationCode   *string                       `json:"authorisationCode,omitempty"`
	FraudResults        *definitions.CardFraudResults `json:"fraudResults,omitempty"`
	Network             *string                       `json:"network,omitempty"`
	PaymentData         *MobilePaymentData            `json:"paymentData,omitempty"`
	PaymentProductID    *int32                        `json:"paymentProductId,omitempty"`
	ThreeDSecureResults *ThreeDSecureResults          `json:"threeDSecureResults,omitempty"`
}

// NewMobilePaymentMethodSpecificOutput constructs a new MobilePaymentMethodSpecificOutput
func NewMobilePaymentMethodSpecificOutput() *MobilePaymentMethodSpecificOutput {
	return &MobilePaymentMethodSpecificOutput{}
}

// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package dispute

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/file"
)

// Output represents class DisputeOutput
type Output struct {
	AmountOfMoney   *definitions.AmountOfMoney `json:"amountOfMoney,omitempty"`
	ContactPerson   *string                    `json:"contactPerson,omitempty"`
	CreationDetails *CreationDetail            `json:"creationDetails,omitempty"`
	EmailAddress    *string                    `json:"emailAddress,omitempty"`
	Files           *[]file.HostedFile         `json:"files,omitempty"`
	Reference       *Reference                 `json:"reference,omitempty"`
	ReplyTo         *string                    `json:"replyTo,omitempty"`
	RequestMessage  *string                    `json:"requestMessage,omitempty"`
	ResponseMessage *string                    `json:"responseMessage,omitempty"`
}

// NewOutput constructs a new Output
func NewOutput() *Output {
	return &Output{}
}

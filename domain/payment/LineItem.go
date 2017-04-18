// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// LineItem represents class LineItem
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_LineItem
type LineItem struct {
	AmountOfMoney                *definitions.AmountOfMoney            `json:"amountOfMoney,omitempty"`
	InvoiceData                  *LineItemInvoiceData                  `json:"invoiceData,omitempty"`
	// Deprecated: Use orderLineDetails instead
	Level3InterchangeInformation *LineItemLevel3InterchangeInformation `json:"level3InterchangeInformation,omitempty"`
	OrderLineDetails             *OrderLineDetails                     `json:"orderLineDetails,omitempty"`
}

// NewLineItem constructs a new LineItem
func NewLineItem() *LineItem {
	return &LineItem{}
}

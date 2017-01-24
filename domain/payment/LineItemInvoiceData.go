// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// LineItemInvoiceData represents class LineItemInvoiceData
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_LineItemInvoiceData
type LineItemInvoiceData struct {
	Description        *string `json:"description,omitempty"`
	MerchantLinenumber *string `json:"merchantLinenumber,omitempty"`
	MerchantPagenumber *string `json:"merchantPagenumber,omitempty"`
	NrOfItems          *string `json:"nrOfItems,omitempty"`
	PricePerItem       *int64  `json:"pricePerItem,omitempty"`
}

// NewLineItemInvoiceData constructs a new LineItemInvoiceData
func NewLineItemInvoiceData() *LineItemInvoiceData {
	return &LineItemInvoiceData{}
}

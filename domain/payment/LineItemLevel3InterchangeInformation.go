// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// LineItemLevel3InterchangeInformation represents class LineItemLevel3InterchangeInformation
type LineItemLevel3InterchangeInformation struct {
	DiscountAmount  *int64  `json:"discountAmount,omitempty"`
	LineAmountTotal *int64  `json:"lineAmountTotal,omitempty"`
	ProductCode     *string `json:"productCode,omitempty"`
	ProductPrice    *int64  `json:"productPrice,omitempty"`
	ProductType     *string `json:"productType,omitempty"`
	Quantity        *int64  `json:"quantity,omitempty"`
	TaxAmount       *int64  `json:"taxAmount,omitempty"`
	Unit            *string `json:"unit,omitempty"`
}

// NewLineItemLevel3InterchangeInformation constructs a new LineItemLevel3InterchangeInformation
func NewLineItemLevel3InterchangeInformation() *LineItemLevel3InterchangeInformation {
	return &LineItemLevel3InterchangeInformation{}
}

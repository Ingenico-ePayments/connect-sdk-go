// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package refund

// References represents class RefundReferences
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RefundReferences
type References struct {
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// NewReferences constructs a new References
func NewReferences() *References {
	return &References{}
}

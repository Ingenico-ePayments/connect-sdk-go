// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// BrowserData represents class BrowserData
type BrowserData struct {
	ColorDepth        *int32  `json:"colorDepth,omitempty"`
	InnerHeight       *string `json:"innerHeight,omitempty"`
	InnerWidth        *string `json:"innerWidth,omitempty"`
	JavaEnabled       *bool   `json:"javaEnabled,omitempty"`
	JavaScriptEnabled *bool   `json:"javaScriptEnabled,omitempty"`
	ScreenHeight      *string `json:"screenHeight,omitempty"`
	ScreenWidth       *string `json:"screenWidth,omitempty"`
}

// NewBrowserData constructs a new BrowserData
func NewBrowserData() *BrowserData {
	return &BrowserData{}
}

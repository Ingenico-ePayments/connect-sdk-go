// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/refund"
)

func approveRefundExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function newInt64 to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for this helper functions can be found in file Helper.go

	var body refund.ApproveRequest
	body.Amount = newInt64(199)

	err := client.Merchant("merchantId").Refunds().Approve("refundId", body, nil)

	fmt.Println(err)
}

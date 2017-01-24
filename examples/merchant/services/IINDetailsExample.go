// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/services"
)

func iinDetailsExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for this helper functions can be found in file Helper.go

	var body services.GetIINDetailsRequest
	body.Bin = newString("4567350000427977")

	response, err := client.Merchant("merchantId").Services().GetIINdetails(body, nil)

	fmt.Println(response, err)
}

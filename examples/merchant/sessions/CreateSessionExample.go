// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/sessions"
)

func createSessionExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	var tokens []string
	tokens = append(tokens, "126166b16ed04b3ab85fb06da1d7a167")
	tokens = append(tokens, "226166b16ed04b3ab85fb06da1d7a167")
	tokens = append(tokens, "122c5b4d-dd40-49f0-b7c9-3594212167a9")
	tokens = append(tokens, "326166b16ed04b3ab85fb06da1d7a167")
	tokens = append(tokens, "426166b16ed04b3ab85fb06da1d7a167")

	var body sessions.SessionRequest
	body.Tokens = &tokens

	response, err := client.Merchant("merchantId").Sessions().Create(body, nil)

	fmt.Println(response, err)
}

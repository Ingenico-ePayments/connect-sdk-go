package defaultimpl

import (
	"strconv"
	"testing"
)

type jsonToken struct {
	Date int
	Iban string
}

type extendedToken struct {
	Date   int
	Iban   string
	Amount int
}

func TestUnmarshalWithExtraFields(t *testing.T) {
	defaultMarshaller, _ := NewDefaultMarshaller()

	iban := "barbarbarbarfoo"
	date := 123

	token := extendedToken{date, iban, 1337}

	json, _ := defaultMarshaller.Marshal(token)

	var jToken jsonToken
	defaultMarshaller.Unmarshal(json, &jToken)

	if token.Date != jToken.Date || token.Iban != jToken.Iban {
		t.Fatal("TestUnmarshalWithExtraFields : expected (" + strconv.Itoa(token.Date) + ", " + token.Iban + ") got (" + strconv.Itoa(jToken.Date) + ", " + jToken.Iban + ")")
	}
}

package logging

import (
	"testing"
)

func CheckObfuscateHeaderWithMatch(t *testing.T, name, originalValue, expectedObfuscatedValue string) {
	obfuscatedValue, err := ObfuscateHeader(name, originalValue)

	if err != nil {
		t.Fatal(err)
	}

	if obfuscatedValue != expectedObfuscatedValue {
		t.Fatalf("CheckObfuscateHeaderWithMatch : expected '%s' got '%s'", expectedObfuscatedValue, obfuscatedValue)
	}
}

func CheckObfuscateHeaderWithNoMatch(t *testing.T, name, originalValue string) {
	obfuscatedValue, err := ObfuscateHeader(name, originalValue)

	if err != nil {
		t.Fatal(err)
	}

	if obfuscatedValue != originalValue {
		t.Fatalf("CheckObfuscateHeaderWithNoMatch : expected '%s' got '%s'", originalValue, obfuscatedValue)
	}
}

func TestObfuscateHeader(t *testing.T) {
	CheckObfuscateHeaderWithMatch(t, "Authorization", "Basic QWxhZGRpbjpPcGVuU2VzYW1l", "********")
	CheckObfuscateHeaderWithMatch(t, "authorization", "Basic QWxhZGRpbjpPcGVuU2VzYW1l", "********")
	CheckObfuscateHeaderWithMatch(t, "AUTHORIZATION", "Basic QWxhZGRpbjpPcGVuU2VzYW1l", "********")

	CheckObfuscateHeaderWithMatch(t, "X-GCS-Authentication-Token", "foobar", "********")
	CheckObfuscateHeaderWithMatch(t, "x-gcs-authentication-token", "foobar", "********")
	CheckObfuscateHeaderWithMatch(t, "X-GCS-AUTHENTICATION-TOKEN", "foobar", "********")

	CheckObfuscateHeaderWithMatch(t, "X-GCS-CallerPassword", "foobar", "********")
	CheckObfuscateHeaderWithMatch(t, "x-gcs-callerpassword", "foobar", "********")
	CheckObfuscateHeaderWithMatch(t, "X-GCS-CALLERPASSWORD", "foobar", "********")

	CheckObfuscateHeaderWithNoMatch(t, "Content-Type", "application/json")
	CheckObfuscateHeaderWithNoMatch(t, "content-type", "application/json")
	CheckObfuscateHeaderWithNoMatch(t, "CONTENT-TYPE", "application/json")
}

func CheckObfuscateBodyWithEmptyBody(t *testing.T) {
	emptyString := ""
	obfuscatedBody, err := ObfuscateBody(emptyString)

	if err != nil {
		t.Fatal(err)
	}

	if obfuscatedBody != emptyString {
		t.Fatalf("CheckObfuscateBodyWithEmptyBody : expected '%s' got '%s'", emptyString, obfuscatedBody)
	}
}

func CheckObfuscateBodyWithCard(t *testing.T) {
	cardObfuscated := `{
    "cardPaymentMethodSpecificInput": {
        "card": {
            "cardNumber": "************3456",
            "cvv": "***",
            "expiryDate": "**20"
        },
        "paymentProductId": 1
    },
    "order": {
        "amountOfMoney": {
            "amount": 2345,
            "currencyCode": "CAD"
        },
        "customer": {
            "billingAddress": {
                "countryCode": "CA"
            }
        }
    }
}`

	cardUnObfuscated := `{
	    "order": {
	        "amountOfMoney": {
	            "currencyCode": "CAD",
	            "amount": 2345
	        },
	        "customer": {
	            "billingAddress": {
	                "countryCode": "CA"
	            }
	        }
	    },
	    "cardPaymentMethodSpecificInput": {
	        "paymentProductId": 1,
	        "card": {
	            "cvv": "123",
	            "cardNumber": "1234567890123456",
	            "expiryDate": "1220"
	        }
	    }
	}`

	obfuscatedBody, err := ObfuscateBody(cardUnObfuscated)

	if err != nil {
		t.Fatal(err)
	}

	if cardObfuscated != obfuscatedBody {
		t.Fatalf("CheckObfuscateBodyWithCard : expected \n%s\ngot\n%s", cardObfuscated, obfuscatedBody)
	}
}

func CheckObfuscateBodyWithIban(t *testing.T) {
	ibanObfuscated := `{
    "paymentProductId": 770,
    "sepaDirectDebit": {
        "customer": {
            "billingAddress": {
                "countryCode": "NL"
            }
        },
        "mandate": {
            "bankAccountIban": {
                "iban": "**************4567"
            },
            "debtor": {
                "surname": "Jones"
            },
            "isRecurring": false
        }
    }
}`
	ibanUnobfuscated := `{
    "sepaDirectDebit": {
        "mandate": {
            "bankAccountIban": {
                "iban": "NL00INGB0001234567"
            },
            "debtor": {
                "surname": "Jones"
            },
            "isRecurring": false
        },
        "customer": {
            "billingAddress": {
                "countryCode": "NL"
            }
        }
    },
    "paymentProductId": 770
}`
	obfuscatedBody, err := ObfuscateBody(ibanUnobfuscated)

	if err != nil {
		t.Fatal(err)
	}

	if ibanObfuscated != obfuscatedBody {
		t.Fatalf("CheckObfuscateBodyWithIban : expected \n%s\ngot\n%s", ibanObfuscated, obfuscatedBody)
	}
}

func CheckObfuscateBodyWithBin(t *testing.T) {
	binObfuscated := `{
    "bin": "123456**"
}`
	binUnobfuscated := `{
    "bin": "12345678"
}`

	obfuscatedBody, err := ObfuscateBody(binUnobfuscated)

	if err != nil {
		t.Fatal(err)
	}

	if binObfuscated != obfuscatedBody {
		t.Fatalf("CheckObfuscateBodyWithBin : expected \n%s\ngot\n%s", binObfuscated, obfuscatedBody)
	}
}

func CheckObfuscateBodyWithNoMatches(t *testing.T) {
	noObfuscation := `{
    "order": {
        "amountOfMoney": {
            "currencyCode": "EUR",
            "amount": 1000
        },
        "customer": {
            "locale": "nl_NL",
            "billingAddress": {
                "countryCode": "NL"
            }
        }
    },
    "bankTransferPaymentMethodSpecificInput": {
        "paymentProductId": 11
    }
}`
	postObfuscation := `{
    "bankTransferPaymentMethodSpecificInput": {
        "paymentProductId": 11
    },
    "order": {
        "amountOfMoney": {
            "amount": 1000,
            "currencyCode": "EUR"
        },
        "customer": {
            "billingAddress": {
                "countryCode": "NL"
            },
            "locale": "nl_NL"
        }
    }
}`

	obfuscatedBody, err := ObfuscateBody(noObfuscation)

	if err != nil {
		t.Fatal(err)
	}

	if postObfuscation != obfuscatedBody {
		t.Fatalf("CheckObfuscateBodyWithNoMatches : expected \n%s\ngot\n%s", postObfuscation, obfuscatedBody)
	}
}

func TestObfuscateProperty(t *testing.T) {
	CheckObfuscateBodyWithEmptyBody(t)
	CheckObfuscateBodyWithCard(t)
	CheckObfuscateBodyWithIban(t)
	CheckObfuscateBodyWithBin(t)
	CheckObfuscateBodyWithNoMatches(t)
}

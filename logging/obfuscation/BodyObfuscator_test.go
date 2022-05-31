package obfuscation

import (
	"bytes"
	"testing"
	"unicode/utf8"
)

func TestObfuscateBodyWithEmptyBody(t *testing.T) {
	emptyString := ""
	obfuscatedBody, err := DefaultBodyObfuscator().ObfuscateBody(emptyString)

	if err != nil {
		t.Fatal(err)
	}

	if obfuscatedBody != emptyString {
		t.Fatalf("TestObfuscateBodyWithEmptyBody : expected '%s' got '%s'", emptyString, obfuscatedBody)
	}
}

func TestObfuscateBodyWithCard(t *testing.T) {
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

	obfuscatedBody, err := DefaultBodyObfuscator().ObfuscateBody(cardUnObfuscated)

	if err != nil {
		t.Fatal(err)
	}

	if cardObfuscated != obfuscatedBody {
		t.Fatalf("TestObfuscateBodyWithCard : expected \n%s\ngot\n%s", cardObfuscated, obfuscatedBody)
	}
}

func TestObfuscateBodyWithCustomCardRule(t *testing.T) {
	cardObfuscated := `{
    "cardPaymentMethodSpecificInput": {
        "card": {
            "cardNumber": "123456******3456",
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

	rule := func(value string) string {
		valueLength := utf8.RuneCountInString(value)
		var chars bytes.Buffer
		i := 0
		for _, r := range value {
			if i < 6 || i >= valueLength - 4 {
				chars.WriteRune(r)
			} else {
				chars.WriteRune('*')
			}
			i++
		}
		return chars.String()
	}
	bodyObfuscator := NewBodyObfuscator(ruleMap{
		"cardNumber": rule,
	})
	obfuscatedBody, err := bodyObfuscator.ObfuscateBody(cardUnObfuscated)

	if err != nil {
		t.Fatal(err)
	}

	if cardObfuscated != obfuscatedBody {
		t.Fatalf("TestObfuscateBodyWithCustomCardRule : expected \n%s\ngot\n%s", cardObfuscated, obfuscatedBody)
	}
}

func TestObfuscateBodyWithIban(t *testing.T) {
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
	obfuscatedBody, err := DefaultBodyObfuscator().ObfuscateBody(ibanUnobfuscated)

	if err != nil {
		t.Fatal(err)
	}

	if ibanObfuscated != obfuscatedBody {
		t.Fatalf("TestObfuscateBodyWithIban : expected \n%s\ngot\n%s", ibanObfuscated, obfuscatedBody)
	}
}

func TestObfuscateBodyWithBin(t *testing.T) {
	binObfuscated := `{
    "bin": "123456**"
}`
	binUnobfuscated := `{
    "bin": "12345678"
}`

	obfuscatedBody, err := DefaultBodyObfuscator().ObfuscateBody(binUnobfuscated)

	if err != nil {
		t.Fatal(err)
	}

	if binObfuscated != obfuscatedBody {
		t.Fatalf("TestObfuscateBodyWithBin : expected \n%s\ngot\n%s", binObfuscated, obfuscatedBody)
	}
}

func TestObfuscateBodyWithNoMatches(t *testing.T) {
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

	obfuscatedBody, err := DefaultBodyObfuscator().ObfuscateBody(noObfuscation)

	if err != nil {
		t.Fatal(err)
	}

	if postObfuscation != obfuscatedBody {
		t.Fatalf("TestObfuscateBodyWithNoMatches : expected \n%s\ngot\n%s", postObfuscation, obfuscatedBody)
	}
}

func TestObfuscateBodyWithObject(t *testing.T) {
	jsonObfuscated := `[
    {
        "value": "****"
    },
    {
        "value": {}
    }
]`
	jsonUnobfuscated := `[ {
	"value": true
}, {
	"value": {
	}
} ]`

	obfuscatedBody, err := DefaultBodyObfuscator().ObfuscateBody(jsonUnobfuscated)

	if err != nil {
		t.Fatal(err)
	}

	if jsonObfuscated != obfuscatedBody {
		t.Fatalf("TestObfuscateBodyWithObject : expected \n%s\ngot\n%s", jsonObfuscated, obfuscatedBody)
	}
}

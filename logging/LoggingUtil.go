package logging

var globalHeaderMap = valueObfuscatorMap{
	"authorization":              valueObfuscatorWithFixedLength(8),
	"www-authenticate":           valueObfuscatorWithFixedLength(8),
	"proxy-authenticate":         valueObfuscatorWithFixedLength(8),
	"proxy-authorization":        valueObfuscatorWithFixedLength(8),
	"x-gcs-authentication-token": valueObfuscatorWithFixedLength(8),
	"x-gcs-callerpassword":       valueObfuscatorWithFixedLength(8),
}

var globalPropertyMap = valueObfuscatorMap{
	"cardNumber":               valueObfuscatorWithKeepingEndCount(4),
	"expiryDate":               valueObfuscatorWithKeepingEndCount(2),
	"cvv":                      valueObfuscatorWithAll(),
	"iban":                     valueObfuscatorWithKeepingEndCount(4),
	"accountNumber":            valueObfuscatorWithKeepingEndCount(4),
	"reformattedAccountNumber": valueObfuscatorWithKeepingEndCount(4),
	"bin":                      valueObfuscatorWithKeepingStartCount(6),
	"value":                    valueObfuscatorWithAll(),
	"keyId":                    valueObfuscatorWithFixedLength(8),
	"secretKey":                valueObfuscatorWithFixedLength(8),
	"publicKey":                valueObfuscatorWithFixedLength(8),
	"userAuthenticationToken":  valueObfuscatorWithFixedLength(8),
	"encryptedPayload":         valueObfuscatorWithFixedLength(8),
	"decryptedPayload":         valueObfuscatorWithFixedLength(8),
	"encryptedCustomerInput":   valueObfuscatorWithFixedLength(8),
}

var globalPropertyObfuscator = newPropertyObfuscator(globalPropertyMap)

// ObfuscateBody obfuscates the given body as necessary
func ObfuscateBody(body string) (string, error) {
	return globalPropertyObfuscator.obfuscate(body)
}

// ObfuscateHeader obfuscates the value for the given header as necessary.
func ObfuscateHeader(name, header string) (string, error) {
	return globalHeaderMap.obfuscateValue(name, header)
}

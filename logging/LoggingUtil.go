package logging

import "github.com/Ingenico-ePayments/connect-sdk-go/logging/obfuscation"

// ObfuscateBody obfuscates the given body as necessary
//
// Deprecated: use BodyObfuscator instead
func ObfuscateBody(body string) (string, error) {
	return obfuscation.DefaultBodyObfuscator().ObfuscateBody(body)
}

// ObfuscateHeader obfuscates the value for the given header as necessary.
//
// Deprecated: use HeaderObfuscator instead
func ObfuscateHeader(name, header string) (string, error) {
	return obfuscation.DefaultHeaderObfuscator().ObfuscateHeader(name, header), nil
}

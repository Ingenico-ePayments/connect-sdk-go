package obfuscation

import "strings"

// HeaderObfuscator can be used to obfuscate headers.
type HeaderObfuscator struct {
	rules ruleMap
}

// ObfuscateHeader obfuscates the value for the given header as necessary.
func (o HeaderObfuscator) ObfuscateHeader(name, value string) string {
	name = strings.ToLower(name)
	return o.rules.obfuscateValue(name, value)
}

// NewHeaderObfuscator returns a header obfuscator.
// This will contain some pre-defined obfuscation rules, as well as any provided custom rules.
func NewHeaderObfuscator(customRules map[string]Rule) HeaderObfuscator {
	rules := ruleMap{
		"authorization":              FixedLength(8),
		"www-authenticate":           FixedLength(8),
		"proxy-authenticate":         FixedLength(8),
		"proxy-authorization":        FixedLength(8),
		"x-gcs-authentication-token": FixedLength(8),
		"x-gcs-callerpassword":       FixedLength(8),
	}

	for name, rule := range customRules {
		name = strings.ToLower(name)
		rules[name] = rule
	}

	return HeaderObfuscator{rules}
}

var defaultHeaderObfuscator = NewHeaderObfuscator(ruleMap{})

// DefaultHeaderObfuscator returns a default header obfuscator.
// This will be equivalent to calling NewHeaderObfuscator with an empty rule map.
func DefaultHeaderObfuscator() HeaderObfuscator {
	return defaultHeaderObfuscator
}

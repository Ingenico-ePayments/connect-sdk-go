package obfuscation

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

var errInvalidDataType = errors.New("invalid data type - should never happen")

// BodyObfuscator can be used to obfuscate properties in JSON bodies.
type BodyObfuscator struct {
	rules ruleMap
}

func (o BodyObfuscator) obfuscateValue(value interface{}, rule Rule) (string, error) {
	if strVal, ok := value.(string); ok {
		return rule(strVal), nil
	}
	if intVal, ok := value.(int32); ok {
		return rule(strconv.FormatInt(int64(intVal), 10)), nil
	}
	if intVal, ok := value.(int64); ok {
		return rule(strconv.FormatInt(intVal, 10)), nil
	}
	if floatVal, ok := value.(float32); ok {
		return rule(strconv.FormatFloat(float64(floatVal), 'f', -1, 32)), nil
	}
	if floatVal, ok := value.(float64); ok {
		return rule(strconv.FormatFloat(floatVal, 'f', -1, 64)), nil
	}
	if boolVal, ok := value.(bool); ok {
		return rule(strconv.FormatBool(boolVal)), nil
	}
	return "", errInvalidDataType
}

func (o BodyObfuscator) navigateJSON(content interface{}) error {
	if content == nil {
		return nil
	}

	if contentMap, ok := content.(map[string]interface{}); ok {
		for name, obj := range contentMap {
			_, isMap := obj.(map[string]interface{})
			if rule, ok := o.rules[name]; ok && !isMap {
				obfuscatedValue, oErr := o.obfuscateValue(obj, rule)
				if oErr != nil {
					return oErr
				}

				contentMap[name] = obfuscatedValue

			} else {
				nErr := o.navigateJSON(obj)
				if nErr != nil {
					return nErr
				}
			}
		}
	}

	if contentSlice, ok := content.([]interface{}); ok {
		for _, obj := range contentSlice {
			nErr := o.navigateJSON(obj)
			if nErr != nil {
				return nErr
			}
		}
	}

	return nil
}

// ObfuscateBody obfuscates the given body as necessary
func (o BodyObfuscator) ObfuscateBody(body string) (string, error) {
	if strings.TrimSpace(body) == "" {
		return body, nil
	}

	var parsedJSON interface{}
	pErr := json.Unmarshal([]byte(body), &parsedJSON)
	if _, ok := pErr.(*json.SyntaxError); ok {
		return body, nil
	}
	if pErr != nil {
		return body, pErr
	}

	nErr := o.navigateJSON(parsedJSON)
	if nErr != nil {
		return body, nErr
	}

	obfuscatedBody, mErr := json.MarshalIndent(parsedJSON, "", "    ")
	if mErr != nil {
		return body, mErr
	}

	return string(obfuscatedBody), nil
}

// NewBodyObfuscator returns a body obfuscator.
// This will contain some pre-defined obfuscation rules, as well as any provided custom rules.
func NewBodyObfuscator(customRules map[string]Rule) BodyObfuscator {
	rules := ruleMap{
		"cardNumber":               KeepingEndCount(4),
		"expiryDate":               KeepingEndCount(2),
		"cvv":                      All(),
		"iban":                     KeepingEndCount(4),
		"accountNumber":            KeepingEndCount(4),
		"reformattedAccountNumber": KeepingEndCount(4),
		"bin":                      KeepingStartCount(6),
		"value":                    All(),
		"keyId":                    FixedLength(8),
		"secretKey":                FixedLength(8),
		"publicKey":                FixedLength(8),
		"userAuthenticationToken":  FixedLength(8),
		"encryptedPayload":         FixedLength(8),
		"decryptedPayload":         FixedLength(8),
		"encryptedCustomerInput":   FixedLength(8),
	}

	for name, rule := range customRules {
		rules[name] = rule
	}

	return BodyObfuscator{rules}
}

var defaultBodyObfuscator = NewBodyObfuscator(ruleMap{})

// DefaultBodyObfuscator returns a default body obfuscator.
// This will be equivalent to calling NewBodyObfuscator with an empty rule map.
func DefaultBodyObfuscator() BodyObfuscator {
	return defaultBodyObfuscator
}

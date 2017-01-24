package logging

import (
	"strings"
)

func (o valueObfuscatorMap) obfuscateValue(key, value string) (string, error) {
	key = strings.ToLower(key)

	obfuscator, ok := o[key]
	if ok == true {
		return obfuscator.obfuscateValue(value)
	}

	return value, nil
}

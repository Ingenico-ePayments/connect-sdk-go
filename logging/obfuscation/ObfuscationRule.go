package obfuscation

// Rule defines how a single value is obfuscated.
// This can be as simple as returning a fixed mask, or as complex as necessary.
type Rule func(value string) string

type ruleMap map[string]Rule

func (o ruleMap) obfuscateValue(key, value string) string {
	rule, ok := o[key]
	if ok == true {
		return rule(value)
	}

	return value
}

// All returns an obfuscation rule that will replace all characters with *.
func All() Rule {
	obfuscator := valueObfuscatorWithAll()
	return obfuscator.obfuscateValue
}

// FixedLength returns an obfuscation rule that will replace values with a fixed length string containing only *.
func FixedLength(fixedLength int) Rule {
	obfuscator := valueObfuscatorWithFixedLength(fixedLength)
	return obfuscator.obfuscateValue
}

// KeepingStartCount returns an obfuscation rule that will keep a fixed number of characters at the start,
// then replaces all other characters with *.
func KeepingStartCount(count int) Rule {
	obfuscator := valueObfuscatorWithKeepingStartCount(count)
	return obfuscator.obfuscateValue
}

// KeepingEndCount returns an obfuscation rule that will keep a fixed number of characters at the end,
// then replaces all other characters with *.
func KeepingEndCount(count int) Rule {
	obfuscator := valueObfuscatorWithKeepingEndCount(count)
	return obfuscator.obfuscateValue
}

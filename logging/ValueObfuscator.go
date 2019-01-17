package logging

import (
	"bytes"
	"unicode/utf8"
)

type valueObfuscator struct {
	maskCharacter  rune
	fixedLength    int
	keepStartCount int
	keepEndCount   int
}

type valueObfuscatorMap map[string]valueObfuscator

func valueObfuscatorWithAll() valueObfuscator {
	return newValueObfuscator(0, 0, 0)
}

func valueObfuscatorWithFixedLength(fixedLength int) valueObfuscator {
	return newValueObfuscator(fixedLength, 0, 0)
}

func valueObfuscatorWithKeepingStartCount(count int) valueObfuscator {
	return newValueObfuscator(0, count, 0)
}

func valueObfuscatorWithKeepingEndCount(count int) valueObfuscator {
	return newValueObfuscator(0, 0, count)
}

func (vo valueObfuscator) obfuscateValue(value string) (string, error) {
	valueLength := utf8.RuneCountInString(value)

	if valueLength == 0 {
		return value, nil
	}
	if vo.fixedLength > 0 {
		return vo.repeatMask(vo.fixedLength), nil
	}
	if valueLength < vo.keepStartCount || valueLength < vo.keepEndCount {
		return value, nil
	}

	var chars bytes.Buffer
	i := 0
	for _, r := range value {
		if i < vo.keepStartCount || i >= valueLength - vo.keepEndCount {
			chars.WriteRune(r)
		} else {
			chars.WriteRune(vo.maskCharacter)
		}
		i++
	}

	return chars.String(), nil
}

func (vo valueObfuscator) repeatMask(count int) string {
	var buffer bytes.Buffer

	for i := vo.keepStartCount; i < count; i++ {
		buffer.WriteRune(vo.maskCharacter)
	}

	return buffer.String()
}

func newValueObfuscator(fixedLength, keepStartCount, keepEndCount int) valueObfuscator {
	return valueObfuscator{'*', fixedLength, keepStartCount, keepEndCount}
}

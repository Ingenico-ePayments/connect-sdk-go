package obfuscation

import "testing"

func TestObfuscateValueWithAll(t *testing.T) {
	value := "Hello, world"
	expected := "************"
	obfuscator := valueObfuscatorWithAll()

	actual := obfuscator.obfuscateValue(value)

	if actual != expected {
		t.Fatalf("TestObfuscateValueWithAll : expected '%s' got '%s'", expected, actual)
	}
}

func TestObfuscateValueWithKeepStartAndEnd(t *testing.T) {
	value := "Hello, world"
	expected := "He********ld"
	obfuscator := newValueObfuscator(0, 2, 2)

	actual := obfuscator.obfuscateValue(value)

	if actual != expected {
		t.Fatalf("TestObfuscateValueWithKeepStartAndEnd : expected '%s' got '%s'", expected, actual)
	}
}

func TestObfuscateValueWithAllUnicode(t *testing.T) {
	value := "Hello, 世界"
	expected := "*********"
	obfuscator := valueObfuscatorWithAll()

	actual := obfuscator.obfuscateValue(value)

	if actual != expected {
		t.Fatalf("TestObfuscateValueWithAllUnicode : expected '%s' got '%s'", expected, actual)
	}
}

func TestObfuscateValueWithKeepStartAndEndUnicode(t *testing.T) {
	value := "Hello, 世界"
	expected := "He*****世界"
	obfuscator := newValueObfuscator(0, 2, 2)

	actual := obfuscator.obfuscateValue(value)

	if actual != expected {
		t.Fatalf("TestObfuscateValueWithKeepStartAndEndUnicode : expected '%s' got '%s'", expected, actual)
	}
}

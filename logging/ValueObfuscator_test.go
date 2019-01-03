package logging

import "testing"

func TestObfuscateUnicodeValue(t *testing.T) {
	value := "Hello, 世界"

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("TestObfuscateUnicodeValue : panic on obfuscate value '%v': %v", value, r)
		}
	}()

	vo := valueObfuscatorWithAll()
	actual, err := vo.obfuscateValue(value)
	if err != nil {
		t.Fatal(err)
	}

	expected := "*********"
	if actual != expected {
		t.Fatalf("TestObfuscateUnicodeValue : expected '%s' got '%s'", expected, actual)
	}
}

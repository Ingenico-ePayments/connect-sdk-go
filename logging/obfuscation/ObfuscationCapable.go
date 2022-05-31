package obfuscation

// Capable indicates support for obfuscating bodies and headers.
type Capable interface {
	// SetBodyObfuscator sets the body obfuscator to use.
	SetBodyObfuscator(bodyObfuscator BodyObfuscator)
	// SetHeaderObfuscator sets the header obfuscator to use.
	SetHeaderObfuscator(headerObfuscator HeaderObfuscator)
}

package morse_test

import (
	"morse/morse"
	"testing"
)

func TestDecryptMorse(t *testing.T) {
	cases := []struct {
		input    string
		expected string
		err      error
	}{
		{
			input:    ".-",
			expected: "A",
			err:      nil,
		},
		{
			input:    "-...",
			expected: "B",
			err:      nil,
		},
		{
			input:    "-..",
			expected: "D",
			err:      nil,
		},
		{
			input:    ".",
			expected: "E",
			err:      nil,
		},
		{
			input:    ".----",
			expected: "1",
			err:      nil,
		},
		{
			input:    ".... . .-.. .-.. --- / .-- --- .-. .-.. -..",
			expected: "HELLO WORLD",
			err:      nil,
		},
		{
			input:    "_",
			expected: "",
			err:      morse.ErrorInvalidCharacterFromMorse,
		},
		{
			input:    "..-..-..-..-..-..",
			expected: "",
			err:      morse.ErrorInvalidCharacterFromMorse,
		},
		{
			input:    ".- .- A",
			expected: "",
			err:      morse.ErrorInvalidCharacterFromMorse,
		},
	}

	for _, c := range cases {
		value, err := morse.DecryptMorse(morse.MorseCode(c.input))
		if value != c.expected {
			t.Errorf("Expected '%s' from '%s' but got '%s' with error '%s'", c.expected, c.input, value, err)
		}

		if err != c.err {
			t.Errorf("Expected error '%s' but got error '%s' for input '%s'", c.err, err, c.input)
		}
	}
}

func TestEncryptMorse(t *testing.T) {
	cases := []struct {
		input    string
		expected morse.MorseCode
		err      error
	}{
		{
			input:    "A",
			expected: ".-",
			err:      nil,
		},
		{
			input:    "B",
			expected: "-...",
			err:      nil,
		},
		{
			input:    "D",
			expected: "-..",
			err:      nil,
		},
		{
			input:    "E",
			expected: ".",
			err:      nil,
		},
		{
			input:    "1",
			expected: ".----",
			err:      nil,
		},
		{
			input:    "Hello World",
			expected: ".... . .-.. .-.. --- / .-- --- .-. .-.. -..",
			err:      nil,
		},
		{
			input:    "A sentace.",
			expected: "",
			err:      morse.ErrorInvalidCharacterToMorse,
		},
		{
			input:    "@#$%",
			expected: "",
			err:      morse.ErrorInvalidCharacterToMorse,
		},
		{
			input:    "_",
			expected: "",
			err:      morse.ErrorInvalidCharacterToMorse,
		},
		{
			input:    "—",
			expected: "",
			err:      morse.ErrorInvalidCharacterToMorse,
		},
		{
			input:    "−",
			expected: "",
			err:      morse.ErrorInvalidCharacterToMorse,
		},
		{
			input:    "  ",
			expected: "",
			err:      morse.ErrorInvalidDoubleSpaces,
		},
	}

	for _, c := range cases {
		value, err := morse.EncryptMorse(c.input)
		if value != c.expected {
			t.Errorf("Expected '%s' from '%s' but got '%s' with error '%s'", c.expected, c.input, value, err)
		}

		if err != c.err {
			t.Errorf("Expected error '%s' but got error '%s' for input '%s'", c.err, err, c.input)
		}
	}
}

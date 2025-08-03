package morse_test

import (
	"morse/morse"
	"testing"
)

func TestDecryptMorse(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{
			name:     "Morse A encrypts as 'A'",
			input:    ".-",
			expected: "A",
			err:      nil,
		},
		{
			name:     "Morse B encrypts as 'B'",
			input:    "-...",
			expected: "B",
			err:      nil,
		},
		{
			name:     "Morse D encrypts as 'D'",
			input:    "-..",
			expected: "D",
			err:      nil,
		},
		{
			name:     "Morse E encrypts as 'E'",
			input:    ".",
			expected: "E",
			err:      nil,
		},
		{
			name:     "Morse 1 encrypts as '1'",
			input:    ".----",
			expected: "1",
			err:      nil,
		},
		{
			name:     "Morse 'Hello World' encrypts as 'HELLO WORLD'",
			input:    ".... . .-.. .-.. --- / .-- --- .-. .-.. -..",
			expected: "HELLO WORLD",
			err:      nil,
		},
		{
			name:     "Underscore should error with invalid character error",
			input:    "_",
			expected: "",
			err:      morse.ErrorInvalidCharacterFromMorse,
		},
		{
			name:     "Unknown morse returns invalid character for morse error",
			input:    "..-..-..-..-..-..",
			expected: "",
			err:      morse.ErrorInvalidCharacterFromMorse,
		},
		{
			name:     "Morse input with non-morse character returns error",
			input:    ".- .- A",
			expected: "",
			err:      morse.ErrorInvalidCharacterFromMorse,
		},
		{
			name:     "Empty space decrypting returns error",
			input:    "  ",
			expected: "",
			err:      morse.ErrorInvalidDoubleSpaces,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			value, err := morse.DecodeMorse(morse.MorseCode(c.input))
			if value != c.expected {
				t.Errorf("Expected '%s' from '%s' but got '%s' with error '%s'", c.expected, c.input, value, err)
			}

			if err != c.err {
				t.Errorf("Expected error '%s' but got error '%s' for input '%s'", c.err, err, c.input)
			}
		})
	}
}

func TestEncryptMorse(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected morse.MorseCode
		err      error
	}{
		{
			name:     "A encrypts as '.-'",
			input:    "A",
			expected: ".-",
			err:      nil,
		},
		{
			name:     "B encrypts as '-...'",
			input:    "B",
			expected: "-...",
			err:      nil,
		},
		{
			name:     "D encrypts as '-..'",
			input:    "D",
			expected: "-..",
			err:      nil,
		},
		{
			name:     "E encrypts as '.'",
			input:    "E",
			expected: ".",
			err:      nil,
		},
		{
			name:     "1 encrypts as '.----'",
			input:    "1",
			expected: ".----",
			err:      nil,
		},
		{
			name:     "String hello world encrypts correctly",
			input:    "Hello World",
			expected: ".... . .-.. .-.. --- / .-- --- .-. .-.. -..",
			err:      nil,
		},
		{
			name:     "Invalid . in encrypt source results in error",
			input:    "A sentace.",
			expected: "",
			err:      morse.ErrorInvalidCharacterToMorse,
		},
		{
			name:     "Invalid morse characters results in error",
			input:    "@#$%",
			expected: "",
			err:      morse.ErrorInvalidCharacterToMorse,
		},
		{
			name:     "Invalid morse character '_' results in error",
			input:    "_",
			expected: "",
			err:      morse.ErrorInvalidCharacterToMorse,
		},
		{
			name:     "Invalid morse character '—' results in error",
			input:    "—",
			expected: "",
			err:      morse.ErrorInvalidCharacterToMorse,
		},
		{
			name:     "Invalid morse character '−' results in error",
			input:    "−",
			expected: "",
			err:      morse.ErrorInvalidCharacterToMorse,
		},
		{
			name:     "double space for encryption results in error",
			input:    "  ",
			expected: "",
			err:      morse.ErrorInvalidDoubleSpaces,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			value, err := morse.EncodeMorse(c.input)
			if value != c.expected {
				t.Errorf("Expected '%s' from '%s' but got '%s' with error '%s'", c.expected, c.input, value, err)
			}

			if err != c.err {
				t.Errorf("Expected error '%s' but got error '%s' for input '%s'", c.err, err, c.input)
			}
		})
	}
}

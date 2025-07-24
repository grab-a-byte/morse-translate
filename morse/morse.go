package morse

import (
	"errors"
	"strings"
)

type MorseCode string

var ErrorInvalidCharacterToMorse = errors.New("invalid character to be translated to morse")

var ErrorInvalidCharacterFromMorse = errors.New("invalid charcter to be translated from morse")

var ErrorInvalidDoubleSpaces = errors.New("double space should not be in morse code input")

var encryptMap map[rune]MorseCode = make(map[rune]MorseCode)
var decryptMap map[MorseCode]rune = make(map[MorseCode]rune)

func init() {
	codes := map[rune]MorseCode{
		'A': ".-",
		'B': "-...",
		'C': "-.-.",
		'D': "-..",
		'E': ".",
		'F': "..-.",
		'G': "--.",
		'H': "....",
		'I': "..",
		'J': ".---",
		'K': "-.-",
		'L': ".-..",
		'M': "--",
		'N': "-.",
		'O': "---",
		'P': ".--.",
		'Q': "--.-",
		'R': ".-.",
		'S': "...",
		'T': "-",
		'U': "..-",
		'V': "...-",
		'W': ".--",
		'X': "-..-",
		'Y': "-.--",
		'Z': "--..",
		'0': "-----",
		'1': ".----",
		'2': "..---",
		'3': "...--",
		'4': "....-",
		'5': ".....",
		'6': "-....",
		'7': "--...",
		'8': "---..",
		'9': "----.",
		' ': "/",
	}

	for c, v := range codes {
		encryptMap[c] = v
		decryptMap[v] = c
	}
}

func Translate(input string) (string, error) {
	parts := strings.Split(input, " ")
	for _, c := range parts[0] {
		if c != '.' && c != '-' {
			out, err := EncryptMorse(input)
			return string(out), err
		}
	}

	return DecryptMorse(MorseCode(input))
}

func EncryptMorse(input string) (MorseCode, error) {
	if strings.Contains(input, "  ") {
		return "", ErrorInvalidDoubleSpaces
	}
	out := make([]string, len(input))
	for i, c := range strings.ToUpper(input) {
		val, ok := encryptMap[c]
		if !ok {
			return "", ErrorInvalidCharacterToMorse
		}
		out[i] = string(val)
	}

	return MorseCode(strings.Join(out, " ")), nil
}

func DecryptMorse(input MorseCode) (string, error) {
	if strings.Contains(string(input), "  ") {
		return "", ErrorInvalidDoubleSpaces
	}

	codes := strings.Split(string(input), " ")
	out := make([]string, len(codes))
	for i, m := range codes {
		val, ok := decryptMap[MorseCode(m)]
		if !ok {
			return "", ErrorInvalidCharacterFromMorse
		}
		out[i] = string(val)
	}

	return strings.Join(out, ""), nil
}

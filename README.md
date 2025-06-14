Input parameters
Standard string of your language for translating to Morse Code
Morse Code inputs must have a single ' ' between letters and a '/' between words


Test Cases:
Input => Output

.- => "A"
-... => "B"
-.. => "D"
. => "E"
.---- => "1"
.... . .-.. .-.. --- / .-- --- .-. .-.. -.. => "Hello World"
"  "(Double space) => Should error as should only be single space

..-..-..-..-..-.. => Should error as invalid character

.- .- A => Should error as invalid char in morse code input

"A" => .-
"B" => -...
"D" => -..
"E" => .
"1" => .----
"Hello World" => .... . .-.. .-.. --- / .-- --- .-. .-.. -..
"A sentance." => Should error ('.' not an encodable charater)
"@#$%" => Should error as non-valid charaters to be converted to Morse Code


"_" => Should error both from and to morse code
"—" => Should error (May need to copy string as it is a character that looks like a -)
"−" => Should error (May need to copy string as it is a character that looks like a -)
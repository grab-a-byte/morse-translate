# Challenge paramters
Input parameters
Standard string of your language for translating to Morse Code
Morse Code inputs must have a single ' ' between letters and a '/' between words


## Test Cases:
Input => Output

From Morse
* .- => "A"
* -... => "B"
* -.. => "D"
* . => "E"
* .---- => "1"
* .... . .-.. .-.. --- / .-- --- .-. .-.. -.. => "Hello World"

From Morse Errors
* "  "(Double space) => Should error as should only be single space
* ..-..-..-..-..-.. => Should error as invalid character
* .- .- A => Should error as invalid char in morse code input

From Text
* "A" => .-
* "B" => -...
* "D" => -..
* "E" => .
* "1" => .----
* "Hello World" => .... . .-.. .-.. --- / .-- --- .-. .-.. -..
* "A sentance." => Should error ('.' not an encodable charater)
* "@#$%" => Should error as non-valid charaters to be converted to Morse Code

From Text Errors
* "_" => Should error both from and to morse code
* "—" => Should error (May need to copy string as it is a character that looks like a -)
* "−" => Should error (May need to copy string as it is a character that looks like a -)

# Building and usage
## Cli

Build : `go build -o cli cmd/cmd/cmd.go`
Usage : `./cli -type="Text" 'Hello World'`
- Valid options for type are "Text" and "MorseCode"
- Positional arguemnt used as input
- If no type if specified, the application will use the contetn of the first block of test before a space to determine and try and encode/decode based on that
    - E.g. (Hello World: Hello: Text mode used) (--. ..-: --.: Morse mode used)

## Api
Build :`go build cmd/api/api.go`
- Request is in format of
    ```
    {
        "message": "Hello World",
        "type":"Text"
    }
    ```
- Response in format of
    ```
    {
        "message": "Hello World",
        "requestId": "d148451d-32ed-4150-8325-04c7c74d187c",
        "translation": ".... . .-.. .-.. --- / .-- --- .-. .-.. -..",
        "type": "Text"
    }
    ```

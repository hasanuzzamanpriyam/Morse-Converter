package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var morseCodeMap = map[string]string{
	".-": "A", "-...": "B", "-.-.": "C", "-..": "D",
	".": "E", "..-.": "F", "--.": "G", "....": "H",
	"..": "I", ".---": "J", "-.-": "K", ".-..": "L",
	"--": "M", "-.": "N", "---": "O", ".--.": "P",
	"--.-": "Q", ".-.": "R", "...": "S", "-": "T", "..-": "U",
	"...-": "V", ".--": "W", "-..-": "X", "-.--": "Y",
	"--..": "Z", ".----": "1", "..---": "2", "...--": "3", "....-": "4",
	".....": "5", "-....": "6", "--...": "7", "---..": "8", "----.": "9", "-----": "0",
}

func MorseToText(morseCode string) string {
	words := strings.Split(morseCode, " ")
	var text strings.Builder

	for _, word := range words {
		if val, ok := morseCodeMap[word]; ok {
			text.WriteString(val)
		} else {
			text.WriteString(" ")
		}
	}

	return text.String()
}

func TextToMorse(text string) string {
	var morseCode strings.Builder

	for _, char := range text {
		if char == ' ' {
			morseCode.WriteString(" ")
		} else {
			for k, v := range morseCodeMap {
				if strings.ToUpper(string(char)) == v {
					morseCode.WriteString(k)
					morseCode.WriteString(" ")
					break
				}
			}
		}
	}

	return morseCode.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter '1' to convert Morse code to text, '2' to convert text to Morse code:")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("Enter Morse code:")
		morseCode, _ := reader.ReadString('\n')
		morseCode = strings.TrimSpace(morseCode)
		text := MorseToText(morseCode)
		fmt.Println("Text:", text)
	case "2":
		fmt.Println("Enter text:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		morseCode := TextToMorse(text)
		fmt.Println("Morse Code:", morseCode)
	default:
		fmt.Println("Invalid choice")
	}
}

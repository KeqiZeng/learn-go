package main

import (
	"fmt"
	"strings"
)

func decipher(message string, keyword string) string {
	cipherText := message
	var decipherText string
	var i int
	for _, c := range cipherText {
		cb := byte(c)
		var cipcode byte
		cipcode = keyword[i] - 'A' + 1
		i++
		if i == len(keyword)-1 {
			i = 0
		}
		cb -= cipcode
		if cb < 'A' {
			cb += 26
		}
		decipherText += string(cb)
	}
	return decipherText
}

func cipher(message string, keyword string) string {
	decipherText := strings.ToUpper(strings.Replace(message, " ", "", -1))
	var cipherText string
	var i int
	for _, c := range decipherText {
		cb := byte(c)
		var decipcode byte
		decipcode = keyword[i] - 'A' + 1
		i++
		if i == len(keyword)-1 {
			i = 0
		}
		cb += decipcode
		if cb > 'Z' {
			cb -= 26
		}
		cipherText += string(cb)
	}
	return cipherText
}

func main() {
	message := "your message"
	keyword := "GOLANG"
	fmt.Println(cipher(message, keyword))
	fmt.Println(decipher(cipher(message, keyword), keyword))
}

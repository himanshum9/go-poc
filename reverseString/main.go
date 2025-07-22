package main

import "fmt"

// func reversedString(str string) string {
// 	newStr := ""
// 	for i := len(str) - 1; i > 0; i-- {
// 		newStr = newStr + string(str[i])
// 	}
// 	return newStr
// }
func main() {
	// str := "Hello World"
	// result := reversedString(str)
	// fmt.Println(result)

	// str := "hello"
	// fmt.Printf("byte: %v\n", str[0])
	// fmt.Println(str[0])
	// fmt.Printf("rune: %v\n", []rune(str)[0])

	// str2 := "世界"
	// fmt.Printf("byte: %v\n", str2[0])
	// fmt.Printf("rune: %v\n", []rune(str2)[0])

	englishText := "Hello"

	// String with mixed characters
	mixedText := "Hello 世界"

	// 1. Simple ASCII processing
	fmt.Println("English text, character by character:")
	for i := 0; i < len(englishText); i++ {
		fmt.Printf("%c ", englishText[i])
	}

	// 2. Unicode text processing
	fmt.Println("\nMixed text, character by character:")
	for _, char := range mixedText {
		fmt.Printf("%c ", char)
	}

}

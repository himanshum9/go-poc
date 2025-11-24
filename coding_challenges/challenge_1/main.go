//Goal: Produce an 8-character alphanumeric username made of two 4-character segments (first-name + last-name). Implement a concise Go program that generates such usernames.
//Notes: Decide how to handle non-alphanumeric input and whether to seed randomness for reproducible tests.

package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var alphaNum = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func clean(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile(`[^a-z0-9]`)
	return re.ReplaceAllString(s, "")
}

func fillRandom(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNum[rand.Intn(len(alphaNum))]
	}
	return string(b)
}

func username(first, last string) string {
	first = clean(first)
	last = clean(last)

	if len(first) < 4 {
		first += fillRandom(4 - len(first))
	} else {
		first = first[:4]
	}

	if len(last) < 4 {
		last += fillRandom(4 - len(last))
	} else {
		last = last[:4]
	}

	return first + last
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(username("Him@n$hu", "Me#t@a"))
}

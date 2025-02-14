package strings

import "fmt"

func PrintAll(s string) {
	if s == "" {
		return
	}
	fmt.Printf("%c ", s[0])
	PrintAll(s[1:])
}

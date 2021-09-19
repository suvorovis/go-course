package main

import (
	"fmt"
)

func main() {
	fmt.Println(Unpack(`qwe\4\5`))
}

func Unpack(s string) string {
	var escape uint8
	var symbol rune
	var result []rune
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z' || escape == 1:
			if symbol != 0 {
				result = append(result, symbol)
			}
			symbol = r
			escape = 0
		case r >= '2' && r <= '9' && symbol != 0:
			for i := 0; i < int(r-'0'); i++ {
				result = append(result, symbol)
			}
			symbol = 0
		case r == '\\':
			escape = 1
		}
	}
	if symbol != 0 {
		result = append(result, symbol)
	}
	return string(result)
}

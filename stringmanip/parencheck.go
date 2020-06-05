package main

import (
	"container/list"
	"fmt"
	"os"
)

func main() {
	var expr = os.Args[1]

	var tokens = map[byte]int{
		'0': 1,
		'1': 1,
		'2': 1,
		'3': 1,
		'4': 1,
		'5': 1,
		'6': 1,
		'7': 1,
		'8': 1,
		'9': 1,
		'+': 1,
		'-': 1,
		'*': 1,
		'/': 1,
		'(': 1,
		')': 1,
		'{': 1,
		'}': 1,
		'[': 1,
		']': 1,
	}

	var openp = map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	var closep = map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	exprstack := list.New()

	for i := range expr {

		c := expr[i]
		_, present := tokens[c]
		if !present {
			fmt.Printf("Encountered non-arithmetic token%c\n", c)
			os.Exit(1)
		}

		_, isopen := openp[c]
		if isopen {
			exprstack.PushFront(c)
			continue
		}

		cp, isclose := closep[c]
		if isclose {
			if exprstack.Len() == 0 {
				fmt.Printf("Encountered unbalanced close paren %c\n", c)
				os.Exit(1)
			}
			lastopen := exprstack.Front()
			if lastopen.Value != cp {
				fmt.Printf("Mismatched parens, open %c and close %c\n", lastopen, cp)
				os.Exit(1)
			}
			// pop open p
			exprstack.Remove(exprstack.Front())
		}
	}

	if exprstack.Len() > 0 {
		fmt.Printf("Encountered unbalanced open paren %c\n", exprstack.Front())
		os.Exit(1)
	}

	fmt.Println("Valid parentheses in expression")
	os.Exit(0)
}

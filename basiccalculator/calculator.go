package basiccalculator

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	var stack []string
	for _, c := range []byte(s) {
		if c == ' ' {
			continue
		} else if len(stack) == 0 || stack[len(stack)-1] == "(" ||
			stack[len(stack)-1] == "+" || stack[len(stack)-1] == "-" ||
			c == '+' || c == '-' || c == '(' {
			stack = append(stack, string([]byte{c}))
			continue
		} else if c >= '0' && c <= '9' {
			num, err := strconv.Atoi(stack[len(stack)-1])
			if err != nil {
				panic("illegal")
			}
			stack[len(stack)-1] = strconv.Itoa(num*10 + int(c-'0'))
			continue
		} else if c != ')' {
			panic(fmt.Sprintf("illegal: %c", c))
		}
		sum := 0
		sum, stack = eval(stack)
		stack = append(stack, strconv.Itoa(sum))
	}
	sum, _ := eval(stack)
	return sum
}

func eval(stack []string) (int, []string) {
	sum := 0
	lastNum := 0
	for len(stack) != 0 && stack[len(stack)-1] != "(" {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if t == "+" {
			continue
		} else if t == "-" {
			sum -= 2 * lastNum
			continue
		}
		num, err := strconv.Atoi(t)
		if err != nil {
			panic("illegal")
		}
		sum += num
		lastNum = num
	}
	if len(stack) != 0 {
		stack = stack[:len(stack)-1]
	}
	return sum, stack
}

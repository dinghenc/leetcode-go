package main

import (
	"fmt"
	"log"

	perfect "github.com/dinghenc/leetcode-go/perfect_number"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("This is a project for leetcode.")

	perfectNumbers := findPerfectNumber()
	log.Println(perfectNumbers)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.Errorf("number by divided must be not zero")
	}
	return a / b, nil
}

func findPerfectNumber() []int {
	var perfectNumbers []int
	for i := 1; i <= 1e8; i++ {
		if perfect.CheckPerfectNumber(i) {
			perfectNumbers = append(perfectNumbers, i)
		}
	}
	return perfectNumbers
}

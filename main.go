package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	perfect "github.com/dinghenc/leetcode-go/perfect_number"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("This is a project for leet-code.")
	log.Println(div(10, 5))

	perfectNumbers := findPerfectNumber()
	log.Println(perfectNumbers)

	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(`{"id":"123","name":"ding-he"}`), &m); err != nil {
		log.Fatal(err)
	}
	log.Println(m)

	httpServer()
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.Errorf("number by divided must be not zero")
	}
	return a / b, nil
}

func findPerfectNumber() []int {
	var perfectNumbers []int
	for i := 1; i <= 100; i++ {
		if perfect.CheckPerfectNumber(i) {
			perfectNumbers = append(perfectNumbers, i)
		}
	}
	return perfectNumbers
}

func httpServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if _, err := io.WriteString(writer, `{"id":"15043216","name":"ding-he"}`); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	})
	log.Fatal(http.ListenAndServe(":9999", nil))
}

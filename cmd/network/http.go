package network

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "hello")
}

func RunNetwork() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}

package service

import (
	"fmt"
	"io"
	"net/http"
)

func SayHello() {
	response, err := http.Get("http://localhost:8080/hello")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	fmt.Println("response from server:", string(body))
}

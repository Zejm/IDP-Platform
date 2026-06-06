package network

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type AddUserRequest struct {
	Name string `json:"name"`
}

type User struct {
	Name string `json:"name"`
	Id string `json:"id"`
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	var req AddUserRequest

	decoderErr := json.NewDecoder(r.Body).Decode(&req)
	if decoderErr != nil {
		http.Error(w, decoderErr.Error(), http.StatusBadRequest)
		return
	}

	fileData, fileErr := os.ReadFile("users.json")

	var users []User

	newUser := User{
		Name: req.Name,
		Id: uuid.New().String(),
	}

	if fileErr == nil {
		json.Unmarshal(fileData, &users)
	}

	users = append(users, newUser)
	updatedData, updErr := json.MarshalIndent(users, "", " ")
	if updErr != nil {
		http.Error(w, updErr.Error(), http.StatusInternalServerError)
		return
	}

	writeErr := os.WriteFile("users.json", updatedData, 0644)
	if writeErr != nil {
		http.Error(w, writeErr.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RunNetwork() {
	http.HandleFunc("/add-user", addUserHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

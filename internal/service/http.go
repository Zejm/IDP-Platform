package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AddUserRequest struct {
	Name string `json:"name"`
}

func AddUser(name string) string {
	data, _ := json.Marshal(AddUserRequest{
		Name: name,
	})
	
	response, err := http.Post("http://localhost:8080/add-user", "application/json", bytes.NewBuffer(data))

	if err != nil {
		return fmt.Sprintf("Error saving user: %s", err)
	}
	if response.StatusCode != http.StatusOK {
		return "Error saving user: request canceled"
	}

	defer response.Body.Close()
	return fmt.Sprintf("User '%s' successfully added", name)
}
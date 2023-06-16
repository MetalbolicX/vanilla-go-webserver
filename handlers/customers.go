package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MetalbolicX/vanilla-go-webserver/repository"
)

type userRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type userResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

/* Example:
curl -X POST -H "Content-Type: application/json" -d '{
  "name": "John Doe",
  "email": "johndoe@example.com"
}' http://localhost:3000/user

*/

func NewCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a User struct
	var user userRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert the user into the database
	err = repository.Post(r.Context(), "INSERT INTO customers (name, email) VALUES ($1, $2)", user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Return a success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse{
		Message: "Created successfully",
		Status:  true,
	})
}

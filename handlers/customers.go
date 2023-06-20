package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MetalbolicX/vanilla-go-webserver/repository"
	"github.com/MetalbolicX/vanilla-go-webserver/utils"
)

type customerRequest struct {
	ID    *int   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type customerResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

// Add a new customer to the database.
// For example:
// curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "johndoe@example.com"}' http://localhost:3000/customer
func NewCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a User struct
	var user customerRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert the user into the database
	err = repository.Post(r.Context(), `
		INSERT INTO customers (name, email)
		VALUES ($1, $2)`,
		user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Return a success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customerResponse{
		Message: "Created successfully",
		Status:  true,
	})
}

// Get information of the customer by the id.
// For example:
// curl localhost:3000/customer/1
func GetCustomerByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the request URL
	idStr := utils.GetIdentifier(r.URL.Path)
	customerID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch data from DB
	customer, err := repository.Get(r.Context(), `
		SELECT
			id
			, name
			, email
		FROM customers
		WHERE id = $1
	`, customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the info to JSON
	customerJSON, err := json.Marshal(customer)
	if err != nil {
		http.Error(w, "Failed to convert customer to JSON", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(customerJSON)
}

// Update the information of a customer by its id.
// For example:
// curl -X PUT -H "Content-Type: application/json" -d '{"name": "Jose", "email": "josee@example.com"}' http://localhost:3000/customer/1
func UpdateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the request URL
	idStr := utils.GetIdentifier(r.URL.Path)
	customerId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	// Parse the request body
	var customer customerRequest
	err = json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// If the ID field is nil, set it to the extracted customer ID
	if customer.ID == nil {
		customer.ID = &customerId //&userID
	}

	// Update the user in the database
	rowsUpdated, err := repository.Put(r.Context(), `
		UPDATE customers
		SET name = $1, email = $2
		WHERE id = $3
	`, customer.Name, customer.Email, &customer.ID)
	if err != nil {
		http.Error(w, "Failed to update customer", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Return a success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customerResponse{
		Message: fmt.Sprintf("%d customer updated successfully", rowsUpdated),
		Status:  true,
	})
}

// Delete a customer by id in the database.
// For example:
// curl -X DELETE localhost:3000/customer/1
func DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the request URL
	idStr := utils.GetIdentifier(r.URL.Path)
	customerId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	rowsDeleted, err := repository.Put(r.Context(),
		"DELETE FROM customers WHERE id = $1",
		customerId)
	if err != nil {
		http.Error(w, "Failed to delete customer", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Return a success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customerResponse{
		Message: fmt.Sprintf("%d customer deleted successfully", rowsDeleted),
		Status:  true,
	})
}

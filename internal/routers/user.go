package routers

import (
	models "apirest/internal/models"
	tools "apirest/tools"
	"encoding/json"
	"net/http"
)

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		http.Error(w, "Error"+err.Error(), 400)

	}

	if len(user.Email) == 0 {
		http.Error(w, "the email is required", 400)

	}

	_, find, _ := tools.CheckUserAlreadyExist(user.Email)
	if find == true {
		http.Error(w, "user already exist", 400)

	}

	_, status, err := tools.InsertUser(user)

	if err != nil {
		http.Error(w, "error inserting new user"+err.Error(), 400)

	}

	if status == false {
		http.Error(w, "error inserting", 400)

	}

	json.NewEncoder(w).Encode(&user)
	w.WriteHeader(http.StatusCreated)

}

// GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) {

	result, status := tools.GetAllUsers()
	if status == false {
		http.Error(w, "eror geting users", http.StatusBadRequest)

	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&result)

}

// GetOneUser ...
func GetOneUser(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "could send the id parameter", http.StatusBadRequest)

	}
	user, err := tools.GetUser(ID)

	if err != nil {
		http.Error(w, "error"+err.Error(), 400)

	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

// UpdateOneUser ...
func UpdateOneUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var ID string

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "data incorrect"+err.Error(), 400)

	}
	var status bool
	status, err = tools.UpdateUser(user, ID)
	if err != nil {
		http.Error(w, "error Updating user"+err.Error(), 400)

	}

	if status == false {
		http.Error(w, "error Updating the user"+err.Error(), 400)

	}

	w.WriteHeader(http.StatusAccepted)

}

// DeleteOneUser ...
func DeleteOneUser(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "could send id parameter", http.StatusBadRequest)

	}

	err := tools.DeleteUser(ID)
	if err != nil {
		http.Error(w, "error", http.StatusBadRequest)

	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

}

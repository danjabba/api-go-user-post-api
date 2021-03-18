package routers

import (
	models "apirest/internal/models"
	tools "apirest/tools"
	"encoding/json"
	"net/http"
)

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) error {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		http.Error(w, "Error"+err.Error(), 400)
		return err
	}

	if len(user.Email) == 0 {
		http.Error(w, "the email is required", 400)
		return err
	}

	_, find, _ := tools.CheckUserAlreadyExist(user.Email)
	if find == true {
		http.Error(w, "user already exist", 400)
		return err
	}

	_, status, err := tools.InsertUser(user)

	if err != nil {
		http.Error(w, "error inserting new user"+err.Error(), 400)
		return err
	}

	if status == false {
		http.Error(w, "error inserting", 400)
		return err

	}

	json.NewEncoder(w).Encode(&user)
	w.WriteHeader(http.StatusCreated)

	return nil
}

// GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) error {

	result, status := tools.GetAllUsers()
	if status == false {
		http.Error(w, "eror geting users", http.StatusBadRequest)
		return nil
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&result)

	return nil
}

// GetOneUser ...
func GetOneUser(w http.ResponseWriter, r *http.Request) error {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "could send the id parameter", http.StatusBadRequest)
		return nil
	}
	user, err := tools.GetUser(ID)

	if err != nil {
		http.Error(w, "error"+err.Error(), 400)
		return err

	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	return nil
}

// UpdateOneUser ...
func UpdateOneUser(w http.ResponseWriter, r *http.Request) error {
	var user models.User
	var ID string

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "data incorrect"+err.Error(), 400)
		return err
	}
	var status bool
	status, err = tools.UpdateUser(user, ID)
	if err != nil {
		http.Error(w, "error Updating user"+err.Error(), 400)
		return err
	}

	if status == false {
		http.Error(w, "error Updating the user"+err.Error(), 400)
		return nil
	}

	w.WriteHeader(http.StatusAccepted)
	return nil

}

// DeleteOneUser ...
func DeleteOneUser(w http.ResponseWriter, r *http.Request) error {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "could send id parameter", http.StatusBadRequest)
		return nil
	}

	err := tools.DeleteUser(ID)
	if err != nil {
		http.Error(w, "error", http.StatusBadRequest)
		return err
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	return err
}

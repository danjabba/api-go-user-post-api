// Package routers Provides all the functions and methods.
package routers

import (
	models "apirest/internal/models"
	tools "apirest/tools"
	"encoding/json"
	"net/http"
)

//  CreatePost ...
func CreatePost(w http.ResponseWriter, r *http.Request) {

	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {

		http.Error(w, "Error"+err.Error(), 400)

	}

	if len(post.Body) == 0 {
		http.Error(w, "the body is required", 400)

	}

	_, find, _ := tools.CheckPostAlreadyExist(post.Slug)
	if find == true {
		http.Error(w, "post already exist", 400)

	}

	_, status, err := tools.InsertPost(post)

	if err != nil {
		http.Error(w, "error inserting new post"+err.Error(), 400)

	}

	if status == false {
		http.Error(w, "error inserting", 400)

	}

	json.NewEncoder(w).Encode(&post)
	w.WriteHeader(http.StatusCreated)

}

// GetPosts ...
func GetPosts(w http.ResponseWriter, r *http.Request) {

	result, status := tools.GetAllPosts()
	if status == false {
		http.Error(w, "eror geting posts", http.StatusBadRequest)

	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&result)

}

// GetOnePost ...
func GetOnePost(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "could send the id parameter", http.StatusBadRequest)

	}
	post, err := tools.GetPost(ID)

	if err != nil {
		http.Error(w, "error"+err.Error(), 400)

	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)

}

// UpdateOnePost ...
func UpdateOnePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	var ID string

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		http.Error(w, "data incorrect"+err.Error(), 400)
		return
	}
	var status bool
	status, err = tools.UpdatePost(post, ID)
	if err != nil {
		http.Error(w, "error Updating post"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "error Updating the post"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

// DeleteOnePost ...
func DeleteOnePost(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "could send id parameter", http.StatusBadRequest)

	}

	err := tools.DeletePost(ID)
	if err != nil {
		http.Error(w, "error", http.StatusBadRequest)

	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

}

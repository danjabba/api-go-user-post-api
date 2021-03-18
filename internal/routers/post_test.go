package routers_test

import (
	models "apirest/internal/models"
	"apirest/internal/routers"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CreatePostResponseStructure ...
type CreatePostResponseStructure struct {
	Message string      `bson:"message" json:"message"`
	Success bool        `bson:"success" json:"success"`
	Data    models.Post `bson:"data" json:"data"`
}

// GetPostResponseStructure ...
type GetPostResponseStructure struct {
	Message string      `bson:"message" json:"message"`
	Success bool        `bson:"success" json:"success"`
	Data    models.Post `bson:"data" json:"data"`
}

// DeletePostResponseStructure ...
type DeletePostResponseStructure struct {
	Message string      `bson:"message" json:"message"`
	Success bool        `bson:"success" json:"success"`
	Data    models.Post `bson:"data" json:"data"`
}

// ListPostResponseStructure...
type ListPostResponseStructure struct {
	Message string        `bson:"message" json:"message"`
	Success bool          `bson:"success" json:"success"`
	Data    []models.Post `bson:"data" json:"data"`
}

// UpdatePostResponseStructure ...
type UpdatePostResponseStructure struct {
	Message string      `bson:"message" json:"message"`
	Success bool        `bson:"success" json:"success"`
	Data    models.Post `bson:"data" json:"data"`
}

// CreatePostByRequest ...
func CreatePostByRequest(input models.Post) (string, error) {

	path := "/createpost"
	payload := fmt.Sprintf(`{"body": "%s",
							"slug": "%s"}`,
		input.Body, input.Slug)
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(payload))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	routers.CreateUser(rec, req)
	var post CreatePostResponseStructure
	json.Unmarshal(rec.Body.Bytes(), &post)
	return post.Data.Slug, nil

}

// TestCreatepost1: should create a post
func TestCreatepost1(t *testing.T) {
	path := "/createpost"

	payload := `{"body": "hola" , "slug" : "primer"}`
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(payload))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()

	if assert.NoError(t, routers.CreatePost(rec, req), rec.Code) {
		var response CreateUserResponseStructure
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NotEqual(t, "", response.Data.ID)
	}

}

// TestCreatepost2: should not create a post, slug already exist
func TestCreatepost2(t *testing.T) {
	path := "/createpost"

	payload := `{"body": "hola" , "slug" : "primer"}`
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(payload))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()

	if assert.NoError(t, routers.CreatePost(rec, req)) {

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

}

// TestCreatepost3: should not create a post, body is missing
func TestCreatepost3(t *testing.T) {
	path := "/createpost"

	payload := `{"body": "" , "slug" : "primer"}`
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(payload))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()

	if assert.NoError(t, routers.CreatePost(rec, req)) {

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

}

// TestCreatepost4: should not create  post, model is missing
func TestCreatepost4(t *testing.T) {
	path := "/createpost"

	payload := `{}`
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(payload))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()

	if assert.NoError(t, routers.CreatePost(rec, req)) {

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

}

// TestGetPost1: should not get a post, id is missing
func TestGetPost1(t *testing.T) {

	path := "/post"
	req := httptest.NewRequest(http.MethodGet, path, strings.NewReader(""))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()

	if assert.NoError(t, routers.GetOnePost(rec, req)) {

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		var user GetPostResponseStructure
		json.Unmarshal(rec.Body.Bytes(), &user)

	}

}

// TestGetPost2: should  get a  post
func TestGetPost2(t *testing.T) {

	path := "/post?id=6053ba2091b3d911eb0da159"
	req := httptest.NewRequest(http.MethodGet, path, strings.NewReader(""))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()

	if assert.NoError(t, routers.GetOnePost(rec, req)) {

		assert.Equal(t, http.StatusOK, rec.Code)
		var user GetPostResponseStructure
		json.Unmarshal(rec.Body.Bytes(), &user)

	}

}

// TestGetPosts1: should  get all  posts
func TestGetPosts1(t *testing.T) {

	path := "/getposts"
	req := httptest.NewRequest(http.MethodGet, path, strings.NewReader(""))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()

	if assert.NoError(t, routers.GetPosts(rec, req)) {

		assert.Equal(t, http.StatusAccepted, rec.Code)
		var user ListPostResponseStructure
		json.Unmarshal(rec.Body.Bytes(), &user)

	}

}

// TestDeletePost1: should  delete a post
func TestDeletePost1(t *testing.T) {

	path := "/post?id=60535788ea5bbbbe1baa5fa9"
	req := httptest.NewRequest(http.MethodDelete, path, strings.NewReader(""))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()

	if assert.NoError(t, routers.DeleteOnePost(rec, req)) {

		assert.Equal(t, http.StatusOK, rec.Code)
		var user DeleteUserResponseStructure
		json.Unmarshal(rec.Body.Bytes(), &user)

	}

}

package routers_test

// import (
// 	models "apirest/internal/models"
// 	"apirest/internal/routers"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// // CreateUserResponseStructure ...
// type CreateUserResponseStructure struct {
// 	Message string      `bson:"message" json:"message"`
// 	Success bool        `bson:"success" json:"success"`
// 	Data    models.User `bson:"data" json:"data"`
// }

// // GetUserResponseStructure ...
// type GetUserResponseStructure struct {
// 	Message string      `bson:"message" json:"message"`
// 	Success bool        `bson:"success" json:"success"`
// 	Data    models.User `bson:"data" json:"data"`
// }

// // DeleteUserResponseStructure ...
// type DeleteUserResponseStructure struct {
// 	Message string      `bson:"message" json:"message"`
// 	Success bool        `bson:"success" json:"success"`
// 	Data    models.User `bson:"data" json:"data"`
// }

// // ListUserResponseStructure ...
// type ListUserResponseStructure struct {
// 	Message string        `bson:"message" json:"message"`
// 	Success bool          `bson:"success" json:"success"`
// 	Data    []models.User `bson:"data" json:"data"`
// }

// // UpdateUserResponseStructure ...
// type UpdateUserResponseStructure struct {
// 	Message string      `bson:"message" json:"message"`
// 	Success bool        `bson:"success" json:"success"`
// 	Data    models.User `bson:"data" json:"data"`
// }

// // CreateUserByRequest ...
// func CreateUserByRequest(input models.User) (string, error) {

// 	path := "/createuser"
// 	payload := fmt.Sprintf(`{"name": "%s",
// 							"las_name": "%s",
// 							"email": "%s"}`,
// 		input.Name, input.LastName, input.Email)
// 	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(payload))
// 	req.Header.Set("content-type", "application/json")
// 	rec := httptest.NewRecorder()
// 	routers.CreateUser(rec, req)
// 	var user CreateUserResponseStructure
// 	json.Unmarshal(rec.Body.Bytes(), &user)
// 	return user.Data.Email, nil

// }

// // TestCreateUser1: should create an user
// func TestCreateUser1(t *testing.T) {
// 	path := "/createuser"

// 	payload := `{"name": "enriques" , "last_name" : "jabba" , "email":"test@test.com"}`
// 	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(payload))
// 	req.Header.Set("content-type", "application/json")
// 	rec := httptest.NewRecorder()

// 	if assert.NoError(t, routers.CreateUser(rec, req), rec.Code) {
// 		var response CreateUserResponseStructure
// 		json.Unmarshal(rec.Body.Bytes(), &response)
// 		assert.NotEqual(t, "", response.Data.ID)
// 	}

// }

// // TestCreateUser2: should not create an user because the email already exist
// func TestCreateUser2(t *testing.T) {

// 	path := "/createuser"

// 	payload := `{"name": "enriques" , "last_name" : "jabba" , "email":"test@test.com"}`
// 	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(payload))
// 	req.Header.Set("content-type", "application/json")
// 	rec := httptest.NewRecorder()

// 	if assert.NoError(t, routers.CreateUser(rec, req)) {
// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
// 	}

// }

// // TestCreateUser3: should not create email is missing
// func TestCreateUser3(t *testing.T) {

// 	path := "/createuser"

// 	payload := `{"name": "enriques" , "last_name" : "jabba" , "email":""}`
// 	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(payload))
// 	req.Header.Set("content-type", "application/json")
// 	rec := httptest.NewRecorder()

// 	if assert.NoError(t, routers.CreateUser(rec, req)) {
// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
// 	}

// }

// // TestCreateUser4: should not create user is missing
// func TestCreateUser4(t *testing.T) {

// 	path := "/createuser"

// 	payload := `{}`
// 	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(payload))
// 	req.Header.Set("content-type", "application/json")
// 	rec := httptest.NewRecorder()

// 	if assert.NoError(t, routers.CreateUser(rec, req)) {
// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
// 	}

// }

// // TestCreateUser5: should not create an user
// func TestCreateUser5(t *testing.T) {
// 	path := "/createuser"

// 	//payload := `{"name": "enriques" , "last_name" : "jabba" , "email":"test@test.com"}`
// 	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader("{}"))
// 	req.Header.Set("content-type", "application/json")
// 	rec := httptest.NewRecorder()

// 	if assert.NoError(t, routers.CreateUser(rec, req)) {
// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
// 	}

// }

// // TestGetUser1: should not get an user, id is missing
// func TestGetUser1(t *testing.T) {

// 	path := "/users"
// 	req := httptest.NewRequest(http.MethodGet, path, strings.NewReader(""))
// 	req.Header.Set("content-type", "application/json")
// 	rec := httptest.NewRecorder()

// 	if assert.NoError(t, routers.GetOneUser(rec, req)) {

// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
// 		var user GetUserResponseStructure
// 		json.Unmarshal(rec.Body.Bytes(), &user)

// 	}

// }

// // TestGetUser2: should  get an user
// func TestGetUser2(t *testing.T) {

// 	path := "/users?id=605355fd5e94881aa7ae0029"
// 	req := httptest.NewRequest(http.MethodGet, path, strings.NewReader(""))
// 	req.Header.Set("content-type", "application/json")
// 	rec := httptest.NewRecorder()

// 	if assert.NoError(t, routers.GetOneUser(rec, req)) {

// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		var user GetUserResponseStructure
// 		json.Unmarshal(rec.Body.Bytes(), &user)

// 	}

// }

// // TestDeleteUser1: should  delete an user
// func TestDeleteUser1(t *testing.T) {

// 	path := "/users?id=60535788ea5bbbbe1baa5fa9"
// 	req := httptest.NewRequest(http.MethodDelete, path, strings.NewReader(""))
// 	req.Header.Set("content-type", "application/json")
// 	rec := httptest.NewRecorder()

// 	if assert.NoError(t, routers.DeleteOneUser(rec, req)) {

// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		var user DeleteUserResponseStructure
// 		json.Unmarshal(rec.Body.Bytes(), &user)

// 	}

// }

// // TestGetUsers1: should  get  users
// func TestGetUsers1(t *testing.T) {

// 	path := "/getusers"
// 	req := httptest.NewRequest(http.MethodGet, path, strings.NewReader(""))
// 	req.Header.Set("content-type", "application/json")
// 	rec := httptest.NewRecorder()

// 	if assert.NoError(t, routers.GetUsers(rec, req)) {

// 		assert.Equal(t, http.StatusAccepted, rec.Code)
// 		var user GetUserResponseStructure
// 		json.Unmarshal(rec.Body.Bytes(), &user)

// 	}

// }

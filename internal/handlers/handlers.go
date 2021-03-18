package handlers

import (
	middlew "apirest/internal/middlew"
	routers "apirest/internal/routers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers ...
func Handlers() {

	router := mux.NewRouter()

	// User Routes
	router.HandleFunc("/createuser", middlew.CheckDataBase(routers.CreateUser)).Methods("POST")
	router.HandleFunc("/user", middlew.CheckDataBase(routers.GetOneUser)).Methods("GET")
	router.HandleFunc("/updateuser", middlew.CheckDataBase(routers.UpdateOneUser)).Methods("PUT")
	router.HandleFunc("/getusers", middlew.CheckDataBase(routers.GetUsers)).Methods("GET")
	router.HandleFunc("/deleteuser", middlew.CheckDataBase(routers.DeleteOneUser)).Methods("DELETE")

	// Post Routes
	router.HandleFunc("/createpost", middlew.CheckDataBase(routers.CreatePost)).Methods("POST")
	router.HandleFunc("/post", middlew.CheckDataBase(routers.GetOnePost)).Methods("GET")
	router.HandleFunc("/updatepost", middlew.CheckDataBase(routers.UpdateOnePost)).Methods("PUT")
	router.HandleFunc("/getposts", middlew.CheckDataBase(routers.GetPosts)).Methods("GET")
	router.HandleFunc("/deletepost", middlew.CheckDataBase(routers.DeleteOnePost)).Methods("DELETE")

	// PORT ...
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

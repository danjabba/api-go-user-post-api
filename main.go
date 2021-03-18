package main

import (
	handler "apirest/internal/handlers"
	tools "apirest/tools"
	"log"
)

func main() {

	if tools.CheckConection() == false {
		log.Fatal("no database conection")
		return
	}
	handler.Handlers()

}

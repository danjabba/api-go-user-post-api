// package main

// import (
// 	handler "apirest/internal/handlers"
// 	tools "apirest/tools"
// 	"log"
// )

// func main() {

// 	if tools.CheckConection() == false {
// 		log.Fatal("no database conection")
// 		return
// 	}
// 	handler.Handlers()

// }

package main

import (
	u "apirest/pdfConvert"
	"fmt"
)

func main() {

	r := u.NewRequestPdf("")

	//html template path
	templatePath := "template/sample.html"

	//path for download pdf
	outputPath := "storage/example.pdf"

	//html template data
	templateData := struct {
		Title       string
		Description string
		Company     string
		Contact     string
		Country     string
	}{
		Title:       "HTML to PDF generator",
		Description: "This is the simple HTML to PDF file.",
		Company:     "Jhon Lewis",
		Contact:     "Maria Anders",
		Country:     "Germany",
	}

	if err := r.ParseTemplate(templatePath, templateData); err == nil {
		ok, _ := r.GeneratePDF(outputPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}

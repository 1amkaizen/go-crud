package main

import (
	"fmt"
	"net/http"

	"github.com/1amkaizen/go_crud/controllers/pasienController"
)

func main() {
	fmt.Println("http://localhost:3000")

	http.HandleFunc("/", pasienController.Index)
	http.HandleFunc("/pasien", pasienController.Index)
	http.HandleFunc("/pasien/index", pasienController.Index)
	http.HandleFunc("/pasien/add", pasienController.Add)
	http.HandleFunc("/pasien/edit", pasienController.Edit)
	http.HandleFunc("/pasien/delete", pasienController.Delete)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("bootstrap"))))
	http.ListenAndServe(":3000", nil)

}

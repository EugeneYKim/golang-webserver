package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func handleMain(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("template/main.html")
	if err != nil {
		fmt.Println("Error generating main page: " + err.Error())

	}

	if err := template.Execute(w, nil); err != nil {
		fmt.Println("Error while executing template: " + err.Error())
	}

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Port not specified, defaulting to Port 8090")
		port = "8090"
	}
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handleMain)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

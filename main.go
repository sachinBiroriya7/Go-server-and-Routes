package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func main() {
	fmt.Println()

	fileServer := http.FileServer(http.Dir("./static")) //searches for index.html itself
	http.Handle("/", fileServer)

	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/exit", exitHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("starting the server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error in starting the server : %v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" && r.Method != "GET" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "Hello Ho gya Ji!! Balle balle!!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("error in parsing form, :", err)
	}
	fmt.Fprintln(w, "POST request success, Data sent okay!")

	fmt.Fprintln(w, "username :=", r.FormValue("name"))
	fmt.Fprintln(w, "Address :=", r.FormValue("address"))
}

// parsing files using ParseFiles & Execute
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/about.html")
	if err != nil {
		fmt.Println("error in parsing about.html,", err)
		return
	}
	tmpl.Execute(w, nil)
}

// parsing files using ParseGlob & ExecuteTemplate for passing data to static file
func exitHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./static/*html")
	if err != nil {
		fmt.Println("error in parsing byebye.html,", err)
		return
	}
	tmpl.ExecuteTemplate(w, "bye.html", "Messiiiii")
}

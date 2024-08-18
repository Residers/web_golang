package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Переход на index")
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	log.Println("Переход на index")
	if err != nil {
		fmt.Fprintf(w, err.Error())

	}
	t.ExecuteTemplate(w, "index", nil)
	log.Println("Запуск ExecuteTemplate")
}
func contacts_page(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/contacts_page.html", "templates/header.html", "templates/footer.html")
	t.ExecuteTemplate(w, "contacts_page", nil)
}
func create(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
	log.Println("Запуск save_article")
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")
	log.Println("Запуск save_article", title, anons, full_text)
}
func handleFunc() {
	log.Println("Запуск handleFunc")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/contacts/", contacts_page)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/save_article", save_article)
	log.Println("Запуск сервера на http://127.0.0.1:8082/")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	log.Println("Запуск ")
	handleFunc()
}

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Переход на index")
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	log.Println("Переход на index")
	if err != nil {
		fmt.Fprintf(w, err.Error())

	}
	db := openDb()
	defer db.Close()
	posts := getAllArticles(db)
	t.ExecuteTemplate(w, "index", posts)
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
func openDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang") //root:root - login:password

	if err != nil {
		panic(err)
	}

	return db
}
func save_article(w http.ResponseWriter, r *http.Request) {
	log.Println("Запуск save_article")
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {

		http.Error(w, "Заполните все поля", http.StatusBadRequest)
		return
	}
	fmt.Println("Start db")
	db := openDb()

	defer db.Close()
	//установка данных
	insertInDb(db, title, anons, full_text)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func insertInDb(db *sql.DB, title string, anons string, full_text string) {

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('%s', '%s', '%s') ", title, anons, full_text))
	if err != nil {
		panic(err)
	}
	defer insert.Close()
}

// new code in main
type Article struct {
	Id        uint16 `json:"id"`
	Title     string `json:"title"`
	Anons     string `json:"anons"`
	Full_text string `json:"full_text"`
}

var posts = []Article{}

func getAllArticles(db *sql.DB) []Article {
	res, err := db.Query("SELECT * FROM `articles`")

	if err != nil {
		panic(err)
	}

	defer res.Close()
	posts = []Article{}

	for res.Next() {
		var post Article
		err := res.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text)

		if err != nil {
			panic(err)
		}
		posts = append(posts, post)

	}
	return posts
}

// new comment
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

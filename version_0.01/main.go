package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	Avg_point, Happiness float32
	Hobbies              []string
}

func home_page(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Main page")
	bob := User{
		Name:      "Bob",
		Age:       32,
		Money:     1000,
		Avg_point: 10.4,
		Happiness: 0.8,
		Hobbies:   []string{"swim", "run"}}

	// fmt.Fprintf(w, "Hello, %s\n", bob.Name)
	// bob.setNewName("Tom")
	// fmt.Fprint(w, bob.getFullInfo())
	tmpl, _ := template.ParseFiles("src/templates/home_page.html")
	tmpl.Execute(w, bob)
}
func (u *User) getFullInfo() string {
	return fmt.Sprintf("Name: %s, age: %d, money: %d, avg_point: %f "+
		"happiness: %f", u.Name, u.Age, u.Money, u.Avg_point, u.Happiness)
}
func (u *User) setNewName(newName string) {
	u.Name = newName
}
func contacts_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("src/templates/contacts_page.html")
	tmpl.Execute(w, nil)
}

func handleRequest() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home_page)
	mux.HandleFunc("/contacts/", contacts_page)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Println("Запуск сервера на http://127.0.0.1:8081")
	err := http.ListenAndServe(":8081", mux)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Нет ошибки")
}
func main() {
	handleRequest()
}

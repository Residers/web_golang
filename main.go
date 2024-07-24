package main

import (
	"fmt"
	"net/http"
)

type User struct{
	name string
	age uint16
	money int16
	avg_point, happiness float32

}
func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Main page")
}
func (u *User) getFullInfo() string{
	return fmt.Sprintf("Name: %s, age: %d, money: %d, avg_point: %f, happiness: %f", u.name, u.age, u.money, u.avg_point, u.happiness)
}
func (u *User) setNewName(newName string) {
	u.name = newName
}
func contacts_page(w http.ResponseWriter, r *http.Request) {
	bob := User{
		name: "Bob",
		age: 32,
		money: 1000,
		avg_point: 10.4,
		happiness: 0.8	}

	fmt.Fprintf(w, "Hello, %s\n", bob.name) 
	bob.setNewName("Tom")
	fmt.Fprint(w, bob.getFullInfo()) 
	fmt.Fprintf(w, "\n Contacts page")
	
}

func handleRequest(){
	

http.HandleFunc("/",home_page)
http.HandleFunc("/contacts/",contacts_page)
http.ListenAndServe(":8080", nil)
}
func main() {
	fmt.Println("Hello, World!")
	handleRequest()
}


package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Promo struct {
	Nom             string
	Filière         string
	Niveau          int
	NombreEtudiants int
	Users           []User
}

type User struct {
	FirstName string
	LastName  string
	Age       int
	Gender    string
}
type PageInit struct {
	Nom           string
	Prenom        string
	DateNaissance string
	Gender        string
}

var count int = 0
var logs PageInit

func main() {
	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Printf(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		dataPage := Promo{"Mentor'ac", "Informatique", 5, 3, []User{
			{"Cyril", "RODRIGUES", 22, "Male"},
			{"Kheir-eddine", "MEDERREG", 22, "Male"},
			{"Alan", "PHILIPIERT", 26, "Male"}}}
		temp.ExecuteTemplate(w, "promo", dataPage)
	})

	type PageChange struct {
		Count      int
		CheckValue bool
	}

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		count++
		data := PageChange{count, false}
		if count%2 == 0 {
			data.CheckValue = true
		}
		temp.ExecuteTemplate(w, "change", data)

	})

	http.HandleFunc("/user/init", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "init", nil)

	})

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		logs = PageInit{
			r.FormValue("user_nom"),
			r.FormValue("user_prenom"),
			r.FormValue("user_date"),
			r.FormValue("user_sexe")}
		fmt.Println(logs)
		http.Redirect(w, r, "/user/display", http.StatusMovedPermanently)
	})

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "display", logs)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Init serv
	http.ListenAndServe("localhost:8080", nil)
}

package router

import (
	"go-login-system/db"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))
var isLogged bool = false

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("func_Index page")

	tmpl.ExecuteTemplate(w, "Index", isLogged)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	log.Println("func_SignIn page")

	tmpl.ExecuteTemplate(w, "SignIn", nil)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	log.Println("func_NewUser")

	if r.Method == "POST" {
		db := db.Connect()
		defer db.Close()

		name := r.FormValue("user_name")
		rawPassword := r.FormValue("user_password")

		log.Println("Encrypting the given password")
		password, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 8)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("INSERT INTO users VALUES ($1, $2)", name, password)
		if err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/", 301)
	}
}

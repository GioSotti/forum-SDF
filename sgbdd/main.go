package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
	Email    string
	Password string
}

func main() {
	// Établir une connexion à la base de données
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/basededonnees")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Handler pour la page d'inscription
	http.HandleFunc("/inscription", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Récupérer les données du formulaire d'inscription
			username := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")

			// Insérer les données dans la base de données
			_, err := db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, password)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintln(w, "Inscription réussie !")
		} else {
			// Afficher le formulaire d'inscription
			tmpl := template.Must(template.ParseFiles("inscription.html"))
			tmpl.Execute(w, nil)
		}
	})

	// Handler pour la page de connexion
	http.HandleFunc("/connexion", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Récupérer les données du formulaire de connexion
			usernameOrEmail := r.FormValue("usernameOrEmail")
			password := r.FormValue("password")

			// Vérifier les informations de connexion dans la base de données
			// ... votre logique de vérification des informations de connexion ici ...

			fmt.Fprintln(w, "Connexion réussie !")
		} else {
			// Afficher le formulaire de connexion
			tmpl := template.Must(template.ParseFiles("connexion.html"))
			tmpl.Execute(w, nil)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

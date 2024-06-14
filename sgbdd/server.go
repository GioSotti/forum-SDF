package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
)

type User struct {
	Pseudo          string `json:"pseudo"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

var user User

func registerHandlertest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Récupération des données du corps de la requête
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		println("posted")

		// Validation des données (vous pouvez ajouter des vérifications supplémentaires ici)

		// Connexion à la base de données MySQL
		db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/forum")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Insertion des données dans la base de données
		insertQuery := `
			INSERT INTO users (pseudo, mail, password)
			VALUES (?, ?, ?)`
		_, err = db.Exec(insertQuery, user.Pseudo, user.Email, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Réponse JSON indiquant que l'inscription a réussi
		response := map[string]interface{}{
			"message": "Inscription réussie",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}

	custTemplate, err := template.ParseFiles("register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Data := user
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Récupération des données du corps de la requête
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validation des données (vous pouvez ajouter des vérifications supplémentaires ici)

		// Connexion à la base de données MySQL
		db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/forum")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Insertion des données dans la base de données
		insertQuery := `
			INSERT INTO users (pseudo, email, password)
			VALUES (?, ?, ?)`
		_, err = db.Exec(insertQuery, user.Pseudo, user.Email, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Réponse JSON indiquant que l'inscription a réussi
		response := map[string]interface{}{
			"message": "Inscription réussie",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/register", registerHandlertest)

	fmt.Println("Tout roule sur le port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

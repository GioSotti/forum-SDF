package main

import (
	/* "database/sql" */
	"html/template"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
}

func main() {
	// Connexion à la base de données PostgreSQL
	dsn := "host=localhost user=forum password=forum dbname=forum port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migration des modèles (création de la table si elle n'existe pas)
	db.AutoMigrate(&User{})

	// Route pour afficher la liste des utilisateurs
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var users []User
		db.Find(&users)

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, users)
	})

	// Démarrage du serveur HTTP
	log.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
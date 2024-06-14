package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	ID         int
	Pseudo     string
	Mail       string
	MotDePasse string
}

func main() {
	// Création de la connexion à la base de données
	db, err := sql.Open("postgres", "user=forum password=forum dbname=forum sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Création de la table si elle n'existe pas déjà
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			pseudo VARCHAR(50) NOT NULL,
			mail VARCHAR(100) NOT NULL,
			mot_de_passe VARCHAR(100) NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Configuration du routeur Gin
	router := gin.Default()

	// Gestion de la requête POST pour insérer les données dans la base de données
	router.POST("/insert", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Erreur: %s", err.Error()))
			return
		}

		// Insertion des données dans la base de données
		_, err := db.Exec(`
			INSERT INTO users (pseudo, mail, mot_de_passe)
			VALUES ($1, $2, $3)
		`, user.Pseudo, user.Mail, user.MotDePasse)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Erreur lors de l'insertion des données: %s", err.Error()))
            return
        }

        c.String(http.StatusOK, "Données insérées avec succès !")
        })


         router.Run(":8080")

         }
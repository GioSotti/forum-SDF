package Func

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func AddLike(username string, likeNumber string) error {
	// Lire le contenu du fichier JSON
	data, err := ioutil.ReadFile("database/account.json")
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier JSON : %v", err)
	}

	var users []User

	// Décoder le contenu JSON dans une structure de données
	err = json.Unmarshal(data, &users)
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion du JSON : %v", err)
	}

	// Rechercher l'utilisateur avec le nom d'utilisateur donné
	for i, user := range users {
		if user.Username == username {
			// Ajouter le numéro au tableau de likes de l'utilisateur
			users[i].Like = append(users[i].Like, likeNumber)
			break
		}
	}

	// Convertir la structure de données en JSON
	jsonData, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion en JSON : %v", err)
	}

	// Écrire le JSON dans le fichier
	err = ioutil.WriteFile("database/account.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture dans le fichier JSON : %v", err)
	}

	return nil
}
func AddDislike(username string, dislikeNumber string) error {
	// Lire le contenu du fichier JSON
	data, err := ioutil.ReadFile("database/account.json")
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier JSON : %v", err)
	}

	var users []User

	// Décoder le contenu JSON dans une structure de données
	err = json.Unmarshal(data, &users)
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion du JSON : %v", err)
	}

	// Rechercher l'utilisateur avec le nom d'utilisateur donné
	for i, user := range users {
		if user.Username == username {
			// Ajouter le numéro au tableau de likes de l'utilisateur
			users[i].Dislike = append(users[i].Dislike, dislikeNumber)
			break
		}
	}

	// Convertir la structure de données en JSON
	jsonData, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion en JSON : %v", err)
	}

	// Écrire le JSON dans le fichier
	err = ioutil.WriteFile("database/account.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture dans le fichier JSON : %v", err)
	}

	return nil
}
func Addlier(username string, lierNumber string) error {
	// Lire le contenu du fichier JSON
	data, err := ioutil.ReadFile("database/account.json")
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier JSON : %v", err)
	}

	var users []User

	// Décoder le contenu JSON dans une structure de données
	err = json.Unmarshal(data, &users)
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion du JSON : %v", err)
	}

	// Rechercher l'utilisateur avec le nom d'utilisateur donné
	for i, user := range users {
		if user.Username == username {
			// Ajouter le numéro au tableau de likes de l'utilisateur
			users[i].Lier = append(users[i].Lier, lierNumber)
			break
		}
	}

	// Convertir la structure de données en JSON
	jsonData, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion en JSON : %v", err)
	}

	// Écrire le JSON dans le fichier
	err = ioutil.WriteFile("database/account.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture dans le fichier JSON : %v", err)
	}

	return nil
}

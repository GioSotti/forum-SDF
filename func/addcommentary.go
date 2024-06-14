package Func

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type StoryAllData struct {
	ID          int      `json:"id"`
	Type        string   `json:"type"`
	Image       string   `json:"image"`
	Story       string   `json:"story"`
	Like        int      `json:"like"`
	Lier        int      `json:"lier"`
	Dislike     int      `json:"dislike"`
	Commentaire []string `json:"commentaire`
	Catégorie   []string `json:"catégorie`
}

func AddCommentary(id string, com string) error {
	// Lire le contenu du fichier JSON
	data, err := ioutil.ReadFile("database/data.json")
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier JSON : %v", err)
	}

	var stories []StoryAllData

	// Décoder le contenu JSON dans une structure de données
	err = json.Unmarshal(data, &stories)
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion du JSON : %v", err)
	}

	// Rechercher l'utilisateur avec le nom d'utilisateur donné
	for i, story := range stories {
		Id, _ := strconv.Atoi(id)
		if story.ID == Id {
			// Ajouter le numéro au tableau de likes de l'utilisateur
			stories[i].Commentaire = append(stories[i].Commentaire, com)
			break
		}
	}

	// Convertir la structure de données en JSON
	jsonData, err := json.MarshalIndent(stories, "", "    ")
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion en JSON : %v", err)
	}

	err = ioutil.WriteFile("database/data.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture dans le fichier JSON : %v", err)
	}

	return nil
}

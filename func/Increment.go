package Func

import (
	"encoding/json"
	"io/ioutil"
)

type Data struct {
	ID          int      `json:"id"`
	Type        string   `json:"type"`
	Image       string   `json:"image"`
	Story       string   `json:"story"`
	Like        int      `json:"like"`
	Lier        int      `json:"lier"`
	Dislike     int      `json:"dislike"`
	Commentaire []string `json:"Commentaire"`
	Categorie   []string `json:"Catégorie"`
}

func loadDataFromJSON(filename string) ([]Data, error) {
	// Lire le contenu du fichier JSON
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Décoder les données JSON dans un slice de structures Data
	var StoryData []Data
	err = json.Unmarshal(data, &StoryData)
	if err != nil {
		return nil, err
	}

	return StoryData, nil
}

func saveDataToJSON(filename string, StoryData []Data) error {
	// Encodage des données en JSON
	jsonData, err := json.MarshalIndent(StoryData, "", "  ")
	if err != nil {
		return err
	}

	// Écriture du JSON dans le fichier
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
func IncrementLike(id int) error {
	// Lecture des données à partir du fichier JSON
	StoryData, err := loadDataFromJSON("database/data.json")
	if err != nil {
		return err
	}

	// Recherche de l'élément avec l'ID correspondant
	for i := range StoryData {
		if StoryData[i].ID == id {
			StoryData[i].Like++
			break
		}
	}

	// Sauvegarde des données mises à jour dans le fichier JSON
	err = saveDataToJSON("database/data.json", StoryData)
	if err != nil {
		return err
	}

	return nil
}
func IncrementDislike(id int) error {
	// Lecture des données à partir du fichier JSON
	StoryData, err := loadDataFromJSON("database/data.json")
	if err != nil {
		return err
	}

	// Recherche de l'élément avec l'ID correspondant
	for i := range StoryData {
		if StoryData[i].ID == id {
			StoryData[i].Dislike++
			break
		}
	}

	// Sauvegarde des données mises à jour dans le fichier JSON
	err = saveDataToJSON("database/data.json", StoryData)
	if err != nil {
		return err
	}

	return nil
}
func IncrementLier(id int) error {
	// Lecture des données à partir du fichier JSON
	StoryData, err := loadDataFromJSON("database/data.json")
	if err != nil {
		return err
	}

	// Recherche de l'élément avec l'ID correspondant
	for i := range StoryData {
		if StoryData[i].ID == id {
			StoryData[i].Lier++
			break
		}
	}

	// Sauvegarde des données mises à jour dans le fichier JSON
	err = saveDataToJSON("database/data.json", StoryData)
	if err != nil {
		return err
	}

	return nil
}
func DecrementLike(id int) error {
	// Lecture des données à partir du fichier JSON
	StoryData, err := loadDataFromJSON("database/data.json")
	if err != nil {
		return err
	}

	// Recherche de l'élément avec l'ID correspondant
	for i := range StoryData {
		if StoryData[i].ID == id {
			StoryData[i].Like--
			break
		}
	}

	// Sauvegarde des données mises à jour dans le fichier JSON
	err = saveDataToJSON("database/data.json", StoryData)
	if err != nil {
		return err
	}

	return nil
}
func DecrementDislike(id int) error {
	// Lecture des données à partir du fichier JSON
	StoryData, err := loadDataFromJSON("database/data.json")
	if err != nil {
		return err
	}

	// Recherche de l'élément avec l'ID correspondant
	for i := range StoryData {
		if StoryData[i].ID == id {
			StoryData[i].Dislike--
			break
		}
	}

	// Sauvegarde des données mises à jour dans le fichier JSON
	err = saveDataToJSON("database/data.json", StoryData)
	if err != nil {
		return err
	}

	return nil
}
func DecrementLier(id int) error {
	// Lecture des données à partir du fichier JSON
	StoryData, err := loadDataFromJSON("database/data.json")
	if err != nil {
		return err
	}

	// Recherche de l'élément avec l'ID correspondant
	for i := range StoryData {
		if StoryData[i].ID == id {
			StoryData[i].Lier--
			break
		}
	}

	// Sauvegarde des données mises à jour dans le fichier JSON
	err = saveDataToJSON("database/data.json", StoryData)
	if err != nil {
		return err
	}

	return nil
}

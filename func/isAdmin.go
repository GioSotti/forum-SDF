package Func

import (
	"encoding/json"
	"io/ioutil"
)

func UserInAdminAccounts(username string) bool {
	filePath := "database/adminAccount.json"

	// Lire le contenu du fichier JSON
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false
	}

	// Définir la structure du fichier JSON
	var adminAccounts []User

	// Analyser le contenu JSON dans la structure
	err = json.Unmarshal(fileContent, &adminAccounts)
	if err != nil {
		return false
	}

	// Vérifier si l'utilisateur actuel est présent dans la liste des comptes administrateur
	for _, admin := range adminAccounts {
		if admin.Username == username {
			return true
		}
	}

	return false
}

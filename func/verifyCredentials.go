package Func

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func LoadUsersFromFile(filename string) ([]User, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func VerifyCredentials(username, password string) (bool, string) {
	users, err := LoadUsersFromFile("./database/account.json")
	if err != nil {
		// Gérer l'erreur lors du chargement des utilisateurs depuis account.json
		fmt.Println("Erreur lors du chargement des utilisateurs:", err)
		return false, "Une erreur s'est produite lors de la vérification des identifiants."
	}

	adminUsers, err := LoadUsersFromFile("./database/adminAccount.json")
	if err != nil {
		// Gérer l'erreur lors du chargement des utilisateurs depuis adminAccount.json
		fmt.Println("Erreur lors du chargement des utilisateurs administrateurs:", err)
		return false, "Une erreur s'est produite lors de la vérification des identifiants."
	}

	for _, user := range users {
		if (user.Username == username || user.Email == username) && user.Password == password {
			return true, ""
		}
	}

	for _, adminUser := range adminUsers {
		if (adminUser.Username == username || adminUser.Email == username) && adminUser.Password == password {
			return true, ""
		}
	}

	return false, "Les identifiants fournis sont incorrects."
}

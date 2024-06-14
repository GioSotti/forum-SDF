package Func

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Like     []string `json:"like"`
	Dislike  []string `json:"dislike"`
	Lier     []string `json:"lier"`
	PostList []string `json:"postlist"`
	ban      bool     `json:"ban"`
}

var userIDCounter int = 0

func NewUser(username, password, email string) User {
	return User{
		ID:       generateUserID(),
		Username: username,
		Password: password,
		Email:    email,
		Like:     []string{},
		Dislike:  []string{},
		Lier:     []string{},
		PostList: []string{},
		ban:      false,
	}
}

func generateUserID() int {
	file, err := os.Open("./database/account.json")
	if err != nil {
	}
	defer file.Close()

	// Lire les données du fichier JSON "account"
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
	}

	// Convertir les données JSON en une liste d'utilisateurs
	var users []User
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
	}

	adminFile, err := os.Open("./database/adminAccount.json")
	if err != nil {
	}
	defer adminFile.Close()

	adminJsonData, err := ioutil.ReadAll(adminFile)
	if err != nil {
	}

	// Convertir les données JSON en une liste d'utilisateurs administrateurs
	var adminUsers []User
	err = json.Unmarshal(adminJsonData, &adminUsers)
	if err != nil {
	}

	// Mettre à jour le compteur d'ID en fonction du nombre de comptes existants
	userIDCounter = len(users) + len(adminUsers)

	userIDCounter++
	return userIDCounter
}

func RegisterUser(username, password, confirmPassword, email, adminKey string) error {
	if password != confirmPassword {
		return fmt.Errorf("Le mot de passe et la confirmation du mot de passe ne correspondent pas")
	}

	file, err := os.Open("./database/account.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Lire les données du fichier JSON "account"
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// Convertir les données JSON en une liste d'utilisateurs
	var users []User
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		return err
	}

	// Vérifier si le nom d'utilisateur existe déjà dans "account" ou "adminAccount"
	for _, user := range users {
		if user.Username == username {
			return fmt.Errorf("Le nom d'utilisateur est déjà pris")
		}
		if user.Email == email {
			return fmt.Errorf("L'adresse e-mail est déjà utilisée")
		}
	}

	// Vérifier si l'adresse e-mail est valide
	if !isEmailValid(email) {
		return fmt.Errorf("L'adresse e-mail n'est pas valide")
	}

	if HashPassword(adminKey) == "967520ae23e8ee14888bae72809031b98398ae4a636773e18fff917d77679334" {
		adminFile, err := os.Open("./database/adminAccount.json")
		if err != nil {
			return err
		}
		defer adminFile.Close()

		adminJsonData, err := ioutil.ReadAll(adminFile)
		if err != nil {
			return err
		}

		// Convertir les données JSON en une liste d'utilisateurs administrateurs
		var adminUsers []User
		err = json.Unmarshal(adminJsonData, &adminUsers)
		if err != nil {
			return err
		}

		// Vérifier si le nom d'utilisateur existe déjà dans "adminAccount"
		for _, user := range adminUsers {
			if user.Username == username {
				return fmt.Errorf("Le nom d'utilisateur est déjà pris dans la liste des administrateurs")
			}
			if user.Email == email {
				return fmt.Errorf("L'adresse e-mail est déjà utilisée dans la liste des administrateurs")
			}
		}

		newUser := NewUser(username, password, email)
		adminUsers = append(adminUsers, newUser)

		adminNewData, err := json.Marshal(adminUsers)
		if err != nil {
			return err
		}

		// Écrire les données JSON dans le fichier "adminAccount"
		err = ioutil.WriteFile("./database/adminAccount.json", adminNewData, 0644)
		if err != nil {
			return err
		}

	} else {
		adminFile, err := os.Open("./database/adminAccount.json")
		if err != nil {
			return err
		}
		defer adminFile.Close()

		adminJsonData, err := ioutil.ReadAll(adminFile)
		if err != nil {
			return err
		}

		// Convertir les données JSON en une liste d'utilisateurs administrateurs
		var adminUsers []User
		err = json.Unmarshal(adminJsonData, &adminUsers)
		if err != nil {
			return err
		}

		for _, user := range adminUsers {
			if user.Username == username {
				return fmt.Errorf("Le nom d'utilisateur est déjà pris dans la liste des administrateurs")
			}
		}

		// Vérifier si le nom d'utilisateur existe déjà dans "account"
		for _, user := range users {
			if user.Username == username {
				return fmt.Errorf("Le nom d'utilisateur est déjà pris")
			}
			if user.Email == email {
				return fmt.Errorf("L'adresse e-mail est déjà utilisée")
			}
		}

		newUser := NewUser(username, password, email)
		users = append(users, newUser)

		newData, err := json.Marshal(users)
		if err != nil {
			return err
		}

		// Écrire les données JSON dans le fichier "account"
		err = ioutil.WriteFile("./database/account.json", newData, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func isEmailValid(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		return false
	}
	return match
}

package Func

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func DelLike(username string, like string) error {
	data, err := ioutil.ReadFile("database/account.json")
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %v", err)
	}

	var users []User

	err = json.Unmarshal(data, &users)
	if err != nil {
		return fmt.Errorf("failed to decode JSON data: %v", err)
	}

	for i, user := range users {
		if user.Username == username {
			index := -1
			for j, existingLike := range user.Like {
				if existingLike == like {
					index = j
					break
				}
			}

			if index >= 0 {
				users[i].Like = append(user.Like[:index], user.Like[index+1:]...)
			}

			break
		}
	}

	updatedData, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to encode JSON data: %v", err)
	}

	err = ioutil.WriteFile("database/account.json", updatedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %v", err)
	}

	return nil
}
func DelDislike(username string, dislike string) error {
	data, err := ioutil.ReadFile("database/account.json")
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %v", err)
	}

	var users []User

	err = json.Unmarshal(data, &users)
	if err != nil {
		return fmt.Errorf("failed to decode JSON data: %v", err)
	}

	for i, user := range users {
		if user.Username == username {
			index := -1
			for j, existingDisLike := range user.Dislike {
				if existingDisLike == dislike {
					index = j
					break
				}
			}

			if index >= 0 {
				users[i].Dislike = append(user.Dislike[:index], user.Dislike[index+1:]...)
			}

			break
		}
	}

	updatedData, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to encode JSON data: %v", err)
	}

	err = ioutil.WriteFile("database/account.json", updatedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %v", err)
	}

	return nil
}
func DelLier(username string, Lier string) error {
	data, err := ioutil.ReadFile("database/account.json")
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %v", err)
	}

	var users []User

	err = json.Unmarshal(data, &users)
	if err != nil {
		return fmt.Errorf("failed to decode JSON data: %v", err)
	}

	for i, user := range users {
		if user.Username == username {
			index := -1
			for j, existingLier := range user.Lier {
				if existingLier == Lier {
					index = j
					break
				}
			}

			if index >= 0 {
				users[i].Lier = append(user.Lier[:index], user.Lier[index+1:]...)
			}

			break
		}
	}

	updatedData, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to encode JSON data: %v", err)
	}

	err = ioutil.WriteFile("database/account.json", updatedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %v", err)
	}

	return nil
}

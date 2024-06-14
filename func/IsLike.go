package Func

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func IsLike(username string, like string) (bool, error) {
	data, err := ioutil.ReadFile("database/account.json")
	if err != nil {
		return false, fmt.Errorf("failed to read JSON file: %v", err)
	}

	var users []User

	err = json.Unmarshal(data, &users)
	if err != nil {
		return false, fmt.Errorf("failed to decode JSON data: %v", err)
	}

	for _, user := range users {
		if user.Username == username {
			for _, existingLike := range user.Like {
				if existingLike == like {
					return true, nil
				}
			}
			break
		}
	}

	return false, nil
}
func IsDislike(username string, dislike string) (bool, error) {
	data, err := ioutil.ReadFile("database/account.json")
	if err != nil {
		return false, fmt.Errorf("failed to read JSON file: %v", err)
	}

	var users []User

	err = json.Unmarshal(data, &users)
	if err != nil {
		return false, fmt.Errorf("failed to decode JSON data: %v", err)
	}

	for _, user := range users {
		if user.Username == username {
			for _, existingDisLike := range user.Dislike {
				if existingDisLike == dislike {
					return true, nil
				}
			}
			break
		}
	}

	return false, nil
}
func IsLier(username string, lier string) (bool, error) {
	data, err := ioutil.ReadFile("database/account.json")
	if err != nil {
		return false, fmt.Errorf("failed to read JSON file: %v", err)
	}

	var users []User

	err = json.Unmarshal(data, &users)
	if err != nil {
		return false, fmt.Errorf("failed to decode JSON data: %v", err)
	}

	for _, user := range users {
		if user.Username == username {
			for _, existingLier := range user.Lier {
				if existingLier == lier {
					return true, nil
				}
			}
			break
		}
	}

	return false, nil
}

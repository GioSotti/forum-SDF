package Func

import (
	"crypto/rand"
	"encoding/base64"
)

func CreateSession() string {
	// Générer un tableau de bytes aléatoires pour l'identifiant de session
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	// Encoder les bytes aléatoires en une chaîne Base64 pour l'identifiant de session
	sessionID := base64.StdEncoding.EncodeToString(b)

	return sessionID
}

package server

import (
	Func "ForumSdf/func"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/sessions"
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

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Like     []int  `json:"like"`
	Dislike  []int  `json:"dislike"`
	Lier     []int  `json:"lier"`
	PostList []int  `json:"postlist"`
	ban      bool   `json:"ban"`
}

type UserSession struct {
	Session   *sessions.Session // Session de l'utilisateur
	Connected bool              // Indique si l'utilisateur est connecté ou non
	Username  string            // Nom d'utilisateur de l'utilisateur connecté

}

var StoryData []StoryAllData
var userSessions map[string]*UserSession
var store = sessions.NewCookieStore([]byte("your-secret-key"))
var UserData []User

func LoadDataFromJSON(filename string) error {
	// Lire le contenu du fichier JSON
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		// Vérifier si le fichier est vide
		if os.IsNotExist(err) {
			// Créer un tableau vide pour StoryData
			StoryData = []StoryAllData{}
			return nil
		}
		return err
	}

	// Décoder les données JSON dans la variable StoryData
	err = json.Unmarshal(data, &StoryData)
	if err != nil {
		return err
	}

	return nil
}
func LoadUserFromJSON(filename string) ([]User, error) {
	// Lire le contenu du fichier JSON
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		// Vérifier si le fichier est vide
		if os.IsNotExist(err) {
			// Créer un tableau vide pour StoryData
			UserData = []User{}
			return UserData, err
		}
		return UserData, err
	}

	// Décoder les données JSON dans la variable StoryData
	err = json.Unmarshal(data, &UserData)
	if err != nil {
		return UserData, err
	}

	return UserData, nil
}

func filterStories(stories []StoryAllData, searchTerm string) []StoryAllData {
	var filtered []StoryAllData

	for _, story := range stories {
		if strings.Contains(story.Story, searchTerm) {
			filtered = append(filtered, story)
		}
	}

	return filtered
}

func filterByCategory(stories []StoryAllData, category string) []StoryAllData {
	var filtered []StoryAllData

	for _, story := range stories {
		for _, cat := range story.Catégorie {
			if cat == category {
				filtered = append(filtered, story)
				break
			}
		}
	}

	return filtered
}
func filteredById(stories []StoryAllData, id int) []StoryAllData {
	var filtered []StoryAllData

	for _, story := range stories {
		if story.ID == id {
			filtered = append(filtered, story)
		}
	}

	return filtered
}
func filtreDecroissant(stories []StoryAllData) []StoryAllData {
	// Sort the stories by ID in descending order
	sort.Slice(stories, func(i, j int) bool {
		return stories[i].ID > stories[j].ID
	})

	return stories
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Récupérer la session de l'utilisateur
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Supprimer la session de l'utilisateur
	delete(userSessions, session.ID)

	// Détruire la session
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Rediriger l'utilisateur vers une autre page (par exemple, la page d'accueil)
	http.Redirect(w, r, "/index", http.StatusFound)
}
func DisplayAccueil(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		SubmitLike(w, r)
		SubmitCommentary(w, r)
	}

	err := LoadDataFromJSON("./database/data.json")
	if err != nil {
		println("Erreur lors du chargement des données :", err.Error())
		return
	}

	custTemplate, err := template.ParseFiles("./template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	searchTerm := r.FormValue("search")
	selectedCategory := r.FormValue("dropdown")

	// Filtrer les histoires si le terme de recherche n'est pas vide
	filteredStories := StoryData
	if searchTerm != "" {
		filteredStories = filterStories(StoryData, searchTerm)
	}
	if selectedCategory != "" {
		filteredStories = filterByCategory(filteredStories, selectedCategory)
	}
	filteredStories = filtreDecroissant(filteredStories)

	// Vérifier si l'utilisateur est connecté
	connected := false
	username := ""
	// Récupérer la session de l'utilisateur
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est déjà connecté
	userSession, exists := userSessions[session.ID]
	if exists && userSession.Connected {
		connected = true
		username = userSession.Username
	}

	Data := struct {
		Connected bool
		Username  string
		Stories   []StoryAllData
	}{
		Connected: connected,
		Username:  username,
		Stories:   filteredStories,
	}

	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func DisplayConnexion(w http.ResponseWriter, r *http.Request) {
	custTemplate, err := template.ParseFiles("./template/connexion.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupérer la session de l'utilisateur
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Vérifier si l'utilisateur est déjà connecté
	userSession, exists := userSessions[session.ID]
	if exists && userSession.Connected {
		// L'utilisateur est déjà connecté, on le rediriges vers une autre page
		http.Redirect(w, r, "/index", http.StatusFound)
		return
	}

	if r.Method == "POST" {
		// Récupérer les données du formulaire
		username := r.FormValue("username")
		password := Func.HashPassword(r.FormValue("password"))
		isValid, errorMessage := Func.VerifyCredentials(username, password)
		if isValid {
			// Les identifiants sont valides, on crée nouvelle session utilisateur
			newSession, err := store.New(r, "session-name")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Mettre à jour les informations de session utilisateur
			userSession = &UserSession{
				Session:   newSession,
				Connected: true,
				Username:  username,
			}
			userSessions[newSession.ID] = userSession

			// Rediriger l'utilisateur vers la page d'accueil
			http.Redirect(w, r, "/index", http.StatusFound)
			return
		} else {
			data := struct {
				ErrorMessage string
			}{
				ErrorMessage: errorMessage,
			}
			// Les identifiants sont invalides
			err := custTemplate.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	} else {
		// Afficher la page de connexion
		err := custTemplate.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
func DisplayInscription(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/inscription.html"))

	if r.Method == http.MethodPost {
		// Récupérer les données du formulaire
		username := r.FormValue("username")
		password := Func.HashPassword(r.FormValue("password"))
		confirmPassword := Func.HashPassword(r.FormValue("confirm-password"))
		email := r.FormValue("email")
		adminKey := r.FormValue("admin-key")

		// Appeler la fonction RegisterUser pour gérer l'inscription
		err := Func.RegisterUser(username, password, confirmPassword, email, adminKey)
		if err != nil {
			data := struct {
				ErrorMessage string
			}{
				ErrorMessage: err.Error(),
			}
			tmpl.Execute(w, data)
			return

		}

		// Rediriger vers la page de confirmation d'inscription
		http.Redirect(w, r, "/inscription-confirm", http.StatusFound)
		return
	}

	// Afficher le formulaire d'inscription

	tmpl.Execute(w, nil)
}
func SubmitCommentary(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Vérifier si l'utilisateur est connecté
	connected := false
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userSession, exists := userSessions[session.ID]
	if exists && userSession.Connected {
		connected = true
	}
	if !connected {
		http.Redirect(w, r, "/connexion", http.StatusFound)
		return
	}
	// Récupérer les données du commentaire
	com := r.FormValue("message")
	Id := r.FormValue("id")
	Func.AddCommentary(Id, com)

}
func SubmitLike(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Vérifier si l'utilisateur est connecté
	connected := false
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userSession, exists := userSessions[session.ID]
	if exists && userSession.Connected {
		connected = true
	}
	if !connected {
		http.Redirect(w, r, "/connexion", http.StatusFound)
		return
	}
	// Récupérer les données du formulaire
	style := r.FormValue("style")
	Id := r.FormValue("id")

	username := userSession.Username

	Func.Click(username, Id, style)

}
func SubmitFormulaire(w http.ResponseWriter, r *http.Request) {
	// Vérifier si l'utilisateur est connecté
	connected := false
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userSession, exists := userSessions[session.ID]
	if exists && userSession.Connected {
		connected = true
	}
	if !connected {
		http.Redirect(w, r, "/connexion", http.StatusFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err = r.ParseMultipartForm(1 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupérer les données du formulaire
	categories := r.Form["difficulty"]
	story := r.FormValue("story")
	mediaFile, mediaHeader, err := r.FormFile("media")

	if err != nil && err != http.ErrMissingFile {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediaFile != nil {
		defer mediaFile.Close()

		ext := filepath.Ext(mediaHeader.Filename)

		// Créer le nom de fichier unique en utilisant le format "preuve_ID.extension"
		mediaFilename := "preuve_" + strconv.Itoa(len(StoryData)+1) + ext

		// Lire le contenu du fichier média
		mediaData, err := ioutil.ReadAll(mediaFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Enregistrer le fichier média sur le serveur
		err = ioutil.WriteFile(filepath.Join("template", "media", mediaFilename), mediaData, 0644)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mettre à jour l'URL de l'image ou de la vidéo dans l'objet StoryAllData
		var mediaType string
		if strings.HasPrefix(mediaHeader.Header.Get("Content-Type"), "image/") {
			mediaType = "image"
		} else if strings.HasPrefix(mediaHeader.Header.Get("Content-Type"), "video/") {
			mediaType = "video"
		}

		newStory := StoryAllData{
			ID:          len(StoryData) + 1,
			Type:        mediaType,
			Image:       "media/" + mediaFilename,
			Story:       story,
			Like:        0,
			Lier:        0,
			Dislike:     0,
			Commentaire: []string{},
			Catégorie:   categories,
		}

		StoryData = append(StoryData, newStory)
	} else {
		// Pas de fichier média, créer un objet StoryAllData sans URL d'image
		newStory := StoryAllData{
			ID:          len(StoryData) + 1,
			Type:        "",
			Image:       "",
			Story:       story,
			Like:        0,
			Lier:        0,
			Dislike:     0,
			Commentaire: []string{},
			Catégorie:   categories,
		}

		// Ajouter la nouvelle histoire à StoryData
		StoryData = append(StoryData, newStory)
	}

	// Convertir StoryData en JSON
	jsonData, err := json.Marshal(StoryData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enregistrer les données JSON dans le fichier
	err = ioutil.WriteFile("database/data.json", jsonData, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := userSession.Username
	postNumber := strconv.Itoa(len(StoryData))
	Func.AddpostedLike(username, postNumber)

	// Rediriger vers la page d'accueil ou une autre page appropriée
	http.Redirect(w, r, "/index", http.StatusFound)
}

func DisplayFormulaire(w http.ResponseWriter, r *http.Request) {
	// Vérifier si l'utilisateur est connecté
	connected := false
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userSession, exists := userSessions[session.ID]
	if exists && userSession.Connected {
		connected = true
	}

	// Rediriger vers la page de connexion si l'utilisateur n'est pas connecté
	if !connected {
		http.Redirect(w, r, "/connexion", http.StatusFound)
		return
	}

	if r.Method == http.MethodPost {
		SubmitFormulaire(w, r)
		return
	}

	custTemplate, err := template.ParseFiles("./template/formulaire.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Data := StoryData
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func DisplayConfirm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		SubmitFormulaire(w, r)
		return
	}

	custTemplate, err := template.ParseFiles("./template/inscription-confirm.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Data := StoryData
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func DisplayMdpperdu(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		return
	}

	custTemplate, err := template.ParseFiles("./template/mdpperdu.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Data := StoryData
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func DisplayAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Vérifier si l'utilisateur est connecté
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userSession, exists := userSessions[session.ID]
	if !exists || !userSession.Connected {
		http.Redirect(w, r, "/connexion", http.StatusFound)
		return
	}

	// Charger les données de l'utilisateur depuis le fichier JSON
	filename := "database/account.json"
	userData, _ := LoadUserFromJSON(filename)
	if err != nil {
		http.Error(w, "Erreur lors du chargement des données utilisateur", http.StatusInternalServerError)
		return
	}

	// Rechercher l'utilisateur connecté dans les données chargées
	var currentUser User
	for _, user := range userData {
		if user.Username == userSession.Username {
			currentUser = user
			break
		}
	}

	if currentUser.ID == 0 {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	// Afficher les détails du compte utilisateur dans le template
	custTemplate, err := template.ParseFiles("./template/moncompte.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = custTemplate.Execute(w, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func DisplayAdmin(w http.ResponseWriter, r *http.Request) {
	// Vérifier si l'utilisateur est connecté
	connected := false
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userSession, exists := userSessions[session.ID]
	if exists && userSession.Connected {
		connected = true
	}
	username := userSession.Username
	userInAdmin := Func.UserInAdminAccounts(username)
	if !userInAdmin {
		http.Redirect(w, r, "/connexion", http.StatusFound)
		return
	}
	if !connected {
		http.Redirect(w, r, "/connexion", http.StatusFound)
		return
	}
	if r.Method == http.MethodPost {
		return
	}

	custTemplate, err := template.ParseFiles("./template/admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Data := StoryData
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Server() {

	println("server started on http://localhost:8080/index")
	userSessions = make(map[string]*UserSession)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./template/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./template/js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./assets/fonts"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./template/images"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./template/assets"))))
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("./template/media"))))
	http.HandleFunc("/index", DisplayAccueil)
	http.HandleFunc("/connexion", DisplayConnexion)
	http.HandleFunc("/inscription", DisplayInscription)
	http.HandleFunc("/formulaire", DisplayFormulaire)
	http.HandleFunc("/inscription-confirm", DisplayConfirm)
	http.HandleFunc("/mdpperdu", DisplayMdpperdu)
	http.HandleFunc("/account", DisplayAccount)
	http.HandleFunc("/admin", DisplayAdmin)
	http.HandleFunc("/logout", Logout)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

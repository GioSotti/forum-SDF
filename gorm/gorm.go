package gorm
import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
  )
  
  db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/forum?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
  if err != nil {
	println("ereur connexion base de donnee")
  }

  type User struct {
	gorm.Model
	Username     string
	Email        string
	PasswordHash string
  }

  router := mux.NewRouter()

router.HandleFunc("/inscription", func(w http.ResponseWriter, r *http.Request) {
    // Récupérer les données du formulaire
    username := r.FormValue("username")
    email := r.FormValue("email")
    password := r.FormValue("password")
    confirmPassword := r.FormValue("confirm_password")

    // Vérifier que les mots de passe correspondent
    if password != confirmPassword {
	println("erreur mot de passe")
        // Gérer l'erreur de non-correspondance des mots de passe
        // Rediriger l'utilisateur vers une page d'erreur ou de nouveau formulaire, par exemple
        return
    }

    // Créer une instance de modèle User avec les données du formulaire
    user := User{
        Username: username,
        Email:    email,
        // Vous devrez probablement hacher le mot de passe avant de le stocker dans la base de données
        PasswordHash: hashPassword(password),
    }

    // Insérer l'utilisateur dans la base de données à l'aide de GORM
    result := db.Create(&user)
    if result.Error != nil {
        // Gérer l'erreur d'insertion dans la base de données
        // Rediriger l'utilisateur vers une page d'erreur ou de nouveau formulaire, par exemple
        return
    }

    // Rediriger l'utilisateur vers une page de succès ou une autre page appropriée
    http.Redirect(w, r, "/#", http.StatusFound)
}).Methods("POST")
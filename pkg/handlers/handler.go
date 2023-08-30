package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	validUsername = "karthik"
	validPassword = "password"
	store         = sessions.NewFilesystemStore("", []byte("secret-key"))
)

// login page handler
func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.page.html")
	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

// login handler
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	session, _ := store.Get(r, "session")
	// log.Println(username, password)
	if username == validUsername && password == validPassword {

		session.Values["authenticated"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else if username == validUsername && password != validPassword {

		session.Values["authenticated"] = false
		session.Save(r, w)

		tmpl, err := template.ParseFiles("templates/login.page.html")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, "Mr "+username+" name and password is incorrect")
	} else {
		session, _ = store.Get(r, "session")

		tmpl, err := template.ParseFiles("templates/login.page.html")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, "Mr/Miss "+username+" the user name and password is incorrect")
	}
}

// home page handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl, err := template.ParseFiles("templates/home.page.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// logout handler
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

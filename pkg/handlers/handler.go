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
	SetNoCacheHeaders(w)
	tmpl, err := template.ParseFiles("templates/login.page.html")
	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

// login handler
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	SetNoCacheHeaders(w)
	r.ParseForm()
	// go func() {
	// 	ctx := r.Context()
	// 	<-ctx.Done()
	// 	log
	// }()
	// time.Sleep(10 * time.Second)
	//
	// w.Header().Add()
	username := r.FormValue("username")
	password := r.FormValue("password")
	session, _ := store.Get(r, "session")
	// log.Println(username, password)
	// session.Options = &sessions.Options{MaxAge: -1}

	if username == validUsername && password == validPassword {

		session.Values["authenticated"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/home", http.StatusSeeOther)

	} else {

		session.Save(r, w)

		tmpl, err := template.ParseFiles("templates/login.page.html")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		var errorMessage string
		if username == validUsername {
			errorMessage = "The password is incorrect"
		} else {
			errorMessage = "The username and password are incorrect"
		}

		tmpl.Execute(w, "Mr "+username+" "+errorMessage)

	}
}

// home page handler
func HomePage(w http.ResponseWriter, r *http.Request) {

	SetNoCacheHeaders(w)
	// ... rest of the code

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

// disable caching
func SetNoCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

// logout handler
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	SetNoCacheHeaders(w)
	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusMovedPermanently)

}

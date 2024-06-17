package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
var users = map[string]string{"pranav": "radar", "admin": "password"}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	if session.Values["authenticated"] != nil && session.Values["authenticated"] != false {
		w.Write([]byte(time.Now().String()))
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	fmt.Println(users[username])
	if originalPassword, ok := users[username]; ok {
		if password == originalPassword {
			session.Values["authenticated"] = true
			session.Save(r, w)
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("Logged in!"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/health_check", HealthCheckHandler)
	r.HandleFunc("/logout", LogoutHandler)

	srvr := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      r,
	}

	srvr.ListenAndServe()
}

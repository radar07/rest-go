package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

var users = map[string]string{"pranav": "radar"}

var secret = []byte("mysecretkey")

type Response struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

// func validateJWT(t string) (*jwt.Token, error) {
// 	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("invalid signing method: %v", t.Header["alg"])
// 		}
// 		return []byte(secret), nil
// 	})
// }

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get("Authorization")
	tokenString := strings.Split(bearerToken, " ")[1]
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		w.Write([]byte("Access Denied"))
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		response := map[string]string{"user": claims["username"].(string), "time": time.Now().String()}
		responseJSON, _ := json.Marshal(response)
		w.Header().Add("Content-Type", "application/json")
		w.Write(responseJSON)
	}
}

func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "could not able to parse", http.StatusBadRequest)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if originalPassword, ok := users[username]; ok {
		if originalPassword == password {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username":  username,
				"ExpiresAt": 100,
			})
			tokenString, err := token.SignedString(secret)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusBadGateway)
			}

			response := Response{Token: tokenString, Status: "success"}
			responseJSON, _ := json.Marshal(response)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)
			w.WriteHeader(http.StatusAccepted)
			return
		} else {
			w.Write([]byte(secret))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	} else {
		w.Write([]byte("Not found"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getToken", getTokenHandler)
	r.HandleFunc("/healthcheck", healthCheckHandler)

	srvr := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srvr.ListenAndServe())
}

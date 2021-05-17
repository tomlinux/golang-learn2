package helper

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func Add(n1, n2 int) int {
	return n1 + n2
}

func CheckAuth2(w http.ResponseWriter, r http.Request) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	return pair[0] == "admin" && pair[1] == "123456"

}

func Check(u, p string) bool {
	if u == "admin" && p == "123456" {
		return true
	} else {
		return false
	}
}

func CheckAuth(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()

		log.Println(user, pass)
		if !Check(user, pass) {
			http.Error(w, "Unauthorized.", 401)
			return
		}
		//username, password, ok := r.BasicAuth()
		//log.Println(username,password)
		//if !ok {
		//	log.Println("Error parsing basic auth")
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}
		//
		//if username != "admin" {
		//	log.Println("Username is not correct")
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}
		//
		//if password != "123456" {
		//	log.Println("Password is not correct")
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}
		fn(w, r)
	}
}

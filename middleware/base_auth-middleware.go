package middleware

import (
	"log"
	"net/http"
)

type BasicAuthMiddleware struct {
	Next http.Handler
}

// 实现这个接口	ServeHTTP(ResponseWriter, *Request)

func (b *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if b.Next == nil {
		b.Next = http.DefaultServeMux
	}
	username, password, ok := r.BasicAuth()
	if !ok {
		log.Println("Error parsing basic auth")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if username != "admin" {
		log.Println("Username is not correct")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}


	if password != "123456" {
		log.Println("Password is not correct")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	b.Next.ServeHTTP(w,r)
}

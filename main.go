package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
	)
}

func main() {

	port, envExists := os.LookupEnv("IPAPP_PORT")
	if envExists {
		port = ":" + port
	} else {
		port = ":80"
	}

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		w.Write([]byte(ip + "\n"))
	})

	log.Fatal(http.ListenAndServe(port, router))

}

package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	port, envExists := os.LookupEnv("IPAPP_PORT")
	if envExists {
		port = ":" + port
	} else {
		port = ":80"
	}

	router.GET("/*any", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		w.Write([]byte(ip + "\n"))
	})

	log.Printf("Starting app on %s", port)
	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(port, handlers.RecoveryHandler()(loggedRouter)))

}

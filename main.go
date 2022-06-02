package main

import (
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"github.com/nanih98/golang-rest-api/router"
	"github.com/sirupsen/logrus"
)

func main() {
	router := router.ConfigureRouter()

	// srv := &http.Server{
	//     Handler:      router,
	//     Addr:         ":3000",
	//     // Good practice: enforce timeouts for servers you create!
	//     WriteTimeout: 15 * time.Second,
	//     ReadTimeout:  15 * time.Second,
	// }
	addr := ":3000"

	logrus.WithField("addr", addr).Info("Starting server...")

	if err := http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, router)); err != nil {
		logrus.WithField("event", "start server").Fatal(err)
	}
	//log.Fatal(http.ListenAndServe(":3000", router))
}
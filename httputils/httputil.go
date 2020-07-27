package httputils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ServeHTTPS inicializa escuta e serve HTTP/2
func ServeHTTPS() {

	r := mux.NewRouter()
	r.HandleFunc("/", rootHTTPHandler)

	tlsConfig := &tls.Config{}
	tlsConfig.NextProtos = []string{"h2"}
	tlsConfig.MinVersion = tls.VersionTLS12

	fmt.Print("HTTPS: Servindo HTTP/2...")

	srv := &http.Server{
		Addr:      ":8443",
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	log.Fatal(srv.ListenAndServeTLS(
		"/home/f8272397/certs/cert.pem",
		"/home/f8272397/certs/key.pem",
	))
}

// ServeHTTP inicializa escuta e serve HTTP/2
func ServeHTTP() {

	r := mux.NewRouter()
	r.HandleFunc("/", rootHTTPHandler)

	tlsConfig := &tls.Config{}
	tlsConfig.NextProtos = []string{"h2"}

	fmt.Print("HTTP: Servindo HTTP/2...")

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	log.Fatal(srv.ListenAndServe())
}

func rootHTTPHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "protocolo: %s", req.Proto)
}

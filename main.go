package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var (
	host = flag.String("host", "0.0.0.0", "server host")
	port = flag.String("port", "8080", "server port")
)

func main() {
	flag.Parse()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.NotFound(w, r)
			return
		}

		logfile := "/tmp/" + strings.TrimPrefix(r.URL.Path, "/") + ".log"

		// GET /:file - Serve log file
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, logfile)
			return
		}

		// DELETE /:file - Delete log file
		if r.Method == http.MethodDelete {
			err := os.Remove(logfile)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "%v\n", err)
				return
			}
			fmt.Fprintf(w, "ok\n")
			return
		}

		// POST/PUT /:file - Appen to log file
		f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		logger := log.New(f, "", log.LstdFlags)

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%v\n", err)
			return
		}
		logger.Printf("%s\n", data)
	})

	addr := net.JoinHostPort(*host, *port)
	log.Println("Listening on", addr)
	panic(http.ListenAndServe(addr, handler))
}

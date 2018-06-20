package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	host    = flag.String("host", "0.0.0.0", "server host")
	port    = flag.String("port", "8080", "server port")
	logfile = flag.String("f", "/tmp/mirror.log", "file path for log file")
)

func main() {
	flag.Parse()

	f, err := os.Create(*logfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, *logfile)
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Println("got err", err)
		}
		logger.Printf("[%s] %s\n", r.Method, data)
	})
	addr := net.JoinHostPort(*host, *port)
	logger.Println("Listening on", addr)
	http.ListenAndServe(addr, handler)
}

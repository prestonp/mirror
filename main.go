package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var (
	host = flag.String("host", "0.0.0.0", "server host")
	port = flag.String("port", "8080", "server port")
)

func main() {
	flag.Parse()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("got err", err)
		}
		log.Printf("[%s] %s\n", r.Method, data)
	})
	addr := net.JoinHostPort(*host, *port)
  log.Println("Listening on", addr)
	http.ListenAndServe(addr, handler)
}

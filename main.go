package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func createReverseProxy(target string) (http.Handler, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return proxy, nil
}

func main() {
	var port int
	flag.IntVar(&port, "p", 9999, "port to serve on")
	flag.Parse()

	target := "http://127.0.0.1:55235"
	proxy, err := createReverseProxy(target)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", proxy)

	log.Printf("Starting reverse proxy server on port %d, forwarding requests to %s", port, target)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}

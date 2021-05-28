package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	host = "localhost"
	uri  = "api/_health"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <port>\n", os.Args[0])
	}

	port := os.Args[1]
	endpoint := fmt.Sprintf("http://%s:%s/%s", host, port, uri)

	log.Infof("Querying %s", endpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		log.Fatal("Status code returned was ", resp.StatusCode)
	}
}

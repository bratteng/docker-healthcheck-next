package main

import (
	"flag"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var Flags struct {
	https bool
	port  string
	host  string
	path  string
}

func main() {
	flag.BoolVar(&Flags.https, "https", false, "Connect via HTTPS")
	flag.StringVar(&Flags.port, "port", "3000", "Which port to connect to")
	flag.StringVar(&Flags.host, "host", "localhost", "Which host to connect to")
	flag.StringVar(&Flags.path, "path", "api/_health", "Which path to check")

	flag.Parse()

	protocol := "http"
	if Flags.https {
		protocol = "https"
	}

	endpoint := fmt.Sprintf("%s://%s:%s/%s", protocol, Flags.host, Flags.port, Flags.path)

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

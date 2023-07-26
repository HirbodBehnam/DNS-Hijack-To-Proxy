package main

import (
	"AllToOneDns/server"
	"os"
)
import log "github.com/sirupsen/logrus"

func main() {
	log.SetLevel(log.TraceLevel)
	if len(os.Args) != 2 {
		log.Fatalln("Pass local address of current computer (the one you want to DNS queries to point to) as the first argument.")
	}
	go server.SetupDNSServer(os.Args[1])
	go server.SetupHTTPForwarder()
	go server.SetupTLSForwarder()
	select {} // wait...
}

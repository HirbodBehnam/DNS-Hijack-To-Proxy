package main

import "AllToOneDns/server"
import log "github.com/sirupsen/logrus"

func main() {
	log.SetLevel(log.TraceLevel)
	go server.SetupDNSServer()
	go server.SetupHTTPForwarder()
	go server.SetupTLSForwarder()
	select {} // wait...
}

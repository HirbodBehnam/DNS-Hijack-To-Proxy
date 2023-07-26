package server

import (
	"AllToOneDns/proxy"
	"net"

	log "github.com/sirupsen/logrus"
)

func handleTLSConnection(c net.Conn) {
	defer c.Close()
	// Read the client hello
	buffer := make([]byte, 1024) // TLS client hello is small
	n, err := c.Read(buffer)
	if err != nil {
		log.Println("cannot read client hello:", err)
		return
	}
	// Get servername
	hostname, err := getHostname(buffer[:n])
	if err != nil {
		log.Println("cannot parse client hello:", err)
		return
	}
	// Resolve hostname
	addr, err := net.LookupIP(hostname)
	if err != nil || len(addr) == 0 {
		log.Println("cannot resolve ip")
		return
	}
	log.WithField("hostname", hostname).Trace("TLS estabilished")
	// Proxy
	proxy.ProxyOverSocks5(c, buffer[:n], addr[0].String()+":443")
}

// SetupTLSForwarder starts a TLS serever on port 443
func SetupTLSForwarder() {
	l, err := net.Listen("tcp", ":443")
	if err != nil {
		log.Fatalln("cannot start the tls forwarder listener:", err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("Cannot accept connection:", err)
		}
		go handleTLSConnection(c)
	}
}

package server

import (
	"AllToOneDns/proxy"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

func getHTTPHost(packet []byte) (string, error) {
	const prefix = "Host: "
	scanner := bufio.NewScanner(bytes.NewReader(packet))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, prefix) {
			return line[len(prefix):], nil
		}
	}
	return "", errors.New("host not found")
}

func handleHTTPConnection(c net.Conn) {
	defer c.Close()
	// Read the client hello
	buffer := make([]byte, 32*1024) // HTTP client hello could be large
	n, err := c.Read(buffer)
	if err != nil {
		log.Println("cannot read client hello:", err)
		return
	}
	// Get servername
	hostname, err := getHTTPHost(buffer[:n])
	if err != nil {
		log.Println("cannot find host:", err)
		return
	}
	// Resolve hostname
	addr, err := net.LookupIP(hostname)
	if err != nil || len(addr) == 0 {
		log.Println("cannot resolve ip")
		return
	}
	fmt.Println("proxy http", hostname)
	// Proxy
	proxy.ProxyOverSocks5(c, buffer[:n], addr[0].String()+":80")
}

func SetupHTTPForwarder() {
	l, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("cannot start the tcp forwarder listener:", err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("Cannot accept connection:", err)
		}
		go handleHTTPConnection(c)
	}
}

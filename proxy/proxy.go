package proxy

import (
	"io"
	"log"
	"net"
	"sync"
	"time"

	netProxy "golang.org/x/net/proxy"
)

// ProxyOverSocks5 will start a
func ProxyOverSocks5(main net.Conn, firstPacket []byte, dest string) {
	// Start the proxy
	dialer, err := netProxy.SOCKS5("tcp", "127.0.0.1:10808", nil, &net.Dialer{
		Timeout:   60 * time.Second,
		KeepAlive: 30 * time.Second,
	})
	if err != nil {
		log.Println("cannot initail socks5 connection:", err)
		return
	}
	proxiedConnection, err := dialer.Dial("tcp", dest)
	if err != nil {
		log.Println("cannot dial proxy:", err)
		return
	}
	defer proxiedConnection.Close()
	// Send the first packet
	proxiedConnection.Write(firstPacket)
	proxyConnection(main, proxiedConnection)
}

func proxyConnection(a, b net.Conn) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		io.Copy(a, b)
		wg.Done()
	}()
	go func() {
		io.Copy(b, a)
		wg.Done()
	}()

	wg.Wait()
}

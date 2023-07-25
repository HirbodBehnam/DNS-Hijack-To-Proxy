package server

import (
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

var thisIP = net.ParseIP("192.168.1.100")

type dnsHandler struct{}

func (dnsHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		fmt.Println("resolved", domain)
		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
			A:   thisIP,
		})
	}
	w.WriteMsg(&msg)
}

func SetupDNSServer() {
	srv := dns.Server{Addr: listenAddress + ":53", Net: "udp"}
	srv.Handler = dnsHandler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}

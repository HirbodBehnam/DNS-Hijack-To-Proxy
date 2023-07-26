# DNS Hijack To Proxy

Hijack DNS queries to redirect them to proxy all of its connections through a SOCKS5 server.

## Why?

Suppose that you have a device which does not accept HTTP nor SOCKS5 proxies. So what you should do if you wanted to
pass device's traffic through a proxy? One way to do it is to share a VPN connection over a hotspot. Doing such with an
android devices requires root access.

Instead, we can trick the device; If we can set the DNS server address to a server which always returns a specific
address for all queries; Thus, all requests are redirected to that server. The HTTP(S) server consumes the first packet
and determines the destination address of the packet based on `Host` header or `SNI` of TLS hello and the traffic is
forwarded over a SOCKS5 proxy.

This is how this app works. It currently has 3 components:

1. DNS server
2. HTTP forwarder server
3. HTTPS forwarder server

## Building

To build the app, simply just clone this repo and run `go build` to build the app.

### Running

To run the app, you need to supply the program with the IP which will be answered for all DNS queries. Supply this IP as
the first argument of the program.
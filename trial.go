package grpc

import "google.golang.org/grpc/internal/resolver/dns"

func NewFormatIp(addr string) (string, error) {
	host, err := dns.FormatIP(addr)
	return host, err
}
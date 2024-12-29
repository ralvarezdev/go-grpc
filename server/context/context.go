package context

import (
	"context"
	"google.golang.org/grpc/peer"
	"net"
	"strings"
)

// GetClientIP extracts the client IP address from the context
func GetClientIP(ctx context.Context) (string, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return "", ErrFailedToGetPeerFromContext
	}

	// Get the IP address from the peer address
	addr := p.Addr.String()
	ip, _, err := net.SplitHostPort(addr)
	if err != nil {
		return "", err
	}

	// Remove any surrounding brackets from IPv6 addresses
	ip = strings.Trim(ip, "[]")

	return ip, nil
}

package context

import (
	"context"
	"net"
	"strings"

	"google.golang.org/grpc/peer"
)

// GetClientIP extracts the client IP address from the context
//
// Parameters:
//
//   - ctx: The context from which to extract the client IP address
//
// Returns:
//
//   - string: The client IP address
//   - error: An error if the IP address could not be extracted
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

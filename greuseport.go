package greuseport

import (
	"context"
	"net"
)

var listenConfig = net.ListenConfig{
	Control: control,
}

// Listen listens at the given network and address.
// Returns a net.Listener with SO_REUSEPORT and SO_REUSEADDR option set.
func Listen(network string, address string) (net.Listener, error) {
	return listenConfig.Listen(context.Background(), network, address)
}

// ListenPacket listens at the given network and address.
// Returns a net.Listener with SO_REUSEPORT and SO_REUSEADDR option set.
func ListenPacket(network string, address string) (net.PacketConn, error) {
	return listenConfig.ListenPacket(context.Background(), network, address)
}

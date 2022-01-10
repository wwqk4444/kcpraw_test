//go:build !faketcp
// +build !faketcp

package ss

import (
	"net"

	"github.com/zhengying/kcpraw/utils"
)

func listenUDP(c *Config) (net.PacketConn, error) {
	return utils.NewUDPListener(c.Localaddr)
}

func dialUDP(c *Config) (net.Conn, error) {
	return net.Dial("udp", c.Remoteaddr)
}

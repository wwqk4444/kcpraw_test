package main

import (
	ss "github.com/wwqk4444/kcpraw_test/shadowsocks-go/shadowsocks"
	"github.com/wwqk4444/ccsexyz_utils"
)

func RunSocksProxyServer(c *ss.Config) {
	RunTCPServer(c.Localaddr, c, ss.ListenSocks5, socksProxyHandler)
}

func socksProxyHandler(conn ss.Conn, c *ss.Config) {
	defer conn.Close()
	target := GetDstOfConn(conn)
	if c.LogHTTP {
		conn = ss.NewHttpLogConn(conn, c)
	}
	buf := utils.GetBuf(1024)
	n, err := conn.Read(buf)
	if err != nil {
		utils.PutBuf(buf)
		return
	}
	rconn, err := ss.DialSSWithOptions(&ss.DialOptions{
		Target: target,
		C:      c,
		Data:   buf[:n],
	})
	utils.PutBuf(buf)
	if err != nil {
		c.Log(err)
		return
	}
	defer rconn.Close()
	c.Log("proxy", target, "from", conn.RemoteAddr(), "->", conn.LocalAddr(),
		"to", rconn.LocalAddr(), "->", rconn.RemoteAddr())
	ss.Pipe(conn, rconn, c)
}

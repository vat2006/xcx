package control

import "net"

func ConnClose(conn net.Conn) (re bool) {
	err := conn.Close()
	if err == nil {
		return true
	} else {
		return false
	}
}

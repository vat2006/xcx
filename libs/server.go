package libs

import (
	"net"
	"runtime"
	"strings"
	"time"
	"xcx/control"
)

type Server struct {
	listener *net.TCPListener
}

func (self *Server) Start() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8282")
	self.listener, _ = net.ListenTCP("tcp", addr)
	defer self.listener.Close()
	for {
		client, _ := self.listener.Accept()
		go self.handler(client)
	}
}
func (self *Server) handler(conn net.Conn) {
	b := make([]byte, 1024)
	n, _ := conn.Read(b)
	if n <= 9 {
		control.ConnClose(conn)
	}
	control.ConnConnect(conn, parser(string(b[0:n])))
	time.Sleep(10000)
}
func parser(s string) (param map[string]string) {
	var ss, ps []string
	if runtime.GOOS != "linux" {
		ss = strings.Split(s, "\r\n")
	}
	param = make(map[string]string)
	for _, v := range ss {
		ps = strings.Split(v, ":")
		if len(ps) == 1 {
			pps := strings.Fields(ps[0])
			if len(pps) > 1 {
				switch pps[0] {
				case "GET":
					param["Type"] = pps[0]
					param["Path"] = pps[1]
					param["Version"] = pps[2]
				}
			}
		} else {
			param[ps[0]] = ps[1]
		}
	}
	return param
}

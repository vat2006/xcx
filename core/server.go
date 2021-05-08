package core

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net"
	"runtime"
	"strings"
	"xcx/libs"
)

type Server struct {
	listener *net.TCPListener
}
type clientChan chan string

func (self *Server) Start() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8282")
	self.listener, _ = net.ListenTCP("tcp", addr)
	mainOutChan := make(chan map[string]clientChan, 2048)
	mainInChan := make(chan map[string]string, 4096)
	go StartCore(mainOutChan, mainInChan)
	for {
		client, _ := self.listener.Accept()
		go self.handler(client, mainOutChan, mainInChan)
	}
}
func (self *Server) handler(conn net.Conn, mainOutChan chan map[string]clientChan, mainInChan chan map[string]string) {
	b := make([]byte, 1024)
	n, _ := conn.Read(b)
	if n <= 9 {
		doClose(conn)
	}
	rHeader := parser(string(b[0:n]))
	if doConnect(conn, rHeader) {
		mainInChan <- map[string]string{rHeader["Sec-WebSocket-Key"]: strings.Replace(rHeader["path"], "/", "")}
		inChan := make(clientChan,100)
		mainOutChan <- map[string]clientChan{rHeader["Sec-WebSocket-Key"]: inChan}
		b := make([]byte, 256)
		var getMsg string
		for {
			n, _ := conn.Read(b)
			if n > 0 {
				getMsg = libs.UnCode(b[0:n])
			}
			mainInChan <- map[string]string{rHeader["Sec-WebSocket-Key"]: getMsg}
		}

	}
	//time.Sleep(10000)
}
func parser(s string) (param map[string]string) {
	var ss, ps []string
	if runtime.GOOS != "linux" {
		ss = strings.Split(s, "\r\n")
	}
	param = make(map[string]string)
	for _, v := range ss {
		ps = strings.SplitN(v, ":", 2)
		if len(ps) == 1 {
			pps := strings.Fields(ps[0])
			if len(pps) > 1 {
				switch pps[0] {
				case "GET":
					param["Method"] = pps[0]
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

func doClose(conn net.Conn) (re bool) {
	err := conn.Close()
	if err == nil {
		return true
	} else {
		return false
	}
}

func doConnect(conn net.Conn, param map[string]string) (isConnected bool) {
	sha := sha1.New()
	sha.Write([]byte(strings.TrimSpace(param["Sec-WebSocket-Key"])))
	sha.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	newKey := base64.StdEncoding.EncodeToString(sha.Sum(nil))
	wbuf := bufio.NewWriter(conn)
	//rbuf := bufio.NewReader(conn)
	wbuf.WriteString("HTTP/1.1 101 Switching Protocols\r\n")
	wbuf.WriteString("Upgrade: websocket\r\n")
	wbuf.WriteString(fmt.Sprintf("Sec-WebSocket-Version:%s\r\n", param["Sec-WebSocket-Version"]))
	wbuf.WriteString("Connection: Upgrade\r\n")
	wbuf.WriteString(fmt.Sprintf("Sec-WebSocket-Accept:%s\r\n\r\n", newKey))
	wbuf.Flush()
	return true
}

func clinetWrite(conn net.Conn, inChan clientChan) {
	select {
	case msg := <-inChan:
		wbuf := bufio.NewWriter(conn)
		wbuf.WriteString(msg)
		wbuf.Flush()
	}

}

package control

import (
	"bufio"
	"fmt"
	"net"
	"xcx/libs"
)

type Servicer struct {
	conn    net.Conn
	inChan  chan string
	outChan chan map[string]string
	header  map[string]string
}

func (self *Servicer) Init(conn net.Conn, header map[string]string, listChan chan map[string]chan string, outChan chan map[string]string) {
	self.conn = conn
	self.inChan = make(chan string, 2048)
	self.outChan = outChan
	self.header = header
	listChan <- map[string]chan string{header["Sec-WebSocket-Key"]: self.inChan}
	msg := <-self.inChan
	fmt.Println(msg)
	for {
		if msg == "ok" {

			self.inChan <- "servicer"
			go self.clinetWrite()
			break
		} else {
			msg = <-self.inChan
		}
	}
}

func (self *Servicer) Run() {
	rb := make([]byte, 32)
	for {
		n, _ := self.conn.Read(rb)
		msg := libs.UnCode(rb[0:n])
		self.outChan <- map[string]string{self.header["Sec-WebSocket-Key"]: msg}
	}
}
func (self *Servicer) clinetWrite() {
	select {
	case msg := <-self.inChan:
		wbuf := bufio.NewWriter(self.conn)
		wbuf.WriteString(msg)
		wbuf.Flush()
	}

}

package control

import (
	"bufio"
	"fmt"
	"net"
	"xcx/libs"
)

type Consumer struct {
	conn    net.Conn
	inChan  chan string
	outChan chan map[string]string
	header  map[string]string
}

func (self *Consumer) Init(conn net.Conn, header map[string]string, listChan chan map[string]chan string, outChan chan map[string]string) {
	self.conn = conn
	self.inChan = make(chan string, 2048)
	self.outChan = outChan
	self.header = header
	listChan <- map[string]chan string{header["Sec-WebSocket-Key"]: self.inChan}
	msg := <-self.inChan
	fmt.Println("server:" + msg)
	for {
		if msg == "ok" {
			self.inChan <- "counsumer"
			go self.clinetWrite()
			break
		} else {
			msg = <-self.inChan
		}
	}
	self.Run()
}

func (self *Consumer) Run() {
	rb := make([]byte, 32)
	for {
		n, _ := self.conn.Read(rb)
		msg := libs.UnCode(rb[0:n])
		fmt.Println("socket get:" + msg)
		self.outChan <- map[string]string{self.header["Sec-WebSocket-Key"]: msg}

	}
}
func (self *Consumer) clinetWrite() {
	for {
		select {
		case msg := <-self.inChan:
			fmt.Println("chan get:" + msg)
			wbuf := bufio.NewWriter(self.conn)
			wbuf.WriteString(msg + "\r\n\r\n")
			wbuf.Flush()
		}
	}

}

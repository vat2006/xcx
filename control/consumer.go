package control

import "net"

type ConsumerActon struct {
}

func (self *ConsumerActon) Start(conn net.Conn, outChan chan map[string]string, inChan chan map[string]string) {

}

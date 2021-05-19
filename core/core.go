package core

import (
	"container/list"
	"fmt"
)


type core struct {
	listChan    chan map[string]chan string
	inChan      chan map[string]string
	clientList  
	chanMap     map[string]chan string
	outChanList map[string]chan string
	servicerList list.List
	counsumerList list.List
}
type clientInfo struct {
	id           string
	inChan       chan string
	servicer     string
	consumerList list.List
	clientType   string
}

func StartCore(listChan chan map[string]chan string, inChan chan map[string]string) {
	c := new(core)
	c.listChan = listChan
	c.inChan = inChan
	c.clientList = make(map[string]*clientInfo)
	c.chanMap = make(map[string]chan string)
	c.outChanList = make(map[string]chan string)
	c.servicerList=list.New()
	c.counsumerList=list.New()
	c.start()
}
func (self *core) start() {
	select {
	case msg := <-self.inChan:
		for k, v := range msg {
			switch v{
			case "close":
			default:
				cInfo := self.clientList[k]
				switch cInfo.clientType {
				case "counsumer":
					if cInfo.servicer==""{
						
					}
				case "servicer":
					
				}
			}			
		}

	case chanInfo := <-self.listChan:
		for k, v := range chanInfo {
			fmt.Println(k)
			self.chanMap[k] = v
			v <- "ok"
			msg := <-v
			cInfo := new (clientInfo)
			cInfo.id = k
			cInfo.inChan = v
			switch msg {
			case "counsumer":
				cInfo.clientType = "counsumer"
				for e:=self.servicerList.Front();e!=nil;e=e.Next(){
					if(self.clientList[e.Value].)
				}
				self.counsumerList.PushBack(k)
			case "servicer":
				cInfo.clientType = "servicer"
				self.servicerList.PushBack(k)
			}
			self.clientList[k] = cInfo
		}
	}
}

package core

type core struct {
	listChan     chan map[string]chan string
	inChan       chan map[string]string
	consumerList map[string]*consumerInfo
	servicerList map[string]*servicerInfo
	chanMap      map[string]chan string
	outChanList  map[string]chan string
}
type consumerInfo struct {
	id       string
	inChan   chan string
	servicer string
}

type servicerInfo struct {
	id           string
	inChan       chan string
	consumerList []string
}

func StartCore(listChan chan map[string]chan string, inChan chan map[string]string) {
	c := new(core)
	c.listChan = listChan
	c.inChan = inChan
	c.servicerList = make(map[string]*servicerInfo)
	c.consumerList = make(map[string]*consumerInfo)
	c.chanMap = make(map[string]chan string)
	c.outChanList = make(map[string]chan string)
	c.start()
}
func (self *core) start() {
	select {
	case msg := <-self.inChan:
		for k, v := range msg {
			switch v {
			case "consumer":
				c := new(consumerInfo)
				c.id = k
				if _, ok := self.outChanList[k]; ok {
					c.inChan = self.outChanList[k]
				}
				self.consumerList[k] = c
			case "servicer":
				s := new(servicerInfo)
				s.id = k
				if _, ok := self.outChanList[k]; ok {
					s.inChan = self.outChanList[k]
				}
				self.servicerList[k] = s
			default:

			}

		}
	case chanInfo := <-self.listChan:
		for k, v := range chanInfo {
			self.outChanList[k] = v
			if _, ok := self.servicerList[k]; ok {
				self.servicerList[k].inChan = v
			}
			if _, ok := self.consumerList[k]; ok {
				self.consumerList[k].inChan = v
			}
		}
	}
}

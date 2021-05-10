package core

type core struct {
	listChan chan map[string]chan string
	inChan   chan map[string]string
}

func StartCore(listChan chan map[string]chan string, inChan chan map[string]string) {
	c := new(core)
	c.listChan = listChan
	c.inChan = inChan
	c.start()
}
func (self *core) start() {

}

package core

type core struct{

}

func StartCore(outChan chan map[string]*clientChan, inChan chan map[string]string) {
	c := new(core)
	c.start(outChan, inChan)
}
func (self *core) start(outChan chan map[string]*clientChan, inChan chan map[string]string) {
	select{
	case msg:=<-inChan:
		for k,v range msg{
			if v="consumer":
			
		}
	}
}

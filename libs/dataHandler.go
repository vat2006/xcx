package libs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
)

type header struct {
	lastFrame  bool
	opCode     int //0 连续帧，1 文本帧，2 二进制帧，3-7 非控制预留，8 关闭握手帧，9 ping帧，10 pong帧，11-15 控制预留
	haveMask   bool
	dataLen    int
	maskingKey [4]byte
	extData    []byte
	appData    []string
}

func getBit(bt byte, index int) (i int) {
	bt = (bt << index) >> 7
	i, _ = strconv.Atoi(fmt.Sprintf("%b", bt))
	return i
}

func UnCode(bts []byte) string {
	h := header{}
	var len int
	nextIndex := 0
	if bts[nextIndex] < 128 {
		h.lastFrame = false //第一位为0表示还有后续帧
	} else {
		h.lastFrame = true //第一位为1表示最后帧
	}
	h.opCode = int(math.Abs(float64(bts[nextIndex] - 128)))
	nextIndex = 1
	if bts[nextIndex] < 128 {
		h.haveMask = false
	} else {
		h.haveMask = true
	}
	len = int(math.Abs(float64(bts[nextIndex] - 128)))
	nextIndex = 2
	switch len {
	case 126:
		b_buf := bytes.NewBuffer([]byte{bts[2], bts[3]})
		var x uint16
		binary.Read(b_buf, binary.BigEndian, &x)
		h.dataLen = int(x)
		nextIndex = 4
	case 127:
		b_buf := bytes.NewBuffer([]byte{bts[2], bts[3], bts[4], bts[5]})
		var x uint32
		binary.Read(b_buf, binary.BigEndian, &x)
		h.dataLen = int(x)
		nextIndex = 6
	default:
		h.dataLen = len
	}
	h.maskingKey[0] = bts[nextIndex]
	h.maskingKey[1] = bts[nextIndex+1]
	h.maskingKey[2] = bts[nextIndex+2]
	h.maskingKey[3] = bts[nextIndex+3]
	nextIndex = nextIndex + 4
	var msg string
	for i := 0; i < h.dataLen; i++ {
		msg += string(rune(h.maskingKey[i%4] ^ bts[nextIndex+i]))
	}
	return msg
}

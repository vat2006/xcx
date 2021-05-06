package control

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net"
)

func ConnConnect(conn net.Conn, param map[string]string) {
	key := param["Sec-WebSocket-Key"]
	sha := sha1.New()
	sha.Write([]byte(key + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	h := sha.Sum(nil)
	newKey := base64.StdEncoding.EncodeToString(h)
	msg := "HTTP/1.1 101 Switching Protocols\r\n"
	msg = msg + "Upgrade: websocket\r\n"
	msg = msg + "Sec-WebSocket-Version:" + param["Sec-WebSocket-Version"] + "\r\n"
	msg = msg + "Connection: Upgrade\r\n"
	msg = msg + "Sec-WebSocket-Accept: " + newKey + "\r\n\r\n"
	fmt.Println(msg)
	n, err := conn.Write([]byte(msg))
	fmt.Println(n)
	if err != nil {
		fmt.Println(err)
	}
}

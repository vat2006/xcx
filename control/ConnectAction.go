package control

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net"
	"strings"
)

func DoConnect(conn net.Conn, param map[string]string) (isConnected bool) {
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
	rb := make([]byte, 1024)
	var n int
	for {
		n, _ = conn.Read(rb)
		if n > 0 {
			break
		}
	}
	fmt.Println(rb[0:n])
	return true
}
func deCode(bs []byte) {

}

package mazeconnect

import (
	"bufio"
	"log"
	"net"
)

var serverAddr = "192.168.88.225:2390" // TODO: make this dynamic

// connect to maze simulator
func (mc *MazeConnection) hw_init() (err error) {
	udpnet, err := net.Dial("udp", serverAddr)
	if err != nil {
		return
	}
	mc.udpconn = udpnet

	return
}

// step forward
func (mc *MazeConnection) hw_stepforward() (err error) {

	_, err = mc.udpconn.Write([]byte("df"))
	if err != nil {
		return
	}

	data, err := bufio.NewReader(mc.udpconn).ReadString('\n')
	if err != nil {
		return
	}
	log.Printf("Received: %v\n", data)
	return
}

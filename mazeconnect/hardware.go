package mazeconnect

import (
	"bufio"
	"errors"
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

func (mc *MazeConnection) hw_observewalls() (wallstate []int, err error) {
	wallstate = make([]int, 4)

	_, err = mc.udpconn.Write([]byte("sw"))
	if err != nil {
		return
	}

	data, err := bufio.NewReader(mc.udpconn).ReadString('\n')
	if err != nil {
		return
	}
	if len(data) != 4 {
		err = errors.New("arduino returned string of wrong length")
		return
	}
	for i := 0; i < 4; i++ {
		if data[i] == '1' {
			wallstate[i] = 1
		} else {
			wallstate[i] = 0
		}
	}
	log.Printf("Received: %v\n", data)

	return
}

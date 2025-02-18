package mazeconnect

import (
	"bufio"
	"log"
	"net"
)

var serverAddr = "192.168.88.225:2390" // TODO: make this dynamic

func (mc *MazeConnection) SetAddr(newaddr string) (err error) {
	serverAddr = newaddr
	err = nil
	return
}

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
	//if len(data) != 5 {
	//	err = errors.New("arduino returned string of wrong length")
	//	return
	//}
	for i := 0; i < 4; i++ {
		if data[i] == '1' {
			wallstate[i] = 1
		} else {
			wallstate[i] = 0
		}
	}
	//log.Printf("Return string length: %v\n", len(data))

	return
}

func (mc *MazeConnection) hw_turnleft() (err error) {

	_, err = mc.udpconn.Write([]byte("tl"))
	if err != nil {
		return
	}

	_, err = bufio.NewReader(mc.udpconn).ReadString('\n')
	if err != nil {
		return
	}
	return
}

func (mc *MazeConnection) hw_turnright() (err error) {

	_, err = mc.udpconn.Write([]byte("tr"))
	if err != nil {
		return
	}

	_, err = bufio.NewReader(mc.udpconn).ReadString('\n')
	if err != nil {
		return
	}
	return
}

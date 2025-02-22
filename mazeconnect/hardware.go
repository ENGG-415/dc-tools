package mazeconnect

import (
	"bufio"
	"log"
	"net"
	"os"
)

var serverAddr = "192.168.88.225:2390" // TODO: make this dynamic
var serverAddrOverride = false

func (mc *MazeConnection) SetAddr(newaddr string) (err error) {
	serverAddr = newaddr
	serverAddrOverride = true
	err = nil
	return
}

// connect to maze simulator
func (mc *MazeConnection) hw_init() (err error) {

	// try to get robot ip address from local file
	// TODO: make this an acutal configuration file (YAML, JSON, or similar)
	if !serverAddrOverride {
		file, err := os.Open("/etc/engg415-robot-ip.txt")
		if err != nil {
			log.Println("Couldn't open IP address file... using default.")
		} else {
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			scanner.Scan()
			serverAddr = scanner.Text()
			log.Println("Read server address from file: ", serverAddr)
		}
	}

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

	_, err = bufio.NewReader(mc.udpconn).ReadString('\n')
	if err != nil {
		return
	}

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

func (mc *MazeConnection) hw_retrievedebugdata() (debugstr string, err error) {
	err = nil

	_, err = mc.udpconn.Write([]byte("debug"))
	if err != nil {
		return
	}

	debugstr, err = bufio.NewReader(mc.udpconn).ReadString('\n')
	if err != nil {
		return
	}

	return
}

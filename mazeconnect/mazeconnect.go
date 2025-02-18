// the mazeconnect package is your interface to both:
// * maze simulator for testing algorithms and displaying locally on your screen
// * hardware robots navigating the physical maze (this functionality coming soon!)
package mazeconnect

import (
	"errors"
	"image/color"
	"net"
	"net/rpc"
)

type ConnectMode int

type RemoteParams struct {
	AgentID  int
	StepSize int // [# cells] +: forward, -:backward, 0: no step, just return
	TurnDeg  int // 90 or -90 ONLY
	Color    color.RGBA
	Data     []float32
}

type RemoteAgentState struct {
	AgentID  int
	CellIdx  int
	ThetaDeg int
}

type RemoteReturn struct {
	CellID int
}

const (
	M_simulator ConnectMode = iota
	M_hardware
)

type MazeConnection struct {
	mode      ConnectMode
	rpcclient *rpc.Client
	udpconn   net.Conn // TODO add checks to mke sure we've connected
}

func (mc *MazeConnection) Init(mode ConnectMode) (err error) {
	err = nil

	// validate mode
	if mode != M_simulator && mode != M_hardware {
		return errors.New("invalid mode: only M_simulator implemented currently")
	}
	mc.mode = mode

	// now handle actual initialization
	switch mc.mode {
	case M_simulator:
		err = mc.sim_init()
	case M_hardware:
		err = mc.hw_init()
	default:
		err = errors.New("cannot initialize connection in requested mode")
	}
	return
}

func (mc *MazeConnection) SetAgentState(agentid int, cellid int, thetadeg int) (err error) {
	err = nil
	switch mc.mode {
	case M_simulator:
		err = mc.sim_setagentstate(agentid, cellid, thetadeg)
	default:
		err = errors.New("cannot set agent state in this mode")
	}
	return
}

func (mc *MazeConnection) ObserveWalls(agentid int) (wallstate []int, err error) {
	err = nil
	wallstate = make([]int, 0)
	switch mc.mode {
	case M_simulator:
		wallstate, err = mc.sim_observewalls(agentid)
	case M_hardware:
		wallstate, err = mc.hw_observewalls()
	default:
		err = errors.New("cannot observe walls in this mode")
	}
	return
}

func (mc *MazeConnection) SetCellValues(cv []float32) (err error) {
	switch mc.mode {
	case M_simulator:
		err = mc.sim_setcellvalues(cv)
	default:
		err = errors.New("cannot set cell values in this mode")
	}
	return
}

func (mc *MazeConnection) StepForward(agentid int) (err error) {
	switch mc.mode {
	case M_simulator:
		err = mc.sim_stepforward(agentid)
	case M_hardware:
		err = mc.hw_stepforward()
	default:
		err = errors.New("cannot step forward in this mode")
	}
	return
}

func (mc *MazeConnection) TurnLeft(agentid int) (err error) {
	switch mc.mode {
	case M_simulator:
		err = mc.sim_turnleft(agentid)
	case M_hardware:
		err = mc.hw_turnleft()
	default:
		err = errors.New("cannot turn left in this mode")
	}
	return
}

func (mc *MazeConnection) TurnRight(agentid int) (err error) {
	switch mc.mode {
	case M_simulator:
		err = mc.sim_turnright(agentid)
	case M_hardware:
		err = mc.hw_turnright()
	default:
		err = errors.New("cannot turn right in this mode")
	}
	return
}

func (mc *MazeConnection) AddPath(cellindices []int, color color.RGBA) (err error) {
	switch mc.mode {
	case M_simulator:
		err = mc.sim_addpath(cellindices, color)
	default:
		err = errors.New("cannot add a path in this mode")
	}
	return
}

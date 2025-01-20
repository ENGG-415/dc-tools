// simulator RPC calls
package mazeconnect

import (
	"image/color"
	"net/rpc"
)

// connect to maze simulator
func (mc *MazeConnection) sim_init() (err error) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		return
	}
	mc.rpcclient = client
	return
}

// set agent state
func (mc *MazeConnection) sim_setagentstate(agentid int, cellid int, thetadeg int) (err error) {
	var j int
	state := RemoteAgentState{
		AgentID:  agentid,
		CellIdx:  cellid,
		ThetaDeg: thetadeg,
	}
	err = mc.rpcclient.Call("Game.RemoteSetAgentState", state, &j)
	return
}

func (mc *MazeConnection) sim_observewalls(agentid int) (wallstate []int, err error) {
	wallstate = make([]int, 4)
	params := RemoteParams{
		AgentID: agentid,
	}
	err = mc.rpcclient.Call("Game.RemoteObserveAgentWalls", params, &wallstate)
	return
}

func (mc *MazeConnection) sim_setcellvalues(cv []float32) (err error) {
	var j int
	params := RemoteParams{
		Data: cv,
	}
	err = mc.rpcclient.Call("Game.RemoteSetCellValues", params, &j)
	return
}

func (mc *MazeConnection) sim_stepforward(agentid int) (err error) {
	var j int
	params := RemoteParams{
		AgentID:  agentid,
		StepSize: 1,
	}
	err = mc.rpcclient.Call("Game.RemoteStepAgent", params, &j)
	return
}

func (mc *MazeConnection) sim_turnleft(agentid int) (err error) {
	var j int
	params := RemoteParams{
		AgentID: agentid,
		TurnDeg: -90,
	}
	err = mc.rpcclient.Call("Game.RemoteTurnAgent", params, &j)
	return
}

func (mc *MazeConnection) sim_turnright(agentid int) (err error) {
	var j int
	params := RemoteParams{
		AgentID: agentid,
		TurnDeg: 90,
	}
	err = mc.rpcclient.Call("Game.RemoteTurnAgent", params, &j)
	return
}

func (mc *MazeConnection) sim_addpath(cellindices []int, color color.RGBA) (err error) {
	var j int
	floatpath := make([]float32, len(cellindices))
	for i, cidx := range cellindices {
		floatpath[i] = float32(cidx)
	}
	params := RemoteParams{
		Color: color,
		Data:  floatpath,
	}
	err = mc.rpcclient.Call("Game.RemoteAddPath", params, &j)
	return

}

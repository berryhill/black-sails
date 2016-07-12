package main

import (
	"fmt"

	"github.com/black-sails/osc"
	"github.com/black-sails/session"
	//"github.com/black-sails/monome"
	"time"
)

func init() {
	//M := monome.NewMonome(1)
	//go M.TestScan()
	//go M.TestFlicker()
}

// TODO: Revise the client!
func main() {
	S := session.NewSession()

	go func () {
		for k:= 0; k < 128; k++ {
			go S.Tracks[0].Stems[0].SetPlayPosition(k)
			time.Sleep(150 * time.Millisecond)
		}
	}()

	OSC := osc.NewOsc("127.0.0.1", 4444, "127.0.0.1:6666")
	fmt.Println(OSC.Client)
}

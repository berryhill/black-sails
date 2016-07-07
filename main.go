package main

import (
	"fmt"

	"github.com/black-sails/osc"
	"github.com/black-sails/monome"
)

func init() {
	M := monome.NewMonome(1)
	go M.TestScan()
}

// TODO: Revise the client!
func main() {
	OSC := osc.NewOsc("127.0.0.1", 4444, "127.0.0.1:6666")
	fmt.Println(OSC.Client)
}

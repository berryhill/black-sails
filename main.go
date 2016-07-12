package main

import (
	"fmt"
	"bufio"
	"os"

	//"github.com/black-sails/osc"
	"github.com/black-sails/session"
	//"github.com/black-sails/monome"
	//"time"
)

func init() {
	//M := monome.NewMonome(1)
	//go M.TestScan()
	//go M.TestFlicker()
}

// TODO: Revise the client!
func main() {
	S := session.NewSession()
	fmt.Println(S)

	//go func () {
	//	for k:= 0; k < 128; k++ {
	//		go S.Tracks[0].Stems[0].SetPlayPosition(k)
	//		time.Sleep(150 * time.Millisecond)
	//	}
	//}()

	reader := bufio.NewReader(os.Stdin)
	for {
		c, err := reader.ReadByte()
		if err != nil {
			os.Exit(0)
		}

		if c == 'q' {
			os.Exit(0)
		}
	}
}

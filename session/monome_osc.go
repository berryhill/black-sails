package session

import (
	"fmt"

	osc "github.com/kward/go-osc"
)

type MOsc struct {
	Session 		*Session
	Client 			*osc.Client
}

func NewMOsc(ip string, port_out int, addr string, S *Session) *MOsc {
	o := new(MOsc)
	//client := osc.NewClient(ip, port_out)

	//message := osc.NewMessage("/m/grid/led/all")
	//message.Append(int32(1))
	//client.Send(message)

	o.Session = S
	o.setupMOsc(addr, S)

	return o
}

func (o *MOsc) setupMOsc(addr string, S *Session) {
	server := &osc.Server{Addr: addr}

	server.Handle("/m/grid/key", func(msg *osc.Message) {
		var seek_position_input int32
		var stem int

		column := msg.Arguments[0].(int32)
		row := msg.Arguments[1].(int32)
		state := msg.Arguments[2].(int32)
		//fmt.Println(row, column, state)

		if row % 2 == 1 && state == 1 {
			stem = (int(row) - 1) / 2
			seek_position_input = column + 16
			S.Tracks[0].Stems[stem].SetSeekPosition(int(seek_position_input))
		} else if row % 1 == 0 && state == 1 {
			stem = int(row) / 2
			seek_position_input = column
			S.Tracks[0].Stems[stem].SetSeekPosition(int(seek_position_input))
		}

		//fmt.Println(seek_position_input, stem)

	})

	fmt.Println("### Welcome to go-osc receiver demo")
	fmt.Println("Press \"q\" to exit")

	go server.ListenAndServe()
}

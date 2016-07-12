package session

import (
	//"bufio"
	//"net"
	"fmt"
	//"os"
	//"time"

	osc "github.com/kward/go-osc"
	//"golang.org/x/net/context"
	//"bufio"
	//"github.com/black-sails/session"
)

type Osc struct {
	Client 			*osc.Client
}

func NewOsc(ip string, port_out int, addr string, S *Session) *Osc {
	o := new(Osc)
	//client := osc.NewClient(ip, port_out)

	//message := osc.NewMessage("/m/grid/led/all")
	//message.Append(int32(1))
	//client.Send(message)

	o.setupOsc(addr, S)

	return o
}

func (o *Osc) setupOsc(addr string, S *Session) {
	server := &osc.Server{Addr: addr}

	server.Handle("/position", func(msg *osc.Message) {
		track := msg.Arguments[0].(int32)
		position := msg.Arguments[1].(int32)
		S.Tracks[0].Stems[track - 1].SetPlayPosition(int(position))
	})

	server.Handle("/m/grid/key", func(msg *osc.Message) {
		message := msg.Arguments[0].(int32)
		fmt.Println(message)
	})

	fmt.Println("### Welcome to go-osc receiver demo")
	fmt.Println("Press \"q\" to exit")

	go server.ListenAndServe()
}

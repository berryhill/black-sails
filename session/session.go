package session

import (
	"fmt"

	"github.com/black-sails/monome"
	"github.com/rakyll/portmidi"
)

var OSC *Osc
var MOSC *MOsc


type Session struct {
	Monome 					*monome.Monome
	Tracks	 				[]*Track
	Osc 					*Osc
}

func NewSession() *Session {
	s := new(Session)
	s.Monome = monome.NewMonome(1)
	s.Tracks = append(s.Tracks, NewTrack(s.Monome))

	OSC = NewOsc("127.0.0.1", 4444, "127.0.0.1:6666", s)
	MOSC = NewMOsc("127.0.0.1", 7777, "127.0.0.1:4444", s)

	fmt.Println(OSC.Client)

	s.initMidi()

	go func() {
		for {
			s.handleMidiQuantization()
		}
	}()

	return s
}

func (s *Session) initMidi() {
	portmidi.Initialize()
}

func (s *Session) handleMidiQuantization () {
	in, err := portmidi.NewInputStream(0, 1024)
	if err != nil {
		fmt.Println(err)
	}

	ch := in.Listen()
	event := <-ch

	if event.Data1 == 1 {
		for k := 0; k < 4; k++ {
			go s.Tracks[0].Stems[k].QuantizedWork()
		}
	}

	in.Close()
}
package session

import (
	"fmt"

	"github.com/black-sails/monome"
)

type Session struct {
	Monome 					*monome.Monome
	Tracks	 				[]*Track
	Osc 					*Osc
}

func NewSession() *Session {
	s := new(Session)
	s.Monome = monome.NewMonome(1)
	s.Tracks = append(s.Tracks, NewTrack(s.Monome))

	OSC := NewOsc("127.0.0.1", 4444, "127.0.0.1:6666", s)
	fmt.Println(OSC.Client)

	return s
}



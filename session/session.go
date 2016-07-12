package session

import "github.com/black-sails/monome"

type Session struct {
	Monome 					*monome.Monome
	Tracks	 				[]*Track
}

func NewSession() *Session {
	s := new(Session)
	s.Monome = monome.NewMonome(1)
	s.Tracks = append(s.Tracks, NewTrack(s.Monome))

	return s
}



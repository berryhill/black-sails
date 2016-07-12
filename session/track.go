package session

import "github.com/black-sails/monome"

type Track struct {
	Stems 				[]*Stem
}

func NewTrack(monome *monome.Monome) *Track {
	t := new(Track)
	t.Stems = append(t.Stems, NewStem("test", 1, monome))

	return t
}

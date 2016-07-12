package session

import "github.com/black-sails/monome"

type Track struct {
	Stems 				[]*Stem
}

func NewTrack(monome *monome.Monome) *Track {
	t := new(Track)

	for k := 1; k < 5; k++ {
		t.Stems = append(t.Stems, NewStem("test", k, monome))
	}

	return t
}

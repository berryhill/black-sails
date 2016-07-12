package session

import (
	"github.com/black-sails/monome"
)

type Stem struct {
	Name 						string
	Index						int
	PlayPosition 				int
	ReflectRow 					*monome.ButtonRow
	ControlRow					*monome.ButtonRow
}

func NewStem(name string, index int, monome *monome.Monome) *Stem {
	s := new(Stem)
	s.Name = name
	s.Index = index
	s.PlayPosition = 0
	s.ReflectRow = monome.Rows[0]
	s.ControlRow = monome.Rows[1]

	return s
}

func (s *Stem) getParsedPlayPosition() (int, int) {
	course_position := (s.PlayPosition / 16) % 8
	fine_position := s.PlayPosition % 16
	return course_position, fine_position
}

func (s *Stem) sendMonomePlayPosition() {
	cp, fp := s.getParsedPlayPosition()
	s.ReflectRow.SetOneLed(fp)
	s.ControlRow.SetBarLed(cp)
}

func (s *Stem) SetPlayPosition(position int) {
	s.PlayPosition = position
	s.sendMonomePlayPosition()
}


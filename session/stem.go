package session

import (
	"github.com/black-sails/monome"
)

type Stem struct {
	Name 						string
	Index						int
	PlayPosition 				int
	ReflectRow1 				*monome.ButtonRow
	ReflectRow2					*monome.ButtonRow

}

func NewStem(name string, index int, monome *monome.Monome) *Stem {
	s := new(Stem)
	s.Name = name
	s.Index = index
	s.PlayPosition = 0
	if index == 1 {
		s.ReflectRow1 = monome.Rows[0]
		s.ReflectRow2 = monome.Rows[1]
	} else if index == 2 {
		s.ReflectRow1 = monome.Rows[2]
		s.ReflectRow2 = monome.Rows[3]
	} else if index == 3 {
		s.ReflectRow1 = monome.Rows[4]
		s.ReflectRow2 = monome.Rows[5]
	} else if index == 4 {
		s.ReflectRow1 = monome.Rows[6]
		s.ReflectRow2 = monome.Rows[7]
	}

	return s
}

func (s *Stem) getParsedPlayPosition() int {
	//course_position := (s.PlayPosition / 16) % 8
	fine_position := s.PlayPosition % 32
	return fine_position
}

func (s *Stem) sendMonomePlayPosition() {
	fp := s.getParsedPlayPosition()
	if fp < 16 {
		s.ReflectRow1.SetOneLed(fp)
		s.ReflectRow2.SetLedOff()
	} else {
		s.ReflectRow2.SetOneLed(fp - 16)
		s.ReflectRow1.SetLedOff()
	}
}

func (s *Stem) SetPlayPosition(position int) {
	s.PlayPosition = position
	s.sendMonomePlayPosition()
}
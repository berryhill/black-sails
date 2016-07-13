package session

import (
	"github.com/black-sails/monome"
	"github.com/kward/go-osc"
)

type Stem struct {
	Name 						string
	Index						int
	PlayPosition 				int
	SeekPosition 				int
	ReflectRow1 				*monome.ButtonRow
	ReflectRow2					*monome.ButtonRow
}

func NewStem(name string, index int, monome *monome.Monome) *Stem {
	s := new(Stem)
	s.Name = name
	s.Index = index
	s.PlayPosition = 0
	s.SeekPosition = -1

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

func (s *Stem) QuantizedWork() {
	//fmt.Println("click")

	if s.SeekPosition != -1 {
		s.sendLiveSeekPosition()
		s.SeekPosition = -1
	}
}

func (s *Stem) getParsedPlayPosition() int {
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

func (s *Stem) sendLiveSeekPosition() {
	seek := (s.SeekPosition - s.PlayPosition) - 1

	client := osc.NewClient("127.0.0.1", 8888)
	message := osc.NewMessage("/seek")
	message.Append(int32(s.Index))
	message.Append(int32(seek))

	client.Send(message)
}

func (s *Stem) SetPlayPosition(position int) {
	s.PlayPosition = position
	s.sendMonomePlayPosition()
}

func (s *Stem) SetSeekPosition(position int) {
	s.SeekPosition = position
	s.sendMonomePlayPosition()
}

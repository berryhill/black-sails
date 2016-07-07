package monome

import (
	"time"

	//"github.com/kward/go-osc"
)

type Monome struct {
	//Client 		*osc.Client
	Id 			int
	Rows			[]*ButtonRow
}

func NewMonome(id int) *Monome {
	m := new(Monome)
	//osc.NewClient("127.0.0.1", 5555)
	m.Id = id
	for k := 0; k < 8; k++ {
		br := NewButtonRow(k)
		m.Rows = append(m.Rows, br)
	}

	m.listenOutput()
	m.listenInput()

	return m
}

func (m *Monome) listenInput() {
	go func() {
		for {
			for k := 0; k < 8; k++ {
				for j := 0; j < 16; j++ {
					if m.Rows[k].Buttons[j].Led == true {

					}
				}
			}
			time.Sleep(time.Millisecond)
			//println("monome-out")
		}
	}()
}

func (m *Monome) listenOutput() {
	go func() {
		for {
			for k := 0; k < 8; k++ {
				for j := 0; j < 16; j++ {
					if m.Rows[k].Buttons[j].Led == true {
						m.Rows[k].Buttons[j].SendMessage(1)
					} else {
						m.Rows[k].Buttons[j].SendMessage(0)
					}
				}
			}
			time.Sleep(time.Millisecond)
			//println("monome-in")
		}
	}()
}

func (m *Monome) TestScan() {
	for k := 0; k < 8; k++ {
		for j := 0; j < 16; j++ {
			m.Rows[k].Buttons[j].LedOn()
			time.Sleep(50 * time.Millisecond)
			m.Rows[k].Buttons[j].LedOff()
		}
	}
}


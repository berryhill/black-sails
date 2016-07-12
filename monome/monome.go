package monome

import (
	"time"

	//"github.com/kward/go-osc"
)

type Monome struct {
	//Client 				*osc.Client
	Id 						int
	Rows					[]*ButtonRow
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
						if m.Rows[k].Buttons[j].LedPast == false {
							m.Rows[k].Buttons[j].SendMessage(1)
							m.Rows[k].Buttons[j].LedPast = true
						}
					} else {
						if m.Rows[k].Buttons[j].LedPast == true {
							m.Rows[k].Buttons[j].SendMessage(0)
							m.Rows[k].Buttons[j].LedPast = false
						}
					}
				}
			}
			time.Sleep(time.Millisecond)
			//println("monome-in")
		}
	}()
}

func (m *Monome) TestFlicker() {
	for {
		for k := 0; k < 8; k++ {
			m.Rows[k].Buttons[0].LedFlicker()
			time.Sleep(time.Millisecond)
		}
	}

}

func (m *Monome) TestScan() {
	for {
		for k := 0; k < 16; k++ {
			for j := 0; j < 8; j++ {
				m.Rows[j].Buttons[k].LedOn()
			}

			time.Sleep(250 * time.Millisecond)
			for j := 0; j < 8; j++ {
				m.Rows[j].Buttons[k].LedOff()
			}
		}
	}
}


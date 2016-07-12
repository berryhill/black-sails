package monome

type ButtonRow struct {
	Index 			int
	Buttons			[]*Button
}

func NewButtonRow(index int) *ButtonRow {
	br := new(ButtonRow)
	br.Index = index
	for k := 0; k < 16; k++ {
		b := NewButton(index, k)
		br.Buttons = append(br.Buttons, b)
	}

	return br
}

func (br *ButtonRow) SetLedOff() {
	for k := 0; k < 16; k++ {
		br.Buttons[k].LedOff()
	}
}

func (br *ButtonRow) SetOneLed(index int) {
	for k := 0; k<16; k++ {
		if k == index {
		br.Buttons[k].LedOn()
	} else {
		br.Buttons[k].LedOff()
		}
	}
}

func (br *ButtonRow) SetBarLed(index int) {
	for k := 0; k<8; k++ {
		if k == index {
			br.Buttons[k].LedOn()
		} else {
			br.Buttons[k].LedOff()
		}
	}
}
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
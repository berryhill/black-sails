package monome

import (
	osc "github.com/kward/go-osc"
)

type Button struct {
	Row 		int
	Index 		int
	Led 		bool
	Pressed 	bool
}

func NewButton(row int, index int) *Button {
	b := new(Button)
	b.Row = row
	b.Index = index
	b.Led = false
	b.Pressed = false
	//fmt.Println("Made Button %i %i", row, index)

	return b
}

func (b *Button) LedOn() {
	b.Led = true
}

func (b *Button) LedOff() {
	b.Led = false
}

func (b *Button) SendMessage(bool int) {
	client := osc.NewClient("127.0.0.1", 5555)
	message := osc.NewMessage("/m/grid/led/set")
	message.Append(int32(b.Row))
	message.Append(int32(b.Index))
	message.Append(int32(bool))

	client.Send(message)
}

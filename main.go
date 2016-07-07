package main

import (
	"bufio"
	"net"
	"fmt"
	"os"

	osc "github.com/kward/go-osc"
	"golang.org/x/net/context"
)

// TODO: Revise the client!
func main() {
	client := osc.NewClient("127.0.0.1", 6666)
	client.Send(osc.NewMessage("/bundle/message/1"))

	addr := "127.0.0.1:5555"
	server := &osc.Server{}
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		fmt.Println("Couldn't listen: ", err)
	}
	defer conn.Close()

	fmt.Println("### Welcome to go-osc receiver demo")
	fmt.Println("Press \"q\" to exit")

	go func() {
		fmt.Println("Start listening on", addr)
		for {
			packet, remote, err := server.ReceivePacket(context.Background(), conn)
			if err != nil {
				fmt.Println("Server error: " + err.Error())
				os.Exit(1)
			}

			if packet != nil {
				switch packet.(type) {
				default:
					fmt.Println("Unknown packet type!")
				case *osc.Message:
					fmt.Printf("-- OSC Message from %v: ", remote)
					osc.PrintMessage(packet.(*osc.Message))
				case *osc.Bundle:
					fmt.Println("-- OSC Bundle from %v:", remote)
					bundle := packet.(*osc.Bundle)
					for i, message := range bundle.Messages {
						fmt.Printf("  -- OSC Message #%d: ", i+1)
						osc.PrintMessage(message)
					}
				}
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		c, err := reader.ReadByte()
		if err != nil {
			os.Exit(0)
		}

		if c == 'q' {
			os.Exit(0)
		}
	}
}

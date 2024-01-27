package main

import (
	"machine"
	"time"
)

func main() {
	rx := machine.D2
	tx := machine.D3

	uart := machine.UART{}
	uart.Configure(machine.UARTConfig{
		TX:       tx,
		RX:       rx,
		BaudRate: 115200,
	})

	_, err := uart.Write([]byte("AT\r\n"))
	if err != nil {
		println(err.Error())
	}

	for {
		res, _ := uart.ReadByte()
		println(string(res))
		time.Sleep(time.Millisecond * 500)
	}
}

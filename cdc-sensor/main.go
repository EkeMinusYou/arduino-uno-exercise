package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	machine.InitADC()
	cdcSensor := machine.ADC{Pin: machine.ADC0}

	for {
		if cdcSensor.Get() > 30000 {
			led.Low()
		} else {
			led.High()
		}
		time.Sleep(time.Millisecond * 1000)
	}
}

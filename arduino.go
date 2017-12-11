package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
	led := gpio.NewLedDriver(firmataAdaptor, "13")
	led2 := gpio.NewLedDriver(firmataAdaptor, "8")

	setupBlink := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	setupBlink2 := func() {
		gobot.Every(3*time.Second, func() {
			led2.Toggle()
		})
	}

	work := func() {
		setupBlink()
		setupBlink2()
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led, led2},
		work,
	)
	robot.Start()
}

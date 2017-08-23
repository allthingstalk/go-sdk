//    _   _ _ _____ _    _              _____     _ _     ___ ___  _  __
//   /_\ | | |_   _| |_ (_)_ _  __ _ __|_   _|_ _| | |__ / __|   \| |/ /
//  / _ \| | | | | | ' \| | ' \/ _` (_-< | |/ _` | | / / \__ \ |) | ' <
// /_/ \_\_|_| |_| |_||_|_|_||_\__, /__/ |_|\__,_|_|_\_\ |___/___/|_|\_\
//                             |___/
//
// Copyright 2017 AllThingsTalk
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/allthingstalk/go-sdk"
	"github.com/allthingstalk/go-sdk/profile"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

const (
	deviceID    = "<DEVICE_ID>"
	deviceToken = "<DEVICE_TOKEN>"
)

// IoT Hello World Example. Toggle the LED remotely from Maker platform.
// Connect your LED to GPIO4 (Pin 7)
var r = raspi.NewAdaptor()
var led = gpio.NewLedDriver(r, "7")

func onMessage(command sdk.Command) {
	if command.Name == "LED" {
		if command.Value == true {
			fmt.Printf("Turning LED On.\n")
			led.On()
		} else {
			fmt.Printf("Turning LED Off.\n")
			led.Off()
		}
	}
}

func main() {
	device, err := sdk.NewDevice(deviceID, deviceToken)
	if err != nil {
		panic(err)
	}

	device.SetCommandHandler(onMessage)
	device.Add(sdk.NewActuator("LED", profile.Boolean()))

	robot := gobot.NewRobot("led_demo",
		[]gobot.Connection{r},
		[]gobot.Device{led},
	)

	robot.Start()
}

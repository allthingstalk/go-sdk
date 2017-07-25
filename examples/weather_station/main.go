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
	"github.com/allthingstalk/go-sdk"
	"math/rand"
	"time"
)

const (
	deviceID    = "<DEVICE_ID>"
	deviceToken = "<DEVICE_TOKEN>"
)

// A virtual weather station device
func main() {
	device, err := sdk.NewDevice(deviceID, deviceToken)
	if err != nil {
		panic(err)
	}

	// add some environmental sensors
	temperature, _ := device.AddNumber("temperature")
	humidity, _ := device.AddNumber("humidity")
	pressure, _ := device.AddNumber("pressure")

	// and just push some random data to them, indefinitely
	rand.Seed(time.Now().UnixNano())
	for {
		time.Sleep(1 * time.Second)
		device.Publish(temperature, rand.Int31n(30))
		device.Publish(humidity, rand.Int31n(100))
		device.Publish(pressure, rand.Int31n(1000))
	}
}

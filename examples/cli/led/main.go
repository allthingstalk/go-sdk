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
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	deviceID    = "<DEVICE_ID>"
	deviceToken = "<DEVICE_TOKEN>"
)

func onMessage(command sdk.Command) {
	fmt.Printf("%s is now %v\n", command.Name, command.Value)
}

// A device with a virtual LED which you can turn on and off
func main() {
	device, err := sdk.NewDevice(deviceID, deviceToken)
	if err != nil {
		panic(err)
	}

	device.SetCommandHandler(onMessage)

	led := sdk.NewActuator("LED", profile.Boolean())
	device.Add(led)

	// loop until SigTerm
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Exiting")
		os.Exit(0)
	}()

	fmt.Println("Waiting for sigterm.")
	for {
		time.Sleep(1 * time.Second)
	}
}

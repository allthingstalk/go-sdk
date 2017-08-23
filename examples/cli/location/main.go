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
	"github.com/allthingstalk/go-sdk/profile"
	"time"
)

const (
	deviceID    = "<DEVICE_ID>"
	deviceToken = "<DEVICE_TOKEN>"
)

// Location represents Geo Location
type Location struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"long"`
}

const locationProfile = `
{
   "type":"object",
   "properties":{
      "lat":  {"type": "number"},
      "long": {"type": "number"}
   }
}
`

// A device which publishes a location
func main() {
	device, err := sdk.NewDevice(deviceID, deviceToken)
	if err != nil {
		panic(err)
	}

	p, _ := profile.JSON(locationProfile)
	location := sdk.NewAsset("location", sdk.Sensor, p)

	device.Add(location)
	time.Sleep(1 * time.Second)

	device.Publish(location, Location{Latitude: 51.48, Longitude: 0})
}

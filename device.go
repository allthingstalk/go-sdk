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

package sdk

import (
	"github.com/allthingstalk/go-sdk/profile"
	"time"
)

// CommandHandler can be attached to function to listen to incoming commands.
type CommandHandler func(command Command)

// Device is a representation of a device on AllThingsTalk platform.
type Device struct {
	id                    string
	token                 string
	options               *Options
	httpClient            *httpClient
	mqttClient            *mqttClient
	defaultCommandHandler CommandHandler
}

// NewDevice creates a new device using a DeviceId and token.
func NewDevice(deviceID string, token string) (*Device, error) {
	return NewDeviceWithOptions(deviceID, token, NewOptions())
}

// NewDeviceWithOptions creates a new device using a DeviceId and Token, with
// custom options being supplied.
func NewDeviceWithOptions(deviceID string, token string, options *Options) (*Device, error) {
	initLogging()

	device := &Device{
		id:      deviceID,
		token:   token,
		options: options,
	}

	httpC, err := newHTTPClient(nil, device)
	if err != nil {
		return nil, err
	}

	mqttC, err := newMqttClient(device)
	if err != nil {
		return nil, err
	}

	device.httpClient = httpC
	device.mqttClient = mqttC
	device.listen()

	return device, nil
}

// Add adds an asset to a device.
func (device *Device) Add(a *Asset) error {
	return device.httpClient.AssetService.addAsset(device, a)
}

// AddInteger adds an Integer sensor to a device.
func (device *Device) AddInteger(name string) (*Asset, error) {
	asset := NewSensor(name, profile.Integer())
	return asset, device.Add(asset)
}

// AddNumber adds an Number sensor to a device.
func (device *Device) AddNumber(name string) (*Asset, error) {
	asset := NewSensor(name, profile.Number())
	return asset, device.Add(asset)
}

// AddBoolean adds an Boolean sensor to a device.
func (device *Device) AddBoolean(name string) (*Asset, error) {
	asset := NewSensor(name, profile.Boolean())
	return asset, device.Add(asset)
}

// AddString adds an String asset to a device.
func (device *Device) AddString(name string) (*Asset, error) {
	asset := NewSensor(name, profile.String())
	return asset, device.Add(asset)
}

// Publish publishes raw asset value. Timestamp is set to current UTC time.
func (device *Device) Publish(a *Asset, value interface{}) {
	state := State{Timestamp: time.Now().UTC(), Value: value}
	device.mqttClient.publish(device, a, state)
}

// PublishState publishes asset state. Client can supply value and timestamp.
func (device *Device) PublishState(asset *Asset, state State) {
	device.mqttClient.publish(device, asset, state)
}

// SetCommandHandler allows for setting a function to handle incoming commands.
func (device *Device) SetCommandHandler(handler CommandHandler) {
	device.defaultCommandHandler = handler
}

func (device *Device) onCommand(command Command) {
	if device.defaultCommandHandler != nil {
		device.defaultCommandHandler(command)
	}
}

func (device *Device) listen() {
	device.mqttClient.subscribe(device)
}

type setAssetRequest struct {
	Type    string `json:"is"`
	Profile string `json:"profile"`
}

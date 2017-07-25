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

package sdk

import (
	"github.com/allthingstalk/go-sdk/profile"
	"time"
)

// Kind of an asset. Can be `Sensor` or `Actuator`
type Kind string

// Existing asset kinds. Can be either `sensor` or an `actuator`.
const (
	Sensor   Kind = "sensor"
	Actuator Kind = "actuator"
)

// Asset defines a structure holding some basic information about an asset,
// like it's kind and supported data type (profile).
type Asset struct {
	Kind    Kind        `json:"is"`
	Name    string      `json:"name"`
	Profile interface{} `json:"profile"`
}

// State message describing asset state in time.
type State struct {
	Timestamp time.Time   `json:"at"`
	Value     interface{} `json:"value"`
}

// Command message describing command intended for an asset.
type Command struct {
	Name      string
	Timestamp time.Time   `json:"at"`
	Value     interface{} `json:"value"`
}

// NewAsset creates an asset with a name, kind and profile.
func NewAsset(name string, kind Kind, profile profile.Profile) *Asset {
	return &Asset{Kind: kind, Name: name, Profile: profile.Definition}
}

// NewSensor creates a new Sensor asset with a given profile.
func NewSensor(name string, profile profile.Profile) *Asset {
	return NewAsset(name, Sensor, profile)
}

// NewActuator creates a new Actuator asset with a given profile.
func NewActuator(name string, profile profile.Profile) *Asset {
	return NewAsset(name, Actuator, profile)
}

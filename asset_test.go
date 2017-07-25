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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAsset(t *testing.T) {
	a := assert.New(t)

	var asset = NewAsset("Name", Sensor, profile.New(profile.StringType))
	a.NotNil(asset)
	a.Equal(asset.Name, "Name")
	a.Equal(asset.Kind, Sensor)
}

func TestNewActuator(t *testing.T) {
	a := assert.New(t)

	var actuator = NewActuator("actuator", profile.String())
	a.NotNil(actuator)
	a.Equal(actuator.Kind, Actuator)
	a.Equal(actuator.Name, "actuator")
}

func TestNewSensor(t *testing.T) {
	a := assert.New(t)

	var actuator = NewSensor("sensor", profile.String())
	a.NotNil(actuator)
	a.Equal(actuator.Kind, Sensor)
	a.Equal(actuator.Name, "sensor")
}

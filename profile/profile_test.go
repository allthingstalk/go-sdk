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

package profile

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProfile(t *testing.T) {
	a := assert.New(t)

	a.Equal("{\"type\":\"string\"}", New(StringType).json())
	a.Equal("{\"type\":\"integer\"}", New(IntegerType).json())
	a.Equal("{\"type\":\"number\"}", New(NumberType).json())
	a.Equal("{\"type\":\"boolean\"}", New(BooleanType).json())
}

func TestPresetProfiles(t *testing.T) {
	a := assert.New(t)

	a.Equal(New(StringType), String())
	a.Equal(New(NumberType), Number())
	a.Equal(New(IntegerType), Integer())
	a.Equal(New(BooleanType), Boolean())
}
func TestNewComplexProfile(t *testing.T) {
	a := assert.New(t)

	var json = "{\"type\": \"object\"}"
	var p, err = JSON(json)

	a.Nil(err)
	a.NotNil(p)
	a.Equal("{\"type\":\"object\"}", p.json())
}

func TestNewComplexProfile_InvalidJson(t *testing.T) {
	a := assert.New(t)

	var json = "not valid"
	var _, err = JSON(json)

	a.Error(err, "Invalid JSON")
}

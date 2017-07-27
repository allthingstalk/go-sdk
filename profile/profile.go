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
	"encoding/json"
	"errors"
	"fmt"
)

// Profile is a JSON Schema Draft-v4 object which describes
// asset data type.
type Profile struct {
	Definition interface{}
}

// SchemaType is a helper for most used profile types.
type SchemaType string

// Available schemas
const (
	StringType  SchemaType = "string"
	IntegerType SchemaType = "integer"
	NumberType  SchemaType = "number"
	BooleanType SchemaType = "boolean"
)

// Errors
var (
	// ErrInvalidProfile is returned when profile is invalid
	ErrInvalidProfile = errors.New("Profile JSON is invalid")
)

// New creates a profile for a given schema type.
func New(schemaType SchemaType) Profile {
	p, _ := JSON(fmt.Sprintf(`{"type": "%s"}`, string(schemaType)))
	return p
}

// JSON creates a profile from a given JSON Schema.
func JSON(j string) (Profile, error) {
	if isJSON(j) {
		var mapped interface{}
		json.Unmarshal([]byte(j), &mapped)
		return Profile{Definition: mapped}, nil
	}
	return Profile{}, ErrInvalidProfile
}

func (p Profile) json() string {
	profile, _ := json.Marshal(p.Definition)
	return string(profile)
}

func (p Profile) String() string {
	return p.json()
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

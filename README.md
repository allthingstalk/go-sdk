# Go SDK [![Build Status](https://travis-ci.org/allthingstalk/go-sdk.svg?branch=master)](https://travis-ci.org/allthingstalk/go-sdk) [![Go Report Card](https://goreportcard.com/badge/github.com/allthingstalk/go-sdk)](https://goreportcard.com/report/github.com/allthingstalk/go-sdk) [![RELEASE](https://img.shields.io/github/release/allthingstalk/go-sdk.svg)](https://github.com/allthingstalk/go-sdk/releases/1.0.0) [![GoDoc](https://godoc.org/github.com/allthingstalk/go-sdk?status.svg)](https://godoc.org/github.com/allthingstalk/go-sdk) [![LICENSE](https://img.shields.io/github/license/allthingstalk/go-sdk.svg)](LICENSE)

AllThingsTalk Golang SDK provides APIs to implement AllThingsTalk devices.

## Installation

Install [Go][5a0ab892], and then:

```bash
go get -u github.com/allthingstalk/go-sdk
```

Import:

```go
import "github.com/allthingstalk/go-sdk"
```

## Quickstart

You should create an [AllThingsTalk Maker][7447b4f9] account, and a simple device. All examples require `DeviceID` and `DeviceToken`, which can be found under Device Settings.

```go
package main

import (
	"github.com/allthingstalk/go-sdk"
	"time"
)

const (
	deviceID    = "<YOUR_DEVICE_ID>"
	deviceToken = "<YOUR_DEVICE_TOKEN>"
)

// A device which counts from 1 to 10 with a 1-second interval
func main() {
	// initialize a device
	device, err := sdk.NewDevice(deviceID, deviceToken)
	if err != nil {
		panic(err)
	}

	// add an asset named `counter` of `Integer` type
	counter, err := device.AddInteger("counter")
	if err != nil {
		panic(err)
	}
	
	// just count from 1 to 10 and send that value to Maker
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		device.Publish(counter, i)
	}
}
```

## HowTo

Please go through [examples][cfdc356c] for more comprehensive tutorials on how to use this package.

### Initializing a device

To manage a device, you must obtain a DeviceID (which is an unique device identifier within [AllThingsTalk Maker][7447b4f9]) and DeviceToken (which this SDK uses to obtain access to APIs), both of which can be found on Device Settings page within [AllThingsTalk Maker][7447b4f9].

Device can be initialized by:

```go
device, _ := sdk.NewDevice(deviceID, deviceToken)
```

You can also customize behind-the-scenes functionality, in this case - HTTP and MQTT clients, by optionally providing endpoints for them. By default these are securely connecting to Maker:

```go
device, _ := sdk.NewDevice("<DEVICE_ID>", "<DEVICE_TOKEN>", 
	sdk.WithHTTP("https://api.allthingstalk.io"),
	sdk.WithMQTT("ssl://api.allthingstalk.io:8883"))
```

### Creating Assets

Assets represent a typed property of a device which can be used either to sense some physical value (a `Sensor`), or send a value to a device (an `Actuator`). Each asset is described by `name` (unique identifier within a device) and `profile` (a [JSON Schema][61d54ea9] which describes data format).

#### Creating simple Sensors

Sensors of simple types (`integers`, `number`, `boolean`, `string`) can be created by `Add*` functions of a `Device`:

```go
// create a Sensor named `hello` with `integer` profile.
sensor, _ := device.AddInteger("hello")
```

Please check API reference for all available methods.

#### Creating Sensors and Actuators

`Sensors` and `Actuators` can be created with [NewSensor][1cb40ec7] and [NewActuator][db6f36c9] functions. These functions expect `name` (unique identifier within a device) and `profile` (JSON Schema describing data).

Afterwards, just to `device.Add(...)` to create that asset on the AllThingsTalk maker:

```go
numberSensor := sdk.NewSensor("number", profile.Number())
device.Add(numberSensor)
```

#### Profiles

```go
import "github.com/allthingstalk/go-sdk/profile"
```

Profiles are JSON Schemas which describe kind of data an Asset is working with. When creating an asset, we have to specify it's profile as well. Simplest way of doing this is using one of the pre-set types in `profile` package:

```go
profile := profile.String()
```

These functions just wrap `profile.New(...)` and provide schema type.

Alternatively, you can specify a complete JSON schema for more complex profile types:

```go

// a profile describing a Geo location
const locationProfile = `
{
   "type":"object",
   "properties":{
	  "lat":  {"type": "number"},
	  "long": {"type": "number"}
   }
}`

// try creating it from JSON string
prof, err := profile.JSON(locationProfile)
if err != nil {
	panic(err)
}

// create a location sensor asset with a location profile
location := sdk.NewAsset("location", sdk.Sensor, prof)
```

Check out [profile package][5d088019] for more details.

### Publishing states

You can publish sensor values using a `device.Publish(...)` function:

```go
device.Publish(sensor, "value")
```

Since all sensor state changes carry a timestamp, by default it's set to current UTC time if it's not provided. In order to set a custom one, just use `device.PublishState(...)`:

```go
device.PublishState(sensor, State{Value: "hello", Timestamp: time.Now().UTC()})
```

### Listening to commands

AllThingsTalk Maker can also send commands to a device. You can set a global command handler which will be invoked every time command is received:

```go
func onMessage(command sdk.Command) {
	fmt.Printf("Received command for %s: %v with timestamp %s\n", command.Name, command.Value, command.Timestamp)
}

func main() {
	// ...
	device.SetCommandHandler(onMessage)
}
```

### Error handling

Most of the SDK APIs return an `error` in case something goes wrong. As per Go's idiomatic usage, it's always recommended to use something like:

```go
if err != nil {
	// handle error...
}
```

To check for any error conditions. This library uses plain `error` model. If you need stack traces, or extended functionality, your production application should probably use [pkg/errors][2f9b8fc1] or similar.

### Logging

You can attach loggers for debugging purposes:

```go
sdk.ERROR = log.New(os.Stdout, "", 0)
```

Available loggers are `ERROR`, `CRITICAL`, `WARN`, `INFO` and `DEBUG`. Go's logging infrastructure being limited as is, it's generally recommended that [go-logging][5f643e4e] or similar is used in production applications.

[1cb40ec7]: https://godoc.org/github.com/allthingstalk/go-sdk/#NewSensor "NewSensor"
[2f9b8fc1]: https://github.com/pkg/errors "pkg/errors"
[5a0ab892]: https://golang.org/doc/install "Go"
[5d088019]: https://godoc.org/github.com/github.com/allthingstalk/go-sdk/profile "profile package"
[5f643e4e]: https://github.com/op/go-logging "go-logging"
[61d54ea9]: http://json-schema.org/ "JSON Schema"
[7447b4f9]: https://maker.allthingstalk.com "AllThingsTalk Maker"
[cfdc356c]: examples/ "examples"
[db6f36c9]: https://godoc.org/github.com/allthingstalk/go-sdk/#NewActuator "NewActuator"

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

/*
Package sdk implements a device SDK for AllThingsTalk IoT Platform.

AllThingsTalk Maker is a IoT device prototyping platform and this SDK aims to help developers who prefer Go language use it.

Let's start with the simplest example possible - creating a device which counts from 1 to 10. Start off by going to AllThingsTalk Maker and creating your account and a device (a Custom device will do just fine). Next, let's write some code:

	import (
		"fmt"
		"time"
		"github.com/allthingstalk/go-sdk"
	)

	func main() {
		device, err := sdk.NewDevice("<DEVICE_ID>", "<DEVICE_TOKEN>")
		if err != nil {
			panic(err)
		}

		counter, _ := device.AddInteger("counter")

		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println(i)
			device.Publish(counter, i)
		}
	}

You should replace <DEVICE_ID> and <DEVICE_TOKEN> with values that can be found in Device Settings page. If you run the example, the 'counter' asset should be created, and you should see it counting from 1 to 10.

That's it! For more comprehensive examples please check out the repository's README.

*/
package sdk

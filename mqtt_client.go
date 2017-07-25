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
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"regexp"
	"time"
)

type mqttClient struct {
	username string
	password string
	Client   mqtt.Client
	device   *Device
}

var topicParser = regexp.MustCompile(`device/(?P<device_id>\w+?)/asset/(?P<asset_name>\w+?)/command`)

func newMqttClient(device *Device) (*mqttClient, error) {
	username, password := device.token, device.token

	DEBUG.Printf("Using MQTT Username/Password: %s\n", username)

	opts := mqtt.NewClientOptions().AddBroker(device.options.mqttServer.String())
	opts.SetKeepAlive(10 * time.Second)
	opts.SetUsername(username)
	opts.SetPassword(password)
	client := mqtt.NewClient(opts)

	// maybe not panic, but return an error instead
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("Unable to connect to MQTT: %v", token.Error())
	}

	return &mqttClient{
		username: username,
		password: password,
		Client:   client,
		device:   device,
	}, nil
}

func (client *mqttClient) publish(device *Device, asset *Asset, state State) {
	topic := formatStateTopic(device, asset)
	payload := formatPayload(state)

	DEBUG.Printf("Publishings %v of asset %v to %s", payload, asset, topic)
	t := client.Client.Publish(topic, 0, false, payload)
	t.Wait()
}

func (client *mqttClient) subscribe(device *Device) {
	t := client.Client.Subscribe(formatCommandTopic(device), 0, client.onMessageReceived)
	t.Wait()
}

func (client *mqttClient) onMessageReceived(mqtt mqtt.Client, message mqtt.Message) {
	DEBUG.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())

	// topic must match expected regex
	ok, parts := matchInboundTopic(message.Topic())
	if ok != true {
		DEBUG.Printf("Inbound topic %s does not match expected topic.", message.Topic())
		return
	}

	command := Command{}
	if err := json.Unmarshal(message.Payload(), &command); err == nil {
		command.Name = parts["asset_name"]
		client.device.onCommand(command)
	}
}

func formatPayload(state State) string {
	payload, _ := json.Marshal(state)
	return string(payload)
}

func formatCommandTopic(device *Device) string {
	return fmt.Sprintf("device/%s/asset/+/command", device.id)
}

func formatStateTopic(device *Device, asset *Asset) string {
	return fmt.Sprintf("device/%s/asset/%s/state", device.id, asset.Name)
}

func matchInboundTopic(topic string) (bool, map[string]string) {
	if !topicParser.MatchString(topic) {
		return false, nil
	}

	match := topicParser.FindStringSubmatch(topic)
	result := make(map[string]string)
	for i, name := range topicParser.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}

	return true, result
}

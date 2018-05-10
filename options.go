package sdk

import (
	"net/url"
)

// Options represents general device options, like URL endpoints.
type Options struct {
	httpServer *url.URL
	mqttServer *url.URL
}

// Option is a single option configuration of a device
type Option interface {
	apply(*Device)
}

type optionFunc func(*Device)

func (f optionFunc) apply(device *Device) {
	f(device)
}

const httpEndpoint string = "https://api.allthingstalk.io"
const mqttEndpoint string = "ssl://api.allthingstalk.io:8883"

// NewOptions creates default options.
func newOptions() *Options {
	return newCustomOptions(httpEndpoint, mqttEndpoint)
}

// NewCustomOptions creates options with HTTP and MQTT endpoints.
func newCustomOptions(http string, mqtt string) *Options {
	httpURL, _ := url.Parse(http)
	mqttURL, _ := url.Parse(mqtt)
	return &Options{
		httpServer: httpURL,
		mqttServer: mqttURL,
	}
}

// WithHTTP sets HTTP REST API endpoint.
func WithHTTP(endpoint string) Option {
	return optionFunc(func(device *Device) {
		u, _ := url.Parse(endpoint)
		device.options.httpServer = u
	})
}

// WithMQTT sets MQTT API endpoint.
func WithMQTT(endpoint string) Option {
	return optionFunc(func(device *Device) {
		u, _ := url.Parse(endpoint)
		device.options.mqttServer = u
	})
}

// SetAPI sets HTTP REST API endpoint.
func (o *Options) SetAPI(endpoint string) *Options {
	u, _ := url.Parse(endpoint)
	o.httpServer = u

	return o
}

// SetMqtt sets MQTT API endpoint.
func (o *Options) SetMqtt(endpoint string) *Options {
	u, _ := url.Parse(endpoint)
	o.mqttServer = u

	return o
}

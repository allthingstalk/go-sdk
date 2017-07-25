package sdk

import (
	"net/url"
)

// Options represents general device options, like URL endpoints.
type Options struct {
	httpServer *url.URL
	mqttServer *url.URL
}

const httpEndpoint string = "http://api.allthingstalk.io"
const mqttEndpoint string = "ssl://api.allthingstalk.io:8883"

// NewOptions creates default options.
func NewOptions() *Options {
	return NewCustomOptions(httpEndpoint, mqttEndpoint)
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

// NewCustomOptions creates options with HTTP and MQTT endpoints.
func NewCustomOptions(http string, mqtt string) *Options {
	httpURL, _ := url.Parse(http)
	mqttURL, _ := url.Parse(mqtt)
	return &Options{
		httpServer: httpURL,
		mqttServer: mqttURL,
	}
}

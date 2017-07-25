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
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type httpClient struct {
	AssetService *assetService
}

func newHTTPClient(c *http.Client, device *Device) (*httpClient, error) {
	return &httpClient{
		AssetService: newAssetService(c, device),
	}, nil
}

type assetService struct {
	sling *sling.Sling
}

type apiError struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func newAssetService(httpClient *http.Client, device *Device) *assetService {
	apiEndpoint := device.options.httpServer.String()
	DEBUG.Printf("Using API endpoint: %s", apiEndpoint)

	return &assetService{
		sling: sling.New().Client(httpClient).Base(apiEndpoint),
	}
}

func (service *assetService) addAsset(device *Device, asset *Asset) error {
	path := fmt.Sprintf("device/%s/asset/%s", device.id, asset.Name)
	requestError := new(apiError)

	resp, err := service.sling.New().
		Set("Authorization", fmt.Sprintf("Bearer %s", device.token)).
		Put(path).
		BodyJSON(asset).
		Receive(nil, requestError)

	if err != nil {
		ERROR.Printf("Unable to add asset due to an error: %s\n", err)
		return fmt.Errorf("Unable to add asset due to an error: %s", err)
	}

	if !isResponseSuccess(resp) {
		ERROR.Printf("API rejected AddAsset with code %d: %s\n", resp.StatusCode, requestError.Error)
		return fmt.Errorf("Could not add asset, api returned %d, %v", resp.StatusCode, requestError.Error)
	}

	INFO.Printf("Added asset %v\n", asset)
	return nil
}

func isResponseSuccess(resp *http.Response) bool {
	return resp.StatusCode/100 == 2
}

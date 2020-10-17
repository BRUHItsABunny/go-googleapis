package go_googleapis

import (
	"errors"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"net/http"
	"net/url"
	"strconv"
)

func GetMapsClient(licenseKey string) *MapsClient {
	httpClient := gokhttp.GetHTTPClient(nil)
	return &MapsClient{BaseClient{Client: &httpClient}, License{
		LicenseType: 2,
		LicenseKey:  licenseKey,
	}}
}

func (mc *MapsClient) Directions(srcLat, srcLng, dstLat, dstLng string) ([]DirectionsRoute, error) {
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = new(MapsResponse)
	)

	params := url.Values{
		"origin":      []string{srcLat + "," + srcLng},
		"destination": []string{dstLat + "," + dstLng},
		"sensor":      []string{"false"},
		"key":         []string{mc.License.LicenseKey},
	}

	req, err = mc.Client.MakeRawPOSTRequest(urlMaps+endpointMapsDirections, params, nil, map[string]string{"User-Agent": headerMapsUserAgent})

	if err == nil {
		resp, err = mc.Client.Do(req)
		if err == nil {
			err = resp.Object(result)
			if err == nil {
				if result.Status == "OK" {
					return result.Routes, nil
				} else {
					err = errors.New(*result.ErrorMessage)
				}
			}
		}
	}
	return nil, err
}

func (mc *MapsClient) Nearby(latitude, longitude, placeType string, radius int) ([]NearbyResult, string, error) {
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = new(MapsResponse)
	)

	params := url.Values{
		"location": []string{latitude + "," + longitude},
		"radius":   []string{strconv.Itoa(radius)},
		"type":     []string{placeType},
		"sensor":   []string{"dalse"},
		"key":      []string{mc.License.LicenseKey},
	}

	req, err = mc.Client.MakeRawPOSTRequest(urlMaps+endpointMapsNearBy, params, nil, map[string]string{"User-Agent": headerMapsUserAgent})

	if err == nil {
		resp, err = mc.Client.Do(req)
		if err == nil {
			// JSON, not sure yet but will parse as structs
			err = resp.Object(result)
			if err == nil {
				if result.Status == "OK" {
					return result.Results, *result.NextPageToken, nil
				} else {
					err = errors.New(*result.ErrorMessage)
				}
			}
		}
	}
	return nil, "", err
}

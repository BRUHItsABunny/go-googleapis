package go_googleapis

import (
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"net/http"
	"net/url"
)

func GetAutoCompleteClient() *AutoCompleteClient {
	httpClient := gokhttp.GetHTTPClient(nil)
	return &AutoCompleteClient{BaseClient{Client: &httpClient}}
}

func (acc *AutoCompleteClient) Suggest(text, language string) ([]string, error) {
	// https://clients1.google.com/complete/search?ds=translate&hjson=t&q=tes&requiredfields=tl%3Anl&hl=en&ie=UTF-8&oe=UTF-8&client=translate-android
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = make([]PurpleAutoCompleteResponse, 0)
	)

	params := url.Values{
		"ds":             []string{"translate"},         // destination
		"hjson":          []string{"t"},                 // no idea
		"q":              []string{text},                // query, just the text you are going to translate
		"requiredfields": []string{"tl:" + language},    // no idea, it includes language though
		"hl":             []string{"en"},                // Locale
		"ie":             []string{"UTF-8"},             // IN encoding
		"oe":             []string{"UTF-8"},             // OUT encoding
		"client":         []string{"translate-android"}, // client, at: android translate?
	}

	req, err = acc.Client.MakeGETRequest(urlComplete+endpointComplete, params, map[string]string{"User-Agent": headerTranslateUserAgent})

	if err == nil {
		resp, err = acc.Client.Do(req)
		if err == nil {
			// JSON, arrays
			r := make([]string, 0)
			err = resp.Object(&result)
			for _, e := range result[1].UnionArrayArray {
				r = append(r, *e[0].String)
			}
			return r, err
		}
	}
	return nil, err
}

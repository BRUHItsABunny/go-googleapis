package go_googleapis

import (
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetTranslateClient() *TranslateClient {
	httpClient := gokhttp.GetHTTPClient(nil)
	return &TranslateClient{BaseClient{Client: &httpClient}}
}

func (tc *TranslateClient) Translate(text, srcLang, dstLang string) (*Translated, error) {
	// https://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&tl=en&dt=t&q=konijn
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = new(TranslateResult)
	)

	params := url.Values{
		"dj":     []string{"1"},                                                   // no idea
		"q":      []string{text},                                                  // query, just the text you are going to translate
		"sl":     []string{srcLang},                                               // Source language
		"tl":     []string{dstLang},                                               // Target language
		"hl":     []string{"en-US"},                                               // Locale
		"ie":     []string{"UTF-8"},                                               // IN encoding
		"oe":     []string{"UTF-8"},                                               // OUT encoding
		"client": []string{"at"},                                                  // client, at: android translate?
		"dt":     []string{"t", "ld", "qca", "rm", "bd", "md", "ss", "ex", "sos"}, // no idea, probably what gives me better JSON results than with just "t", so scope(s)?
	}

	req, err = tc.Client.MakeGETRequest(urlTranslate+endpointTranslate, params, map[string]string{"User-Agent": headerTranslateUserAgent})

	if err == nil {
		resp, err = tc.Client.Do(req)
		if err == nil {
			err = resp.Object(result)
			if err == nil {
				trans := ""
				// Loop through sentences
				for _, e := range result.Sentences {
					trans = trans + *e.Trans
				}
				return &Translated{
					SrcLang:   *result.Src,
					DstLang:   dstLang,
					Origin:    text,
					Translate: trans,
				}, err
			}
		}
	}
	return nil, err
}

func (tc *TranslateClient) TTS(text, srcLang string) (string, error) {
	// https://translate.googleapis.com/translate_tts?client=gtx&tl=nl&dt=t&q=konijn&textlen=6&total=1&idx=0

	var (
		err      error
		req      *http.Request
		resp     *gokhttp.HttpResponse
		fileOut  *os.File
		mp3Bytes []byte
	)

	strSplit := strings.Split(text, "\n")
	fileName := "tts_" + strconv.FormatInt(time.Now().Unix(), 10) + ".mp3"
	fileOut, err = os.Create(fileName)

	if err == nil {
		for i, e := range strSplit {
			params := url.Values{
				"ie":      []string{"UTF-8"},                     // IN encoding
				"client":  []string{"at"},                        // client, at: android translate?
				"q":       []string{e},                           // query, just the text you are going to translate
				"tl":      []string{srcLang},                     // Target language
				"total":   []string{strconv.Itoa(len(strSplit))}, // Constant 1 for some reason?
				"idx":     []string{strconv.Itoa(i)},
				"textlen": []string{strconv.Itoa(len(e))},
				"prev":    []string{"input"},
			}

			req, err = tc.Client.MakeGETRequest(urlTranslate+endpointTranslateTTS, params, map[string]string{"User-Agent": headerTranslateUserAgent})

			if err == nil {
				resp, err = tc.Client.Do(req)
				if err == nil {
					// MP3, can we concat these? yes and it will be playable but headers will be wrong
					mp3Bytes, err = resp.Bytes()
					if err == nil {
						// Should I just return the array of bytes or *os.File object instead?
						_, err = fileOut.Write(mp3Bytes)
						if err != nil {
							break
						}
					}
				}
			}
		}
	}

	if err != nil {
		return "", err
	}
	return fileName, err
}

package go_googleapis

import (
	"bytes"
	"encoding/json"
	"errors"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"net/http"
	"net/url"
)

func GetVisionClient(licenseKey, androidPackage, androidCert string) *VisionClient {
	httpClient := gokhttp.GetHTTPClient(nil)
	return &VisionClient{BaseClient{Client: &httpClient}, License{
		LicenseType:    1,
		LicenseKey:     licenseKey,
		LicensePackage: androidPackage,
		LicenseCert:    androidCert,
	}}
}

func (vc *VisionClient) DetectLabel(imageJob *ImageJob) ([]Annotation, error) {
	var (
		err error
		//data string
		req           *http.Request
		resp          *gokhttp.HttpResponse
		result        = new(VisionResponse)
		postBodyBytes []byte
	)

	if imageJob != nil {
		if imageJob.Content == "" {
			err = imageJob.GetImageBase64()
		}
		if err == nil {
			params := url.Values{
				"key": []string{vc.License.LicenseKey},
			}

			postBodyBytes, err = vc.makePOSTBody(imageJob, "LABEL_DETECTION")

			if err == nil {
				req, err = vc.Client.MakeRawPOSTRequest(urlVision+endpointVisionImages, params, bytes.NewReader(postBodyBytes), map[string]string{"User-Agent": headerVisionUserAgent, "x-goog-api-client": headerVisionGoogleClient, "x-android-package": vc.License.LicensePackage, "x-android-cert": vc.License.LicenseCert, "Content-Type": "application/json; charset=UTF-8"})

				if err == nil {
					resp, err = vc.Client.Do(req)
					if err == nil {
						err = resp.Object(result)
						if err == nil {
							if result.Error == nil {
								return result.Responses[0].LabelAnnotations, nil
							}
							err = errors.New(result.Error.Message)
						}
					}
				}
			}
		}
	} else {
		err = errors.New("ImageJob empty")
	}

	return nil, err
}

func (vc *VisionClient) DetectSafeSearch(imageJob *ImageJob) (*SafeSearchAnnotation, error) {
	var (
		err error
		//data string
		req           *http.Request
		resp          *gokhttp.HttpResponse
		result        = new(VisionResponse)
		postBodyBytes []byte
	)

	if imageJob != nil {
		if imageJob.Content == "" {
			err = imageJob.GetImageBase64()
		}
		if err == nil {
			params := url.Values{
				"key": []string{vc.License.LicenseKey},
			}

			postBodyBytes, err = vc.makePOSTBody(imageJob, "SAFE_SEARCH_DETECTION")

			if err == nil {
				req, err = vc.Client.MakeRawPOSTRequest(urlVision+endpointVisionImages, params, bytes.NewReader(postBodyBytes), map[string]string{"User-Agent": headerVisionUserAgent, "x-goog-api-client": headerVisionGoogleClient, "x-android-package": vc.License.LicensePackage, "x-android-cert": vc.License.LicenseCert, "Content-Type": "application/json; charset=UTF-8"})

				if err == nil {
					resp, err = vc.Client.Do(req)
					if err == nil {
						err = resp.Object(result)
						if err == nil {
							if result.Error == nil {
								return result.Responses[0].SafeSearchAnnotation, nil
							}
							err = errors.New(result.Error.Message)
						}
					}
				}
			}
		}
	} else {
		err = errors.New("ImageJob empty")
	}

	return nil, err
}

func (vc *VisionClient) ImageProperties(imageJob *ImageJob) (*ImagePropertiesAnnotation, error) {
	var (
		err error
		//data string
		req           *http.Request
		resp          *gokhttp.HttpResponse
		result        = new(VisionResponse)
		postBodyBytes []byte
	)

	if imageJob != nil {
		if imageJob.Content == "" {
			err = imageJob.GetImageBase64()
		}
		if err == nil {
			params := url.Values{
				"key": []string{vc.License.LicenseKey},
			}

			postBodyBytes, err = vc.makePOSTBody(imageJob, "IMAGE_PROPERTIES")

			if err == nil {
				req, err = vc.Client.MakeRawPOSTRequest(urlVision+endpointVisionImages, params, bytes.NewReader(postBodyBytes), map[string]string{"User-Agent": headerVisionUserAgent, "x-goog-api-client": headerVisionGoogleClient, "x-android-package": vc.License.LicensePackage, "x-android-cert": vc.License.LicenseCert, "Content-Type": "application/json; charset=UTF-8"})

				if err == nil {
					resp, err = vc.Client.Do(req)
					if err == nil {
						err = resp.Object(result)
						if err == nil {
							if result.Error == nil {
								return result.Responses[0].ImagePropertiesAnnotation, nil
							}
							err = errors.New(result.Error.Message)
						}
					}
				}
			}
		}
	} else {
		err = errors.New("ImageJob empty")
	}

	return nil, err
}

func (vc *VisionClient) DetectFace(imageJob *ImageJob) ([]FaceAnnotation, error) {
	var (
		err error
		//data string
		req           *http.Request
		resp          *gokhttp.HttpResponse
		result        = new(VisionResponse)
		postBodyBytes []byte
	)

	if imageJob != nil {
		if imageJob.Content == "" {
			err = imageJob.GetImageBase64()
		}
		if err == nil {
			params := url.Values{
				"key": []string{vc.License.LicenseKey},
			}

			postBodyBytes, err = vc.makePOSTBody(imageJob, "FACE_DETECTION")

			if err == nil {
				req, err = vc.Client.MakeRawPOSTRequest(urlVision+endpointVisionImages, params, bytes.NewReader(postBodyBytes), map[string]string{"User-Agent": headerVisionUserAgent, "x-goog-api-client": headerVisionGoogleClient, "x-android-package": vc.License.LicensePackage, "x-android-cert": vc.License.LicenseCert, "Content-Type": "application/json; charset=UTF-8"})

				if err == nil {
					resp, err = vc.Client.Do(req)
					if err == nil {
						err = resp.Object(result)
						if err == nil {
							if result.Error == nil {
								return result.Responses[0].FaceAnnotations, nil
							}
							err = errors.New(result.Error.Message)
						}
					}
				}
			}
		}
	} else {
		err = errors.New("ImageJob empty")
	}

	return nil, err
}

func (vc *VisionClient) DetectLandMark(imageJob *ImageJob) ([]LandmarkAnnotation, error) {
	var (
		err error
		//data string
		req           *http.Request
		resp          *gokhttp.HttpResponse
		result        = new(VisionResponse)
		postBodyBytes []byte
	)

	if imageJob != nil {
		if imageJob.Content == "" {
			err = imageJob.GetImageBase64()
		}
		if err == nil {
			params := url.Values{
				"key": []string{vc.License.LicenseKey},
			}

			postBodyBytes, err = vc.makePOSTBody(imageJob, "LANDMARK_DETECTION")

			if err == nil {
				req, err = vc.Client.MakeRawPOSTRequest(urlVision+endpointVisionImages, params, bytes.NewReader(postBodyBytes), map[string]string{"User-Agent": headerVisionUserAgent, "x-goog-api-client": headerVisionGoogleClient, "x-android-package": vc.License.LicensePackage, "x-android-cert": vc.License.LicenseCert, "Content-Type": "application/json; charset=UTF-8"})

				if err == nil {
					resp, err = vc.Client.Do(req)
					if err == nil {
						err = resp.Object(result)
						if err == nil {
							if result.Error == nil {
								return result.Responses[0].LandmarkAnnotations, nil
							}
							err = errors.New(result.Error.Message)
						}
					}
				}
			}
		}
	} else {
		err = errors.New("ImageJob empty")
	}

	return nil, err
}

func (vc *VisionClient) DetectLogo(imageJob *ImageJob) ([]Annotation, error) {
	var (
		err error
		//data string
		req           *http.Request
		resp          *gokhttp.HttpResponse
		result        = new(VisionResponse)
		postBodyBytes []byte
	)

	if imageJob != nil {
		if imageJob.Content == "" {
			err = imageJob.GetImageBase64()
		}
		if err == nil {
			params := url.Values{
				"key": []string{vc.License.LicenseKey},
			}

			postBodyBytes, err = vc.makePOSTBody(imageJob, "LOGO_DETECTION")

			if err == nil {
				req, err = vc.Client.MakeRawPOSTRequest(urlVision+endpointVisionImages, params, bytes.NewReader(postBodyBytes), map[string]string{"User-Agent": headerVisionUserAgent, "x-goog-api-client": headerVisionGoogleClient, "x-android-package": vc.License.LicensePackage, "x-android-cert": vc.License.LicenseCert, "Content-Type": "application/json; charset=UTF-8"})

				if err == nil {
					resp, err = vc.Client.Do(req)
					if err == nil {
						err = resp.Object(result)
						if err == nil {
							if result.Error == nil {
								return result.Responses[0].LogoAnnotations, nil
							}
							err = errors.New(result.Error.Message)
						}
					}
				}
			}
		}
	} else {
		err = errors.New("ImageJob empty")
	}

	return nil, err
}

func (vc *VisionClient) DetectDocumentText(imageJob *ImageJob) ([]TextAnnotation, *FullTextAnnotation, error) {
	var (
		err error
		//data string
		req           *http.Request
		resp          *gokhttp.HttpResponse
		result        = new(VisionResponse)
		postBodyBytes []byte
	)

	if imageJob != nil {
		if imageJob.Content == "" {
			err = imageJob.GetImageBase64()
		}
		if err == nil {
			params := url.Values{
				"key": []string{vc.License.LicenseKey},
			}

			postBodyBytes, err = vc.makePOSTBody(imageJob, "DOCUMENT_TEXT_DETECTION")

			if err == nil {
				req, err = vc.Client.MakeRawPOSTRequest(urlVision+endpointVisionImages, params, bytes.NewReader(postBodyBytes), map[string]string{"User-Agent": headerVisionUserAgent, "x-goog-api-client": headerVisionGoogleClient, "x-android-package": vc.License.LicensePackage, "x-android-cert": vc.License.LicenseCert, "Content-Type": "application/json; charset=UTF-8"})

				if err == nil {
					resp, err = vc.Client.Do(req)
					if err == nil {
						err = resp.Object(result)
						if err == nil {
							if result.Error == nil {
								return result.Responses[0].TextAnnotations, result.Responses[0].FullTextAnnotation, nil
							}
							err = errors.New(result.Error.Message)
						}
					}
				}
			}
		}
	} else {
		err = errors.New("ImageJob empty")
	}

	return nil, nil, err
}

func (vc *VisionClient) DetectText(imageJob *ImageJob) ([]TextAnnotation, *FullTextAnnotation, error) {
	var (
		err error
		//data string
		req           *http.Request
		resp          *gokhttp.HttpResponse
		result        = new(VisionResponse)
		postBodyBytes []byte
	)

	if imageJob != nil {
		if imageJob.Content == "" {
			err = imageJob.GetImageBase64()
		}
		if err == nil {
			params := url.Values{
				"key": []string{vc.License.LicenseKey},
			}

			postBodyBytes, err = vc.makePOSTBody(imageJob, "TEXT_DETECTION")

			if err == nil {
				req, err = vc.Client.MakeRawPOSTRequest(urlVision+endpointVisionImages, params, bytes.NewReader(postBodyBytes), map[string]string{"User-Agent": headerVisionUserAgent, "x-goog-api-client": headerVisionGoogleClient, "x-android-package": vc.License.LicensePackage, "x-android-cert": vc.License.LicenseCert, "Content-Type": "application/json; charset=UTF-8"})

				if err == nil {
					resp, err = vc.Client.Do(req)
					if err == nil {
						err = resp.Object(result)
						if err == nil {
							if result.Error == nil {
								return result.Responses[0].TextAnnotations, result.Responses[0].FullTextAnnotation, nil
							}
							err = errors.New(result.Error.Message)
						}
					}
				}
			}
		}
	} else {
		err = errors.New("ImageJob empty")
	}

	return nil, nil, err
}

func (vc *VisionClient) DetectWeb(imageJob *ImageJob) (*WebDetectionAnnotation, error) {
	var (
		err error
		//data string
		req           *http.Request
		resp          *gokhttp.HttpResponse
		result        = new(VisionResponse)
		postBodyBytes []byte
	)

	if imageJob != nil {
		if imageJob.Content == "" {
			err = imageJob.GetImageBase64()
		}
		if err == nil {
			params := url.Values{
				"key": []string{vc.License.LicenseKey},
			}

			postBodyBytes, err = vc.makePOSTBody(imageJob, "WEB_DETECTION")

			if err == nil {
				req, err = vc.Client.MakeRawPOSTRequest(urlVision+endpointVisionImages, params, bytes.NewReader(postBodyBytes), map[string]string{"User-Agent": headerVisionUserAgent, "x-goog-api-client": headerVisionGoogleClient, "x-android-package": vc.License.LicensePackage, "x-android-cert": vc.License.LicenseCert, "Content-Type": "application/json; charset=UTF-8"})

				if err == nil {
					resp, err = vc.Client.Do(req)
					if err == nil {
						err = resp.Object(result)
						if err == nil {
							if result.Error == nil {
								return result.Responses[0].WebDetection, nil
							}
							err = errors.New(result.Error.Message)
						}
					}
				}
			}
		}
	} else {
		err = errors.New("ImageJob empty")
	}

	return nil, err
}

func (vc *VisionClient) makePOSTBody(job *ImageJob, actionType string) ([]byte, error) {
	visionReq := VisionRequest{
		Features: []VisionFeature{{
			MaxResults: 10,
			Type:       actionType,
		}},
		Image: *job,
	}

	postBody := struct {
		Requests []VisionRequest `json:"requests"`
	}{Requests: []VisionRequest{visionReq}}
	return json.Marshal(&postBody)
}

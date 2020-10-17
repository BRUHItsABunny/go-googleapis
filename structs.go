package go_googleapis

import (
	"bytes"
	"encoding/json"
	"errors"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"io"
)

type License struct {
	LicenseType    int    `json:"type"`
	LicenseKey     string `json:"key"`
	LicensePackage string `json:"package"`
	LicenseCert    string `json:"cert"`
}

type BaseClient struct {
	Client *gokhttp.HttpClient
}

type TranslateClient struct {
	BaseClient
}

type AutoCompleteClient struct {
	BaseClient
}

type VisionClient struct {
	BaseClient
	License License
}

type MapsClient struct {
	BaseClient
	License License
}

// Translation
type Translated struct {
	SrcLang   string
	DstLang   string
	Origin    string
	Translate string
}

type TranslateResult struct {
	Sentences  []Sentence `json:"sentences,omitempty"`
	Src        *string    `json:"src,omitempty"`
	Confidence *float64   `json:"confidence,omitempty"`
	Spell      *Spell     `json:"spell,omitempty"`
	LdResult   *LdResult  `json:"ld_result,omitempty"`
}

type LdResult struct {
	Srclangs            []string  `json:"srclangs,omitempty"`
	SrclangsConfidences []float64 `json:"srclangs_confidences,omitempty"`
	ExtendedSrclangs    []string  `json:"extended_srclangs,omitempty"`
}

type Sentence struct {
	Trans                      *string                      `json:"trans,omitempty"`
	Orig                       *string                      `json:"orig,omitempty"`
	Backend                    *int64                       `json:"backend,omitempty"`
	ModelSpecification         []ModelSpecification         `json:"model_specification,omitempty"`
	TranslationEngineDebugInfo []TranslationEngineDebugInfo `json:"translation_engine_debug_info,omitempty"`
}

type ModelSpecification struct {
}

type TranslationEngineDebugInfo struct {
	ModelTracking *ModelTracking `json:"model_tracking,omitempty"`
}

type ModelTracking struct {
	CheckpointMd5 *string `json:"checkpoint_md5,omitempty"`
	LaunchDoc     *string `json:"launch_doc,omitempty"`
}

type Spell struct {
	SpellHTMLRes   *string `json:"spell_html_res,omitempty"`
	SpellRes       *string `json:"spell_res,omitempty"`
	CorrectionType []int64 `json:"correction_type,omitempty"`
	Confident      *bool   `json:"confident,omitempty"`
}

// AutoComplete
type AutoCompleteResponse []PurpleAutoCompleteResponse

func UnmarshalAutoCompleteResponse(data []byte) (AutoCompleteResponse, error) {
	var r AutoCompleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AutoCompleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AutoCompleteResponseClass struct {
	Q string `json:"q"`
}

type PurpleAutoCompleteResponse struct {
	AutoCompleteResponseClass *AutoCompleteResponseClass
	String                    *string
	UnionArrayArray           [][]AutoCompleteResponseAutoCompleteResponseUnion
}

func (x *PurpleAutoCompleteResponse) UnmarshalJSON(data []byte) error {
	x.UnionArrayArray = nil
	x.AutoCompleteResponseClass = nil
	var c AutoCompleteResponseClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.UnionArrayArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.AutoCompleteResponseClass = &c
	}
	return nil
}

func (x *PurpleAutoCompleteResponse) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.UnionArrayArray != nil, x.UnionArrayArray, x.AutoCompleteResponseClass != nil, x.AutoCompleteResponseClass, false, nil, false, nil, false)
}

type AutoCompleteResponseAutoCompleteResponseUnion struct {
	Integer *int64
	String  *string
}

func (x *AutoCompleteResponseAutoCompleteResponseUnion) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *AutoCompleteResponseAutoCompleteResponseUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}

// Vision

type VisionError struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Errors  []ErrorElement `json:"errors"`
	Status  string         `json:"status"`
}

type ErrorElement struct {
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
}

type ImageJob struct {
	Bytes   []byte    `json:"-"`
	Reader  io.Reader `json:"-"`
	Content string    `json:"content"`
}

type VisionFeature struct {
	MaxResults int    `json:"maxResults"`
	Type       string `json:"type"`
}

type VisionRequest struct {
	Features []VisionFeature `json:"features"`
	Image    ImageJob        `json:"image"`
}

type VisionResponse struct {
	Responses []Response   `json:"responses,omitempty"`
	Error     *VisionError `json:"error,omitempty"`
}

type Response struct {
	LabelAnnotations          []Annotation               `json:"labelAnnotations,omitempty"`
	SafeSearchAnnotation      *SafeSearchAnnotation      `json:"safeSearchAnnotation,omitempty"`
	ImagePropertiesAnnotation *ImagePropertiesAnnotation `json:"imagePropertiesAnnotation,omitempty"`
	CropHintsAnnotation       *CropHintsAnnotation       `json:"cropHintsAnnotation,omitempty"`
	WebDetection              *WebDetectionAnnotation    `json:"webDetection,omitempty"`
	LandmarkAnnotations       []LandmarkAnnotation       `json:"landmarkAnnotations,omitempty"`
	FaceAnnotations           []FaceAnnotation           `json:"faceAnnotations,omitempty"`
	LogoAnnotations           []Annotation               `json:"logoAnnotations,omitempty"`
	TextAnnotations           []TextAnnotation           `json:"textAnnotations,omitempty"`
	FullTextAnnotation        *FullTextAnnotation        `json:"fullTextAnnotation,omitempty"`
}

type CropHintsAnnotation struct {
	CropHints []CropHint `json:"cropHints"`
}

type CropHint struct {
	BoundingPoly       BoundingPoly `json:"boundingPoly"`
	Confidence         float64      `json:"confidence"`
	ImportanceFraction float64      `json:"importanceFraction"`
}

type BoundingPoly struct {
	Vertices []Vertex `json:"vertices"`
}

type Vertex struct {
	X int64  `json:"x"`
	Y *int64 `json:"y,omitempty"`
}

type FaceAnnotation struct {
	BoundingPoly           BoundingPoly `json:"boundingPoly"`
	FdBoundingPoly         BoundingPoly `json:"fdBoundingPoly"`
	Landmarks              []Landmark   `json:"landmarks"`
	RollAngle              float64      `json:"rollAngle"`
	PanAngle               float64      `json:"panAngle"`
	TiltAngle              float64      `json:"tiltAngle"`
	DetectionConfidence    float64      `json:"detectionConfidence"`
	LandmarkingConfidence  float64      `json:"landmarkingConfidence"`
	JoyLikelihood          string       `json:"joyLikelihood"`
	SorrowLikelihood       string       `json:"sorrowLikelihood"`
	AngerLikelihood        string       `json:"angerLikelihood"`
	SurpriseLikelihood     string       `json:"surpriseLikelihood"`
	UnderExposedLikelihood string       `json:"underExposedLikelihood"`
	BlurredLikelihood      string       `json:"blurredLikelihood"`
	HeadwearLikelihood     string       `json:"headwearLikelihood"`
}

type Landmark struct {
	Type     string   `json:"type"`
	Position Position `json:"position"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type FullTextAnnotation struct {
	Pages []Page `json:"pages"`
	Text  string `json:"text"`
}

type Page struct {
	Property ParagraphProperty `json:"property"`
	Width    int64             `json:"width"`
	Height   int64             `json:"height"`
	Blocks   []Block           `json:"blocks"`
}

type Block struct {
	Property    ParagraphProperty `json:"property"`
	BoundingBox BoundingPoly      `json:"boundingBox"`
	Paragraphs  []Paragraph       `json:"paragraphs"`
	BlockType   string            `json:"blockType"`
	Confidence  *float64          `json:"confidence,omitempty"`
}

type Paragraph struct {
	Property    ParagraphProperty `json:"property"`
	BoundingBox BoundingPoly      `json:"boundingBox"`
	Words       []Word            `json:"words"`
	Confidence  *float64          `json:"confidence,omitempty"`
}

type ParagraphProperty struct {
	DetectedLanguages []PurpleDetectedLanguage `json:"detectedLanguages"`
}

type PurpleDetectedLanguage struct {
	LanguageCode string  `json:"languageCode"`
	Confidence   float64 `json:"confidence"`
}

type Word struct {
	Property    WordProperty `json:"property"`
	BoundingBox BoundingPoly `json:"boundingBox"`
	Symbols     []Symbol     `json:"symbols"`
	Confidence  *float64     `json:"confidence,omitempty"`
}

type WordProperty struct {
	DetectedLanguages []FluffyDetectedLanguage `json:"detectedLanguages"`
}

type FluffyDetectedLanguage struct {
	LanguageCode string `json:"languageCode"`
}

type Symbol struct {
	Property    SymbolProperty `json:"property"`
	BoundingBox BoundingPoly   `json:"boundingBox"`
	Text        string         `json:"text"`
	Confidence  *float64       `json:"confidence,omitempty"`
}

type SymbolProperty struct {
	DetectedLanguages []FluffyDetectedLanguage `json:"detectedLanguages"`
	DetectedBreak     *DetectedBreak           `json:"detectedBreak,omitempty"`
}

type DetectedBreak struct {
	Type string `json:"type"`
}

type ImagePropertiesAnnotation struct {
	DominantColors DominantColors `json:"dominantColors"`
}

type DominantColors struct {
	Colors []ColorElement `json:"colors"`
}

type ColorElement struct {
	Color         ColorColor `json:"color"`
	Score         float64    `json:"score"`
	PixelFraction float64    `json:"pixelFraction"`
}

type ColorColor struct {
	Red   int64 `json:"red"`
	Green int64 `json:"green"`
	Blue  int64 `json:"blue"`
}

type Annotation struct {
	Mid          string        `json:"mid"`
	Description  string        `json:"description"`
	Score        float64       `json:"score"`
	Topicality   *float64      `json:"topicality,omitempty"`
	BoundingPoly *BoundingPoly `json:"boundingPoly,omitempty"`
}

type LandmarkAnnotation struct {
	Mid          string           `json:"mid"`
	Description  string           `json:"description"`
	Score        float64          `json:"score"`
	BoundingPoly BoundingPoly     `json:"boundingPoly"`
	Locations    []VisionLocation `json:"locations"`
}

type VisionLocation struct {
	LatLng LatLng `json:"latLng"`
}

type LatLng struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type SafeSearchAnnotation struct {
	Adult    string `json:"adult"`
	Spoof    string `json:"spoof"`
	Medical  string `json:"medical"`
	Violence string `json:"violence"`
	Racy     string `json:"racy"`
}

type TextAnnotation struct {
	Locale       *string      `json:"locale,omitempty"`
	Description  string       `json:"description"`
	BoundingPoly BoundingPoly `json:"boundingPoly"`
}

type WebDetectionAnnotation struct {
	WebEntities             []WebEntity              `json:"webEntities"`
	FullMatchingImages      []Image                  `json:"fullMatchingImages"`
	PartialMatchingImages   []Image                  `json:"partialMatchingImages"`
	PagesWithMatchingImages []PagesWithMatchingImage `json:"pagesWithMatchingImages"`
	VisuallySimilarImages   []Image                  `json:"visuallySimilarImages"`
	BestGuessLabels         []BestGuessLabel         `json:"bestGuessLabels"`
}

type BestGuessLabel struct {
	Label        string `json:"label"`
	LanguageCode string `json:"languageCode"`
}

type Image struct {
	URL string `json:"url"`
}

type PagesWithMatchingImage struct {
	URL                string  `json:"url"`
	PageTitle          string  `json:"pageTitle"`
	FullMatchingImages []Image `json:"fullMatchingImages"`
}

type WebEntity struct {
	EntityID    string  `json:"entityId"`
	Score       float64 `json:"score"`
	Description string  `json:"description"`
}

// Maps
type MapsResponse struct {
	ErrorMessage      *string            `json:"error_message,omitempty"`
	Routes            []DirectionsRoute  `json:"routes,omitempty"`
	Status            string             `json:"status"`
	HTMLAttributions  []interface{}      `json:"html_attributions,omitempty"`
	Results           []NearbyResult     `json:"results,omitempty"`
	GeocodedWaypoints []GeocodedWaypoint `json:"geocoded_waypoints,omitempty"`
	NextPageToken     *string            `json:"next_page_token,omitempty"`
}

type GeocodedWaypoint struct {
	GeocoderStatus string   `json:"geocoder_status"`
	PlaceID        string   `json:"place_id"`
	Types          []string `json:"types"`
}

type NearbyResult struct {
	Geometry         Geometry      `json:"geometry"`
	Icon             string        `json:"icon"`
	Name             string        `json:"name"`
	Photos           []Photo       `json:"photos"`
	PlaceID          string        `json:"place_id"`
	Reference        string        `json:"reference"`
	Scope            string        `json:"scope"`
	Types            []string      `json:"types"`
	Vicinity         string        `json:"vicinity"`
	BusinessStatus   *string       `json:"business_status,omitempty"`
	OpeningHours     *OpeningHours `json:"opening_hours,omitempty"`
	PlusCode         *PlusCode     `json:"plus_code,omitempty"`
	Rating           *float64      `json:"rating,omitempty"`
	UserRatingsTotal *int64        `json:"user_ratings_total,omitempty"`
	PriceLevel       *int64        `json:"price_level,omitempty"`
}

type Geometry struct {
	Location MapsLocation `json:"location"`
	Viewport Bounds       `json:"viewport"`
}

type MapsLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Bounds struct {
	Northeast MapsLocation `json:"northeast"`
	Southwest MapsLocation `json:"southwest"`
}

type OpeningHours struct {
	OpenNow bool `json:"open_now"`
}

type Photo struct {
	Height           int64    `json:"height"`
	HTMLAttributions []string `json:"html_attributions"`
	PhotoReference   string   `json:"photo_reference"`
	Width            int64    `json:"width"`
}

type PlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}

type DirectionsRoute struct {
	Bounds           Bounds        `json:"bounds"`
	Copyrights       string        `json:"copyrights"`
	Legs             []Leg         `json:"legs"`
	OverviewPolyline Polyline      `json:"overview_polyline"`
	Summary          string        `json:"summary"`
	Warnings         []interface{} `json:"warnings"`
	WaypointOrder    []interface{} `json:"waypoint_order"`
}

type Leg struct {
	Distance          Distance      `json:"distance"`
	Duration          Distance      `json:"duration"`
	EndAddress        string        `json:"end_address"`
	EndLocation       MapsLocation  `json:"end_location"`
	StartAddress      string        `json:"start_address"`
	StartLocation     MapsLocation  `json:"start_location"`
	Steps             []Step        `json:"steps"`
	TrafficSpeedEntry []interface{} `json:"traffic_speed_entry"`
	ViaWaypoint       []interface{} `json:"via_waypoint"`
}

type Distance struct {
	Text  string `json:"text"`
	Value int64  `json:"value"`
}

type Step struct {
	Distance         Distance     `json:"distance"`
	Duration         Distance     `json:"duration"`
	EndLocation      MapsLocation `json:"end_location"`
	HTMLInstructions string       `json:"html_instructions"`
	Polyline         Polyline     `json:"polyline"`
	StartLocation    MapsLocation `json:"start_location"`
	TravelMode       string       `json:"travel_mode"`
	Maneuver         *string      `json:"maneuver,omitempty"`
}

type Polyline struct {
	Points string `json:"points"`
}

package ridder

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// type
//
type Ridder struct {
	apiURL                string
	apiKey                string
	maxRetries            uint
	secondsBetweenRetries uint32
}

type RidderConfig struct {
	APIURL                string
	APIKey                string
	MaxRetries            *uint
	SecondsBetweenRetries *uint32
}

func NewRidder(config RidderConfig) (*Ridder, *errortools.Error) {
	ridder := new(Ridder)

	if config.APIURL == "" {
		return nil, errortools.ErrorMessage("Ridder API URL not provided")
	}
	ridder.apiURL = strings.TrimRight(config.APIURL, "/")

	if config.APIKey == "" {
		return nil, errortools.ErrorMessage("Ridder API Key not provided")
	}
	ridder.apiKey = config.APIKey

	if config.MaxRetries != nil {
		ridder.maxRetries = *config.MaxRetries
	} else {
		ridder.maxRetries = 0
	}

	if config.SecondsBetweenRetries != nil {
		ridder.secondsBetweenRetries = *config.SecondsBetweenRetries
	} else {
		ridder.secondsBetweenRetries = 3
	}

	return ridder, nil
}

// generic Get method
//
func (r *Ridder) Get(urlPath string, model interface{}) *errortools.Error {
	return r.httpRequest(http.MethodGet, urlPath, nil, model)
}

func (r *Ridder) httpRequest(httpMethod string, urlPath string, body io.Reader, model interface{}) *errortools.Error {
	client := new(http.Client)

	url := fmt.Sprintf("%s/%s", r.apiURL, urlPath)
	fmt.Println(url)

	e := new(errortools.Error)

	request, err := http.NewRequest(httpMethod, url, body)
	e.SetRequest(request)
	if err != nil {
		e.SetMessage(err)
		return e
	}

	// Add authorization token to header
	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-ApiKey", r.apiKey)

	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	// Send out the HTTP request
	response, e := utilities.DoWithRetry(client, request, r.maxRetries, r.secondsBetweenRetries)

	if response != nil {
		// Check HTTP StatusCode
		if response.StatusCode < 200 || response.StatusCode > 299 {
			fmt.Println(fmt.Sprintf("ERROR in %s", httpMethod))
			fmt.Println("url", url)
			fmt.Println("StatusCode", response.StatusCode)

			if e == nil {
				e = new(errortools.Error)
				e.SetRequest(request)
				e.SetResponse(response)
			}

			e.SetMessage(fmt.Sprintf("Server returned statuscode %v", response.StatusCode))
		}
	}

	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		e.SetMessage(err)
		return e
	}

	if e != nil {
		errorString := ""

		err2 := json.Unmarshal(b, &errorString)
		errortools.CaptureInfo(err2)

		if errorString != "" {
			e.SetMessage(errorString)
		}

		return e
	}

	if model != nil {
		err = json.Unmarshal(b, &model)
		if err != nil {
			e.SetMessage(err)
			return e
		}
	}

	return nil
}

package ridder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// type
//
type Ridder struct {
	APIURL string
	APIKey string
}

// Response represents highest level of exactonline api response
//
type Response struct {
	Data     *json.RawMessage `json:"data,omitempty"`
	NextPage *NextPage        `json:"next_page,omitempty"`
	Errors   *[]AsanaError    `json:"errors,omitempty"`
}

// NextPage contains info for batched data retrieval
//
type NextPage struct {
	Offset string `json:"offset"`
	Path   string `json:"path"`
	URI    string `json:"uri"`
}

// AsanaError contains error info
//
type AsanaError struct {
	Message string `json:"message"`
	Help    string `json:"help"`
}

func New(apiURL string, apiKey string) (*Ridder, *errortools.Error) {
	ridder := new(Ridder)

	if apiURL == "" {
		return nil, errortools.ErrorMessage("Ridder API URL not provided")
	}
	if apiKey == "" {
		return nil, errortools.ErrorMessage("Ridder API Key not provided")
	}

	ridder.APIURL = strings.TrimRight(apiURL, "/")
	ridder.APIKey = apiKey

	return ridder, nil
}

// generic Get method
//
func (ridder *Ridder) Get(url string, model interface{}) (*NextPage, *Response, *errortools.Error) {
	client := &http.Client{}

	e := new(errortools.Error)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	e.SetRequest(req)
	if err != nil {
		e.SetMessage(err)
		return nil, nil, e
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-ApiKey", ridder.APIKey)

	// Send out the HTTP request
	res, ee := utilities.DoWithRetry(client, req, 10, 3)
	e.SetResponse(res)
	if err != nil {
		e.SetMessage(ee)
		return nil, nil, e
	}

	if res == nil {
		return nil, nil, nil
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	response := Response{}

	err = json.Unmarshal(b, &response)
	if err != nil {
		e.SetMessage(err)
		return nil, nil, e
	}

	if response.Data != nil {
		err = json.Unmarshal(*response.Data, &model)
		if err != nil {
			e.SetMessage(err)
			return nil, nil, e
		}
	}

	//ridder.captureErrors(res.StatusCode, url, &response)

	return response.NextPage, &response, nil
}

func (a *Ridder) captureErrors(responseStatusCode int, url string, response *Response) {
	if response != nil {
		if response.Errors != nil {
			ee := []string{}
			for _, err := range *response.Errors {
				ee = append(ee, fmt.Sprintf("%s\n%s", err.Message, err.Help))
			}

			e := errortools.ErrorMessage(strings.Join(ee, "\n\n"))
			e.SetExtra("response_status_code", strconv.Itoa(responseStatusCode))
			e.SetExtra("url", url)
			errortools.CaptureInfo(e)
		}
	}
}

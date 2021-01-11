package ridder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	utilities "github.com/leapforce-libraries/go_utilities"
)

const (
	MaxLengthOrganizationName int    = 60
	MaxLengthOpportunityName  int    = 80
	DateTimeFormat            string = "2006-01-02T15:04:05"
)

// type
//
type Service struct {
	apiURL                string
	apiKey                string
	maxRetries            uint
	secondsBetweenRetries uint32
}

type ServiceConfig struct {
	APIURL                string
	APIKey                string
	MaxRetries            *uint
	SecondsBetweenRetries *uint32
}

func NewService(config ServiceConfig) (*Service, *errortools.Error) {
	ridder := new(Service)

	if config.APIURL == "" {
		return nil, errortools.ErrorMessage("Service API URL not provided")
	}
	ridder.apiURL = strings.TrimRight(config.APIURL, "/")

	if config.APIKey == "" {
		return nil, errortools.ErrorMessage("Service API Key not provided")
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
func (service *Service) Get(urlPath string, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, urlPath, nil, responseModel)
}

// generic Post method
//
func (service *Service) Post(urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPost, urlPath, bodyModel, responseModel)
}

func (service *Service) httpRequest(httpMethod string, urlPath string, bodyModel interface{}, responseModel interface{}) (*http.Request, *http.Response, *errortools.Error) {
	client := new(http.Client)

	url := fmt.Sprintf("%s/%s", service.apiURL, urlPath)
	fmt.Println(url)

	e := new(errortools.Error)

	buffer := new(bytes.Buffer)
	buffer = nil

	if bodyModel != nil {

		b, err := json.Marshal(bodyModel)
		if err != nil {
			e.SetMessage(err)
			return nil, nil, e
		}
		fmt.Println(string(b)) //temp
		buffer = bytes.NewBuffer(b)
	}

	request, err := func() (*http.Request, error) {
		// function necessary because a Buffer nil pointer differs from a nil value
		if buffer == nil {
			return http.NewRequest(httpMethod, url, nil)
		}
		return http.NewRequest(httpMethod, url, buffer)
	}()

	e.SetRequest(request)

	if err != nil {
		e.SetMessage(err)
		return request, nil, e
	}

	// Add authorization token to header
	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-ApiKey", service.apiKey)

	if bodyModel != nil {
		request.Header.Set("Content-Type", "application/json-patch+json")
	}

	// Send out the HTTP request
	response, e := utilities.DoWithRetry(client, request, service.maxRetries, service.secondsBetweenRetries)

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

		if response.Body != nil {

			defer response.Body.Close()

			b, err := ioutil.ReadAll(response.Body)
			if err != nil {
				e.SetMessage(err)
				return request, response, e
			}

			if e != nil {
				// try to unmarshal to ErrorResponse
				errorResponse := ErrorResponse{}
				errError := json.Unmarshal(b, &errorResponse)

				if errError == nil {
					if errorResponse.Error != "" {
						e.SetMessage(errorResponse.Error)
					}
				} else {
					// try to unmarshal to string
					errorString := ""
					errError = json.Unmarshal(b, &errorString)

					if errorString != "" {
						e.SetMessage(errorString)
					}
				}

				errortools.CaptureInfo(errError)

				return request, response, e
			}

			if responseModel != nil {
				err = json.Unmarshal(b, &responseModel)
				if err != nil {
					e.SetMessage(err)
					return request, response, e
				}
			}
		}
	}

	return request, response, e
}

func (service *Service) removeSpecialCharacters(test *string) *errortools.Error {
	if test == nil {
		return nil
	}

	re := regexp.MustCompile(`[\\/:*?"<>|]`)

	removedCount := len(*test) - len(string(re.ReplaceAll([]byte(*test), []byte(""))))

	if removedCount == 0 {
		return nil
	}

	message := fmt.Sprintf("%v special characters in '%s' replaced by a dot", removedCount, *test)
	(*test) = string(re.ReplaceAll([]byte(*test), []byte(".")))

	return errortools.ErrorMessage(message)
}

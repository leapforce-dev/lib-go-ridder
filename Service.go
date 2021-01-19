package ridder

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	utilities "github.com/leapforce-libraries/go_utilities"
)

const (
	MaxLengthOrganizationEmail         int    = 255
	MaxLengthOrganizationName          int    = 60
	MaxLengthOrganizationPhone         int    = 50
	MaxLengthOrganizationWebsite       int    = 255
	MaxLengthAddressHouseNumber        int    = 50
	MaxLengthAddressCity               int    = 50
	MaxLengthAddressZipCode            int    = 50
	MaxLengthAddressStreet             int    = 50
	MaxLengthContactPhone              int    = 50
	MaxLengthContactEmail              int    = 255
	MaxLengthContactFunctionName       int    = 150
	MaxLengthContactCellphone          int    = 50
	MaxLengthContactLastName           int    = 127
	MaxLengthContactInitials           int    = 50
	MaxLengthContactFirstName          int    = 127
	MaxLengthOpportunityName           int    = 80
	MaxLengthOpportunityInsightlyState int    = 4000
	DateTimeFormat                     string = "2006-01-02T15:04:05"
)

// type
//
type Service struct {
	apiURL      string
	apiKey      string
	httpService *go_http.Service
}

type ServiceConfig struct {
	APIURL                string
	APIKey                string
	MaxRetries            *uint
	SecondsBetweenRetries *uint32
}

func NewService(config ServiceConfig) (*Service, *errortools.Error) {
	if config.APIURL == "" {
		return nil, errortools.ErrorMessage("Service API URL not provided")
	}

	if config.APIKey == "" {
		return nil, errortools.ErrorMessage("Service API Key not provided")
	}

	httpServiceConfig := go_http.ServiceConfig{
		MaxRetries:            config.MaxRetries,
		SecondsBetweenRetries: config.SecondsBetweenRetries,
	}

	return &Service{
		apiURL:      strings.TrimRight(config.APIURL, "/"),
		apiKey:      config.APIKey,
		httpService: go_http.NewService(httpServiceConfig),
	}, nil
}

func (service *Service) httpRequest(httpMethod string, requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add api key header
	header := http.Header{}
	header.Set("X-ApiKey", service.apiKey)

	if !utilities.IsNil(requestConfig.BodyModel) {
		header.Set("Content-Type", "application/json-patch+json")
	}
	(*requestConfig).NonDefaultHeaders = &header

	// add error model
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.httpService.HTTPRequest(httpMethod, requestConfig)
	if errorResponse.Error != "" {
		e.SetMessage(errorResponse.Error)
	}

	return request, response, e
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", service.apiURL, path)
}

func (service *Service) get(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, requestConfig)
}

func (service *Service) post(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPost, requestConfig)
}

func (service *Service) truncateString(fieldName string, value *string, maxLength int, errors *[]string) {
	if len(*value) > maxLength {
		*value = (*value)[:maxLength]

		*errors = append(*errors, fmt.Sprintf("%s truncated to %v characters.", fieldName, maxLength))
	}
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

package ridder

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	apiName string = "Ridder"
)

type Service struct {
	apiURL      string
	apiKey      string
	httpService *go_http.Service
}

type ServiceConfig struct {
	APIURL string
	APIKey string
}

func NewService(config *ServiceConfig) (*Service, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if config.APIURL == "" {
		return nil, errortools.ErrorMessage("Service API URL not provided")
	}

	if config.APIKey == "" {
		return nil, errortools.ErrorMessage("Service API Key not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		apiURL:      config.APIURL,
		apiKey:      config.APIKey,
		httpService: httpService,
	}, nil
}

func (service *Service) httpRequest(httpMethod string, requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add api key header
	header := http.Header{}
	header.Set("X-ApiKey", service.apiKey)
	(*requestConfig).NonDefaultHeaders = &header

	// add error model
	problemDetails := ProblemDetails{}
	(*requestConfig).ErrorModel = &problemDetails

	request, response, e := service.httpService.HTTPRequest(httpMethod, requestConfig)
	if problemDetails.Title != "" {
		e.SetMessage(problemDetails.Title)
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

func (service *Service) put(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPut, requestConfig)
}

func (service *Service) delete(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodDelete, requestConfig)
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

func (service *Service) parseInt32String(int32String string) (*int32, *errortools.Error) {
	_int64, err := strconv.ParseInt(int32String, 10, 32)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	_int32 := int32(_int64)

	return &_int32, nil
}

func (service *Service) APIName() string {
	return apiName
}

func (service *Service) APIKey() string {
	return service.apiKey
}

func (service *Service) APICallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) APIReset() {
	service.httpService.ResetRequestCount()
}

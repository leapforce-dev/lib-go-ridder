package ridder

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	utilities "github.com/leapforce-libraries/go_utilities"
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

	// truncate strings
	if !utilities.IsNil(requestConfig.BodyModel) {
		service.truncateStrings(requestConfig.BodyModel)
	}

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

func (service *Service) truncateStrings(model interface{}) *errortools.Error {
	if reflect.TypeOf(model).Kind() != reflect.Ptr {
		return errortools.ErrorMessage("The interface is not a pointer.")
	}

	v := reflect.ValueOf(model).Elem()
	if v.Kind() != reflect.Struct {
		return errortools.ErrorMessage("The interface is not a pointer to a struct.")
	}

	errors := []string{}

	structType := reflect.TypeOf(model).Elem()
	value := reflect.ValueOf(model).Elem()

	e := service.truncateStringsValue(structType, value, &errors)
	if e != nil {
		return e
	}

	if len(errors) > 0 {
		e := new(errortools.Error)
		e.SetMessage("One or more fields truncated")
		e.SetExtra("errors", strings.Join(errors, "\n"))
		errortools.CaptureWarning(e)
	}

	return nil
}

func (service *Service) truncateStringsValue(structType reflect.Type, value reflect.Value, errors *[]string) *errortools.Error {
	for i := 0; i < structType.NumField(); i++ {
		fieldName := structType.Field(i).Name
		maxLength := structType.Field(i).Tag.Get("max")

		field := value.FieldByName(fieldName)

		if field.Kind() == reflect.Struct {

			e := service.truncateStringsValue(field.Type(), field, errors)
			if e != nil {
				errortools.CaptureError(e)
			}
			continue
		}

		if field.Kind() == reflect.Ptr {
			if !field.IsNil() {
				_field := field.Elem()

				if _field.Kind() == reflect.Struct {
					e := service.truncateStringsValue(_field.Type(), _field, errors)
					if e != nil {
						errortools.CaptureError(e)
					}
					continue
				}
			}
		}

		if maxLength == "" {
			continue
		}

		maxLenghtInt64, err := strconv.ParseInt(maxLength, 10, 64)
		if err != nil {
			errortools.CaptureError(err)
			continue
		}

		switch field.Type().String() {
		case "string":
			_value := field.String()
			if len(_value) > int(maxLenghtInt64) {
				field.SetString(_value[:maxLenghtInt64])

				*errors = append(*errors, fmt.Sprintf("%s truncated to %v characters.", fieldName, maxLength))
			}

			break
		case "*string":
			if !field.IsNil() {
				_value := field.Elem().String()
				if len(_value) > int(maxLenghtInt64) {
					field.Elem().SetString(_value[:maxLenghtInt64])

					*errors = append(*errors, fmt.Sprintf("%s truncated to %v characters.", fieldName, maxLength))
				}
			}

			break
		default:
			errortools.CaptureError("Field is not a (pointer to a) string")
			break
		}
	}

	return nil
}

/*
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
}*/

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

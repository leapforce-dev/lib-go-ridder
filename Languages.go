package ridder

import (
	"fmt"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type LanguageFormat string

const (
	LanguageFormatISO639_1 LanguageFormat = "ISO639_1"
	LanguageFormatISO639_3 LanguageFormat = "ISO639_3"
)

func (service *Service) GetLanguage(languageFormat LanguageFormat, languageCode string) (*int32, *errortools.Error) {
	var params url.Values
	params.Set("languageIsoFormat", string(languageFormat))
	params.Set("languageCode", languageCode)

	var languageIDString string

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("languages?%s", params.Encode())),
		ResponseModel: &languageIDString,
	}
	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return service.parseInt32String(languageIDString)
}

package ridder

import (
	"fmt"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type CountryFormat string

const (
	CountryFormatISO3166_1_Alpha_2  CountryFormat = "ISO3166_1_ALPHA_2"
	CountryFormatISO3166_1_Alpha_3  CountryFormat = "ISO3166_1_ALPHA_3"
	CountryFormatISO3166_1_Numeric3 CountryFormat = "ISO3166_1_NUMERIC_3"
)

func (service *Service) GetCountry(countryFormat CountryFormat, countryCode string) (*int32, *errortools.Error) {
	params := url.Values{}
	params.Set("countryIsoFormat", string(countryFormat))
	params.Set("countryCode", countryCode)

	var countryIDString string

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("countries?%s", params.Encode())),
		ResponseModel: &countryIDString,
	}
	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return service.parseInt32String(countryIDString)
}

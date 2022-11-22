package ridder

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type ItemGroup struct {
	Id          int32   `json:"Id"`
	Code        string  `json:"Code"`
	Description *string `json:"Description"`
}

func (service *Service) GetItemGroups() (*[]ItemGroup, *errortools.Error) {
	var itemGroups []ItemGroup

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("itemGroups"),
		ResponseModel: &itemGroups,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &itemGroups, nil
}

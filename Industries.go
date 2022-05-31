package ridder

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Industry struct {
	ID          int32   `json:"Id"`
	Code        string  `json:"Code" max:"3"`
	Description *string `json:"Description" max:"80"`
}

func (service *Service) GetIndustries() (*[]Industry, *errortools.Error) {
	var industries []Industry

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("industries"),
		ResponseModel: &industries,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &industries, nil
}

func (service *Service) GetIndustryByID(industryID int32) (*Industry, *errortools.Error) {
	return service.getIndustry(fmt.Sprintf("industries/id/%v", industryID))
}

func (service *Service) GetIndustryByCode(code string) (*Industry, *errortools.Error) {
	return service.getIndustry(fmt.Sprintf("industries/code/%s", code))
}

func (service *Service) getIndustry(urlPath string) (*Industry, *errortools.Error) {
	var industry Industry

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(urlPath),
		ResponseModel: &industry,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &industry, nil
}

func (service *Service) CreateIndustry(industry *Industry) (*int32, *errortools.Error) {
	if industry == nil {
		return nil, nil
	}

	var industryIDString string

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("industries"),
		BodyModel:     industry,
		ResponseModel: &industryIDString,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return service.parseInt32String(industryIDString)
}

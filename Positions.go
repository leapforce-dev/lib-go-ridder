package ridder

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Position struct {
	ID          int32   `json:"Id"`
	Code        string  `json:"Code" max:"3"`
	Description *string `json:"Description" max:"80"`
}

func (service *Service) GetPositions() (*[]Position, *errortools.Error) {
	var positions []Position

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("positions"),
		ResponseModel: &positions,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &positions, nil
}

func (service *Service) GetPositionByID(positionID int32) (*Position, *errortools.Error) {
	return service.getPosition(fmt.Sprintf("positions/id/%v", positionID))
}

func (service *Service) GetPositionByCode(code string) (*Position, *errortools.Error) {
	return service.getPosition(fmt.Sprintf("positions/code/%s", code))
}

func (service *Service) getPosition(urlPath string) (*Position, *errortools.Error) {
	var position Position

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(urlPath),
		ResponseModel: &position,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &position, nil
}

func (service *Service) CreatePosition(position *Position) (*int32, *errortools.Error) {
	if position == nil {
		return nil, nil
	}

	var positionIDString string

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("positions"),
		BodyModel:     position,
		ResponseModel: &positionIDString,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return service.parseInt32String(positionIDString)
}

package ridder

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type State struct {
	ID          int32   `json:"Id"`
	Code        string  `json:"Code" max:"3"`
	Description *string `json:"Description" max:"80"`
	CountryID   int32   `json:"CountryId"`
}

func (service *Service) GetStates() (*[]State, *errortools.Error) {
	var states []State

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("states"),
		ResponseModel: &states,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &states, nil
}

func (service *Service) GetStateByID(stateID int32) (*State, *errortools.Error) {
	return service.getState(fmt.Sprintf("states/id/%v", stateID))
}

func (service *Service) GetStateByCode(code string) (*State, *errortools.Error) {
	return service.getState(fmt.Sprintf("states/code/%s", code))
}

func (service *Service) getState(urlPath string) (*State, *errortools.Error) {
	var state State

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(urlPath),
		ResponseModel: &state,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &state, nil
}

func (service *Service) CreateState(state *State) (*int32, *errortools.Error) {
	if state == nil {
		return nil, nil
	}

	var stateIDString string

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("states"),
		BodyModel:     state,
		ResponseModel: &stateIDString,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return service.parseInt32String(stateIDString)
}

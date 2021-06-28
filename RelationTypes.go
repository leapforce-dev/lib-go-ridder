package ridder

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type RelationType struct {
	ID          int32   `json:"Id"`
	Code        string  `json:"Code" max:"3"`
	Description *string `json:"Description" max:"80"`
}

func (service *Service) GetRelationTypes() (*[]RelationType, *errortools.Error) {
	var relationTypes []RelationType

	requestConfig := go_http.RequestConfig{
		URL:           service.url("relationTypes"),
		ResponseModel: &relationTypes,
	}
	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &relationTypes, nil
}

func (service *Service) GetRelationTypeByID(relationTypeID int32) (*RelationType, *errortools.Error) {
	return service.getRelationType(fmt.Sprintf("relationTypes/id/%v", relationTypeID))
}

func (service *Service) GetRelationTypeByCode(code string) (*RelationType, *errortools.Error) {
	return service.getRelationType(fmt.Sprintf("relationTypes/code/%s", code))
}

func (service *Service) getRelationType(urlPath string) (*RelationType, *errortools.Error) {
	var relationType RelationType

	requestConfig := go_http.RequestConfig{
		URL:           service.url(urlPath),
		ResponseModel: &relationType,
	}
	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &relationType, nil
}

func (service *Service) CreateRelationType(relationType *RelationType) (*int32, *errortools.Error) {
	if relationType == nil {
		return nil, nil
	}

	var relationTypeIDString string

	requestConfig := go_http.RequestConfig{
		URL:           service.url("relationTypes"),
		BodyModel:     relationType,
		ResponseModel: &relationTypeIDString,
	}
	_, _, e := service.post(&requestConfig)
	if e != nil {
		return nil, e
	}

	return service.parseInt32String(relationTypeIDString)
}

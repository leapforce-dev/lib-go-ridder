package ridder

import (
	"net/http"
	"strconv"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Relation struct {
	ID                int32          `json:"Id"`
	ExternalCRMID     string         `json:"ExternalCrmId" max:"50"`
	Name              string         `json:"Name" max:"60"`
	Code              *string        `json:"Code,omitempty" max:"10"`
	CurrencyISOCode   string         `json:"CurrencyIsoCode" max:"3"`
	LanguageISOCode   string         `json:"LanguageIsoCode"`
	LanguageISOFormat LanguageFormat `json:"LanguageIsoFormat"`
	SalesPersonID     *string        `json:"SalesPersonId,omitempty" max:"50"`
	RelationTypeCode  *string        `json:"RelationTypeCode,omitempty"`
	IndustryCode      *string        `json:"IndustryCode,omitempty"`
	Phone1            *string        `json:"Phone1,omitempty" max:"50"`
	Phone2            *string        `json:"Phone2,omitempty" max:"50"`
	Fax               *string        `json:"Fax,omitempty" max:"50"`
	Email             *string        `json:"Email,omitempty" max:"255"`
	Website           *string        `json:"Website,omitempty" max:"255"`
	Memo              *string        `json:"Memo,omitempty"`
	PostalAddress     *Address       `json:"PostalAddress,omitempty"`
	VisitingAddress   *Address       `json:"VisitingAddress,omitempty"`
}

func (service *Service) UpdateRelation(relation *Relation) (*http.Response, *errortools.Error) {
	if relation == nil {
		return nil, nil
	}

	requestConfig := go_http.RequestConfig{
		URL:       service.url("relations"),
		BodyModel: relation,
	}
	_, response, e := service.put(&requestConfig)
	if e != nil {
		return response, e
	}

	return response, nil
}

func (service *Service) CreateRelation(relation *Relation) (*int32, *http.Response, *errortools.Error) {
	if relation == nil {
		return nil, nil, nil
	}

	var relationIDString string

	requestConfig := go_http.RequestConfig{
		URL:           service.url("relations"),
		BodyModel:     relation,
		ResponseModel: &relationIDString,
	}
	_, response, e := service.post(&requestConfig)
	if e != nil {
		return nil, response, e
	}

	relationIDInt64, err := strconv.ParseInt(relationIDString, 10, 64)
	if err != nil {
		return nil, response, errortools.ErrorMessage(err)
	}
	relationIDInt32 := int32(relationIDInt64)

	return &relationIDInt32, response, e
}

func (service *Service) DeleteRelation(id int32) *errortools.Error {
	requestConfig := go_http.RequestConfig{
		URL:       service.url("relations"),
		BodyModel: id,
	}
	_, _, e := service.delete(&requestConfig)

	return e
}

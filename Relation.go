package ridder

import (
	"strconv"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Relation struct {
	ID                int32          `json:"Id"`
	ExternalCRMID     string         `json:"ExternalCrmId"`
	Name              string         `json:"Name"`
	Code              *string        `json:"Code,omitempty"`
	CurrencyISOCode   string         `json:"CurrencyIsoCode"`
	LanguageISOCode   string         `json:"LanguageIsoCode"`
	LanguageISOFormat LanguageFormat `json:"LanguageIsoFormat"`
	SalesPersonID     *string        `json:"SalesPersonId,omitempty"`
	RelationTypeCode  *string        `json:"RelationTypeCode,omitempty"`
	IndustryCode      *string        `json:"IndustryCode,omitempty"`
	Phone1            *string        `json:"Phone1,omitempty"`
	Phone2            *string        `json:"Phone2,omitempty"`
	Fax               *string        `json:"Fax,omitempty"`
	Email             *string        `json:"Email,omitempty"`
	Website           *string        `json:"Website,omitempty"`
	Memo              *string        `json:"Memo,omitempty"`
	PostalAddress     *Address       `json:"PostalAddress,omitempty"`
	VisitingAddress   *Address       `json:"VisitingAddress,omitempty"`
}

func (service *Service) UpdateRelation(relation *Relation) *errortools.Error {
	if relation == nil {
		return nil
	}

	requestConfig := go_http.RequestConfig{
		URL:       service.url("relations"),
		BodyModel: relation,
	}
	_, _, e := service.put(&requestConfig)
	if e != nil {
		return e
	}

	return nil
}

func (service *Service) CreateRelation(relation *Relation) (*int32, *errortools.Error) {
	if relation == nil {
		return nil, nil
	}

	var relationIDString string

	requestConfig := go_http.RequestConfig{
		URL:           service.url("relations"),
		BodyModel:     relation,
		ResponseModel: &relationIDString,
	}
	_, _, e := service.post(&requestConfig)
	if e != nil {
		return nil, e
	}

	relationIDInt64, err := strconv.ParseInt(relationIDString, 10, 64)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	relationIDInt32 := int32(relationIDInt64)

	return &relationIDInt32, e
}

func (service *Service) DeleteRelation(id int32) *errortools.Error {
	requestConfig := go_http.RequestConfig{
		URL:       service.url("relations"),
		BodyModel: id,
	}
	_, _, e := service.delete(&requestConfig)

	return e
}

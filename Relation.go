package ridder

import (
	"fmt"
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

	e := service.validateRelation(relation)
	if e != nil {
		errortools.CaptureWarning(e)
	}

	requestConfig := go_http.RequestConfig{
		URL:       service.url("relations"),
		BodyModel: relation,
	}
	_, _, e = service.put(&requestConfig)
	fmt.Println(e)
	if e != nil {
		return e
	}

	return nil
}

func (service *Service) CreateRelation(relation *Relation) (*int32, *errortools.Error) {
	if relation == nil {
		return nil, nil
	}

	e := service.validateRelation(relation)
	if e != nil {
		errortools.CaptureWarning(e)
	}

	var relationIDString string

	requestConfig := go_http.RequestConfig{
		URL:           service.url("relations"),
		BodyModel:     relation,
		ResponseModel: &relationIDString,
	}
	_, _, e = service.post(&requestConfig)

	relationIDInt64, err := strconv.ParseInt(relationIDString, 10, 64)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	relationIDInt32 := int32(relationIDInt64)

	return &relationIDInt32, e
}

func (service *Service) validateRelation(relation *Relation) *errortools.Error {
	if relation == nil {
		return nil
	}
	/*
		errors := []string{}

		service.truncateString("EmailAddress", &(*relation).EmailAddress, MaxLengthRelationEmail, &errors)
		service.truncateString("RelationName", &(*relation).RelationName, MaxLengthRelationName, &errors)
		service.truncateString("Phone", &(*relation).Phone, MaxLengthRelationPhone, &errors)
		service.truncateString("Website", &(*relation).Website, MaxLengthRelationWebsite, &errors)

		service.truncateString("BillingAddress-HouseNumber", &(*relation).BillingAddress.HouseNumber, MaxLengthAddressHouseNumber, &errors)
		service.truncateString("BillingAddress-City", &(*relation).BillingAddress.City, MaxLengthAddressCity, &errors)
		service.truncateString("BillingAddress-ZipCode", &(*relation).BillingAddress.ZipCode, MaxLengthAddressZipCode, &errors)
		service.truncateString("BillingAddress-Street", &(*relation).BillingAddress.Street, MaxLengthAddressStreet, &errors)

		service.truncateString("ShippingAddress-HouseNumber", &(*relation).ShippingAddress.HouseNumber, MaxLengthAddressHouseNumber, &errors)
		service.truncateString("ShippingAddress-City", &(*relation).ShippingAddress.City, MaxLengthAddressCity, &errors)
		service.truncateString("ShippingAddress-ZipCode", &(*relation).ShippingAddress.ZipCode, MaxLengthAddressZipCode, &errors)
		service.truncateString("ShippingAddress-Street", &(*relation).ShippingAddress.Street, MaxLengthAddressStreet, &errors)

		e := service.removeSpecialCharacters(&(*relation).RelationName)
		if e != nil {
			errors = append(errors, e.Message())
		}

		if len(errors) > 0 {
			return errortools.ErrorMessage(strings.Join(errors, "\n"))
		}*/

	return nil
}

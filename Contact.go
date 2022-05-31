package ridder

import (
	"net/http"
	"strconv"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Contact struct {
	ID                  int32   `json:"Id"`
	RelationID          int32   `json:"RelationId"`
	ExternalCRMID       string  `json:"ExternalCrmId" max:"50"`
	Person              Person  `json:"Person"`
	Email               *string `json:"Email,omitempty" max:"255"`
	Fax                 *string `json:"Fax,omitempty" max:"50"`
	BusinessPhone1      *string `json:"BusinessPhone1,omitempty" max:"50"`
	BusinessPhone2      *string `json:"BusinessPhone2,omitempty" max:"50"`
	BusinessMobilePhone *string `json:"BusinessMobilePhone,omitempty" max:"50"`
	PrivatePhone1       *string `json:"PrivatePhone1,omitempty" max:"50"`
	PrivatePhone2       *string `json:"PrivatePhone2,omitempty" max:"50"`
	PrivateMobilePhone  *string `json:"PrivateMobilePhone,omitempty" max:"50"`
	PositionID          *int32  `json:"PositionId,omitempty"`
	Memo                *string `json:"Memo,omitempty"`
}

func (service *Service) UpdateContact(contact *Contact) (*http.Response, *errortools.Error) {
	if contact == nil {
		return nil, nil
	}

	requestConfig := go_http.RequestConfig{
		Method:    http.MethodPut,
		Url:       service.url("contacts"),
		BodyModel: contact,
	}
	_, response, e := service.httpRequest(&requestConfig)
	if e != nil {
		return response, e
	}

	return response, nil
}

func (service *Service) CreateContact(contact *Contact) (*int32, *http.Response, *errortools.Error) {
	if contact == nil {
		return nil, nil, nil
	}

	var contactIDString string

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("contacts"),
		BodyModel:     contact,
		ResponseModel: &contactIDString,
	}
	_, response, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, response, e
	}

	contactIDInt64, err := strconv.ParseInt(contactIDString, 10, 64)
	if err != nil {
		return nil, response, errortools.ErrorMessage(err)
	}
	contactIDInt32 := int32(contactIDInt64)

	return &contactIDInt32, response, e
}

func (service *Service) DeleteContact(id int32) *errortools.Error {
	requestConfig := go_http.RequestConfig{
		Method:    http.MethodDelete,
		Url:       service.url("contacts"),
		BodyModel: id,
	}
	_, _, e := service.httpRequest(&requestConfig)

	return e
}

package ridder

import (
	"strconv"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Contact struct {
	ID                  int32   `json:"Id"`
	RelationID          int32   `json:"RelationId"`
	ExternalCRMID       string  `json:"ExternalCrmId"`
	Person              Person  `json:"Person"`
	Email               *string `json:"Email,omitempty"`
	Fax                 *string `json:"Fax,omitempty"`
	BusinessPhone1      *string `json:"BusinessPhone1,omitempty"`
	BusinessPhone2      *string `json:"BusinessPhone2,omitempty"`
	BusinessMobilePhone *string `json:"BusinessMobilePhone,omitempty"`
	PrivatePhone1       *string `json:"PrivatePhone1,omitempty"`
	PrivatePhone2       *string `json:"PrivatePhone2,omitempty"`
	PrivateMobilePhone  *string `json:"PrivateMobilePhone,omitempty"`
	PositionID          *int32  `json:"PositionId,omitempty"`
	Memo                *string `json:"Memo,omitempty"`
}

func (service *Service) UpdateContact(contact *Contact) *errortools.Error {
	if contact == nil {
		return nil
	}

	requestConfig := go_http.RequestConfig{
		URL:       service.url("contacts"),
		BodyModel: contact,
	}
	_, _, e := service.put(&requestConfig)
	if e != nil {
		return e
	}

	return nil
}

func (service *Service) CreateContact(contact *Contact) (*int32, *errortools.Error) {
	if contact == nil {
		return nil, nil
	}

	var contactIDString string

	requestConfig := go_http.RequestConfig{
		URL:           service.url("contacts"),
		BodyModel:     contact,
		ResponseModel: &contactIDString,
	}
	_, _, e := service.post(&requestConfig)
	if e != nil {
		return nil, e
	}

	contactIDInt64, err := strconv.ParseInt(contactIDString, 10, 64)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	contactIDInt32 := int32(contactIDInt64)

	return &contactIDInt32, e
}

func (service *Service) DeleteContact(id int32) *errortools.Error {
	requestConfig := go_http.RequestConfig{
		URL:       service.url("contacts"),
		BodyModel: id,
	}
	_, _, e := service.delete(&requestConfig)

	return e
}

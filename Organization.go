package ridder

import (
	"fmt"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Organization struct {
	RidderID         int32   `json:"RidderId"`
	InsightlyID      int32   `json:"InsightlyId"`
	OrganizationName string  `json:"OrganizationName"`
	Phone            string  `json:"Phone"`
	Website          string  `json:"Website"`
	EmailAddress     string  `json:"EmailAddress"`
	AccountManager   int32   `json:"SalesPerson"`
	BillingAddress   Address `json:"BillingAddress"`
	ShippingAddress  Address `json:"ShippingAddress"`
	Expired          bool    `json:"Expired"`
}

func (service *Service) GetOrganization(ridderID int32) (*Organization, *errortools.Error) {
	organization := Organization{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("organizations?ridderid=%v", ridderID)),
		ResponseModel: &organization,
	}
	_, _, e := service.get(&requestConfig)

	return &organization, e
}

func (service *Service) UpdateOrganization(organization *Organization) (*int32, *errortools.Error) {
	if organization == nil {
		return nil, nil
	}

	ev := service.validateOrganization(organization)

	organizationID := new(int32)

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("organizations/%v", organization.RidderID)),
		BodyModel:     organization,
		ResponseModel: organizationID,
	}
	req, res, e := service.post(&requestConfig)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return organizationID, e
}

func (service *Service) CreateOrganization(newOrganization *Organization) (*int32, *errortools.Error) {
	if newOrganization == nil {
		return nil, nil
	}

	ev := service.validateOrganization(newOrganization)

	organizationID := new(int32)

	requestConfig := go_http.RequestConfig{
		URL:           service.url("organizations"),
		BodyModel:     newOrganization,
		ResponseModel: organizationID,
	}
	req, res, e := service.post(&requestConfig)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return organizationID, e
}

func (service *Service) validateOrganization(organization *Organization) *errortools.Error {
	if organization == nil {
		return nil
	}

	errors := []string{}

	service.truncateString("EmailAddress", &(*organization).EmailAddress, MaxLengthOrganizationEmail, &errors)
	service.truncateString("OrganizationName", &(*organization).OrganizationName, MaxLengthOrganizationName, &errors)
	service.truncateString("Phone", &(*organization).Phone, MaxLengthOrganizationPhone, &errors)
	service.truncateString("Website", &(*organization).Website, MaxLengthOrganizationWebsite, &errors)

	service.truncateString("BillingAddress-HouseNumber", &(*organization).BillingAddress.HouseNumber, MaxLengthAddressHouseNumber, &errors)
	service.truncateString("BillingAddress-City", &(*organization).BillingAddress.City, MaxLengthAddressCity, &errors)
	service.truncateString("BillingAddress-ZipCode", &(*organization).BillingAddress.ZipCode, MaxLengthAddressZipCode, &errors)
	service.truncateString("BillingAddress-Street", &(*organization).BillingAddress.Street, MaxLengthAddressStreet, &errors)

	service.truncateString("ShippingAddress-HouseNumber", &(*organization).ShippingAddress.HouseNumber, MaxLengthAddressHouseNumber, &errors)
	service.truncateString("ShippingAddress-City", &(*organization).ShippingAddress.City, MaxLengthAddressCity, &errors)
	service.truncateString("ShippingAddress-ZipCode", &(*organization).ShippingAddress.ZipCode, MaxLengthAddressZipCode, &errors)
	service.truncateString("ShippingAddress-Street", &(*organization).ShippingAddress.Street, MaxLengthAddressStreet, &errors)

	e := service.removeSpecialCharacters(&(*organization).OrganizationName)
	if e != nil {
		errors = append(errors, e.Message())
	}

	if len(errors) > 0 {
		return errortools.ErrorMessage(strings.Join(errors, "\n"))
	}

	return nil
}

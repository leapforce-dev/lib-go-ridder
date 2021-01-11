package ridder

import (
	"fmt"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
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
	url := fmt.Sprintf("organizations?ridderid=%v", ridderID)

	organization := Organization{}
	_, _, e := service.Get(url, &organization)

	return &organization, e
}

func (service *Service) UpdateOrganization(organization *Organization) (*int32, *errortools.Error) {
	url := fmt.Sprintf("organizations/%v", organization.RidderID)

	if organization == nil {
		return nil, nil
	}

	ev := service.validateOrganization(organization)

	organizationID := new(int32)
	req, res, e := service.Post(url, &organization, organizationID)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return organizationID, e
}

func (service *Service) CreateOrganization(newOrganization *Organization) (*int32, *errortools.Error) {
	url := "organizations"

	if newOrganization == nil {
		return nil, nil
	}

	ev := service.validateOrganization(newOrganization)

	organizationID := new(int32)
	req, res, e := service.Post(url, &newOrganization, organizationID)

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

	if len(organization.OrganizationName) > MaxLengthOrganizationName {
		(*organization).OrganizationName = organization.OrganizationName[:MaxLengthOrganizationName]

		errors = append(errors, fmt.Sprintf("OrganizationName truncated to %v characters.", MaxLengthOrganizationName))
	}

	e := service.removeSpecialCharacters(&(*organization).OrganizationName)
	if e != nil {
		errors = append(errors, e.Message())
	}

	if len(errors) > 0 {
		return errortools.ErrorMessage(strings.Join(errors, "\n"))
	}

	return nil
}

package ridder

import (
	"fmt"

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

func (r *Ridder) GetOrganization(ridderID int32) (*Organization, *errortools.Error) {
	url := fmt.Sprintf("organizations?ridderid=%v", ridderID)

	organization := Organization{}
	_, _, e := r.Get(url, &organization)

	return &organization, e
}

func (r *Ridder) UpdateOrganization(organization *Organization) (*int32, *errortools.Error) {
	url := fmt.Sprintf("organizations/%v", organization.RidderID)

	if organization == nil {
		return nil, nil
	}

	ev := organization.validate()

	organizationID := new(int32)
	req, res, e := r.Post(url, &organization, organizationID)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return organizationID, e
}

func (r *Ridder) CreateOrganization(newOrganization *Organization) (*int32, *errortools.Error) {
	url := fmt.Sprintf("organizations")

	if newOrganization == nil {
		return nil, nil
	}

	ev := newOrganization.validate()

	organizationID := new(int32)
	req, res, e := r.Post(url, &newOrganization, organizationID)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return organizationID, e
}

func (organization *Organization) validate() *errortools.Error {
	if organization == nil {
		return nil
	}

	if len(organization.OrganizationName) > MaxLengthOrganizationName {
		(*organization).OrganizationName = organization.OrganizationName[:MaxLengthOrganizationName]

		return errortools.ErrorMessage(fmt.Sprintf("OrganizationName truncated to %v characters.", MaxLengthOrganizationName))
	}

	return nil
}

package ridder

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type Organization struct {
	RidderID         int64   `json:"RidderId"`
	InsightlyID      int64   `json:"InsightlyId"`
	OrganizationName string  `json:"OrganizationName"`
	Phone            string  `json:"Phone"`
	Website          string  `json:"Website"`
	EmailAddress     string  `json:"EmailAddress"`
	AccountManager   *int64  `json:"SalesPerson"`
	BillingAddress   Address `json:"BillingAddress"`
	ShippingAddress  Address `json:"ShippingAddress"`
	Expired          bool    `json:"Expired"`
}

func (r *Ridder) GetOrganization(ridderID int64) (*Organization, *errortools.Error) {
	url := fmt.Sprintf("organizations?ridderid=%v", ridderID)

	organization := Organization{}
	_, _, e := r.Get(url, &organization)

	return &organization, e
}

func (r *Ridder) UpdateOrganization(ridderID int64, organization *Organization) *errortools.Error {
	url := fmt.Sprintf("organizations/%v", ridderID)

	if organization == nil {
		return nil
	}

	ev := organization.validate()

	req, res, e := r.Post(url, &organization, nil)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return e
}

func (r *Ridder) CreateOrganization(ridderID int64, newOrganization *Organization) (*Organization, *errortools.Error) {
	url := fmt.Sprintf("organizations")

	if newOrganization == nil {
		return nil, nil
	}

	ev := newOrganization.validate()

	organization := Organization{}
	req, res, e := r.Post(url, &newOrganization, &organization)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return &organization, e
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

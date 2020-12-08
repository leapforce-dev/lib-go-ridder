package ridder

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type Opportunity struct {
	RidderID           int64   `json:"RidderId"`
	InsightlyID        int64   `json:"InsightlyId"`
	InsightlyState     string  `json:"InsightlyState"`
	OfferNumber        int64   `json:"OfferNumber"`
	OpportunityName    string  `json:"OpportunityName"`
	OrganizationID     int64   `json:"OrganizationId"`
	ContactID          int64   `json:"ContactId"`
	Currency           string  `json:"Currency"`
	OpportunityCreated *string `json:"OpportunityCreated"`
	ForecastCloseDate  *string `json:"ForecastCloseDate"`
	SalesPerson        int64   `json:"SalesPerson"`
}

func (r *Ridder) GetOpportunity(ridderID int64) (*Opportunity, *errortools.Error) {
	url := fmt.Sprintf("opportunities?ridderid=%v", ridderID)

	opportunity := Opportunity{}
	_, _, e := r.Get(url, &opportunity)

	return &opportunity, e
}

func (r *Ridder) UpdateOpportunity(ridderID int64, opportunity *Opportunity) *errortools.Error {
	url := fmt.Sprintf("opportunities/%v", ridderID)

	if opportunity == nil {
		return nil
	}

	ev := opportunity.validate()

	req, res, e := r.Post(url, &opportunity, nil)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return e
}

func (r *Ridder) CreateOpportunity(ridderID int64, newOpportunity *Opportunity) (*Opportunity, *errortools.Error) {
	url := fmt.Sprintf("opportunities")

	if newOpportunity == nil {
		return nil, nil
	}

	ev := newOpportunity.validate()

	opportunity := Opportunity{}
	req, res, e := r.Post(url, &newOpportunity, &opportunity)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return &opportunity, e
}

func (opportunity *Opportunity) validate() *errortools.Error {
	if opportunity == nil {
		return nil
	}

	if len(opportunity.OpportunityName) > MaxLengthOpportunityName {
		(*opportunity).OpportunityName = opportunity.OpportunityName[:MaxLengthOpportunityName]

		return errortools.ErrorMessage(fmt.Sprintf("OpportunityName truncated to %v characters.", MaxLengthOpportunityName))
	}

	return nil
}

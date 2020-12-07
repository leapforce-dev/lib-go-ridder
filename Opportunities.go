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
	ContactID          *int64  `json:"ContactId"`
	Currency           string  `json:"Currency"`
	OpportunityCreated *string `json:"OpportunityCreated"`
	ForecastCloseDate  *string `json:"ForecastCloseDate"`
	SalesPerson        *int64  `json:"SalesPerson"`
}

func (r *Ridder) GetOpportunities(ridderID int64) (*Opportunity, *errortools.Error) {
	url := fmt.Sprintf("opportunities?ridderid=%v", ridderID)

	opportunity := Opportunity{}
	e := r.Get(url, &opportunity)

	return &opportunity, e
}

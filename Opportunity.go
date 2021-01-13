package ridder

import (
	"fmt"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type Workflow string

const (
	WorkflowNone                  Workflow = "None"
	WorkflowReject                Workflow = "Reject"
	WorkflowRejectAndMakeHistoric Workflow = "RejectAndMakeHistoric"
	WorkflowMakeHistoric          Workflow = "MakeHistoric"
	WorkflowReOpen                Workflow = "ReOpen"
)

type Opportunity struct {
	RidderID             int32  `json:"RidderId"`
	InsightlyID          int32  `json:"InsightlyId"`
	InsightlyState       string `json:"InsightlyState"`
	OfferNumber          int32  `json:"OfferNumber"`
	OpportunityName      string `json:"OpportunityName"`
	OrganizationID       int32  `json:"OrganizationId"`
	ContactID            int32  `json:"ContactId"`
	Currency             string `json:"Currency"`
	OpportunityCreated   string `json:"OpportunityCreated"`
	ForecastCloseDate    string `json:"ForecastCloseDate"`
	ProbabilityOfWinning int32  `json:"ProbabilityOfWinning"`
	SalesPerson          int32  `json:"SalesPerson"`
}

type OpportunityResponse struct {
	RidderID    int32 `json:"RidderId"`
	OfferNumber int32 `json:"OfferNumber"`
}

func (service *Service) GetOpportunity(ridderID int32) (*Opportunity, *errortools.Error) {
	url := fmt.Sprintf("opportunities?ridderid=%v", ridderID)

	opportunity := Opportunity{}
	_, _, e := service.Get(url, &opportunity)

	return &opportunity, e
}

func (service *Service) UpdateOpportunity(opportunity *Opportunity) (*OpportunityResponse, *errortools.Error) {
	if opportunity == nil {
		return nil, nil
	}

	url := fmt.Sprintf("opportunities/%v", opportunity.RidderID)

	ev := service.validateOpportunity(opportunity)

	opportunityResponse := OpportunityResponse{}
	req, res, e := service.Post(url, &opportunity, &opportunityResponse)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return &opportunityResponse, e
}

func (service *Service) CreateOpportunity(newOpportunity *Opportunity) (*OpportunityResponse, *errortools.Error) {
	url := fmt.Sprintf("opportunities")

	if newOpportunity == nil {
		return nil, nil
	}

	ev := service.validateOpportunity(newOpportunity)

	opportunityResponse := OpportunityResponse{}
	req, res, e := service.Post(url, &newOpportunity, &opportunityResponse)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return &opportunityResponse, e
}

func (service *Service) WorkflowOpportunity(opportunity *Opportunity, workflow Workflow) *errortools.Error {
	if opportunity == nil {
		return nil
	}

	if workflow == WorkflowNone {
		return nil
	}

	url := fmt.Sprintf("opportunities/%v/%s", opportunity.RidderID, workflow)

	_, _, e := service.Post(url, &opportunity, nil)
	return e
}

func (service *Service) validateOpportunity(opportunity *Opportunity) *errortools.Error {
	if opportunity == nil {
		return nil
	}

	errors := []string{}

	if len(opportunity.OpportunityName) > MaxLengthOpportunityName {
		(*opportunity).OpportunityName = opportunity.OpportunityName[:MaxLengthOpportunityName]

		errors = append(errors, fmt.Sprintf("OpportunityName truncated to %v characters.", MaxLengthOpportunityName))
	}

	e := service.removeSpecialCharacters(&(*opportunity).OpportunityName)
	if e != nil {
		errors = append(errors, e.Message())
	}

	if len(errors) > 0 {
		return errortools.ErrorMessage(strings.Join(errors, "\n"))
	}

	return nil
}

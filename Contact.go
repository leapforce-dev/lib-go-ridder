package ridder

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type Contact struct {
	RidderID             int64  `json:"RidderId"`
	InsightlyID          int64  `json:"InsightlyId"`
	Person               Person `json:"Person"`
	Email                string `json:"Email"`
	Cellphone            string `json:"Cellphone"`
	Phone                string `json:"Phone"`
	Manual               bool   `json:"Manual"`
	MainContact          bool   `json:"MainContact"`
	MainContactCreditor  bool   `json:"MainContactCreditor"`
	MainContactDebtor    bool   `json:"MainContactDebtor"`
	MainContactInvoice   bool   `json:"MainContactInvoice"`
	FunctionName         string `json:"FunctionName"`
	EmploymentTerminated bool   `json:"EmploymentTerminated"`
	OrganizationID       int64  `json:"OrganizationId"`
}

func (r *Ridder) GetContact(ridderID int64) (*Contact, *errortools.Error) {
	url := fmt.Sprintf("contacts?ridderid=%v", ridderID)

	contact := Contact{}
	_, _, e := r.Get(url, &contact)

	return &contact, e
}

func (r *Ridder) UpdateContact(ridderID int64, contact *Contact) *errortools.Error {
	url := fmt.Sprintf("contacts/%v", ridderID)

	_, _, e := r.Post(url, &contact, nil)

	return e
}

func (r *Ridder) CreateContact(ridderID int64, newContact *Contact) (*Contact, *errortools.Error) {
	url := fmt.Sprintf("contacts")

	contact := Contact{}
	_, _, e := r.Post(url, &newContact, &contact)

	return &contact, e
}

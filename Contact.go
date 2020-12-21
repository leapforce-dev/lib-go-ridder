package ridder

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type Contact struct {
	RidderID             int32  `json:"RidderId"`
	InsightlyID          int32  `json:"InsightlyId"`
	Person               Person `json:"Person"`
	Email                string `json:"Email"`
	Cellphone            string `json:"Cellphone"`
	Phone                string `json:"Phone"`
	Manual               bool   `json:"Manual"`
	MainContact          bool   `json:"MainContact"`
	MainContactCreditor  bool   `json:"MainContactCreditor"`
	MainContactDebtor    bool   `json:"MainContactDebtor"`
	FunctionName         string `json:"FunctionName"`
	EmploymentTerminated bool   `json:"EmploymentTerminated"`
	OrganizationID       int32  `json:"OrganizationId"`
}

func (r *Ridder) GetContact(ridderID int32) (*Contact, *errortools.Error) {
	url := fmt.Sprintf("contacts?ridderid=%v", ridderID)

	contact := Contact{}
	_, _, e := r.Get(url, &contact)

	return &contact, e
}

func (r *Ridder) UpdateContact(contact *Contact) (*int32, *errortools.Error) {
	url := fmt.Sprintf("contacts/%v", contact.RidderID)

	contactID := new(int32)
	_, _, e := r.Post(url, &contact, &contactID)

	return contactID, e
}

func (r *Ridder) CreateContact(newContact *Contact) (*int32, *errortools.Error) {
	url := fmt.Sprintf("contacts")

	contactID := new(int32)
	_, _, e := r.Post(url, &newContact, &contactID)

	return contactID, e
}

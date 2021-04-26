package ridder

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Contact struct {
	ID                  int32   `json:"Id"`
	RelationID          int32   `json:"RelationId"`
	ExternalCRMID       string  `json:"ExternalCrmId"`
	Person              Person  `json:"Person"`
	Email               *string `json:"Email"`
	Fax                 *string `json:"Fax"`
	BusinessPhone1      *string `json:"BusinessPhone1"`
	BusinessPhone2      *string `json:"BusinessPhone2"`
	BusinessMobilePhone *string `json:"BusinessMobilePhone"`
	PrivatePhone1       *string `json:"PrivatePhone1"`
	PrivatePhone2       *string `json:"PrivatePhone2"`
	PrivateMobilePhone  *string `json:"PrivateMobilePhone"`
	PositionID          *int32  `json:"PositionId"`
	Memo                *string `json:"Memo"`
}

/*
func (service *Service) GetContact(ridderID int32) (*Contact, *errortools.Error) {
	contact := Contact{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("contacts?ridderid=%v", ridderID)),
		ResponseModel: &contact,
	}
	_, _, e := service.get(&requestConfig)

	return &contact, e
}*/

func (service *Service) UpdateContact(contact *Contact) *errortools.Error {
	if contact == nil {
		return nil
	}

	ev := service.validateContact(contact)

	requestConfig := go_http.RequestConfig{
		URL:       service.url("contacts"),
		BodyModel: contact,
	}
	req, res, e := service.put(&requestConfig)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return e
}

func (service *Service) CreateContact(contact *Contact) (*int32, *errortools.Error) {
	if contact == nil {
		return nil, nil
	}

	ev := service.validateContact(contact)

	contactID := new(int32)

	requestConfig := go_http.RequestConfig{
		URL:           service.url("contacts"),
		BodyModel:     contact,
		ResponseModel: contactID,
	}
	req, res, e := service.post(&requestConfig)

	if ev != nil {
		ev.SetRequest(req)
		ev.SetResponse(res)
		errortools.CaptureWarning(ev)
	}

	return contactID, e
}

func (service *Service) validateContact(contact *Contact) *errortools.Error {
	if contact == nil {
		return nil
	}
	/*
		errors := []string{}

		service.truncateString("Phone", &(*contact).Phone, MaxLengthContactPhone, &errors)
		service.truncateString("Email", &(*contact).Email, MaxLengthContactEmail, &errors)
		service.truncateString("FunctionName", &(*contact).FunctionName, MaxLengthContactFunctionName, &errors)
		service.truncateString("Cellphone", &(*contact).Cellphone, MaxLengthContactCellphone, &errors)
		service.truncateString("LastName", &(*contact).Person.LastName, MaxLengthContactLastName, &errors)
		service.truncateString("Initials", &(*contact).Person.Initials, MaxLengthContactInitials, &errors)
		service.truncateString("FirstName", &(*contact).Person.FirstName, MaxLengthContactFirstName, &errors)

		if len(errors) > 0 {
			return errortools.ErrorMessage(strings.Join(errors, "\n"))
		}*/

	return nil
}

/*
func (c *Contact) MarshalJSON() ([]byte, error) {
	val := reflect.ValueOf(c).Elem()
	for i := 0; i < val.Type().NumField(); i++ {
		fmt.Println(val.Type().Field(i).Tag.Get("json"))
	}
}
*/

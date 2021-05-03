package ridder

import (
	"strconv"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Contact struct {
	ID                  int32   `json:"Id"`
	RelationID          int32   `json:"RelationId"`
	ExternalCRMID       string  `json:"ExternalCrmId"`
	Person              Person  `json:"Person"`
	Email               *string `json:"Email,omitempty"`
	Fax                 *string `json:"Fax,omitempty"`
	BusinessPhone1      *string `json:"BusinessPhone1,omitempty"`
	BusinessPhone2      *string `json:"BusinessPhone2,omitempty"`
	BusinessMobilePhone *string `json:"BusinessMobilePhone,omitempty"`
	PrivatePhone1       *string `json:"PrivatePhone1,omitempty"`
	PrivatePhone2       *string `json:"PrivatePhone2,omitempty"`
	PrivateMobilePhone  *string `json:"PrivateMobilePhone,omitempty"`
	PositionID          *int32  `json:"PositionId,omitempty"`
	Memo                *string `json:"Memo,omitempty"`
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

	e := service.validateContact(contact)
	if e != nil {
		errortools.CaptureWarning(e)
	}

	requestConfig := go_http.RequestConfig{
		URL:       service.url("contacts"),
		BodyModel: contact,
	}
	_, _, e = service.put(&requestConfig)
	if e != nil {
		return e
	}

	return nil
}

func (service *Service) CreateContact(contact *Contact) (*int32, *errortools.Error) {
	if contact == nil {
		return nil, nil
	}

	e := service.validateContact(contact)
	if e != nil {
		errortools.CaptureWarning(e)
	}

	var contactIDString string

	requestConfig := go_http.RequestConfig{
		URL:           service.url("contacts"),
		BodyModel:     contact,
		ResponseModel: &contactIDString,
	}
	_, _, e = service.post(&requestConfig)

	contactIDInt64, err := strconv.ParseInt(contactIDString, 10, 64)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	contactIDInt32 := int32(contactIDInt64)

	return &contactIDInt32, e
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

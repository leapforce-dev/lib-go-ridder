package ridder

import (
	r_types "github.com/leapforce-libraries/go_ridder/types"
)

type Person struct {
	Initials       *string             `json:"Initials"`
	FirstName      string              `json:"FirstName"`
	LastNamePrefix *string             `json:"LastNamePrefix"`
	LastName       string              `json:"LastName"`
	Gender         Gender              `json:"Gender"`
	TitleID        *int32              `json:"TitleId"`
	DateOfBirth    *r_types.DateString `json:"DateOfBirth"`
	PrivateAddress *Address            `json:"PrivateAddress"`
}

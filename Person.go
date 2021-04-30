package ridder

import (
	r_types "github.com/leapforce-libraries/go_ridder/types"
)

type Person struct {
	Initials       *string             `json:"Initials,omitempty"`
	FirstName      string              `json:"FirstName"`
	LastNamePrefix *string             `json:"LastNamePrefix,omitempty"`
	LastName       string              `json:"LastName"`
	Gender         Gender              `json:"Gender"`
	TitleID        *int32              `json:"TitleId,omitempty"`
	DateOfBirth    *r_types.DateString `json:"DateOfBirth,omitempty"`
	PrivateAddress *Address            `json:"PrivateAddress,omitempty"`
}

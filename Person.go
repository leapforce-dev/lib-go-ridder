package ridder

import (
	r_types "github.com/leapforce-libraries/go_ridder/types"
)

type Person struct {
	Initials       *string             `json:"Initials,omitempty" max:"50"`
	FirstName      string              `json:"FirstName" max:"127"`
	LastNamePrefix *string             `json:"LastNamePrefix,omitempty" max:"50"`
	LastName       string              `json:"LastName" max:"127"`
	Gender         Gender              `json:"Gender"`
	TitleID        *int32              `json:"TitleId,omitempty"`
	DateOfBirth    *r_types.DateString `json:"DateOfBirth,omitempty"`
	PrivateAddress *Address            `json:"PrivateAddress,omitempty"`
}

package ridder

type Address struct {
	ID          int32   `json:"Id,omitempty"`
	CountryID   int32   `json:"CountryId,omitempty"`
	City        *string `json:"City,omitempty" max:"50"`
	ZIPCode     *string `json:"ZipCode,omitempty" max:"50"`
	Street      *string `json:"Street,omitempty" max:"50"`
	HouseNumber *int32  `json:"HouseNumber,omitempty"`
	Addition    *string `json:"Addition,omitempty" max:"50"`
	StateID     *int32  `json:"StateId,omitempty"`
}

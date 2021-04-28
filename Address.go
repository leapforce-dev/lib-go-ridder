package ridder

type Address struct {
	ID          int32   `json:"Id,omitempty"`
	CountryID   int32   `json:"CountryId,omitempty"`
	City        *string `json:"City,omitempty"`
	ZIPCode     *string `json:"ZipCode,omitempty"`
	Street      *string `json:"Street,omitempty"`
	HouseNumber *int32  `json:"HouseNumber,omitempty"`
	Addition    *string `json:"Addition,omitempty"`
	StateID     *int32  `json:"StateId,omitempty"`
}

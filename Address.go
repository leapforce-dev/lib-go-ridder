package ridder

type Address struct {
	ID          int32   `json:"Id"`
	CountryID   int32   `json:"CountryId"`
	City        *string `json:"City"`
	ZIPCode     *string `json:"ZipCode"`
	Street      *string `json:"Street"`
	HouseNumber *int32  `json:"HouseNumber"`
	Addition    *string `json:"Addition"`
	StateID     *int32  `json:"StateId"`
}

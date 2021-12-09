package main

type Address struct {
	HouseNumber   string `json:"house_number,omitempty"`
	Road          string `json:"road,omitempty"`
	Residential   string `json:"residential,omitempty"`
	Borough       string `json:"borough,omitempty"`
	Neighbourhood string `json:"neighbourhood,omitempty"`
	Quarter       string `json:"quarter,omitempty"`
	Hamlet        string `json:"hamlet,omitempty"`
	Suburb        string `json:"suburb,omitempty"`
	Island        string `json:"island,omitempty"`
	Village       string `json:"village,omitempty"`
	Town          string `json:"town,omitempty"`
	City          string `json:"city,omitempty"`
	CityDistrict  string `json:"city_district,omitempty"`
	County        string `json:"county,omitempty"`
	State         string `json:"state,omitempty"`
	StateDistrict string `json:"state_district,omitempty"`
	Postcode      string `json:"postcode,omitempty"`
	Country       string `json:"country,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	StateCode     string `json:"state_code,omitempty"`
}

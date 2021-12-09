package main

// Covid Details
type CovidDetails struct {
	State                 string `json:"state" xml:"state"`
	TotalNoOfCasesInState string `json:"TotalNoOfCasesInState" xml:"TotalNoOfCasesInState"`
	TotalNoOfCasesInIndia string `json:"TotalNoOfCasesInIndia" xml:"TotalNoOfCasesInIndia"`
	LastUpdatedDate       string `json:"LastUpdatedDate" xml:"LastUpdatedDate"`
}

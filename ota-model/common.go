package ota_model

type RequestorID struct {
	ID        string `json:"ID"`
	IDContext string `json:"ID_Context"`
	Type      string `json:"Type"`
}
type Source struct {
	RequestorID RequestorID `json:"RequestorId"`
}
type POS struct {
	Source Source `json:"source"`
}

type StatusApplicationControl struct {
	InvTypeCode  string `json:"InvTypeCode"`
	RatePlanCode string `json:"RatePlanCode"`
	Start        string `json:"Start"`
	End          string `json:"End"`
	Mon          string `json:"Mon"`
	Tue          string `json:"Tue"`
	Weds         string `json:"Weds"`
	Thur         string `json:"Thur"`
	Fri          string `json:"Fri"`
	Sat          string `json:"Sat"`
	Sun          string `json:"Sun"`
}

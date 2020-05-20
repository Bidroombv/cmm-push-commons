package otamodel

type HotelInvCountNotifRQ struct {
	OTAHotelInvCountNotifRQ OTAHotelInvCountNotifRQ `json:"OTA_HotelInvCountNotifRQ"`
}

type InvCount struct {
	Count        string                 `json:"Count"`
	CountType    string                 `json:"CountType"`
	OriginalData map[string]interface{} `json:"original_data"`
}

type InvCounts struct {
	InvCount []InvCount `json:"InvCount"`
}

type Inventory struct {
	StatusApplicationControl StatusApplicationControl `json:"StatusApplicationControl"`
	InvCounts                InvCounts                `json:"InvCounts"`
}

type Inventories struct {
	Inventory []Inventory `json:"Inventory"`
	HotelCode string      `json:"HotelCode"  valid:"required~Hotel code cannot be blank"`
	HotelName string      `xml:"HotelName,attr" json:"HotelName,omitempty"`
}

type OTAHotelInvCountNotifRQ struct {
	Inventories   Inventories `json:"Inventories"`
	EchoToken     string      `json:"EchoToken"`
	PrimaryLangID string      `json:"PrimaryLangID"`
	Target        string      `json:"Target"`
	TimeStamp     string      `json:"TimeStamp"`
	Version       string      `json:"Version"`
}

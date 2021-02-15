package otamodel

type HotelInvCountNotifRQ struct {
	OTAHotelInvCountNotifRQ OTAHotelInvCountNotifRQ `json:"OTA_HotelInvCountNotifRQ"`
}

type InvCount struct {
	Count        string                 `json:"Count"`
	CountType    string                 `json:"CountType"`
	OriginalData map[string]interface{} `json:"original_data,omitempty"`
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
	ChainCode string      `json:"ChainCode,omitempty"`
	HotelName string      `json:"HotelName,omitempty"`
}

type OTAHotelInvCountNotifRQ struct {
	EchoToken   string      `json:"EchoToken"`
	Inventories Inventories `json:"Inventories"`
}

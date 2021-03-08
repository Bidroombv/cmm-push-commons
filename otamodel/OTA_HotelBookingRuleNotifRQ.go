package otamodel

type HotelBookingRuleNotifRQ struct {
	OTAHotelBookingRuleNotifRQ OTAHotelBookingRuleNotifRQ `json:"OTA_HotelBookingRuleNotifRQ"`
}

type RuleMessage struct {
	StatusApplicationControl StatusApplicationControl `json:"StatusApplicationControl"`
	BookingRules             BookingRules             `json:"BookingRules"`
}

type RuleMessages struct {
	HotelCode   string        `json:"HotelCode" valid:"required~Hotel code cannot be blank"`
	ChainCode   string        `json:"ChainCode,omitempty"`
	RuleMessage []RuleMessage `json:"RuleMessage"`
}

type OTAHotelBookingRuleNotifRQ struct {
	EchoToken     string       `json:"EchoToken"`
	PrimaryLangID string       `json:"PrimaryLangID"`
	Target        string       `json:"Target"`
	TimeStamp     string       `json:"TimeStamp"`
	Version       string       `json:"Version"`
	RuleMessages  RuleMessages `json:"RuleMessages"`
}

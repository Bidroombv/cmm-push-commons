package otamodel

type HotelBookingRuleNotifRQ struct {
	OTAHotelBookingRuleNotifRQ OTAHotelBookingRuleNotifRQ `json:"OTA_HotelBookingRuleNotifRQ"`
}

type RestrictionStatus struct {
	Restriction string `json:"Restriction"`
	Status      string `json:"Status"`
}

// DaysOfWeek model
type DaysOfWeek struct {
	Mon  string `json:"Mon"`
	Tue  string `json:"Tue"`
	Weds string `json:"Weds"`
	Thur string `json:"Thur"`
	Fri  string `json:"Fri"`
	Sat  string `json:"Sat"`
	Sun  string `json:"Sun"`
}

type RuleMessage struct {
	StatusApplicationControl StatusApplicationControl `json:"StatusApplicationControl"`
	BookingRules             BookingRules             `json:"BookingRules"`
}

type RuleMessages struct {
	HotelCode   string        `json:"HotelCode" valid:"required~Hotel code cannot be blank"`
	RuleMessage []RuleMessage `json:"RuleMessage"`
}

type OTAHotelBookingRuleNotifRQ struct {
	EchoToken     string       `json:"EchoToken"`
	PrimaryLangID string       `json:"PrimaryLangID"`
	Target        string       `json:"Target"`
	TimeStamp     string       `json:"TimeStamp"`
	Version       string       `json:"Version"`
	POS           POS          `json:"POS"`
	RuleMessages  RuleMessages `json:"RuleMessages"`
}

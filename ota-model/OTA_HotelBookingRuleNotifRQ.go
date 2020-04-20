package ota_model

type HotelBookingRuleNotifRQ struct {
	OTAHotelBookingRuleNotifRQ OTAHotelBookingRuleNotifRQ `json:"OTA_HotelBookingRuleNotifRQ"`
	POS                        POS                        `json:"POS"`
}

type RestrictionStatus struct {
	Restriction string `json:"Restriction"`
	Status      string `json:"Status"`
}

type LengthOfStay struct {
	MinMaxMessageType string `json:"MinMaxMessageType"`
	Time              string `json:"Time"`
}

type LengthsOfStay struct {
	LengthOfStay []LengthOfStay `json:"LengthOfStay"`
}

type BookingRule struct {
	RestrictionStatus []RestrictionStatus    `json:"RestrictionStatus"`
	LengthsOfStay     LengthsOfStay          `json:"LengthsOfStay"`
	Start             string                 `json:"Start"`
	End               string                 `json:"End"`
	OriginalData      map[string]interface{} `json:"original_data"`
	DOWRestrictions   DOWRestrictions        `json:"DOWRestrictions"`
}

type DOWRestrictions struct {
	AvailableDaysOfWeek DaysOfWeek `json:"AvailableDaysOfWeek"`
	ArrivalDaysOfWeek   DaysOfWeek `json:"ArrivalDaysOfWeek"`
	DepartureDaysOfWeek DaysOfWeek `json:"DepartureDaysOfWeek"`
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

type BookingRules struct {
	BookingRule []BookingRule `json:"BookingRule"`
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

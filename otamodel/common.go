package otamodel

type StatusApplicationControl struct {
	InvTypeCode  string `json:"InvTypeCode"`
	RatePlanCode string `json:"RatePlanCode"`
	Start        string `json:"Start"`
	End          string `json:"End"`
	Mon          string `json:"Mon,omitempty"`
	Tue          string `json:"Tue,omitempty"`
	Weds         string `json:"Weds,omitempty"`
	Thur         string `json:"Thur,omitempty"`
	Fri          string `json:"Fri,omitempty"`
	Sat          string `json:"Sat,omitempty"`
	Sun          string `json:"Sun,omitempty"`
}

type LengthOfStay struct {
	MinMaxMessageType string `json:"MinMaxMessageType"`
	Time              string `json:"Time"`
	TimeUnit          string `json:"TimeUnit,omitempty"`
}

type LengthsOfStay struct {
	LengthOfStay []LengthOfStay `json:"LengthOfStay"`
}

type AdditionalGuestAmount struct {
	Amount            string `json:"Amount"`
	DecimalPlaces     string `json:"DecimalPlaces"`
	TaxInclusive      string `json:"TaxInclusive"`
	AgeQualifyingCode string `json:"AgeQualifyingCode"`
	CurrencyCode      string `json:"CurrencyCode"`
}

type AdditionalGuestAmounts struct {
	AdditionalGuestAmount []AdditionalGuestAmount `json:"AdditionalGuestAmount"`
}

type Rates struct {
	Rate []Rate `json:"Rate"`
}

type DOWRestrictions struct {
	AvailableDaysOfWeek DaysOfWeek `json:"AvailableDaysOfWeek"`
	ArrivalDaysOfWeek   DaysOfWeek `json:"ArrivalDaysOfWeek"`
	DepartureDaysOfWeek DaysOfWeek `json:"DepartureDaysOfWeek"`
}

// Guestline supports DOWRestrictions instead of DOW_Restrictions

type BookingRule struct {
	RestrictionStatus        []RestrictionStatus    `json:"RestrictionStatus"`
	LengthsOfStay            LengthsOfStay          `json:"LengthsOfStay"`
	Start                    string                 `json:"Start"`
	End                      string                 `json:"End"`
	OriginalData             map[string]interface{} `json:"original_data"`
	OldDOWRestrictions       DOWRestrictions        `json:"DOWRestrictions"`
	DOWRestrictions          DOWRestrictions        `json:"DOW_Restrictions"`
	MinAdvancedBookingOffset string                 `json:"MinAdvancedBookingOffset,omitempty"`
	MaxAdvancedBookingOffset string                 `json:"MaxAdvancedBookingOffset,omitempty"`
}

type Rate struct {
	BaseByGuestAmts        BaseByGuestAmts        `json:"BaseByGuestAmts"`
	OriginalData           map[string]interface{} `json:"original_data"`
	AdditionalGuestAmounts AdditionalGuestAmounts `json:"AdditionalGuestAmounts"`
	CancelPolicies         CancelPolicies         `json:"CancelPolicies,omitempty"`
	MealsIncluded          MealsIncluded          `json:"MealsIncluded,omitempty"`
	Fees                   Fees                   `json:"Fees"`
	GuaranteePolicies      GuaranteePolicies      `json:"GuaranteePolicies"`
	RateDescription        RateDescription        `json:"RateDescription"`
	NumberOfUnits          string                 `json:"NumberOfUnits"`
	Mon                    string                 `json:"Mon"`
	Tue                    string                 `json:"Tue"`
	Weds                   string                 `json:"Weds"`
	Thur                   string                 `json:"Thur"`
	Fri                    string                 `json:"Fri"`
	Sat                    string                 `json:"Sat"`
	Sun                    string                 `json:"Sun"`
	Start                  string                 `json:"Start"`
	End                    string                 `json:"End"`
	InvTypeCode            string                 `json:"InvTypeCode,omitempty"`
	MaxLOS                 string                 `json:"MaxLOS,omitempty"`
	MinLOS                 string                 `json:"MinLOS,omitempty"`
	RateTimeUnit           string                 `json:"RateTimeUnit,omitempty"`
}

type BaseByGuestAmt struct {
	NumberOfGuests    string `json:"NumberOfGuests"`
	AgeQualifyingCode string `json:"AgeQualifyingCode"`
	AmountAfterTax    string `json:"AmountAfterTax"`
	AmountBeforeTax   string `json:"AmountBeforeTax"`
	DecimalPlaces     string `json:"DecimalPlaces"`
	CurrencyCode      string `json:"CurrencyCode"`
}

type BaseByGuestAmts struct {
	BaseByGuestAmt []BaseByGuestAmt `json:"BaseByGuestAmt"`
}

type BookingRules struct {
	InventoryInfo InventoryInfo `json:"InventoryInfo"`
	BookingRule   []BookingRule `json:"BookingRule"`
}

type RateDescription struct {
	Text Text   `json:"Text"`
	Name string `json:"Name"`
}

type RestrictionStatus struct {
	Restriction              string `json:"Restriction"`
	Status                   string `json:"Status"`
	MinAdvancedBookingOffset string `json:"MinAdvancedBookingOffset,omitempty"`
	MaxAdvancedBookingOffset string `json:"MaxAdvancedBookingOffset,omitempty"`
}

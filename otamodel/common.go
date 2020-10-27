package otamodel

type StatusApplicationControl struct {
	InvTypeCode  string `json:"InvTypeCode"`
	RatePlanCode string `json:"RatePlanCode"`
	Mon          string `json:"Mon,omitempty"`
	Tue          string `json:"Tue,omitempty"`
	Weds         string `json:"Weds,omitempty"`
	Thur         string `json:"Thur,omitempty"`
	Fri          string `json:"Fri,omitempty"`
	Sat          string `json:"Sat,omitempty"`
	Sun          string `json:"Sun,omitempty"`
	Start        string `json:"Start"`
	End          string `json:"End"`
}

type LengthOfStay struct {
	MinMaxMessageType string `json:"MinMaxMessageType"`
	Time              string `json:"Time"`
	TimeUnit          string `json:"TimeUnit,omitempty"`
}

type LengthsOfStay struct {
	LengthOfStay []LengthOfStay `json:"LengthOfStay,omitempty"`
}

type AdditionalGuestAmount struct {
	Amount            string `json:"Amount"`
	DecimalPlaces     string `json:"DecimalPlaces,omitempty"`
	TaxInclusive      string `json:"TaxInclusive,omitempty"`
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
	RestrictionStatus        *[]RestrictionStatus   `json:"RestrictionStatus,omitempty"`
	LengthsOfStay            *LengthsOfStay         `json:"LengthsOfStay,omitempty"`
	Start                    string                 `json:"Start"`
	End                      string                 `json:"End"`
	OriginalData             map[string]interface{} `json:"original_data,omitempty"`
	OldDOWRestrictions       *DOWRestrictions       `json:"DOWRestrictions,omitempty"`
	DOWRestrictions          *DOWRestrictions       `json:"DOW_Restrictions,omitempty"`
	MinAdvancedBookingOffset string                 `json:"MinAdvancedBookingOffset,omitempty"`
	MaxAdvancedBookingOffset string                 `json:"MaxAdvancedBookingOffset,omitempty"`
}

type Rate struct {
	BaseByGuestAmts        BaseByGuestAmts         `json:"BaseByGuestAmts"`
	OriginalData           map[string]interface{}  `json:"original_data,omitempty"`
	AdditionalGuestAmounts *AdditionalGuestAmounts `json:"AdditionalGuestAmounts,omitempty"`
	CancelPolicies         *CancelPolicies         `json:"CancelPolicies,omitempty"`
	MealsIncluded          *MealsIncluded          `json:"MealsIncluded,omitempty"`
	Fees                   *Fees                   `json:"Fees,omitempty"`
	GuaranteePolicies      *GuaranteePolicies      `json:"GuaranteePolicies,omitempty"`
	RateDescription        *RateDescription        `json:"RateDescription,omitempty"`
	NumberOfUnits          string                  `json:"NumberOfUnits,omitempty"`
	Mon                    string                  `json:"Mon,omitempty"`
	Tue                    string                  `json:"Tue,omitempty"`
	Weds                   string                  `json:"Weds,omitempty"`
	Thur                   string                  `json:"Thur,omitempty"`
	Fri                    string                  `json:"Fri,omitempty"`
	Sat                    string                  `json:"Sat,omitempty"`
	Sun                    string                  `json:"Sun,omitempty"`
	Start                  string                  `json:"Start"`
	End                    string                  `json:"End"`
	InvTypeCode            string                  `json:"InvTypeCode,omitempty"`
	MaxLOS                 string                  `json:"MaxLOS,omitempty"`
	MinLOS                 string                  `json:"MinLOS,omitempty"`
	RateTimeUnit           string                  `json:"RateTimeUnit,omitempty"`
}

type BaseByGuestAmt struct {
	NumberOfGuests    string `json:"NumberOfGuests"`
	AgeQualifyingCode string `json:"AgeQualifyingCode"`
	AmountAfterTax    string `json:"AmountAfterTax"`
	AmountBeforeTax   string `json:"AmountBeforeTax,omitempty"`
	DecimalPlaces     string `json:"DecimalPlaces,omitempty"`
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
	Restriction              string `json:"Restriction,omitempty"`
	Status                   string `json:"Status,omitempty"`
	MinAdvancedBookingOffset string `json:"MinAdvancedBookingOffset,omitempty"`
	MaxAdvancedBookingOffset string `json:"MaxAdvancedBookingOffset,omitempty"`
}
type DaysOfWeek struct {
	Mon   string `json:"Mon,omitempty"`
	Tue   string `json:"Tue,omitempty"`
	Weds  string `json:"Weds,omitempty"`
	Thur  string `json:"Thur,omitempty"`
	Fri   string `json:"Fri,omitempty"`
	Sat   string `json:"Sat,omitempty"`
	Sun   string `json:"Sun,omitempty"`
	Start string `json:"Start,omitempty"`
	End   string `json:"End,omitempty"`
}

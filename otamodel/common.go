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
	CancelPolicies         *CancelPolicies        `json:"CancelPolicies,omitempty"`
	MealsIncluded          *MealsIncluded         `json:"MealsIncluded,omitempty"`
	Fees                   *Fees                  `json:"Fees,omitempty"`
	GuaranteePolicies      *GuaranteePolicies     `json:"GuaranteePolicies,omitempty"`
	RateDescription        *RateDescription       `json:"RateDescription,omitempty"`
	NumberOfUnits          string                 `json:"NumberOfUnits,omitempty"`
	Mon                    string                 `json:"Mon,omitempty"`
	Tue                    string                 `json:"Tue,omitempty"`
	Weds                   string                 `json:"Weds,omitempty"`
	Thur                   string                 `json:"Thur,omitempty"`
	Fri                    string                 `json:"Fri,omitempty"`
	Sat                    string                 `json:"Sat,omitempty"`
	Sun                    string                 `json:"Sun,omitempty"`
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

type Contact struct {
	Address *[]Address `xml:"Address,omitempty" valid:"required"`
}

type Address struct {
	AddressLine *[]string `xml:"AddressLine"`
}

type RoomTypes struct {
	Text     string      `xml:",chardata"`
	RoomType *[]RoomType `xml:"RoomType,omitempty" valid:"required"`
}

type RoomType struct {
	Text         string       `xml:",chardata"`
	RoomTypeCode string       `xml:"RoomTypeCode,attr"`
	RoomTypeName string       `xml:"RoomTypeName,attr"`
	MaxOccupancy string       `xml:"MaxOccupancy,attr"`
	Descriptons  []Descripton `xml:"Descriptons"`
}

type Descriptons struct {
	Descripton *[]Descripton `xml:"Descripton,omitempty" valid:"required"`
}

type Descripton struct {
	Image string `xml:"Image,omitempty" valid:"required"`
	URL   string `xml:"URL,omitempty" valid:"required"`
	Text  string `xml:"Text,omitempty" valid:"required"`
}

type RatePlans struct {
	RatePlan *[]RatePlan `xml:"RatePlan,omitempty" valid:"required"`
}

type RatePlan struct {
	RatePlanCode  string          `xml:"RatePlanCode,attr,omitempty" valid:"required"`
	RatePlanName  string          `xml:"RatePlanName,attr,omitempty" valid:"required"`
	RatePlanType  string          `xml:"RatePlanType,attr,omitempty" valid:"required"`
	CurrencyCode  string          `xml:"CurrencyCode,attr,omitempty" valid:"required"`
	BaseOccupancy string          `xml:"BaseOccupancy,attr,omitempty" valid:"required"`
	Descriptons   Descriptons     `xml:"Descriptons,omitempty" valid:"required"`
	StayDateRange []StayDateRange `xml:"StayDateRange,omitempty" valid:"required"`
}

type StayDateRange struct {
	Duration string `xml:"Duration,attr,omitempty" valid:"required"`
	Start    string `xml:"Start,attr,omitempty" valid:"required"`
	End      string `xml:"End,attr,omitempty" valid:"required"`
}

type Taxes struct {
	Tax *[]Tax `xml:"Tax,omitempty" valid:"required"`
}

type Tax struct {
	ChargeFrequency           string `xml:"ChargeFrequency,attr"`
	ChargeFrequencyExempt     string `xml:"ChargeFrequencyExempt,attr"`
	ChargeUnit                string `xml:"ChargeUnit,attr"`
	ChargeUnitExempt          string `xml:"ChargeUnitExempt,attr"`
	MaxChargeFrequencyApplies string `xml:"MaxChargeFrequencyApplies,attr"`
	MaxChargeUnitApplies      string `xml:"MaxChargeUnitApplies,attr"`
	Amount                    string `xml:"Amount,attr"`
	CurrencyCode              string `xml:"CurrencyCode,attr"`
	DecimalPlaces             string `xml:"DecimalPlaces,attr"`
	Percent                   string `xml:"Percent,attr"`
	Type                      string `xml:"Type,attr"`
}

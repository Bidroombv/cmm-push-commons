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
	Address *[]Address `xml:"Address,omitempty" valid:"required" json:"Address,omitempty"`
}

type Address struct {
	AddressLine *[]string `xml:"AddressLine" json:"AddressLine"`
}

type RoomTypes struct {
	RoomType *[]RoomType `xml:"RoomType,omitempty" valid:"required" json:"RoomType,omitempty"`
}

type RoomType struct {
	RoomTypeCode string               `xml:"RoomTypeCode,attr" json:"RoomTypeCode"`
	RoomTypeName string               `xml:"RoomTypeName,attr" json:"RoomTypeName"`
	MaxOccupancy string               `xml:"MaxOccupancy,attr" json:"MaxOccupancy"`
	Descriptions []ProductDescription `xml:"Descriptions" json:"Descriptions"`
}

type Descriptions struct {
	Description *[]ProductDescription `xml:"Description,omitempty" valid:"required" json:"Description,omitempty"`
}

type ProductDescription struct {
	Image string `xml:"Image,omitempty" valid:"required" json:"Image,omitempty"`
	URL   string `xml:"URL,omitempty" valid:"required" json:"URL,omitempty"`
	Text  string `xml:"Text,omitempty" valid:"required" json:"Text,omitempty"`
}

type StayDateRange struct {
	Duration string `xml:"Duration,attr,omitempty" valid:"required" json:"Duration,omitempty"`
	Start    string `xml:"Start,attr,omitempty" valid:"required" json:"Start,omitempty"`
	End      string `xml:"End,attr,omitempty" valid:"required" json:"End,omitempty"`
}

type Taxes struct {
	Tax *[]Tax `xml:"Tax,omitempty" valid:"required" json:"Tax,omitempty"`
}

type Tax struct {
	ChargeFrequency           string `xml:"ChargeFrequency,attr" json:"ChargeFrequency"`
	ChargeFrequencyExempt     string `xml:"ChargeFrequencyExempt,attr" json:"ChargeFrequencyExempt"`
	ChargeUnit                string `xml:"ChargeUnit,attr" json:"ChargeUnit"`
	ChargeUnitExempt          string `xml:"ChargeUnitExempt,attr" json:"ChargeUnitExempt"`
	MaxChargeFrequencyApplies string `xml:"MaxChargeFrequencyApplies,attr" json:"MaxChargeFrequencyApplies"`
	MaxChargeUnitApplies      string `xml:"MaxChargeUnitApplies,attr" json:"MaxChargeUnitApplies"`
	Amount                    string `xml:"Amount,attr" json:"Amount"`
	CurrencyCode              string `xml:"CurrencyCode,attr" json:"CurrencyCode"`
	DecimalPlaces             string `xml:"DecimalPlaces,attr" json:"DecimalPlaces"`
	Percent                   string `xml:"Percent,attr" json:"Percent"`
	Type                      string `xml:"Type,attr" json:"Type"`
}

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
	MinMaxMessageType string `json:"MinMaxMessageType,omitempty"`
	Time              string `json:"Time,omitempty"`
	TimeUnit          string `json:"TimeUnit,omitempty"`
	MinLOS            string `json:"MinLOS,omitempty"`
	MaxLOS            string `json:"MaxLOS,omitempty"`
}

type LengthsOfStay struct {
	LengthOfStay []LengthOfStay `json:"LengthOfStay,omitempty"`
}

type AdditionalGuestAmount struct {
	Amount            string `json:"Amount"`
	DecimalPlaces     string `json:"DecimalPlaces,omitempty"`
	TaxInclusive      string `json:"TaxInclusive,omitempty"`
	AgeQualifyingCode string `json:"AgeQualifyingCode"`
	CurrencyCode      string `json:"CurrencyCode,omitempty"`
}

type Rates struct {
	Rate []Rate `json:"Rate"`
}

type DOWRestrictions struct {
	AvailableDaysOfWeek DaysOfWeek `json:"AvailableDaysOfWeek,omitempty"`
	ArrivalDaysOfWeek   DaysOfWeek `json:"ArrivalDaysOfWeek,omitempty"`
	DepartureDaysOfWeek DaysOfWeek `json:"DepartureDaysOfWeek,omitempty"`
}

// Guestline supports DOWRestrictions instead of DOW_Restrictions

type BookingRule struct {
	RestrictionStatus        *[]RestrictionStatus   `json:"RestrictionStatus,omitempty"`
	LengthsOfStay            *LengthsOfStay         `json:"LengthsOfStay,omitempty"`
	Start                    string                 `json:"Start,omitempty"`
	End                      string                 `json:"End,omitempty"`
	OriginalData             map[string]interface{} `json:"original_data,omitempty"`
	OldDOWRestrictions       *DOWRestrictions       `json:"DOWRestrictions,omitempty"`
	DOWRestrictions          *DOWRestrictions       `json:"DOW_Restrictions,omitempty"`
	MinAdvancedBookingOffset string                 `json:"MinAdvancedBookingOffset,omitempty"`
	MaxAdvancedBookingOffset string                 `json:"MaxAdvancedBookingOffset,omitempty"`
}

type Rate struct {
	BaseByGuestAmts        *BaseByGuestAmts        `json:"BaseByGuestAmts,omitempty"`
	OriginalData           map[string]interface{}  `json:"original_data,omitempty"`
	AdditionalGuestAmounts *AdditionalGuestAmounts `json:"AdditionalGuestAmounts,omitempty"`
	CancelPolicies         *CancelPolicies         `json:"CancelPolicies,omitempty"`
	PaymentPolicies        *PaymentPolicies        `json:"PaymentPolicies,omitempty"`
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
	Start                  string                  `json:"Start,omitempty"`
	End                    string                  `json:"End,omitempty"`
	InvTypeCode            string                  `json:"InvTypeCode,omitempty"`
	InvCode                string                  `json:"InvCode,omitempty"`
	MaxLOS                 string                  `json:"MaxLOS,omitempty"`
	MinLOS                 string                  `json:"MinLOS,omitempty"`
	RateTimeUnit           string                  `json:"RateTimeUnit,omitempty"`
	Status                 string                  `json:"Status,omitempty"`
	CurrencyCode           string                  `json:"CurrencyCode,omitempty"`
	UnitMultiplier         string                  `json:"UnitMultiplier,omitempty"`
	Restrictions           *Restrictions           `json:"Restrictions,omitempty"`
	LengthOfStay           *LengthOfStay           `json:"LengthOfStay,omitempty"`
	StayOfThrough          *StayOfThrough          `json:"StayOfThrough,omitempty"`
	AdvancedBooking        *AdvancedBooking        `json:"AdvancedBooking,omitempty"`
}

type BaseByGuestAmt struct {
	NumberOfGuests    string `json:"NumberOfGuests"`
	AgeQualifyingCode string `json:"AgeQualifyingCode"`
	AmountAfterTax    string `json:"AmountAfterTax"`
	AmountBeforeTax   string `json:"AmountBeforeTax,omitempty"`
	DecimalPlaces     string `json:"DecimalPlaces,omitempty"`
	CurrencyCode      string `json:"CurrencyCode,omitempty"`
}

type AdditionalGuestAmounts struct {
	AdditionalGuestAmount []AdditionalGuestAmount `json:"AdditionalGuestAmount"`
}

type BaseByGuestAmts struct {
	BaseByGuestAmt *[]BaseByGuestAmt `json:"BaseByGuestAmt,omitempty"`
}

type CancelPolicies struct {
	CancelPenalty []CancelPenalty `json:"CancelPenalty,omitempty"`
}

type CancelPenalty struct {
	Start              string                `json:"Start,omitempty"`
	End                string                `json:"End,omitempty"`
	NonRefundable      bool                  `json:"NonRefundable,omitempty"`
	Deadline           *PenaltyDeadline      `json:"Deadline,omitempty"`
	AmountPercent      *PenaltyAmountPercent `json:"AmountPercent,omitempty"`
	PenaltyDescription *[]PenaltyDescription `json:"PenaltyDescription,omitempty"`
	Mon                string                `json:"Mon,omitempty"`
	Tue                string                `json:"Tue,omitempty"`
	Weds               string                `json:"Weds,omitempty"`
	Thur               string                `json:"Thur,omitempty"`
	Fri                string                `json:"Fri,omitempty"`
	Sat                string                `json:"Sat,omitempty"`
	Sun                string                `json:"Sun,omitempty"`
}

type BookingRules struct {
	InventoryInfo *InventoryInfo `json:"InventoryInfo,omitempty"`
	BookingRule   *[]BookingRule `json:"BookingRule,omitempty"`
}

type MealsIncluded struct {
	Breakfast     bool   `json:"Breakfast,omitempty"`
	Lunch         bool   `json:"Lunch,omitempty"`
	Dinner        bool   `json:"Dinner,omitempty"`
	MealPlanCodes string `json:"MealPlanCodes,omitempty"`
}

type Text struct {
	TextFormat string `json:"TextFormat"`
	Text       string `json:"__text"`
}

type InventoryInfo struct {
	InvTypeCode string `json:"InvTypeCode"`
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

type Restrictions struct {
	ClosedToArrival   string `json:"ClosedToArrival,omitempty"`
	ClosedToDeparture string `json:"ClosedToDeparture,omitempty"`
}

type StayOfThrough struct {
	MinST string `json:"MinST,omitempty"`
	MaxST string `json:"MaxST,omitempty"`
}

type AdvancedBooking struct {
	MinAD string `json:"MinAD,omitempty"`
	MaxAD string `json:"MaxAD,omitempty"`
}

type RatePlanTPAExtensions struct {
	TPAExtension RatePlanTPAExtension `json:"TPA_Extension,omitempty"`
}

type RatePlanTPAExtension struct {
	Category     string    `json:"Category,omitempty"`
	RatePlanCode string    `json:"RatePlanCode,omitempty"`
	Extension    Extension `json:"Extension,omitempty"`
}

type Extension struct {
	Name string         `json:"Name,omitempty"`
	Item []RatePlanItem `json:"Item,omitempty"`
}

type RatePlanItem struct {
	Text  string `json:"Text,omitempty"`
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

type PaymentPolicies struct {
	GuaranteePayment []PaymentPoliciesGuaranteePayment `json:"GuaranteePayment,omitempty"`
}

type PaymentPoliciesGuaranteePayment struct {
	Start         string                `json:"Start,omitempty"`
	End           string                `json:"End,omitempty"`
	GuaranteeType string                `json:"GuaranteeType,omitempty"`
	HoldTime      string                `json:"HoldTime,omitempty"`
	Mon           string                `json:"Mon,omitempty"`
	Tue           string                `json:"Tue,omitempty"`
	Weds          string                `json:"Weds,omitempty"`
	Thur          string                `json:"Thur,omitempty"`
	Fri           string                `json:"Fri,omitempty"`
	Sat           string                `json:"Sat,omitempty"`
	Sun           string                `json:"Sun,omitempty"`
	AmountPercent *PenaltyAmountPercent `json:"AmountPercent,omitempty"`
	Description   *Description          `json:"Description,omitempty"`
}

type PenaltyAmountPercent struct {
	TaxInclusive  bool   `json:"TaxInclusive,omitempty"`
	FeesInclusive bool   `json:"FeesInclusive,omitempty"`
	Percent       string `json:"Percent,omitempty"`
	BasisType     string `json:"BasisType,omitempty"`
	Amount        string `json:"Amount,omitempty"`
	NmbrOfNights  string `json:"NmbrOfNights,omitempty"`
	DecimalPlaces string `json:"DecimalPlaces,omitempty"`
	CurrencyCode  string `json:"CurrencyCode,omitempty"`
}

type PenaltyDescription struct {
	Text string `json:"Text,omitempty"`
}

type PenaltyDeadline struct {
	OffsetDropTime       string `json:"OffsetDropTime,omitempty"`
	OffsetTimeUnit       string `json:"OffsetTimeUnit,omitempty"`
	OffsetUnitMultiplier string `json:"OffsetUnitMultiplier,omitempty"`
	AbsoluteDeadline     string `json:"AbsoluteDeadline,omitempty"`
}

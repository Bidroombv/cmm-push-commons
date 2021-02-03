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

type Rates struct {
	Rate []Rate `json:"Rate"`
}

type Rate struct {
	CancelPolicies         *CancelPolicies         `json:"CancelPolicies,omitempty"`
	MealsIncluded          *MealsIncluded          `json:"MealsIncluded,omitempty"`
	Fees                   *Fees                   `json:"Fees,omitempty"`
	GuaranteePolicies      *GuaranteePolicies      `json:"GuaranteePolicies,omitempty"`
	RateDescription        *RateDescription        `json:"RateDescription,omitempty"`
	NumberOfUnits          string                  `json:"NumberOfUnits,omitempty"`
	CurrencyCode           string                  `json:"CurrencyCode,omitempty"`
	PaymentPolicies        *PaymentPolicies        `json:"PaymentPolicies,omitempty"`
	OriginalData           map[string]interface{}  `json:"original_data,omitempty"`
	BaseByGuestAmts        *BaseByGuestAmts        `json:"BaseByGuestAmts,omitempty"`
	AdditionalGuestAmounts *AdditionalGuestAmounts `json:"AdditionalGuestAmounts,omitempty"`
	InvTypeCode            string                  `json:"InvTypeCode,omitempty"`
	InvCode                string                  `json:"InvCode,omitempty"`
	MaxLOS                 string                  `json:"MaxLOS,omitempty"`
	MinLOS                 string                  `json:"MinLOS,omitempty"`
	RateTimeUnit           string                  `json:"RateTimeUnit,omitempty"`
	UnitMultiplier         string                  `json:"UnitMultiplier,omitempty"`
	Status                 string                  `json:"Status,omitempty"`
	Start                  string                  `json:"Start,omitempty"`
	End                    string                  `json:"End,omitempty"`
	Mon                    string                  `json:"Mon,omitempty"`
	Tue                    string                  `json:"Tue,omitempty"`
	Weds                   string                  `json:"Weds,omitempty"`
	Thur                   string                  `json:"Thur,omitempty"`
	Fri                    string                  `json:"Fri,omitempty"`
	Sat                    string                  `json:"Sat,omitempty"`
	Sun                    string                  `json:"Sun,omitempty"`
}

type BaseByGuestAmts struct {
	BaseByGuestAmt *[]BaseByGuestAmt `json:"BaseByGuestAmt,omitempty"`
}

type BaseByGuestAmt struct {
	AmountBeforeTax   string `json:"AmountBeforeTax,omitempty"`
	AmountAfterTax    string `json:"AmountAfterTax,omitempty"`
	NumberOfGuests    string `json:"NumberOfGuests,omitempty"`
	AgeQualifyingCode string `json:"AgeQualifyingCode,omitempty"`
	DecimalPlaces     string `json:"DecimalPlaces,omitempty"`
	CurrencyCode      string `json:"CurrencyCode,omitempty"`
}

type AdditionalGuestAmounts struct {
	AdditionalGuestAmount []AdditionalGuestAmount `json:"AdditionalGuestAmount"`
}

type AdditionalGuestAmount struct {
	Amount            string `json:"Amount"`
	AgeQualifyingCode string `json:"AgeQualifyingCode"`
	TaxInclusive      string `json:"TaxInclusive,omitempty"`
	DecimalPlaces     string `json:"DecimalPlaces,omitempty"`
	CurrencyCode      string `json:"CurrencyCode,omitempty"`
}

type CancelPolicies struct {
	CancelPenalty []CancelPenalty `json:"CancelPenalty,omitempty"`
}

type CancelPenalty struct {
	PolicyCode         string                `json:"PolicyCode,omitempty"`
	NonRefundable      bool                  `json:"NonRefundable,omitempty"`
	Start              string                `json:"Start,omitempty"`
	End                string                `json:"End,omitempty"`
	Mon                string                `json:"Mon,omitempty"`
	Tue                string                `json:"Tue,omitempty"`
	Weds               string                `json:"Weds,omitempty"`
	Thur               string                `json:"Thur,omitempty"`
	Fri                string                `json:"Fri,omitempty"`
	Sat                string                `json:"Sat,omitempty"`
	Sun                string                `json:"Sun,omitempty"`
	PenaltyDescription *[]PenaltyDescription `json:"PenaltyDescription,omitempty"`
	Deadline           *PenaltyDeadline      `json:"Deadline,omitempty"`
	AmountPercent      *PenaltyAmountPercent `json:"AmountPercent,omitempty"`
}

type PaymentPolicies struct {
	GuaranteePayment []PaymentPoliciesGuaranteePayment `json:"GuaranteePayment,omitempty"`
}

type PaymentPoliciesGuaranteePayment struct {
	GuaranteeType    string                `json:"GuaranteeType,omitempty"`
	HoldTime         string                `json:"HoldTime,omitempty"`
	Start            string                `json:"Start,omitempty"`
	End              string                `json:"End,omitempty"`
	Mon              string                `json:"Mon,omitempty"`
	Tue              string                `json:"Tue,omitempty"`
	Weds             string                `json:"Weds,omitempty"`
	Thur             string                `json:"Thur,omitempty"`
	Fri              string                `json:"Fri,omitempty"`
	Sat              string                `json:"Sat,omitempty"`
	Sun              string                `json:"Sun,omitempty"`
	PaymentCode      string                `json:"PaymentCode,omitempty"`
	Description      *Description          `json:"Description,omitempty"`
	AmountPercent    *PenaltyAmountPercent `json:"AmountPercent,omitempty"`
	AcceptedPayments *AcceptedPayments     `json:"AcceptedPayments,omitempty"`
}

type MealsIncluded struct {
	Breakfast     bool   `json:"Breakfast,omitempty"`
	Lunch         bool   `json:"Lunch,omitempty"`
	Dinner        bool   `json:"Dinner,omitempty"`
	MealPlanCodes string `json:"MealPlanCodes,omitempty"`
}

type Fees struct {
	Fee []Fee `json:"Fee,omitempty"`
}

type GuaranteePolicies struct {
	GuaranteePolicy []GuaranteePolicy `json:"GuaranteePolicy"`
}

type GuaranteePolicy struct {
	GuaranteesAccepted GuaranteesAccepted `json:"GuaranteesAccepted"`
}

type GuaranteesAccepted struct {
	GuaranteeAccepted GuaranteeAccepted `json:"GuaranteeAccepted"`
}

type GuaranteeAccepted struct {
	PaymentCard string `json:"PaymentCard"`
}

type RateDescription struct {
	Text Text   `json:"Text"`
	Name string `json:"Name"`
}

type Text struct {
	TextFormat string `json:"TextFormat"`
	Text       string `json:"__text"`
}

type LengthOfStay struct {
	MinMaxMessageType string `json:"MinMaxMessageType,omitempty"`
	Time              string `json:"Time,omitempty"`
	TimeUnit          string `json:"TimeUnit,omitempty"`
}

type BookingRules struct {
	InventoryInfo *InventoryInfo `json:"InventoryInfo,omitempty"`
	BookingRule   *[]BookingRule `json:"BookingRule,omitempty"`
}

type InventoryInfo struct {
	InvTypeCode string `json:"InvTypeCode"`
}

type BookingRule struct {
	OriginalData             map[string]interface{} `json:"original_data,omitempty"`
	LengthsOfStay            *LengthsOfStay         `json:"LengthsOfStay,omitempty"`
	OldDOWRestrictions       *DOWRestrictions       `json:"DOWRestrictions,omitempty"`
	DOWRestrictions          *DOWRestrictions       `json:"DOW_Restrictions,omitempty"`
	RestrictionStatus        *[]RestrictionStatus   `json:"RestrictionStatus,omitempty"`
	MinAdvancedBookingOffset string                 `json:"MinAdvancedBookingOffset,omitempty"`
	MaxAdvancedBookingOffset string                 `json:"MaxAdvancedBookingOffset,omitempty"`
	Start                    string                 `json:"Start,omitempty"`
	End                      string                 `json:"End,omitempty"`
}

type RestrictionStatus struct {
	Restriction              string `json:"Restriction,omitempty"`
	Status                   string `json:"Status,omitempty"`
	MinAdvancedBookingOffset string `json:"MinAdvancedBookingOffset,omitempty"`
	MaxAdvancedBookingOffset string `json:"MaxAdvancedBookingOffset,omitempty"`
}

type LengthsOfStay struct {
	LengthOfStay []LengthOfStay `json:"LengthOfStay,omitempty"`
}

// Guestline supports DOWRestrictions instead of DOW_Restrictions
type DOWRestrictions struct {
	AvailableDaysOfWeek DaysOfWeek `json:"AvailableDaysOfWeek,omitempty"`
	ArrivalDaysOfWeek   DaysOfWeek `json:"ArrivalDaysOfWeek,omitempty"`
	DepartureDaysOfWeek DaysOfWeek `json:"DepartureDaysOfWeek,omitempty"`
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

type Fee struct {
	Amount          string `json:"Amount,omitempty"`
	Code            string `json:"Code,omitempty"`
	DecimalPlaces   string `json:"DecimalPlaces,omitempty"`
	Percent         string `json:"Percent,omitempty"`
	ChargeFrequency string `json:"ChargeFrequency,omitempty"`
	Description     string `json:"Description,omitempty"`
}

type RatePlanTPAExtensions struct {
	TPAExtension RatePlanTPAExtension `json:"TPA_Extension,omitempty"`
}

type RatePlanTPAExtension struct {
	Category     string      `json:"Category,omitempty"`
	RatePlanCode string      `json:"RatePlanCode,omitempty"`
	Extension    []Extension `json:"Extension,omitempty"`
}

type Extension struct {
	Name string         `json:"Name,omitempty"`
	Item []RatePlanItem `json:"Item,omitempty"`
}

type RatePlanItem struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
	Text  string `json:"Text,omitempty"`
	Start string `json:"Start,omitempty"`
	End   string `json:"End,omitempty"`
}

type PenaltyAmountPercent struct {
	Percent       string `json:"Percent,omitempty"`
	NmbrOfNights  string `json:"NmbrOfNights,omitempty"`
	Amount        string `json:"Amount,omitempty"`
	CurrencyCode  string `json:"CurrencyCode,omitempty"`
	TaxInclusive  bool   `json:"TaxInclusive,omitempty"`
	FeesInclusive bool   `json:"FeesInclusive,omitempty"`
	BasisType     string `json:"BasisType,omitempty"`
	DecimalPlaces string `json:"DecimalPlaces,omitempty"`
}

type Description struct {
	Name  string `json:"Name,omitempty"`
	Text  string `json:"Text,omitempty"`
	Image string `json:"Image,omitempty"`
	URL   string `json:"URL,omitempty"`
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

//AcceptedPayments **
type AcceptedPayments struct {
	AcceptedPayment []AcceptedPayment `json:"AcceptedPayment"`
}

//AcceptedPayment **
type AcceptedPayment struct {
	PaymentCard PaymentCard `json:"PaymentCard"`
}

//PaymentCard **
type PaymentCard struct {
	CardCode string `json:"CardCode,attr"`
}

//POS **
type POS struct {
	Source Source `json:"Source"`
}

//Source **
type Source struct {
	BookingChannel BookingChannel `json:"BookingChannel"`
}

//BookingChannel **
type BookingChannel struct {
	Type        string      `json:"Type,omitempty"`
	CompanyName CompanyName `json:"CompanyName"`
}

//CompanyName **
type CompanyName struct {
	Code string `json:"Code,omitempty"`
}

// Taxes **
type Taxes struct {
	Tax []Tax `json:"Tax"`
}

// Tax **
type Tax struct {
	Code                      string `json:"Code"`
	Amount                    string `json:"Amount,omitempty"`
	Percent                   string `json:"Percent,omitempty"`
	Type                      string `json:"Type,omitempty"`
	CurrencyCode              string `json:"CurrencyCode,omitempty"`
	DecimalPlaces             string `json:"DecimalPlaces,omitempty"`
	ChargeFrequency           string `json:"ChargeFrequency,omitempty"`
	ChargeFrequencyExempt     string `json:"ChargeFrequencyExempt,omitempty"`
	ChargeUnit                string `json:"ChargeUnit,omitempty"`
	ChargeUnitExempt          string `json:"ChargeUnitExempt,omitempty"`
	MaxChargeFrequencyApplies string `json:"MaxChargeFrequencyApplies,omitempty"`
	MaxChargeUnitApplies      string `json:"MaxChargeUnitApplies,omitempty"`
}

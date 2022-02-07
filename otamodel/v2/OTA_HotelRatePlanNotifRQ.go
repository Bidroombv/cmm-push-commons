package v1

type HotelRatePlanNotifRQ struct {
	OTAHotelRatePlanNotifRQ OTAHotelRatePlanNotifRQ `json:"OTA_HotelRatePlanNotifRQ"`
}

type OTAHotelRatePlanNotifRQ struct {
	RatePlans     RatePlans      `json:"RatePlans"`
	TPAExtensions *TPAExtensions `json:"TPAExtensions,omitempty"`
}

type RatePlans struct {
	HotelCode string     `json:"HotelCode"`
	ChainCode string     `json:"ChainCode,omitempty"`
	HotelName string     `json:"HotelName,omitempty"`
	RatePlan  []RatePlan `json:"RatePlan"`
}

type RatePlan struct {
	RatePlanID         string                 `json:"RatePlanID"`
	RatePlanCode       string                 `json:"RatePlanCode"`
	RatePlanType       string                 `json:"RatePlanType"`
	Rates              Rates                  `json:"Rates"`
	CurrencyCode       string                 `json:"CurrencyCode,omitempty"`
	OriginalData       string                 `json:"original_data,omitempty"`
	RatePlanNotifType  string                 `json:"RatePlanNotifType,omitempty"`
	RatePlanStatusType string                 `json:"RatePlanStatusType,omitempty"`
	Start              string                 `json:"Start,omitempty"`
	End                string                 `json:"End,omitempty"`
	BookingRules       *BookingRules          `json:"BookingRules,omitempty"`
	SellableProducts   *SellableProducts      `json:"SellableProducts,omitempty"`
	Description        *[]RatePlanDescription `json:"Description,omitempty"`
	RatePlanLevelFee   *RatePlanLevelFee      `json:"RatePlanLevelFee,omitempty"`
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
type PaymentPolicies struct {
	GuaranteePayment []PaymentPoliciesGuaranteePayment `json:"GuaranteePayment,omitempty"`
}

type PaymentPoliciesGuaranteePayment struct {
	PolicyCode       string                `json:"PolicyCode,omitempty"`
	TPAExtensions    *TPAExtensions        `json:"TPA_Extensions,omitempty"`
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

type RateDescription struct {
	Text Text   `json:"Text"`
	Name string `json:"Name"`
}

type Text struct {
	TextFormat string `json:"TextFormat"`
	Text       string `json:"__text"`
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

type Fees struct {
	Fee []Fee `json:"Fee,omitempty"`
}
type Fee struct {
	Amount          string `json:"Amount,omitempty"`
	Code            string `json:"Code,omitempty"`
	DecimalPlaces   string `json:"DecimalPlaces,omitempty"`
	Percent         string `json:"Percent,omitempty"`
	ChargeFrequency string `json:"ChargeFrequency,omitempty"`
	Description     string `json:"Description,omitempty"`
	ChargeUnit      string `json:"ChargeUnit,omitempty"`
	CurrencyCode    string `json:"CurrencyCode,omitempty"`
	EffectiveDate   string `json:"EffectiveDate,omitempty"`
	ExpireDate      string `json:"ExpireDate,omitempty"`
	SequenceNbr     string `json:"SequenceNbr,omitempty"`
}
type MealsIncluded struct {
	Breakfast     bool   `json:"Breakfast,omitempty"`
	Lunch         bool   `json:"Lunch,omitempty"`
	Dinner        bool   `json:"Dinner,omitempty"`
	MealPlanCodes string `json:"MealPlanCodes,omitempty"`
}

type CancelPolicies struct {
	CancelPenalty []CancelPenalty `json:"CancelPenalty,omitempty"`
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
type LengthsOfStay struct {
	ArrivalDateBased string         `json:"ArrivalDateBased,omitempty"`
	LengthOfStay     []LengthOfStay `json:"LengthOfStay,omitempty"`
}
type LengthOfStay struct {
	MinMaxMessageType string `json:"MinMaxMessageType,omitempty"`
	Time              string `json:"Time,omitempty"`
	TimeUnit          string `json:"TimeUnit,omitempty"`
}

type RestrictionStatus struct {
	Restriction              string `json:"Restriction,omitempty"`
	Status                   string `json:"Status,omitempty"`
	MinAdvancedBookingOffset string `json:"MinAdvancedBookingOffset,omitempty"`
	MaxAdvancedBookingOffset string `json:"MaxAdvancedBookingOffset,omitempty"`
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

type SellableProducts struct {
	SellableProduct []SellableProduct `json:"SellableProduct"`
}

type SellableProduct struct {
	InvTypeCode string                    `json:"InvTypeCode,omitempty"`
	IsRoom      string                    `json:"IsRoom,omitempty"`
	Start       string                    `json:"Start,omitempty"`
	End         string                    `json:"End,omitempty"`
	GuestRoom   *SellableProductGuestRoom `json:"GuestRoom,omitempty"`
}

type SellableProductGuestRoom struct {
	Room SellableProductRoom `json:"Room,omitempty"`
}

type SellableProductRoom struct {
	RoomTypeCode string `json:"RoomTypeCode,omitempty"`
}

type RatePlanDescription struct {
	Name string                     `json:"Name,omitempty"`
	Text *[]RatePlanDescriptionText `json:"Text,omitempty"`
}

type RatePlanDescriptionText struct {
	Language           string `json:"Language,omitempty"`
	LastModifyDateTime string `json:"LastModifyDateTime,omitempty"`
	Text               string `json:"Text,omitempty"`
}

type RatePlanLevelFee struct {
	Fee []Fee `json:"Fee"`
}

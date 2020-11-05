package otamodel

type HotelRatePlanNotifRQ struct {
	OTAHotelRatePlanNotifRQ OTAHotelRatePlanNotifRQ `json:"OTA_HotelRatePlanNotifRQ"`
}

type InventoryInfo struct {
	InvTypeCode string `json:"InvTypeCode"`
}

type CancelPenalty struct {
	Start              string                `json:"Start,omitempty"`
	End                string                `json:"End,omitempty"`
	NonRefundable      bool                  `json:"NonRefundable"`
	Deadline           *PenaltyDeadline      `json:"Deadline,omitempty"`
	AmountPercent      *PenaltyAmountPercent `json:"AmountPercent,omitempty"`
	PenaltyDescription *[]PenaltyDescription `json:"PenaltyDescription,omitempty"`
}

type PenaltyAmountPercent struct {
	Percent      string `json:"Percent,omitempty"`
	BasisType    string `json:"BasisType,omitempty"`
	Amount       string `json:"Amount,omitempty"`
	NmbrOfNights string `json:"NmbrOfNights,omitempty"`
}

type PenaltyDescription struct {
	Text string `json:"Text,omitempty"`
}

type PenaltyDeadline struct {
	OffsetDropTime       string `json:"OffsetDropTime,omitempty"`
	OffsetTimeUnit       string `json:"OffsetTimeUnit,omitempty"`
	OffsetUnitMultiplier string `json:"OffsetUnitMultiplier,omitempty"`
}

type CancelPolicies struct {
	CancelPenalty []CancelPenalty `json:"CancelPenalty,omitempty"`
}

type MealsIncluded struct {
	Breakfast     bool   `json:"Breakfast"`
	Dinner        bool   `json:"Dinner"`
	MealPlanCodes string `json:"MealPlanCodes"`
}

type Text struct {
	TextFormat string `json:"TextFormat"`
	Text       string `json:"__text"`
}

type SellableProduct struct {
	InvTypeCode string `json:"InvTypeCode"`
	IsRoom      string `json:"IsRoom"`
}

type SellableProducts struct {
	SellableProduct []SellableProduct `json:"SellableProduct"`
}

type RatePlanLevelFee struct {
	Fee []Fee `json:"Fee"`
}

type Description struct {
	Name string `json:"Name,omitempty"`
	Text string `json:"Text,omitempty"`
}

type RatePlan struct {
	CurrencyCode     string            `json:"CurrencyCode,omitempty"`
	OriginalData     string            `json:"original_data,omitempty"`
	RatePlanID       string            `json:"RatePlanID,omitempty"`
	RatePlanCode     string            `json:"RatePlanCode,omitempty"`
	Rates            Rates             `json:"Rates"`
	BookingRules     *BookingRules     `json:"BookingRules,omitempty"`
	SellableProducts *SellableProducts `json:"SellableProducts,omitempty"`
	Description      []Description     `json:"Description,omitempty"`
	RatePlanLevelFee *RatePlanLevelFee `json:"RatePlanLevelFee,omitempty"`
}

type RatePlans struct {
	HotelCode string     `json:"HotelCode"`
	HotelName string     `json:"HotelName,omitempty"`
	RatePlan  []RatePlan `json:"RatePlan"`
}

type OTAHotelRatePlanNotifRQ struct {
	RatePlans RatePlans `json:"RatePlans"`
}

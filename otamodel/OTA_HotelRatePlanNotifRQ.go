package otamodel

type HotelRatePlanNotifRQ struct {
	OTAHotelRatePlanNotifRQ OTAHotelRatePlanNotifRQ `json:"OTA_HotelRatePlanNotifRQ"`
}

type InventoryInfo struct {
	InvTypeCode string `json:"InvTypeCode"`
}

type CancelPenalty struct {
	NonRefundable bool `json:"NonRefundable"`
}

type CancelPolicies struct {
	CancelPenalty []CancelPenalty `json:"CancelPenalty"`
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
	BookingRules     *BookingRules     `json:"BookingRules,omitempty"`
	Rates            Rates             `json:"Rates"`
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

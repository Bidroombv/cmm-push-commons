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
	Breakfast bool `json:"Breakfast"`
	Dinner    bool `json:"Dinner"`
}

type Text struct {
	TextFormat string `json:"TextFormat"`
	Text       string `json:"__text"`
}

type SellableProduct struct {
	InvCode string `json:"InvCode"`
}

type SellableProducts struct {
	SellableProduct []SellableProduct `json:"SellableProduct"`
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
}

type RatePlans struct {
	HotelCode string     `json:"HotelCode"`
	HotelName string     `json:"HotelName,omitempty"`
	RatePlan  []RatePlan `json:"RatePlan"`
}

type OTAHotelRatePlanNotifRQ struct {
	RatePlans RatePlans `json:"RatePlans"`
}

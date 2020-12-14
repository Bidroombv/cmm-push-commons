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
	NonRefundable      bool                  `json:"NonRefundable,omitempty"`
	NoCancelInd        bool                  `json:"NoCancelInd,omitempty"`
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

type CancelPolicies struct {
	CancelPenalty []CancelPenalty `json:"CancelPenalty,omitempty"`
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

type RatePlanDescription struct {
	Name string                     `json:"Name,omitempty"`
	Text *[]RatePlanDescriptionText `json:"Text,omitempty"`
}

type RatePlanDescriptionText struct {
	Language string `json:"Language,omitempty"`
	Text     string `json:"Text,omitempty"`
}

type RatePlan struct {
	CurrencyCode       string                 `json:"CurrencyCode,omitempty"`
	OriginalData       string                 `json:"original_data,omitempty"`
	RatePlanID         string                 `json:"RatePlanID,omitempty"`
	RatePlanCode       string                 `json:"RatePlanCode,omitempty"`
	RatePlanNotifType  string                 `json:"RatePlanNotifType,omitempty"`
	Rates              Rates                  `json:"Rates"`
	RatePlanStatusType string                 `json:"RatePlanStatusType,omitempty"`
	Start              string                 `json:"Start,omitempty"`
	End                string                 `json:"End,omitempty"`
	BookingRules       *BookingRules          `json:"BookingRules,omitempty"`
	SellableProducts   *SellableProducts      `json:"SellableProducts,omitempty"`
	Description        *[]RatePlanDescription `json:"Description,omitempty"`
	RatePlanLevelFee   *RatePlanLevelFee      `json:"RatePlanLevelFee,omitempty"`
}

type RatePlans struct {
	HotelCode string     `json:"HotelCode"`
	ChainCode string     `json:"ChainCode,omitempty"`
	HotelName string     `json:"HotelName,omitempty"`
	RatePlan  []RatePlan `json:"RatePlan"`
}

type OTAHotelRatePlanNotifRQ struct {
	RatePlans     RatePlans              `json:"RatePlans"`
	TPAExtensions *RatePlanTPAExtensions `json:"TPAExtensions,omitempty"`
}

package otamodel

type HotelRatePlanNotifRQ struct {
	OTAHotelRatePlanNotifRQ OTAHotelRatePlanNotifRQ `json:"OTA_HotelRatePlanNotifRQ"`
}

type OTAHotelRatePlanNotifRQ struct {
	RatePlans     RatePlans              `json:"RatePlans"`
	TPAExtensions *RatePlanTPAExtensions `json:"TPAExtensions,omitempty"`
}

type RatePlans struct {
	HotelCode string     `json:"HotelCode"`
	RatePlan  []RatePlan `json:"RatePlan"`
	ChainCode string     `json:"ChainCode,omitempty"`
	HotelName string     `json:"HotelName,omitempty"`
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

type SellableProducts struct {
	SellableProduct []SellableProduct `json:"SellableProduct"`
}

type RatePlanDescription struct {
	Name string                     `json:"Name,omitempty"`
	Text *[]RatePlanDescriptionText `json:"Text,omitempty"`
}

type RatePlanLevelFee struct {
	Fee []Fee `json:"Fee"`
}

type RatePlanDescriptionText struct {
	Language string `json:"Language,omitempty"`
	Text     string `json:"Text,omitempty"`
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

type Description struct {
	Name string `json:"Name,omitempty"`
	Text string `json:"Text,omitempty"`
}

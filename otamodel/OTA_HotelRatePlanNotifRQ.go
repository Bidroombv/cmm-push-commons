package otamodel

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
	Language string `json:"Language,omitempty"`
	Text     string `json:"Text,omitempty"`
}

type RatePlanLevelFee struct {
	Fee []Fee `json:"Fee"`
}

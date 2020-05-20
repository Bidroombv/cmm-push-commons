package otamodel

type HotelAvailNotifRQ struct {
	OTAHotelAvailNotifRQ OTAHotelAvailNotifRQ `json:"OTA_HotelAvailNotifRQ"`
}

type OTAHotelAvailNotifRQ struct {
	AvailStatusMessages AvailStatusMessages `json:"AvailStatusMessages"`
}

type AvailStatusMessages struct {
	HotelCode          string               `json:"HotelCode"`
	AvailStatusMessage []AvailStatusMessage `json:"AvailStatusMessage"`
}

type AvailStatusMessage struct {
	StatusApplicationControl StatusApplicationControl `json:"StatusApplicationControl"`
	LengthsOfStay            LengthsOfStay            `json:"LengthsOfStay,omitempty"`
	RestrictionStatus        RestrictionStatus        `json:"RestrictionStatus,omitempty"`
}

//type RestrictionStatus struct {
//	Restriction string `json:"_Restriction"`
//	Status      string `json:"_Status"`
//	MinAdvancedBookingOffset string `json:"_MinAdvancedBookingOffset"`
//}

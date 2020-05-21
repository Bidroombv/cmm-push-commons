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
	OriginalData             map[string]interface{}   `json:"original_data,omitempty"`
}

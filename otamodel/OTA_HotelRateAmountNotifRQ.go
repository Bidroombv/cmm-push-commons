package otamodel

type HotelRateAmountNotifRQ struct {
	OTAHotelRateAmountNotifRQ OTAHotelRateAmountNotifRQ `json:"OTA_HotelRateAmountNotifRQ"`
}

type RateAmountMessage struct {
	StatusApplicationControl StatusApplicationControl `json:"StatusApplicationControl,omitempty"`
	Rates                    Rates                    `json:"Rates"`
}

type RateAmountMessages struct {
	RateAmountMessage []RateAmountMessage `json:"RateAmountMessage"`
	HotelCode         string              `json:"HotelCode" valid:"required~Hotel code cannot be blank"`
}

type OTAHotelRateAmountNotifRQ struct {
	RateAmountMessages RateAmountMessages `json:"RateAmountMessages"`
	EchoToken          string             `json:"EchoToken"`
}

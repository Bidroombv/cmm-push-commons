package otamodel

type HotelRateAmountNotifRQ struct {
	OTAHotelRateAmountNotifRQ OTAHotelRateAmountNotifRQ `json:"OTA_HotelRateAmountNotifRQ"`
}

type Fee struct {
	Amount          string `json:"Amount"`
	Code            string `json:"Code"`
	DecimalPlaces   string `json:"DecimalPlaces"`
	Percent         string `json:"Percent"`
	ChargeFrequency string `json:"ChargeFrequency"`
	Description     string `json:"Description"`
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

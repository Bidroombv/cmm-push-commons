package otamodel

type HotelRateAmountNotifRQ struct {
	OTAHotelRateAmountNotifRQ OTAHotelRateAmountNotifRQ `json:"OTA_HotelRateAmountNotifRQ"`
}

type Fees struct {
	Fee []Fee `json:"Fee"`
}

type Fee struct {
	Amount          string `json:"Amount"`
	Code            string `json:"Code"`
	DecimalPlaces   string `json:"DecimalPlaces"`
	Percent         string `json:"Percent"`
	ChargeFrequency string `json:"ChargeFrequency"`
	Description     string `json:"Description"`
}

type GuaranteeAccepted struct {
	PaymentCard string `json:"PaymentCard"`
}

type GuaranteesAccepted struct {
	GuaranteeAccepted GuaranteeAccepted `json:"GuaranteeAccepted"`
}

type GuaranteePolicy struct {
	GuaranteesAccepted GuaranteesAccepted `json:"GuaranteesAccepted"`
}

type GuaranteePolicies struct {
	GuaranteePolicy []GuaranteePolicy `json:"GuaranteePolicy"`
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
	OriginalData       OriginalData       `json:"original_data"`
}

type OriginalData struct {
	OriginalDataForRateAmount OriginalDataForRateAmount `json:"original_data_for_rate_amount"`
}

type OriginalDataForRateAmount struct {
	RateAmountMessages []RateAmountMessage `json:"RateAmountMessage"`
}

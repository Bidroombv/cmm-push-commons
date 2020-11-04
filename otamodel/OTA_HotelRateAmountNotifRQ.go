package otamodel

type HotelRateAmountNotifRQ struct {
	OTAHotelRateAmountNotifRQ OTAHotelRateAmountNotifRQ `json:"OTA_HotelRateAmountNotifRQ"`
}

type UniqueID struct {
	Type     string `json:"Type"`
	Instance string `json:"Instance"`
	ID       string `json:"ID"`
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
	UniqueID           UniqueID           `json:"UniqueID"`
	RateAmountMessages RateAmountMessages `json:"RateAmountMessages"`
	EchoToken          string             `json:"EchoToken"`
	PrimaryLangID      string             `json:"PrimaryLangID"`
	Target             string             `json:"Target"`
	TimeStamp          string             `json:"TimeStamp"`
	Version            string             `json:"Version"`
}

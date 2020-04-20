package ota_model

type HotelRateAmountNotifRQ struct {
	OTAHotelRateAmountNotifRQ OTAHotelRateAmountNotifRQ `json:"OTA_HotelRateAmountNotifRQ"`
	POS                       POS                       `json:"POS"`
}

type UniqueID struct {
	Type     string `json:"Type"`
	Instance string `json:"Instance"`
	ID       string `json:"ID"`
}

type BaseByGuestAmt struct {
	NumberOfGuests    string `json:"NumberOfGuests"`
	AgeQualifyingCode string `json:"AgeQualifyingCode"`
	AmountAfterTax    string `json:"AmountAfterTax"`
	AmountBeforeTax   string `json:"AmountBeforeTax"`
	DecimalPlaces     string `json:"DecimalPlaces"`
	CurrencyCode      string `json:"CurrencyCode"`
}
type BaseByGuestAmts struct {
	BaseByGuestAmt []BaseByGuestAmt `json:"BaseByGuestAmt"`
}
type AdditionalGuestAmount struct {
	Amount            string `json:"Amount"`
	DecimalPlaces     string `json:"DecimalPlaces"`
	TaxInclusive      string `json:"TaxInclusive"`
	AgeQualifyingCode string `json:"AgeQualifyingCode"`
	CurrencyCode      string `json:"CurrencyCode"`
}
type AdditionalGuestAmounts struct {
	AdditionalGuestAmount []AdditionalGuestAmount `json:"AdditionalGuestAmount"`
}
type Fees struct {
	Fee []Fee `json:"Fee"`
}
type Fee struct {
	Amount string `json:"Amount"`
	Code   string `json:"Code"`
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
type RateDescription struct {
	Text string `json:"Text"`
}
type Rate struct {
	BaseByGuestAmts        BaseByGuestAmts        `json:"BaseByGuestAmts"`
	OriginalData           map[string]interface{} `json:"original_data"`
	AdditionalGuestAmounts AdditionalGuestAmounts `json:"AdditionalGuestAmounts"`
	Fees                   Fees                   `json:"Fees"`
	GuaranteePolicies      GuaranteePolicies      `json:"GuaranteePolicies"`
	RateDescription        RateDescription        `json:"RateDescription"`
	NumberOfUnits          string                 `json:"NumberOfUnits"`
	Mon                    string                 `json:"Mon"`
	Tue                    string                 `json:"Tue"`
	Weds                   string                 `json:"Weds"`
	Thur                   string                 `json:"Thur"`
	Fri                    string                 `json:"Fri"`
	Sat                    string                 `json:"Sat"`
	Sun                    string                 `json:"Sun"`
	Start                  string                 `json:"Start"`
	End                    string                 `json:"End"`
}
type Rates struct {
	Rate []Rate `json:"Rate"`
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
	POS                POS                `json:"POS"`
	UniqueID           UniqueID           `json:"UniqueID"`
	RateAmountMessages RateAmountMessages `json:"RateAmountMessages"`
	EchoToken          string             `json:"EchoToken"`
	PrimaryLangID      string             `json:"PrimaryLangID"`
	Target             string             `json:"Target"`
	TimeStamp          string             `json:"TimeStamp"`
	Version            string             `json:"Version"`
}

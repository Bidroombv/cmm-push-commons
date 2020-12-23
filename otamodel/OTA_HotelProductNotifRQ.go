package otamodel

import "encoding/xml"

//OTAHotelProductNotifRQ **
type OTAHotelProductNotifRQ struct {
	XMLName       xml.Name       `xml:"OTA_HotelProductNotifRQ" json:"-,omitempty"`
	HotelProducts *HotelProducts `json:"HotelProducts,omitempty,required"`
	Xmlns         string         `json:"xmlns,omitempty"`
	Xsi           string         `json:"xsi,omitempty"`
	PrimaryLangID string         `json:"PrimaryLangID,omitempty"`
	EchoToken     string         `json:"EchoToken,omitempty"`
	Version       string         `json:"Version,omitempty,required"`
	Target        string         `json:"Target,omitempty"`
}

//Contact **
type Contact struct {
	Address *[]Address `json:"Address,omitempty,required"`
}

//ArrivalDaysOfWeek **
type ArrivalDaysOfWeek struct {
	Mon  string `json:"Mon,omitempty"`
	Tue  string `json:"Tue,omitempty"`
	Weds string `json:"Weds,omitempty"`
	Thur string `json:"Thur,omitempty"`
	Fri  string `json:"Fri,omitempty"`
	Sat  string `json:"Sat,omitempty"`
	Sun  string `json:"Sun,omitempty"`
}

//Deadline **
type Deadline struct {
	Time string `json:"Time,omitempty"`
}

//AmountPercent **
type AmountPercent struct {
	Amount string `json:"Amount,omitempty"`
}

//CancelPolicy **
type CancelPolicy struct {
	CancelPolicyIndicator string          `json:"CancelPolicyIndicator,omitempty"`
	CancelPenalty         []CancelPenalty `json:"CancelPenalty,omitempty"`
}

//Descriptions **
type Descriptions struct {
	Description []Description `json:"Description,omitempty"`
}

//RoomType **
type RoomType struct {
	RoomTypeCode string       `json:"RoomTypeCode,omitempty,required"`
	RoomTypeName string       `json:"RoomTypeName,omitempty,required"`
	MaxOccupancy string       `json:"MaxOccupancy,omitempty,required"`
	Descriptions Descriptions `json:"Descriptions,omitempty"`
}

//RoomTypes **
type RoomTypes struct {
	RoomType *[]RoomType `json:"RoomType,omitempty,required"`
}

//StayDateRange **
type StayDateRange struct {
	Duration string `json:"Duration,omitempty"`
	Start    string `json:"Start,omitempty"`
	End      string `json:"End,omitempty"`
}

//HotelProductRatePlan **
type HotelProductRatePlan struct {
	RatePlanCode  string          `json:"RatePlanCode,omitempty,required"`
	RatePlanName  string          `json:"RatePlanName,omitempty,required"`
	RatePlanType  string          `json:"RatePlanType,omitempty"`
	CurrencyCode  string          `json:"CurrencyCode,omitempty,required"`
	BaseOccupancy string          `json:"BaseOccupancy,omitempty,required"`
	Descriptions  Descriptions    `json:"Descriptions,omitempty,required"`
	StayDateRange []StayDateRange `json:"StayDateRange,omitempty,required"`
}

//HotelProductRatePlans **
type HotelProductRatePlans struct {
	RatePlan *[]HotelProductRatePlan `json:"RatePlan,omitempty,required"`
}

//PrepaymentPolicy **
type PrepaymentPolicy struct {
	EffectiveFrom string `json:"EffectiveFrom,omitempty"`
}

//TPAExtensions **
type TPAExtensions struct {
	PrepaymentPolicy PrepaymentPolicy `json:"PrepaymentPolicy,omitempty"`
}

//GuaranteePayment **
type GuaranteePayment struct {
	TPAExtensions    *TPAExtensions    `json:"TPA_Extensions,omitempty"`
	PolicyCode       string            `json:"PolicyCode,omitempty"`
	AcceptedPayments *AcceptedPayments `json:"AcceptedPayments,omitempty"`
}

//GuaranteePaymentPolicy **
type GuaranteePaymentPolicy struct {
	GuaranteePayment []GuaranteePayment `json:"GuaranteePayment,omitempty"`
}

//HotelProduct **
type HotelProduct struct {
	Contact                *[]Contact              `json:"Contact,omitempty,required"`
	PolicyInfo             *PolicyInfo             `json:"PolicyInfo,omitempty,required"`
	ProductNotifType       string                  `json:"ProductNotifType,omitempty"`
	RoomTypes              RoomTypes               `json:"RoomTypes,omitempty"`
	RatePlans              HotelProductRatePlans   `json:"RatePlans,omitempty"`
	GuaranteePaymentPolicy *GuaranteePaymentPolicy `json:"GuaranteePaymentPolicy,omitempty"`
}

//HotelProducts **
type HotelProducts struct {
	HotelCode    string          `json:"HotelCode,omitempty,required"`
	HotelName    string          `json:"HotelName,omitempty,required"`
	HotelProduct *[]HotelProduct `json:"HotelProduct,omitempty,required"`
}

package otamodel

import "encoding/xml"

type OTAHotelProductNotifRQ struct {
	XMLName       xml.Name       `xml:"OTA_HotelProductNotifRQ" json:"-"`
	Xmlns         string         `xml:"xmlns,attr" json:"xmlns"`
	Xsi           string         `xml:"xsi,attr" json:"xsi"`
	PrimaryLangID string         `xml:"PrimaryLangID,attr" json:"PrimaryLangID"`
	EchoToken     string         `xml:"EchoToken,attr" json:"EchoToken"`
	Version       string         `xml:"Version,attr,omitempty"  valid:"required" json:"Version,omitempty"`
	Target        string         `xml:"Target,attr" json:"Target"`
	HotelProducts *HotelProducts `xml:"HotelProducts,omitempty" valid:"required" json:"HotelProducts,omitempty"`
}

type HotelProducts struct {
	HotelCode    string          `xml:"HotelCode,attr,omitempty" valid:"required" json:"HotelCode,omitempty"`
	HotelName    string          `xml:"HotelName,attr,omitempty" valid:"required" json:"HotelName,omitempty"`
	HotelProduct *[]HotelProduct `xml:"HotelProduct,omitempty" valid:"required" json:"HotelProduct,omitempty"`
}

type HotelProduct struct {
	ProductNotifType string            `xml:"ProductNotifType,attr,omitempty" valid:"required" json:"ProductNotifType,omitempty"`
	Contact          *[]Contact        `xml:"Contact,omitempty" valid:"required" json:"Contact,omitempty"`
	RoomTypes        *RoomTypes        `xml:"RoomTypes,omitempty" valid:"required" json:"RoomTypes,omitempty"`
	RatePlans        *ProductRatePlans `xml:"RatePlans,omitempty" valid:"required" json:"RatePlans,omitempty"`
	PolicyInfo       *PolicyInfo       `xml:"PolicyInfo,omitempty" valid:"required" json:"PolicyInfo,omitempty"`
	Taxes            *Taxes            `xml:"Taxes,omitempty" valid:"required" json:"Taxes,omitempty"`
	Descriptions     *Descriptions     `xml:"Descriptions,omitempty" valid:"required" json:"Descriptions,omitempty"`
}

type PolicyInfo struct {
	BookingRules *ProductBookingRules `xml:"BookingRules,omitempty" valid:"required" json:"BookingRules,omitempty"`
	CancelPolicy *CancelPolicy        `xml:"CancelPolicy,omitempty" valid:"required" json:"CancelPolicy,omitempty"`
}

type ProductBookingRules struct {
	BookingRule *[]ProductBookingRule `xml:"BookingRule,omitempty" valid:"required" json:"BookingRule,omitempty"`
}

type ProductBookingRule struct {
	MinAdvancedBookingOffset string `xml:"MinAdvancedBookingOffset,attr,omitempty" valid:"required" json:"MinAdvancedBookingOffset,omitempty"`
}

type CancelPolicy struct {
	CancelPolicyIndicator string                `xml:"CancelPolicyIndicator,attr,omitempty" valid:"required" json:"CancelPolicyIndicator,omitempty"`
	CancelPenalty         *ProductCancelPenalty `xml:"CancelPenalty,omitempty" valid:"required" json:"CancelPenalty,omitempty"`
}

type ProductCancelPenalty struct {
	PolicyCode    string         `xml:"PolicyCode,attr,omitempty" valid:"required" json:"PolicyCode,omitempty"`
	AmountPercent *AmountPercent `xml:"AmountPercent,omitempty" valid:"required" json:"AmountPercent,omitempty"`
	Deadline      *Deadline      `xml:"Deadline,omitempty" valid:"required" json:"Deadline,omitempty"`
}

type Deadline struct {
	AbsoluteDeadline     string `xml:"AbsoluteDeadline,attr" json:"AbsoluteDeadline"`
	OffsetDropTime       string `xml:"OffsetDropTime,attr" json:"OffsetDropTime"`
	OffsetTimeUnit       string `xml:"OffsetTimeUnit,attr" json:"OffsetTimeUnit"`
	OffsetUnitMultiplier string `xml:"OffsetUnitMultiplier,attr" json:"OffsetUnitMultiplier"`
}

type AmountPercent struct {
	TaxInclusive  string `xml:"TaxInclusive,attr" json:"TaxInclusive"`
	Percent       string `xml:"Percent,attr" json:"Percent"`
	NmbrOfNights  string `xml:"NmbrOfNights,attr" json:"NmbrOfNights"`
	FeesInclusive string `xml:"FeesInclusive,attr" json:"FeesInclusive"`
	Amount        string `xml:"Amount,attr" json:"Amount"`
	CurrencyCode  string `xml:"CurrencyCode,attr" json:"CurrencyCode"`
	DecimalPlaces string `xml:"DecimalPlaces,attr" json:"DecimalPlaces"`
}

type ProductRatePlans struct {
	RatePlan *[]ProductRatePlan `xml:"RatePlan,omitempty" valid:"required" json:"RatePlan,omitempty"`
}

type ProductRatePlan struct {
	RatePlanCode  string          `xml:"RatePlanCode,attr,omitempty" valid:"required" json:"RatePlanCode,omitempty"`
	RatePlanName  string          `xml:"RatePlanName,attr,omitempty" valid:"required" json:"RatePlanName,omitempty"`
	RatePlanType  string          `xml:"RatePlanType,attr,omitempty" valid:"required" json:"RatePlanType,omitempty"`
	CurrencyCode  string          `xml:"CurrencyCode,attr,omitempty" valid:"required" json:"CurrencyCode,omitempty"`
	BaseOccupancy string          `xml:"BaseOccupancy,attr,omitempty" valid:"required" json:"BaseOccupancy,omitempty"`
	Descriptions  Descriptions    `xml:"Descriptions,omitempty" valid:"required" json:"Descriptions,omitempty"`
	StayDateRange []StayDateRange `xml:"StayDateRange,omitempty" valid:"required" json:"StayDateRange,omitempty"`
}

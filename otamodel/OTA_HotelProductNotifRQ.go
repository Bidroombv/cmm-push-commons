package otamodel

import "encoding/xml"

type OTAHotelProductNotifRQ struct {
	XMLName       xml.Name       `xml:"OTA_HotelProductNotifRQ"`
	Text          string         `xml:",chardata"`
	Xmlns         string         `xml:"xmlns,attr"`
	Xsi           string         `xml:"xsi,attr"`
	PrimaryLangID string         `xml:"PrimaryLangID,attr"`
	EchoToken     string         `xml:"EchoToken,attr"`
	Version       string         `xml:"Version,attr,omitempty"  valid:"required"`
	Target        string         `xml:"Target,attr"`
	HotelProducts *HotelProducts `xml:"HotelProducts,omitempty" valid:"required"`
}

type HotelProducts struct {
	HotelCode    string       `xml:"HotelCode,attr,omitempty" valid:"required"`
	HotelName    string       `xml:"HotelName,attr,omitempty" valid:"required"`
	HotelProduct HotelProduct `xml:"HotelProduct,omitempty" valid:"required"`
}

type HotelProduct struct {
	ProductNotifType string       `xml:"ProductNotifType,attr,omitempty" valid:"required"`
	Contact          *[]Contact   `xml:"Contact,omitempty" valid:"required"`
	RoomTypes        *RoomTypes   `xml:"RoomTypes,omitempty" valid:"required"`
	RatePlans        *RatePlans   `xml:"RatePlans,omitempty" valid:"required"`
	PolicyInfo       *PolicyInfo  `xml:"PolicyInfo,omitempty" valid:"required"`
	Taxes            *Taxes       `xml:"Taxes,omitempty" valid:"required"`
	Descriptons      *Descriptons `xml:"Descriptons,omitempty" valid:"required"`
}

type PolicyInfo struct {
	BookingRules *BookingRules `xml:"BookingRules,omitempty" valid:"required"`
	CancelPolicy *CancelPolicy `xml:"CancelPolicy,omitempty" valid:"required"`
}

type BookingRules struct {
	BookingRule *[]BookingRule `xml:"BookingRule,omitempty" valid:"required"`
}

type BookingRule struct {
	MinAdvancedBookingOffset string `xml:"MinAdvancedBookingOffset,attr,omitempty" valid:"required"`
}

type CancelPolicy struct {
	CancelPolicyIndicator string         `xml:"CancelPolicyIndicator,attr,omitempty" valid:"required"`
	CancelPenalty         *CancelPenalty `xml:"CancelPenalty,omitempty" valid:"required"`
}

type CancelPenalty struct {
	PolicyCode    string         `xml:"PolicyCode,attr,omitempty" valid:"required"`
	AmountPercent *AmountPercent `xml:"AmountPercent,omitempty" valid:"required"`
	Deadline      *Deadline      `xml:"Deadline,omitempty" valid:"required"`
}

type Deadline struct {
	AbsoluteDeadline     string `xml:"AbsoluteDeadline,attr"`
	OffsetDropTime       string `xml:"OffsetDropTime,attr"`
	OffsetTimeUnit       string `xml:"OffsetTimeUnit,attr"`
	OffsetUnitMultiplier string `xml:"OffsetUnitMultiplier,attr"`
}

type AmountPercent struct {
	TaxInclusive  string `xml:"TaxInclusive,attr"`
	Percent       string `xml:"Percent,attr"`
	NmbrOfNights  string `xml:"NmbrOfNights,attr"`
	FeesInclusive string `xml:"FeesInclusive,attr"`
	Amount        string `xml:"Amount,attr"`
	CurrencyCode  string `xml:"CurrencyCode,attr"`
	DecimalPlaces string `xml:"DecimalPlaces,attr"`
}

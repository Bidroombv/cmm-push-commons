package otamodel

//HotelProductNotifRQ **
type HotelProductNotifRQ struct {
	OTAHotelProductNotifRQ OTAHotelProductNotifRQ `json:"OTA_HotelProductNotifRQ"`
}

//OTAHotelProductNotifRQ **
type OTAHotelProductNotifRQ struct {
	HotelProducts *HotelProducts `json:"HotelProducts,omitempty"`
}

//HotelProducts **
type HotelProducts struct {
	HotelCode     string                     `json:"HotelCode,omitempty"`
	HotelName     string                     `json:"HotelName,omitempty"`
	ChainCode     string                     `json:"ChainCode,omitempty"`
	HotelProduct  *[]HotelProduct            `json:"HotelProduct,omitempty"`
	TPAExtensions *HotelProductTPAExtensions `json:"TPA_Extentions,omitempty"`
}

//HotelProduct **
type HotelProduct struct {
	ProductNotifType  string                     `json:"ProductNotifType,omitempty"`
	ProductStatusType string                     `json:"ProductStatusType,omitempty"`
	Contacts          *Contacts                  `json:"Contacts,omitempty"`
	RoomTypes         *RoomTypes                 `json:"RoomTypes,omitempty"`
	RatePlans         *HotelProductRatePlans     `json:"RatePlans,omitempty"`
	PolicyInfo        *HotelProductPolicyInfo    `json:"PolicyInfo,omitempty"`
	Taxes             *Taxes                     `json:"Taxes,omitempty"`
	Descriptions      *Descriptions              `json:"Descriptions,omitempty"`
	TPAExtensions     *HotelProductTPAExtensions `json:"TPA_Extentions,omitempty"`
}

//Contacts **
type Contacts struct {
	Contact *[]Contact `json:"Contact,omitempty"`
}

//Contact **
type Contact struct {
	Address *[]HotelProductAddress `json:"Address,omitempty"`
}

// HotelProductAddress **
type HotelProductAddress struct {
	AddressLine *[]string `json:"AddressLine,omitempty"`
}

//RoomTypes **
type RoomTypes struct {
	RoomType *[]RoomType `json:"RoomType,omitempty"`
}

//RoomType **
type RoomType struct {
	RoomTypeCode      string        `json:"RoomTypeCode,omitempty"`
	RoomTypeName      string        `json:"RoomTypeName,omitempty"`
	MaxOccupancy      string        `json:"MaxOccupancy,omitempty"`
	MaxAdultOccupancy string        `json:"MaxAdultOccupancy,omitempty"`
	MaxChildOccupancy string        `json:"MaxChildOccupancy,omitempty"`
	Descriptions      *Descriptions `json:"Descriptions,omitempty"`
}

//Descriptions **
type Descriptions struct {
	Description *[]Description `json:"Description,omitempty"`
}

//HotelProductRatePlans **
type HotelProductRatePlans struct {
	RatePlan *[]HotelProductRatePlan `json:"RatePlan,omitempty"`
}

//HotelProductRatePlan **
type HotelProductRatePlan struct {
	RatePlanCode  string           `json:"RatePlanCode,omitempty"`
	RatePlanName  string           `json:"RatePlanName,omitempty"`
	RatePlanType  string           `json:"RatePlanType,omitempty"`
	CurrencyCode  string           `json:"CurrencyCode,omitempty"`
	BaseOccupancy string           `json:"BaseOccupancy,omitempty"`
	Descriptions  *Descriptions    `json:"Descriptions,omitempty"`
	SellDateRange *SellDateRange   `json:"SellDateRange,omitempty"`
	StayDateRange *[]StayDateRange `json:"StayDateRange,omitempty"`
}

//SellDateRange **
type SellDateRange struct {
	Start string `json:"Start,omitempty"`
	End   string `json:"End,omitempty"`
}

//HotelProductPolicyInfo **
type HotelProductPolicyInfo struct {
	BookingRules           *BookingRules           `json:"BookingRules,omitempty"`
	CancelPolicy           *CancelPolicy           `json:"CancelPolicy,omitempty"`
	GuaranteePaymentPolicy *GuaranteePaymentPolicy `json:"GuaranteePaymentPolicy,omitempty"`
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
	CancelPolicyIndicator string           `json:"CancelPolicyIndicator,omitempty"`
	CancelPenalty         *[]CancelPenalty `json:"CancelPenalty,omitempty"`
}

//StayDateRange **
type StayDateRange struct {
	Duration string `json:"Duration,omitempty"`
	Start    string `json:"Start,omitempty"`
	End      string `json:"End,omitempty"`
}

//GuaranteePaymentPolicy **
type GuaranteePaymentPolicy struct {
	GuaranteePayment *[]PaymentPoliciesGuaranteePayment `json:"GuaranteePayment,omitempty"`
}

// HotelProductTPAExtensions **
type HotelProductTPAExtensions struct {
	TPAExtension *[]HotelProductTPAExtension `json:"TPA_Extension,omitempty"`
}

//HotelProductTPAExtension **
type HotelProductTPAExtension struct {
	Category     string                   `json:"Category,omitempty"`
	RatePlanCode string                   `json:"RatePlanCode,omitempty"`
	HotelCode    string                   `json:"HotelCode,omitempty"`
	Extension    *[]HotelProductExtension `json:"Extension,omitempty"`
}

// HotelProductExtension **
type HotelProductExtension struct {
	Name string             `json:"Name,omitempty"`
	Item []HotelProductItem `json:"Item,omitempty"`
}

// HotelProductItem **
type HotelProductItem struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
	Text  string `json:"Text,omitempty"`
	Start string `json:"Start,omitempty"`
	End   string `json:"End,omitempty"`
}

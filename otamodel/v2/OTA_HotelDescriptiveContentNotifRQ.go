package v1

//HotelDescriptiveContentNotifRQ **
type HotelDescriptiveContentNotifRQ struct {
	OTAHotelDescriptiveContentNotifRQ OTAHotelDescriptiveContentNotifRQ `json:"OTA_HotelDescriptiveContentNotifRQ"`
}

// OTAHotelDescriptiveContentNotifRQ **
type OTAHotelDescriptiveContentNotifRQ struct {
	PrimaryLangID            string                   `json:"PrimaryLangID,omitempty"`
	EchoToken                string                   `json:"EchoToken,omitempty"`
	TimeStamp                string                   `json:"TimeStamp,omitempty"`
	CorrelationID            string                   `json:"CorrelationID,omitempty"`
	Version                  string                   `json:"Version,omitempty"`
	Target                   string                   `json:"Target,omitempty"`
	HotelDescriptiveContents HotelDescriptiveContents `json:"HotelDescriptiveContents"`
}

//HotelDescriptiveContents **
type HotelDescriptiveContents struct {
	HotelDescriptiveContent []HotelDescriptiveContent `json:"HotelDescriptiveContent"`
}

//HotelDescriptiveContent **
type HotelDescriptiveContent struct {
	LanguageCode           string                  `json:"LanguageCode,omitempty"`
	CurrencyCode           string                  `json:"CurrencyCode,omitempty"`
	ChainCode              string                  `json:"ChainCode,omitempty"`
	ChainName              string                  `json:"ChainName,omitempty"`
	HotelName              string                  `json:"HotelName,omitempty"`
	HotelCode              string                  `json:"HotelCode,omitempty"`
	LastUpdated            string                  `json:"LastUpdated,omitempty"`
	BrandName              string                  `json:"BrandName,omitempty"`
	HotelCityCode          string                  `json:"HotelCityCode,omitempty"`
	TimeZone               string                  `json:"TimeZone,omitempty"`
	HotelInfo              *HotelInfo              `json:"HotelInfo,omitempty"`
	FacilityInfo           *FacilityInfo           `json:"FacilityInfo,omitempty"`
	Policies               *Policies               `json:"Policies,omitempty"`
	MultimediaDescriptions *MultimediaDescriptions `json:"MultimediaDescriptions,omitempty"`
	AffiliationInfo        *AffiliationInfo        `json:"AffiliationInfo,omitempty"`
	ContactInfos           *ContactInfos           `json:"ContactInfos,omitempty"`
	TPAExtensions          *TPAExtensions          `json:"TPA_Extensions,omitempty"`
	AreaInfo               *AreaInfo               `json:"AreaInfo,omitempty"`
}
type TPAExtensions struct {
	TPAExtension     *[]TPAExtension   `json:"TPA_Extension,omitempty"`
	PrepaymentPolicy *PrepaymentPolicy `json:"PrepaymentPolicy,omitempty"`
}
type PrepaymentPolicy struct {
	EffectiveFrom string `json:"EffectiveFrom,omitempty"`
}

type TPAExtension struct {
	Category     string       `json:"Category,omitempty"`
	Code         string       `json:"Code,omitempty"`
	RatePlanCode string       `json:"RatePlanCode,omitempty"`
	HotelCode    string       `json:"HotelCode,omitempty"`
	Extension    *[]Extension `json:"Extension,omitempty"`
}

type Extension struct {
	Name string  `json:"Name,omitempty"`
	Item *[]Item `json:"Item,omitempty"`
}

type Item struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
	Text  string `json:"Text,omitempty"`
	Start string `json:"Start,omitempty"`
	End   string `json:"End,omitempty"`
}

//AreaInfo **
type AreaInfo struct {
	Attractions *Attractions `json:"Attractions,omitempty"`
}

//Attractions **
type Attractions struct {
	Attraction []Attraction `json:"Attraction,omitempty"`
}

//Attraction **
type Attraction struct {
	Items Items `json:"items,omitempty"`
}

//Items **
type Items struct {
	AttractionName         string `json:"AttractionName"`
	AttractionCategoryCode string `json:"AttractionCategoryCode"`
	Distance               string `json:"Distance"`
	DistanceUnit           string `json:"DistanceUnit"`
	LanguageCode           string `json:"LanguageCode"`
}

// ContactInfos **
type ContactInfos struct {
	ContactInfo []ContactInfo `json:"ContactInfo"`
}

// ContactInfo **
type ContactInfo struct {
	ContactProfileID   string     `json:"ContactProfileID,omitempty"`
	ContactProfileType string     `json:"ContactProfileType,omitempty"`
	Location           string     `json:"Location,omitempty"`
	Addresses          *Addresses `json:"Addresses,omitempty"`
	Emails             *Emails    `json:"Emails,omitempty"`
	Names              *Names     `json:"Names,omitempty"`
	Phones             *Phones    `json:"Phones,omitempty"`
	URLs               *URLs      `json:"URLs,omitempty"`
}

// Addresses **
type Addresses struct {
	Address []Address `json:"Address,omitempty"`
}

// Address **
type Address struct {
	Language       string     `json:"Language,omitempty"`
	AddressLine    string     `json:"AddressLine,omitempty"`
	CityName       string     `json:"CityName,omitempty"`
	PostalCode     string     `json:"PostalCode,omitempty"`
	CountryName    string     `json:"CountryName,omitempty"`
	CountryCode    string     `json:"CountryCode,omitempty"`
	HotelName      string     `json:"HotelName,omitempty"`
	AddressUseType string     `json:"AddressUseType,omitempty"`
	Country        string     `json:"Country,omitempty"`
	StateProv      *StateProv `json:"StateProv,omitempty"`
}

// StateProv **
type StateProv struct {
	StateCode string `json:"StateCode,omitempty"`
}

// Emails **
type Emails struct {
	Email []Email `json:"Email,omitempty"`
}

type Email struct {
	EmailType string `json:"EmailType,omitempty"`
	Text      string `json:"Text,omitempty"`
}

// Names **
type Names struct {
	Name []Name `json:"Name"`
}

// Name **
type Name struct {
	Gender    string `json:"Gender,omitempty"`
	Language  string `json:"Language,omitempty"`
	GivenName string `json:"GivenName"`
	Surname   string `json:"Surname,omitempty"`
	JobTitle  string `json:"JobTitle,omitempty"`
}

// Phones **
type Phones struct {
	Phone []Phone `json:"Phone,omitempty"`
}

// Phone **
type Phone struct {
	Extension         string `json:"Extension,omitempty"`
	PhoneNumber       string `json:"PhoneNumber"`
	PhoneTechType     string `json:"PhoneTechType,omitempty"`
	AreaCityCode      string `json:"AreaCityCode,omitempty"`
	CountryAccessCode string `json:"CountryAccessCode,omitempty"`
	PIN               string `json:"PIN,omitempty"`
	PhoneLocationType string `json:"PhoneLocationType,omitempty"`
}

// URLs **
type URLs struct {
	URL []string `json:"URL,omitempty"`
}

// AffiliationInfo **
type AffiliationInfo struct {
	LoyalPrograms LoyalPrograms `json:"LoyalPrograms"`
	Awards        Awards        `json:"Awards,omitempty"`
}
type LoyalPrograms struct {
	LoyalProgram []LoyalProgram `json:"LoyalProgram,omitempty"`
}
type LoyalProgram struct {
	ProgramName        string        `json:"ProgramName,omitempty"`
	TravelSector       string        `json:"TravelSector,omitempty"`
	HotelLevel         string        `json:"HotelLevel,omitempty"`
	ProgramDescription []Description `json:"ProgramDescription"`
}

//Awards **
type Awards struct {
	Award []Award `json:"Award,omitempty"`
}

// Award **
type Award struct {
	Date                   string `json:"Date,omitempty"`
	Provider               string `json:"Provider,omitempty"`
	Rating                 string `json:"Rating,omitempty"`
	RatingSymbol           string `json:"RatingSymbol,omitempty"`
	OfficialAppointmentInd string `json:"OfficialAppointmentInd,omitempty"`
}

// Policies **
type Policies struct {
	Policy []Policy `json:"Policy"`
}

// Policy **
type Policy struct {
	Mon                    bool                                `json:"Mon,omitempty"`
	Tue                    bool                                `json:"Tue,omitempty"`
	Wed                    bool                                `json:"Wed,omitempty"`
	Thur                   bool                                `json:"Thur,omitempty"`
	Fri                    bool                                `json:"Fri,omitempty"`
	Sat                    bool                                `json:"Sat,omitempty"`
	Sun                    bool                                `json:"Sun,omitempty"`
	Start                  string                              `json:"Start,omitempty"`
	End                    string                              `json:"End,omitempty"`
	Code                   string                              `json:"Code,omitempty"`
	CancelPolicy           *CancelPolicy                       `json:"CancelPolicy,omitempty"`
	FeePolicies            *FeePolicies                        `json:"FeePolicies,omitempty"`
	PetsPolicies           *PetsPolicies                       `json:"PetsPolicies,omitempty"`
	GuaranteePaymentPolicy *GuaranteePaymentPolicyHotelContent `json:"GuaranteePaymentPolicy,omitempty"`
	PolicyInfo             *PolicyInfo                         `json:"PolicyInfo,omitempty"`
	TaxPolicies            *TaxPolicies                        `json:"TaxPolicies,omitempty"`
	CheckoutCharges        *CheckoutCharges                    `json:"CheckoutCharges,omitempty"`
}

//CancelPolicy **
type CancelPolicy struct {
	CancelPolicyIndicator string           `json:"CancelPolicyIndicator,omitempty"`
	CancelPenalty         *[]CancelPenalty `json:"CancelPenalty,omitempty"`
}
type CancelPenalty struct {
	RoomTypeCode       string                `json:"RoomTypeCode,omitempty"`
	PolicyCode         string                `json:"PolicyCode,omitempty"`
	NonRefundable      bool                  `json:"NonRefundable,omitempty"`
	NoCancelInd        bool                  `json:"NoCancelInd,omitempty"`
	Start              string                `json:"Start,omitempty"`
	End                string                `json:"End,omitempty"`
	Mon                string                `json:"Mon,omitempty"`
	Tue                string                `json:"Tue,omitempty"`
	Weds               string                `json:"Weds,omitempty"`
	Thur               string                `json:"Thur,omitempty"`
	Fri                string                `json:"Fri,omitempty"`
	Sat                string                `json:"Sat,omitempty"`
	Sun                string                `json:"Sun,omitempty"`
	PenaltyDescription *[]PenaltyDescription `json:"PenaltyDescription,omitempty"`
	Deadline           *PenaltyDeadline      `json:"Deadline,omitempty"`
	AmountPercent      *PenaltyAmountPercent `json:"AmountPercent,omitempty"`
}
type PenaltyAmountPercent struct {
	Percent          string            `json:"Percent,omitempty"`
	NmbrOfNights     string            `json:"NmbrOfNights,omitempty"`
	Amount           string            `json:"Amount,omitempty"`
	CurrencyCode     string            `json:"CurrencyCode,omitempty"`
	TaxInclusive     bool              `json:"TaxInclusive,omitempty"`
	FeesInclusive    bool              `json:"FeesInclusive,omitempty"`
	BasisType        string            `json:"BasisType,omitempty"`
	DecimalPlaces    string            `json:"DecimalPlaces,omitempty"`
	AcceptedPayments *AcceptedPayments `json:"AcceptedPayments,omitempty"`
}

//AcceptedPayments **
type AcceptedPayments struct {
	AcceptedPayment []AcceptedPayment `json:"AcceptedPayment"`
}

//AcceptedPayment **
type AcceptedPayment struct {
	PaymentCard PaymentCard `json:"PaymentCard"`
}

//PaymentCard **
type PaymentCard struct {
	CardCode string `json:"CardCode,omitempty"`
}

type PenaltyDeadline struct {
	OffsetDropTime       string `json:"OffsetDropTime,omitempty"`
	OffsetTimeUnit       string `json:"OffsetTimeUnit,omitempty"`
	OffsetUnitMultiplier string `json:"OffsetUnitMultiplier,omitempty"`
	AbsoluteDeadline     string `json:"AbsoluteDeadline,omitempty"`
}

type PenaltyDescription struct {
	Text string `json:"Text,omitempty"`
}

type CheckoutCharges struct {
	CheckoutCharge []CheckoutCharge `json:"CheckoutCharges,omitempty"`
}
type CheckoutCharge struct {
	Amount        string `json:"Amount,omitempty"`
	Percent       string `json:"Percent,omitempty"`
	CurrencyCode  string `json:"CurrencyCode,omitempty"`
	DecimalPlaces string `json:"DecimalPlaces,omitempty"`
	NmbrOfNights  string `json:"NmbrOfNights,omitempty"`
}

//GuaranteePaymentPolicyHotelContent **
type GuaranteePaymentPolicyHotelContent struct {
	GuaranteePayment []GuaranteePaymentHotelContent `json:"GuaranteePayment,omitempty"`
}

//GuaranteePaymentHotelContent **
type GuaranteePaymentHotelContent struct {
	Mon              bool              `json:"Mon,omitempty"`
	Tue              bool              `json:"Tue,omitempty"`
	Wed              bool              `json:"Wed,omitempty"`
	Thur             bool              `json:"Thur,omitempty"`
	Fri              bool              `json:"Fri,omitempty"`
	Sat              bool              `json:"Sat,omitempty"`
	Sun              bool              `json:"Sun,omitempty"`
	Start            string            `json:"Start,omitempty"`
	End              string            `json:"End,omitempty"`
	GuaranteeType    string            `json:"GuaranteeType,omitempty"`
	HoldTime         string            `json:"HoldTime,omitempty"`
	AmountPercent    *AmountPercent    `json:"AmountPercent,omitempty"`
	Deadline         *Deadline         `json:"Deadline,omitempty"`
	AcceptedPayments *AcceptedPayments `json:"AcceptedPayments,omitempty"`
}

//Deadline **
type Deadline struct {
	Time                 string `json:"Time,omitempty"`
	OffsetUnitMultiplier string `json:"OffsetUnitMultiplier,omitempty"`
	OffsetTimeUnits      string `json:"OffsetTimeUnits,omitempty"`
	OffsetDropTime       string `json:"OffsetDropTime,omitempty"`
}

//AmountPercent **
type AmountPercent struct {
	Amount        string `json:"Amount,omitempty"`
	Percent       string `json:"Percent,omitempty"`
	BasisType     string `json:"BasisType,omitempty"`
	NmbrOfNights  string `json:"NmbrOfNights,omitempty"`
	DecimalPlaces string `json:"DecimalPlaces,omitempty"`
}

// PetsPolicies **
type PetsPolicies struct {
	PetsAllowedCode string       `json:"PetsAllowedCode"`
	PetsPolicy      []PetsPolicy `json:"PetsPolicy,omitempty"`
}

// PetsPolicy **
type PetsPolicy struct {
	PetsPolicyCode        string         `json:"PetsPolicyCode,omitempty"`
	MaxPetQuantity        string         `json:"MaxPetQuantity,omitempty"`
	UnitOfMeasureCode     string         `json:"UnitOfMeasureCode,omitempty"`
	NonRefundableFee      string         `json:"NonRefundableFee,omitempty"`
	UnitOfMeasureQuantity string         `json:"UnitOfMeasureQuantity,omitempty"`
	RefundableDeposit     string         `json:"RefundableDeposit,omitempty"`
	CurrencyCode          string         `json:"CurrencyCode,omitempty"`
	DecimalPlaces         string         `json:"DecimalPlaces,omitempty"`
	ChargeCode            string         `json:"ChargeCode,omitempty"`
	Description           *[]Description `json:"Description,omitempty"`
}

// FeePolicies **
type FeePolicies struct {
	FeePolicy []FeePolicy `json:"FeePolicy"`
}

// FeePolicy **
type FeePolicy struct {
	Code            string `json:"Code"`
	Amount          string `json:"Amount,omitempty"`
	Percent         string `json:"Percent,omitempty"`
	Type            string `json:"Type,omitempty"`
	CurrencyCode    string `json:"CurrencyCode,omitempty"`
	DecimalPlaces   string `json:"DecimalPlaces,omitempty"`
	ChargeFrequency string `json:"ChargeFrequency,omitempty"`
	MandatoryInd    string `json:"MandatoryInd,omitempty"`
	ChargeUnit      string `json:"ChargeUnit,omitempty"`
	EffectiveDate   string `json:"EffectiveDate,omitempty"`
	ExpireDate      string `json:"ExpireDate,omitempty"`
	SequenceNbr     string `json:"SequenceNbr,omitempty"`
	Taxes           *Taxes `json:"Taxes,omitempty"`
}

// Taxes **
type Taxes struct {
	Tax []Tax `json:"Tax"`
}

// Tax **
type Tax struct {
	Code                      string `json:"Code"`
	Amount                    string `json:"Amount,omitempty"`
	Percent                   string `json:"Percent,omitempty"`
	Type                      string `json:"Type,omitempty"`
	CurrencyCode              string `json:"CurrencyCode,omitempty"`
	DecimalPlaces             string `json:"DecimalPlaces,omitempty"`
	ChargeFrequency           string `json:"ChargeFrequency,omitempty"`
	ChargeFrequencyExempt     string `json:"ChargeFrequencyExempt,omitempty"`
	ChargeUnit                string `json:"ChargeUnit,omitempty"`
	ChargeUnitExempt          string `json:"ChargeUnitExempt,omitempty"`
	MaxChargeFrequencyApplies string `json:"MaxChargeFrequencyApplies,omitempty"`
	MaxChargeUnitApplies      string `json:"MaxChargeUnitApplies,omitempty"`
	EffectiveDate             string `json:"EffectiveDate,omitempty"`
	ExpireDate                string `json:"ExpireDate,omitempty"`
	SequenceNbr               string `json:"SequenceNbr,omitempty"`
}

// PolicyInfo **
type PolicyInfo struct {
	Description                string `json:"Description,omitempty"`
	MaxChildAge                string `json:"MaxChildAge,omitempty"`
	MinGuestAge                string `json:"MinGuestAge,omitempty"`
	AcceptedGuestType          string `json:"AcceptedGuestType,omitempty"`
	KidsStayFree               string `json:"KidsStayFree,omitempty"`
	CheckInTime                string `json:"CheckInTime,omitempty"`
	CheckOutTime               string `json:"CheckOutTime,omitempty"`
	TotalGuestCount            string `json:"TotalGuestCount,omitempty"`
	RoomGuaranteeLateArrvInd   string `json:"RoomGuaranteeLateArrvInd,omitempty"`
	UsualStayFreeChildPerAdult string `json:"UsualStayFreeChildPerAdult,omitempty"`
	UsualStayFreeCutoffAge     string `json:"UsualStayFreeCutoffAge,omitempty"`
}

// TaxPolicies **
type TaxPolicies struct {
	TaxPolicy []Tax `json:"TaxPolicy"`
}

// FacilityInfo **
type FacilityInfo struct {
	MeetingRooms *MeetingRooms `json:"MeetingRooms,omitempty"`
	GuestRooms   GuestRooms    `json:"GuestRooms"`
}

//MeetingRooms **
type MeetingRooms struct {
	MeetingRoomCount       string `json:"MeetingRoomCount,omitempty"`
	SmallestRoomSpace      string `json:"SmallestRoomSpace,omitempty"`
	LargestRoomSpace       string `json:"LargestRoomSpace,omitempty"`
	UnitOfMeasure          string `json:"UnitOfMeasure,omitempty"`
	UnitOfMeasureCode      string `json:"UnitOfMeasureCode,omitempty"`
	TotalRoomSpace         string `json:"TotalRoomSpace,omitempty"`
	LargestSeatingCapacity string `json:"LargestSeatingCapacity,omitempty"`
}

// GuestRooms **
type GuestRooms struct {
	GuestRoom    []GuestRoom `json:"GuestRoom"`
	MaxOccupancy string      `json:"MaxOccupancy,omitempty"`
}

// GuestRoom **
type GuestRoom struct {
	Language               string                  `json:"Language,omitempty"`
	MaxAdultOccupancy      string                  `json:"MaxAdultOccupancy,omitempty"`
	MaxChildOccupancy      string                  `json:"MaxChildOccupancy,omitempty"`
	MaxOccupancy           string                  `json:"MaxOccupancy,omitempty"`
	MinOccupancy           string                  `json:"MinOccupancy,omitempty"`
	NonsmokingQuantity     string                  `json:"NonsmokingQuantity,omitempty"`
	Quality                string                  `json:"Quality,omitempty"`
	RoomTypeName           string                  `json:"RoomTypeName,omitempty"`
	Code                   string                  `json:"Code"`
	Quantity               string                  `json:"Quantity,omitempty"`
	DescriptiveText        string                  `json:"DescriptiveText,omitempty"`
	Amenities              *Amenities              `json:"Amenities,omitempty"`
	Features               *Features               `json:"Features,omitempty"`
	MultimediaDescriptions *MultimediaDescriptions `json:"MultimediaDescriptions,omitempty"`
	TypeRoom               *[]TypeRoom             `json:"TypeRoom,omitempty"`
}

// Amenities **
type Amenities struct {
	Amenity []Amenity `json:"Amenity"`
}

// Amenity **
type Amenity struct {
	ExistsCode              string `json:"ExistsCode"`
	RoomAmenityCode         string `json:"RoomAmenityCode"`
	IncludedInRateIndicator string `json:"IncludedInRateIndicator,omitempty"`
	CodeDetail              string `json:"CodeDetail,omitempty"`
	DescriptiveText         string `json:"DescriptiveText,omitempty"`
	Quantity                string `json:"Quantity,omitempty"`
}

// TypeRoom **
type TypeRoom struct {
	RoomID            string `json:"RoomID"`
	RoomCategory      string `json:"RoomCategory,omitempty"`
	Quantity          string `json:"Quantity,omitempty"`
	NonSmoking        string `json:"NonSmoking,omitempty"`
	BedTypeCode       string `json:"BedTypeCode"`
	Count             string `json:"Count,omitempty"`
	Name              string `json:"Name,omitempty"`
	Size              string `json:"Size,omitempty"`
	StandardNumBeds   string `json:"StandardNumBeds,omitempty"`
	StandardOccupancy string `json:"StandardOccupancy,omitempty"`
	RoomLocationCode  string `json:"RoomLocationCode,omitempty"`
	RoomViewCode      string `json:"RoomViewCode,omitempty"`
}

//HotelInfo **
type HotelInfo struct {
	TaxID           string             `json:"TaxID,omitempty"`
	WhenBuilt       string             `json:"WhenBuilt,omitempty"`
	HotelStatus     string             `json:"HotelStatus,omitempty"`
	HotelStatusCode string             `json:"HotelStatusCode,omitempty"`
	ClosedSeasons   *ClosedSeasons     `json:"ClosedSeasons,omitempty"`
	CategoryCodes   *CategoryCodes     `json:"CategoryCodes,omitempty"`
	Descriptions    *HotelDescriptions `json:"Descriptions,omitempty"`
	HotelName       *HotelName         `json:"HotelName,omitempty"`
	Languages       *Languages         `json:"Languages,omitempty"`
	Position        *Position          `json:"Position,omitempty"`
	Services        *Services          `json:"Services,omitempty"`
}

type ClosedSeasons struct {
	ClosedSeason []ClosedSeason `json:"ClosedSeason"`
}
type ClosedSeason struct {
	Start string `json:"Start"`
	End   string `json:"End"`
}

// HotelDescriptions **
type HotelDescriptions struct {
	Renovation  *[]Renovation      `json:"Renovation,omitempty"`
	Description []HotelDescription `json:"Description,omitempty"`
}
type Renovation struct {
	RenovationCompletionDate string `json:"RenovationCompletionDate"`
	AreaText                 string `json:"AreaText"`
	DescriptiveText          string `json:"DescriptiveText"`
}

// HotelDescription **
type HotelDescription struct {
	MultimediaDescriptions *MultimediaDescriptions `json:"MultimediaDescriptions,omitempty"`
	DescriptiveText        string                  `json:"DescriptiveText,omitempty"`
}

// Services **
type Services struct {
	Service []Service `json:"Service"`
}

// Service **
type Service struct {
	Code                string              `json:"Code"`
	Quantity            string              `json:"Quantity,omitempty"`
	Included            string              `json:"Included,omitempty"`
	ProximityCode       string              `json:"ProximityCode,omitempty"`
	BusinessServiceCode string              `json:"BusinessServiceCode,omitempty"`
	MeetingRoomCode     string              `json:"MeetingRoomCode,omitempty"`
	ExistsCode          string              `json:"ExistsCode,omitempty"`
	MealPlanCode        string              `json:"MealPlanCode,omitempty"`
	InvCode             string              `json:"InvCode,omitempty"`
	DescriptiveText     string              `json:"DescriptiveText,omitempty"`
	Features            *Features           `json:"Features,omitempty"`
	OperationSchedules  *OperationSchedules `json:"OperationSchedules,omitempty"`
}

type OperationSchedules struct {
	OperationSchedule []OperationSchedule `json:"OperationSchedule,omitempty"`
}
type OperationSchedule struct {
	Charge         []Charge        `json:"Charge,omitempty"`
	OperationTimes *OperationTimes `json:"OperationTimes,omitempty"`
}
type OperationTimes struct {
	OperationTime []OperationTime `json:"Charge,OperationTime"`
}
type OperationTime struct {
	Mon   bool   `json:"Mon,omitempty"`
	Tue   bool   `json:"Tue,omitempty"`
	Wed   bool   `json:"Wed,omitempty"`
	Thur  bool   `json:"Thur,omitempty"`
	Fri   bool   `json:"Fri,omitempty"`
	Sat   bool   `json:"Sat,omitempty"`
	Sun   bool   `json:"Sun,omitempty"`
	Start string `json:"Start,omitempty"`
	End   string `json:"End,omitempty"`
}

// Features **
type Features struct {
	Feature []Feature `json:"Feature"`
}

// Feature **
type Feature struct {
	CodeDetail             string                  `json:"CodeDetail"`
	UnitOfMeasure          string                  `json:"UnitOfMeasure,omitempty"`
	UnitOfMeasureQuantity  string                  `json:"UnitOfMeasureQuantity,omitempty"`
	MultimediaDescriptions *MultimediaDescriptions `json:"MultimediaDescriptions,omitempty"`
	DescriptiveText        string                  `json:"DescriptiveText,omitempty"`
	SecurityCode           string                  `json:"SecurityCode,omitempty"`
	AccessibilityCode      string                  `json:"AccessibilityCode,omitempty"`
	Quantity               string                  `json:"Quantity,omitempty"`
	Charge                 *Charge                 `json:"Charge,omitempty"`
}

// Charge **
type Charge struct {
	Amount        string `json:"Amount"`
	CurrencyCode  string `json:"CurrencyCode,omitempty"`
	DecimalPlaces string `json:"DecimalPlaces,omitempty"`
	Percent       string `json:"Percent,omitempty"`
	ChargeUnit    string `json:"ChargeUnit,omitempty"`
}

// Position **
type Position struct {
	Latitude             string `json:"Latitude"`
	Longitude            string `json:"Longitude"`
	PositionAccuracyCode string `json:"PositionAccuracyCode"`
}

// Languages **
type Languages struct {
	Language []Language `json:"Language"`
}

// Language **
type Language struct {
	Language       string `json:"Language"`
	PrimaryLangInd string `json:"PrimaryLangInd,omitempty"`
}

// HotelName **
type HotelName struct {
	HotelShortName string `json:"HotelShortName"`
}

// CategoryCodes **
type CategoryCodes struct {
	SegmentCategory  *[]SegmentCategory  `json:"SegmentCategory,omitempty"`
	GuestRoomInfo    *[]GuestRoomInfo    `json:"GuestRoomInfo,omitempty"`
	HotelCategory    *[]HotelCategory    `json:"HotelCategory,omitempty"`
	LocationCategory *[]LocationCategory `json:"LocationCategory,omitempty"`
}
type LocationCategory struct {
	Code       string `json:"Code,omitempty"`
	CodeDetail string `json:"CodeDetail,omitempty"`
}

// SegmentCategory **
type SegmentCategory struct {
	Code       string `json:"Code,omitempty"`
	ExistsCode string `json:"ExistsCode,omitempty"`
	CodeDetail string `json:"CodeDetail,omitempty"`
}

// HotelCategory **
type HotelCategory struct {
	Code       string `json:"Code"`
	ExistsCode string `json:"ExistsCode,omitempty"`
	CodeDetail string `json:"CodeDetail,omitempty"`
}

// GuestRoomInfo **
type GuestRoomInfo struct {
	Code                   string                  `json:"Code"`
	ExistsCode             string                  `json:"ExistsCode,omitempty"`
	CodeDetail             string                  `json:"CodeDetail,omitempty"`
	Quantity               string                  `json:"Quantity,omitempty"`
	DescriptiveText        string                  `json:"DescriptiveText,omitempty"`
	MultimediaDescriptions *MultimediaDescriptions `json:"MultimediaDescriptions,omitempty"`
}

// MultimediaDescriptions **
type MultimediaDescriptions struct {
	MultimediaDescription []MultimediaDescription `json:"MultimediaDescription"`
}

// MultimediaDescription **
type MultimediaDescription struct {
	InfoCode             string      `json:"InfoCode,omitempty"`
	AdditionalDetailCode string      `json:"AdditionalDetailCode,omitempty"`
	ImageItems           *ImageItems `json:"ImageItems,omitempty"`
	TextItems            *TextItems  `json:"TextItems,omitempty"`
}

//TextItems **
type TextItems struct {
	TextItem []TextItem `json:"TextItem"`
}

//TextItem **
type TextItem struct {
	Title           string       `json:"Title"`
	ApplicableStart string       `json:"ApplicableStart,omitempty"`
	ApplicableEnd   string       `json:"ApplicableEnd,omitempty"`
	RecordId        string       `json:"RecordId,omitempty"`
	SourceId        string       `json:"SourceId,omitempty"`
	Description     *Description `json:"Description,omitempty"`
}

// ImageItems **
type ImageItems struct {
	ImageItem []ImageItem `json:"ImageItem"`
}

// ImageItem **
type ImageItem struct {
	Category           string         `json:"Category,omitempty"`
	LastModifyDateTime string         `json:"LastModifyDateTime,omitempty"`
	ID                 string         `json:"ID,omitempty"`
	ImageFormat        []ImageFormat  `json:"ImageFormat"`
	TPAExtensions      *TPAExtensions `json:"TPA_Extensions,omitempty"`
	Description        *[]Description `json:"Description,omitempty"`
}

//DescriptionImageItem **
type DescriptionImageItem struct {
	Caption string `json:"Caption"`
}

// ImageFormat **
type ImageFormat struct {
	URL             string `json:"URL"`
	Width           string `json:"Width,omitempty"`
	Height          string `json:"Height,omitempty"`
	Resolution      string `json:"Resolution,omitempty"`
	Format          string `json:"Format,omitempty"`
	FileSize        string `json:"FileSize,omitempty"`
	ApplicableStart string `json:"ApplicableStart,omitempty"`
	ApplicableEnd   string `json:"ApplicableEnd,omitempty"`
	Sort            string `json:"Sort,omitempty"`
	RecordId        string `json:"RecordId,omitempty"`
	SourceId        string `json:"SourceId,omitempty"`
}
type Description struct {
	Name string `json:"Name,omitempty"`
	Text string `json:"Text,omitempty"`
}

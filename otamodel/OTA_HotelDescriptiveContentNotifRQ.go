package otamodel

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
	HotelInfo              *HotelInfo              `json:"HotelInfo,omitempty"`
	FacilityInfo           *FacilityInfo           `json:"FacilityInfo,omitempty"`
	Policies               *Policies               `json:"Policies,omitempty"`
	MultimediaDescriptions *MultimediaDescriptions `json:"MultimediaDescriptions,omitempty"`
	AffiliationInfo        *AffiliationInfo        `json:"AffiliationInfo,omitempty"`
	ContactInfos           *ContactInfos           `json:"ContactInfos,omitempty"`
	TPAExtensions          *TPAExtensions          `json:"TPA_Extensions,omitempty"`
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
	Language    string     `json:"Language,omitempty"`
	AddressLine string     `json:"AddressLine,omitempty"`
	CityName    string     `json:"CityName,omitempty"`
	PostalCode  string     `json:"PostalCode,omitempty"`
	CountryName string     `json:"CountryName,omitempty"`
	CountryCode string     `json:"CountryCode,omitempty"`
	HotelName   string     `json:"HotelName,omitempty"`
	StateProv   *StateProv `json:"StateProv,omitempty"`
}

// StateProv **
type StateProv struct {
	StateCode string `json:"StateCode,omitempty"`
}

// Emails **
type Emails struct {
	Email []string `json:"Email,omitempty"`
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
	PhoneTechType     int    `json:"PhoneTechType,omitempty"`
	AreaCityCode      string `json:"AreaCityCode,omitempty"`
	CountryAccessCode string `json:"CountryAccessCode,omitempty"`
	PIN               string `json:"PIN,omitempty"`
}

// URLs **
type URLs struct {
	URL []string `json:"URL,omitempty"`
}

// AffiliationInfo **
type AffiliationInfo struct {
	Awards []Award `json:"Awards,omitempty"`
}

// Award **
type Award struct {
	Date         string `json:"Date,omitempty"`
	Provider     string `json:"Provider,omitempty"`
	Rating       string `json:"Rating"`
	RatingSymbol string `json:"RatingSymbol,omitempty"`
}

// Policies **
type Policies struct {
	Policy []Policy `json:"Policy"`
}

// Policy **
type Policy struct {
	Code                   string                  `json:"Code"`
	CancelPolicy           *CancelPolicy           `json:"CancelPolicy,omitempty"`
	FeePolicies            *FeePolicies            `json:"FeePolicies,omitempty"`
	PetsPolicies           *PetsPolicies           `json:"PetsPolicies,omitempty"`
	GuaranteePaymentPolicy *GuaranteePaymentPolicy `json:"GuaranteePaymentPolicy,omitempty"`
	PolicyInfo             *PolicyInfo             `json:"PolicyInfo,omitempty"`
	TaxPolicies            *TaxPolicies            `json:"TaxPolicies,omitempty"`
}

// PetsPolicies **
type PetsPolicies struct {
	PetsAllowedCode string        `json:"PetsAllowedCode"`
	PetsPolicy      *[]PetsPolicy `json:"PetsPolicy,omitempty"`
}

// PetsPolicy **
type PetsPolicy struct {
	PetsPolicyCode string `json:"PetsPolicyCode"`
	MaxPetQuantity string `json:"MaxPetQuantity"`
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
	Taxes           *Taxes `json:"Taxes,omitempty"`
}

// Taxes **
type Taxes struct {
	Tax []Tax `json:"Tax"`
}

// Tax **
type Tax struct {
	Code            string `json:"Code"`
	Amount          string `json:"Amount,omitempty"`
	Percent         string `json:"Percent,omitempty"`
	Type            string `json:"Type,omitempty"`
	CurrencyCode    string `json:"CurrencyCode,omitempty"`
	DecimalPlaces   string `json:"DecimalPlaces,omitempty"`
	ChargeFrequency string `json:"ChargeFrequency,omitempty"`
}

// PolicyInfo **
type PolicyInfo struct {
	Description       string `json:"Description,omitempty"`
	MaxChildAge       int    `json:"MaxChildAge,omitempty"`
	MinGuestAge       int    `json:"MinGuestAge,omitempty"`
	AcceptedGuestType string `json:"AcceptedGuestType,omitempty"`
	KidsStayFree      bool   `json:"KidsStayFree,omitempty"`
}

// TaxPolicies **
type TaxPolicies struct {
	TaxPolicy []Tax `json:"TaxPolicy"`
}

// FacilityInfo **
type FacilityInfo struct {
	GuestRooms GuestRooms `json:"GuestRooms"`
}

// GuestRooms **
type GuestRooms struct {
	GuestRoom    []GuestRoom `json:"GuestRoom"`
	MaxOccupancy int         `json:"MaxOccupancy,omitempty"`
}

// GuestRoom **
type GuestRoom struct {
	Language               string                  `json:"Language,omitempty"`
	MaxAdultOccupancy      int                     `json:"MaxAdultOccupancy,omitempty"`
	MaxChildOccupancy      int                     `json:"MaxChildOccupancy,omitempty"`
	MaxOccupancy           int                     `json:"MaxOccupancy,omitempty"`
	MinOccupancy           int                     `json:"MinOccupancy,omitempty"`
	NonsmokingQuantity     int                     `json:"NonsmokingQuantity,omitempty"`
	Quality                string                  `json:"Quality,omitempty"`
	RoomTypeName           string                  `json:"RoomTypeName,omitempty"`
	Code                   string                  `json:"Code"`
	Quantity               int                     `json:"Quantity,omitempty"`
	DescriptiveText        string                  `json:"DescriptiveText,omitempty"`
	Amenities              Amenities               `json:"Amenities"`
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
	RoomAmenityCode         int    `json:"RoomAmenityCode"`
	IncludedInRateIndicator bool   `json:"IncludedInRateIndicator,omitempty"`
	CodeDetail              string `json:"CodeDetail,omitempty"`
	DescriptiveText         string `json:"DescriptiveText,omitempty"`
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
}

//HotelInfo **
type HotelInfo struct {
	HotelStatus     string             `json:"HotelStatus,omitempty"`
	HotelStatusCode string             `json:"HotelStatusCode,omitempty"`
	CategoryCodes   *CategoryCodes     `json:"CategoryCodes,omitempty"`
	Descriptions    *HotelDescriptions `json:"Descriptions,omitempty"`
	HotelName       *HotelName         `json:"HotelName,omitempty"`
	Languages       *Languages         `json:"Languages,omitempty"`
	Position        *Position          `json:"Position,omitempty"`
	Services        *Services          `json:"Services,omitempty"`
}

// HotelDescriptions **
type HotelDescriptions struct {
	Description []HotelDescription `json:"Description,omitempty"`
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
	Code            string    `json:"Code"`
	ExistsCode      string    `json:"ExistsCode,omitempty"`
	MealPlanCode    string    `json:"MealPlanCode,omitempty"`
	InvCode         string    `json:"InvCode,omitempty"`
	DescriptiveText string    `json:"DescriptiveText,omitempty"`
	Features        *Features `json:"Features,omitempty"`
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
	DescriptiveText        string                  `json:"DescriptiveText"`
	Charge                 *Charge                 `json:"Charge,omitempty"`
}

// Charge **
type Charge struct {
	Amount        string `json:"Amount"`
	CurrencyCode  string `json:"CurrencyCode,omitempty"`
	DecimalPlaces string `json:"DecimalPlaces,omitempty"`
}

// Position **
type Position struct {
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
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
	GuestRoomInfo *[]GuestRoomInfo `json:"GuestRoomInfo,omitempty"`
	HotelCategory *[]HotelCategory `json:"HotelCategory,omitempty"`
}

// HotelCategory **
type HotelCategory struct {
	Code       int    `json:"Code"`
	ExistsCode string `json:"ExistsCode,omitempty"`
}

// GuestRoomInfo **
type GuestRoomInfo struct {
	Code                   int                     `json:"Code"`
	ExistsCode             string                  `json:"ExistsCode,omitempty"`
	Quantity               int                     `json:"Quantity,omitempty"`
	DescriptiveText        string                  `json:"DescriptiveText,omitempty"`
	MultimediaDescriptions *MultimediaDescriptions `json:"MultimediaDescriptions,omitempty"`
}

// MultimediaDescriptions **
type MultimediaDescriptions struct {
	MultimediaDescription []MultimediaDescription `json:"MultimediaDescription"`
}

// MultimediaDescription **
type MultimediaDescription struct {
	InfoCode             int         `json:"InfoCode,omitempty"`
	AdditionalDetailCode int         `json:"AdditionalDetailCode,omitempty"`
	ImageItems           *ImageItems `json:"ImageItems,omitempty"`
	TextItems            *TextItems  `json:"TextItems,omitempty"`
}

//TextItems **
type TextItems struct {
	TextItem []TextItem `json:"TextItem"`
}

//TextItem **
type TextItem struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

// ImageItems **
type ImageItems struct {
	ImageItem []ImageItem `json:"ImageItem"`
}

// ImageItem **
type ImageItem struct {
	Category      int                     `json:"Category,omitempty"`
	ImageFormat   []ImageFormat           `json:"ImageFormat"`
	Description   *DescriptionImageItem   `json:"Description,omitempty"`
	TPAExtensions *TPAExtensionsImageItem `json:"TPA_Extensions,omitempty"`
}

//TPAExtensionsImageItem **
type TPAExtensionsImageItem struct {
	Category     string       `json:"Category"`
	TPAExtension TPAExtension `json:"TPAExtension"`
}

//TPAExtension **
type TPAExtension struct {
	Extension ExtensionImageItem `json:"Extension"`
}

// ExtensionImageItem **
type ExtensionImageItem struct {
	Name string `json:"Name,omitempty"`
	Item *Item  `json:"Item,omitempty"`
}

//Item **
type Item struct {
	Key   string `json:"Key"`
	Value bool   `json:"Value"`
}

//DescriptionImageItem **
type DescriptionImageItem struct {
	Caption string `json:"Caption"`
}

// ImageFormat **
type ImageFormat struct {
	URL        string `json:"URL"`
	Width      string `json:"Width,omitempty"`
	Height     string `json:"Height,omitempty"`
	Resolution string `json:"Resolution,omitempty"`
	Format     string `json:"Format,omitempty"`
	FileSize   string `json:"FileSize,omitempty"`
}

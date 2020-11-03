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
	HotelName              string                  `json:"HotelName"`
	HotelCode              string                  `json:"HotelCode"`
	HotelInfo              HotelInfo               `json:"HotelInfo"`
	FacilityInfo           FacilityInfo            `json:"FacilityInfo"`
	Policies               Policies                `json:"Policies"`
	MultimediaDescriptions *MultimediaDescriptions `json:"MultimediaDescriptions,omitempty"`
	AffiliationInfo        *AffiliationInfo        `json:"AffiliationInfo,omitempty"`
	ContactInfos           *ContactInfos           `json:"ContactInfos,omitempty"`
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
	AddressLine string     `json:"AddressLine"`
	CityName    string     `json:"CityName"`
	PostalCode  string     `json:"PostalCode"`
	CountryName string     `json:"CountryName,omitempty"`
	CountryCode string     `json:"CountryCode"`
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
	PhoneTechType     string `json:"PhoneTechType,omitempty"`
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
	Rating       string `json:"Rating,omitempty"`
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
	MaxChildAge       string `json:"MaxChildAge,omitempty"`
	MinGuestAge       string `json:"MinGuestAge,omitempty"`
	AcceptedGuestType string `json:"AcceptedGuestType,omitempty"`
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
	MaxOccupancy string      `json:"MaxOccupancy,omitempty"`
}

// GuestRoom **
type GuestRoom struct {
	Language               string                  `json:"Language,omitempty"`
	MaxAdultOccupancy      int                     `json:"MaxAdultOccupancy"`
	MaxChildOccupancy      int                     `json:"MaxChildOccupancy"`
	MaxOccupancy           int                     `json:"MaxOccupancy"`
	MinOccupancy           int                     `json:"MinOccupancy"`
	NonsmokingQuantity     string                  `json:"NonsmokingQuantity,omitempty"`
	Quality                string                  `json:"Quality,omitempty"`
	RoomTypeName           string                  `json:"RoomTypeName"`
	Code                   string                  `json:"Code"`
	Quantity               string                  `json:"Quantity"`
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
	ExistsCode      string `json:"ExistsCode,omitempty"`
	RoomAmenityCode string `json:"RoomAmenityCode"`
	CodeDetail      string `json:"CodeDetail,omitempty"`
	DescriptiveText string `json:"DescriptiveText,omitempty"`
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
	CategoryCodes   CategoryCodes      `json:"CategoryCodes"`
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
	DescriptiveText string `json:"DescriptiveText,omitempty"`
}

// Services **
type Services struct {
	Service []Service `json:"Service"`
}

// Service **
type Service struct {
	Code            string    `json:"Code"`
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
	DescriptiveText string `json:"DescriptiveText"`
	Charge          Charge `json:"Charge,omitempty"`
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
	GuestRoomInfo []GuestRoomInfo  `json:"GuestRoomInfo"`
	HotelCategory *[]HotelCategory `json:"HotelCategory,omitempty"`
}

// HotelCategory **
type HotelCategory struct {
	Code       string `json:"Code"`
	ExistsCode string `json:"ExistsCode,omitempty"`
}

// GuestRoomInfo **
type GuestRoomInfo struct {
	Code                   string                  `json:"Code"`
	ExistsCode             string                  `json:"ExistsCode,omitempty"`
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
	ImageItems ImageItems `json:"ImageItems"`
}

// ImageItems **
type ImageItems struct {
	ImageItem []ImageItem `json:"ImageItem"`
}

// ImageItem **
type ImageItem struct {
	ImageFormat []ImageFormat `json:"ImageFormat"`
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

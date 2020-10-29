package otamodel

type HotelDescriptiveContentNotifRQ struct {
	OTAHotelDescriptiveContentNotifRQ OTAHotelDescriptiveContentNotifRQ `json:"OTA_HotelDescriptiveContentNotifRQ"`
}

type OTAHotelDescriptiveContentNotifRQ struct {
	PrimaryLangID            string                   `json:"PrimaryLangID,omitempty"`
	EchoToken                string                   `json:"EchoToken,omitempty"`
	TimeStamp                string                   `json:"TimeStamp,omitempty"`
	CorrelationID            string                   `json:"CorrelationID,omitempty"`
	Version                  string                   `json:"Version,omitempty"`
	Target                   string                   `json:"Target,omitempty"`
	HotelDescriptiveContents HotelDescriptiveContents `json:"HotelDescriptiveContents"`
}

type HotelDescriptiveContents struct {
	HotelDescriptiveContent []HotelDescriptiveContent `json:"HotelDescriptiveContent"`
}

type HotelDescriptiveContent struct {
	HotelName              string                  `json:"HotelName"`
	HotelCode              string                  `json:"HotelCode"`
	LanguageCode           string                  `json:"LanguageCode,omitempty"`
	CurrencyCode           string                  `json:"CurrencyCode,omitempty"`
	ContactInfos           ContactInfos            `json:"ContactInfos"`
	HotelInfo              HotelInfo               `json:"HotelInfo"`
	FacilityInfo           FacilityInfo            `json:"FacilityInfo"`
	AreaInfo               *AreaInfo               `json:"AreaInfo,omitempty"`
	Policies               *Policies               `json:"Policies,omitempty"`
	AffiliationInfo        *AffiliationInfo        `json:"AffiliationInfo,omitempty"`
	MultimediaDescriptions *MultimediaDescriptions `json:"MultimediaDescriptions,omitempty"`
}

type ContactInfos struct {
	ContactInfo []ContactInfo `json:"ContactInfo"`
}

type HotelInfo struct {
	CategoryCodes            *CategoryCodes            `json:"CategoryCodes,omitempty"`
	Languages                *Languages                `json:"Languages,omitempty"`
	Position                 Position                  `json:"Position"`
	Services                 *Services                 `json:"Services,omitempty"`
	OwnershipManagementInfos *OwnershipManagementInfos `json:"OwnershipManagementInfos,omitempty"`
	Descriptions             *HotelDescriptions        `json:"Descriptions,omitempty"`
}

type CategoryCodes struct {
	GuestRoomInfo *[]GuestRoomInfo `json:"GuestRoomInfo,omitempty"`
	HotelCategory *[]HotelCategory `json:"HotelCategory,omitempty"`
}

type GuestRoomInfo struct {
	Quantity   string `json:"Quantity,omitempty"`
	Code       string `json:"Code,omitempty"`
	ExistsCode string `json:"ExistsCode,omitempty"`
}

type HotelCategory struct {
	Code       string `json:"Code,omitempty"`
	ExistsCode string `json:"ExistsCode,omitempty"`
}

type Languages struct {
	Language []Language `json:"Language,omitempty"`
}

type Language struct {
	LanguageCode string `json:"LanguageCode,omitempty"`
}

type Position struct {
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

type Services struct {
	Service []Service `json:"Service,omitempty"`
}

type HotelDescriptions struct {
	Description []HotelDescription `json:"Description,omitempty"`
}

type HotelDescription struct {
	DescriptiveText string `json:"DescriptiveText,omitempty"`
}

type Service struct {
	Code           string          `json:"Code,omitempty"`
	Included       string          `json:"Included,omitempty"`
	Price          string          `json:"Price,omitempty"`
	CurrencyCode   string          `json:"CurrencyCode,omitempty"`
	ExistsCode     string          `json:"ExistsCode,omitempty"`
	Quantity       string          `json:"Quantity,omitempty"`
	OperationTimes *OperationTimes `json:"OperationTimes,omitempty"`
	Features       *Features       `json:"Features,omitempty"`
}

type OperationTimes struct {
	OperationTime []DaysOfWeek `json:"OperationTime,omitempty"`
}

type Features struct {
	Feature []Feature `json:"Feature,omitempty"`
}

type Feature struct {
	ID              string `json:"ID,omitempty"`
	DescriptiveText string `json:"DescriptiveText,omitempty"`
	Charge          Charge `json:"Charge,omitempty"`
}

type Charge struct {
	Amount string `json:"Amount,omitempty"`
}

type OwnershipManagementInfos struct {
	OwnershipManagementInfo []OwnershipManagementInfo `json:"OwnershipManagementInfo,omitempty"`
}

type OwnershipManagementInfo struct {
	CompanyName CompanyName `json:"CompanyName,omitempty"`
}

type CompanyName struct {
	Code string `json:"Code,omitempty"`
}

type FacilityInfo struct {
	GuestRooms GuestRooms `json:"GuestRooms,omitempty"`
}

type GuestRooms struct {
	GuestRoom []GuestRoom `json:"GuestRoom,omitempty"`
}

type GuestRoom struct {
	MaxAdultOccupancy      string                  `json:"MaxAdultOccupancy,omitempty"`
	MaxChildOccupancy      string                  `json:"MaxChildOccupancy,omitempty"`
	MaxOccupancy           string                  `json:"MaxOccupancy,omitempty"`
	MinOccupancy           string                  `json:"MinOccupancy,omitempty"`
	RoomTypeName           string                  `json:"RoomTypeName"`
	TypeRoom               *TypeRoom               `json:"TypeRoom,omitempty"`
	Amenities              *Amenities              `json:"Amenities,omitempty"`
	Code                   string                  `json:"Code"`
	DescriptiveText        string                  `json:"DescriptiveText,omitempty"`
	MultimediaDescriptions *MultimediaDescriptions `json:"MultimediaDescriptions,omitempty"`
}

type TypeRoom struct {
	BedTypeCode  string `json:"BedTypeCode,omitempty"`
	RoomTypeCode string `json:"RoomTypeCode,omitempty"`
	NonSmoking   string `json:"NonSmoking,omitempty"`
}

type Amenities struct {
	Amenity []Amenity `json:"Amenity,omitempty"`
}

type Amenity struct {
	RoomAmenityCode string `json:"RoomAmenityCode,omitempty"`
	Quantity        string `json:"Quantity,omitempty"`
}

type Restaurants struct {
	Restaurant []Restaurant `json:"Restaurant"`
}

type Restaurant struct {
	RestaurantName     string             `json:"RestaurantName,omitempty"`
	OfferBreakfast     string             `json:"OfferBreakfast,omitempty"`
	OfferDinner        string             `json:"OfferDinner,omitempty"`
	CuisineCodes       CuisineCodes       `json:"CuisineCodes,omitempty"`
	OperationSchedules OperationSchedules `json:"OperationSchedules,omitempty"`
	Features           Features           `json:"Features,omitempty"`
}

type CuisineCodes struct {
	CuisineCode []CuisineCode `json:"CuisineCode,omitempty"`
}

type CuisineCode struct {
	Code string `json:"Code,omitempty"`
}

type OperationSchedules struct {
	OperationSchedule []OperationSchedule `json:"OperationSchedule,omitempty"`
}

type OperationSchedule struct {
	OperationTimes OperationTimes `json:"OperationTimes,omitempty"`
}

type AreaInfo struct {
	Attractions Attractions `json:"Attractions,omitempty"`
}

type Attractions struct {
	Attraction []Attraction `json:"Attraction,omitempty"`
}

type Attraction struct {
	AttractionName         string `json:"AttractionName,omitempty"`
	AttractionCategoryCode string `json:"AttractionCategoryCode,omitempty"`
	Distance               string `json:"Distance,omitempty"`
	DistanceUnit           string `json:"DistanceUnit,omitempty"`
	LanguageCode           string `json:"LanguageCode,omitempty"`
}

type Policies struct {
	Policy []Policy `json:"Policy,omitempty"`
}

type Policy struct {
	GuaranteePaymentPolicy string        `json:"GuaranteePaymentPolicy,omitempty"`
	PolicyInfo             *PolicyInfo   `json:"PolicyInfo,omitempty"`
	PetsPolicies           *PetsPolicies `json:"PetsPolicies,omitempty"`
	TaxPolicies            *TaxPolicies  `json:"TaxPolicies,omitempty"`
	FeePolicies            *FeePolicies  `json:"FeePolicies,omitempty"`
}

type PolicyInfo struct {
	CheckInTime  string `json:"CheckInTime,omitempty"`
	CheckOutTime string `json:"CheckOutTime,omitempty"`
}

type PetsPolicies struct {
	PetsAllowedCode string       `json:"PetsAllowedCode,omitempty"`
	PetsPolicy      []PetsPolicy `json:"PetsPolicy,omitempty"`
}

type PetsPolicy struct {
	NonRefundableFee string `json:"NonRefundableFee,omitempty"`
}

type TaxPolicies struct {
	TaxPolicy []TaxPolicy `json:"TaxPolicy,omitempty"`
}

type TaxPolicy struct {
	Code            string `json:"Code,omitempty"`
	Percent         string `json:"Percent,omitempty"`
	Type            string `json:"Type,omitempty"`
	Amount          string `json:"Amount,omitempty"`
	ChargeFrequency string `json:"ChargeFrequency,omitempty"`
}

type FeePolicies struct {
	FeePolicy FeePolicy `json:"FeePolicy,omitempty"`
}

type FeePolicy struct {
	Code            string `json:"Code,omitempty"`
	Amount          string `json:"Amount,omitempty"`
	Type            string `json:"Type,omitempty"`
	ChargeFrequency string `json:"ChargeFrequency,omitempty"`
}

type AffiliationInfo struct {
	Awards Awards `json:"Awards,omitempty"`
}

type Awards struct {
	Award Award `json:"Award,omitempty"`
}

type Award struct {
	Provider string `json:"Provider,omitempty"`
	Rating   string `json:"Rating,omitempty"`
}

type MultimediaDescriptions struct {
	MultimediaDescription []MultimediaDescription `json:"MultimediaDescription,omitempty"`
}

type MultimediaDescription struct {
	ImageItems ImageItems `json:"ImageItems,omitempty"`
}

type ImageItems struct {
	ImageItem []ImageItem `json:"ImageItem,omitempty"`
}

type ImageItem struct {
	ImageFormat ImageFormat `json:"ImageFormat,omitempty"`
}

type ImageFormat struct {
	Main string `json:"Main,omitempty"`
	Sort string `json:"Sort,omitempty"`
	URL  string `json:"URL"`
}

type ContactInfo struct {
	ContactProfileType string    `json:"ContactProfileType"`
	Addresses          Addresses `json:"Addresses"`
	Names              *Names    `json:"Names,omitempty"`
	Emails             *Emails   `json:"Emails,omitempty"`
	Phones             *Phones   `json:"Phones,omitempty"`
	URLs               []URLs    `json:"URLs,omitempty"`
}

type Addresses struct {
	Address []Address `json:"Address"`
}

type URLs struct {
	URL string `json:"URL"`
}

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

type StateProv struct {
	StateCode string `json:"StateCode,omitempty"`
}

type Names struct {
	Name []Name `json:"Name"`
}

type Name struct {
	Gender    string `json:"Gender,omitempty"`
	Language  string `json:"Language,omitempty"`
	GivenName string `json:"GivenName"`
	Surname   string `json:"Surname,omitempty"`
	JobTitle  string `json:"JobTitle,omitempty"`
}

type Emails struct {
	Email []string `json:"Email,omitempty"`
}

type Phones struct {
	Phone []Phone `json:"Phone,omitempty"`
}

type Phone struct {
	Extension     string `json:"Extension,omitempty"`
	PhoneNumber   string `json:"PhoneNumber"`
	PhoneTechType string `json:"PhoneTechType,omitempty"`
}

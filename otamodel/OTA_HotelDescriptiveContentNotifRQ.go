package otamodel

type OTAHotelDescriptiveContentNotifRQ struct {
	PrimaryLangID            string                   `json:"PrimaryLangID,attr"`
	EchoToken                string                   `json:"EchoToken,attr"`
	TimeStamp                string                   `json:"TimeStamp,attr"`
	CorrelationID            string                   `json:"CorrelationID,attr"`
	Version                  string                   `json:"Version,attr"`
	Target                   string                   `json:"Target,attr"`
	HotelDescriptiveContents HotelDescriptiveContents `json:"HotelDescriptiveContents"`
}

type HotelDescriptiveContents struct {
	HotelDescriptiveContent []HotelDescriptiveContent `json:"HotelDescriptiveContent"`
}

type HotelDescriptiveContent struct {
	HotelName              string                 `json:"HotelName,attr"`
	HotelCode              string                 `json:"HotelCode,attr"`
	LanguageCode           string                 `json:"LanguageCode,attr"`
	ContactInfos           ContactInfos           `json:"ContactInfos"`
	HotelInfo              HotelInfo              `json:"HotelInfo"`
	FacilityInfo           FacilityInfo           `json:"FacilityInfo"`
	AreaInfo               AreaInfo               `json:"AreaInfo,omitempty"`
	Policies               Policies               `json:"Policies,omitempty"`
	AffiliationInfo        AffiliationInfo        `json:"AffiliationInfo,omitempty"`
	MultimediaDescriptions MultimediaDescriptions `json:"MultimediaDescriptions,omitempty"`
}

type ContactInfos struct {
	ContactInfo []ContactInfo `json:"ContactInfo"`
}

type HotelInfo struct {
	CategoryCodes            CategoryCodes            `json:"CategoryCodes"`
	Languages                Languages                `json:"Languages,omitempty"`
	Position                 Position                 `json:"Position"`
	Services                 Services                 `json:"Services,omitempty"`
	OwnershipManagementInfos OwnershipManagementInfos `json:"OwnershipManagementInfos,omitempty"`
}

type CategoryCodes struct {
	GuestRoomInfo GuestRoomInfo `json:"GuestRoomInfo"`
	HotelCategory HotelCategory `json:"HotelCategory,omitempty"`
}

type GuestRoomInfo struct {
	Quantity string `json:"Quantity,attr"`
}

type HotelCategory struct {
	Code       string `json:"Code,attr"`
	ExistsCode string `json:"ExistsCode,attr"`
}

type Languages struct {
	Language []Language `json:"Language"`
}

type Language struct {
	LanguageCode string `json:"LanguageCode,attr"`
}

type Position struct {
	Latitude  string `json:"Latitude,attr"`
	Longitude string `json:"Longitude,attr"`
}

type Services struct {
	Service []Service `json:"Service"`
}

type Service struct {
	Code           string         `json:"Code,attr"`
	Included       string         `json:"Included,attr,omitempty"`
	Price          string         `json:"Price,attr,omitempty"`
	CurrencyCode   string         `json:"CurrencyCode,attr,omitempty"`
	ExistsCode     string         `json:"ExistsCode,attr,omitempty"`
	Quantity       string         `json:"Quantity,attr,omitempty"`
	OperationTimes OperationTimes `json:"OperationTimes,attr,omitempty"`
	Features       Features       `json:"Features,omitempty"`
}

type OperationTimes struct {
	OperationTime []DaysOfWeek `json:"OperationTime"`
}

type Features struct {
	Feature []Feature `json:"Feature"`
}

type Feature struct {
	ID              string `json:"ID,attr,omitempty"`
	DescriptiveText string `json:"DescriptiveText,omitempty"`
	Charge          Charge `json:"Charge"`
}

type Charge struct {
	Amount string `json:"Amount,attr"`
}

type OwnershipManagementInfos struct {
	OwnershipManagementInfo []OwnershipManagementInfo `json:"OwnershipManagementInfo"`
}

type OwnershipManagementInfo struct {
	CompanyName CompanyName `json:"CompanyName"`
}

type CompanyName struct {
	Code string `json:"Code,attr"`
}

type FacilityInfo struct {
	GuestRooms GuestRooms `json:"GuestRooms"`
}

type GuestRooms struct {
	GuestRoom []GuestRoom `json:"GuestRoom"`
}

type GuestRoom struct {
	MaxAdultOccupancy string    `json:"MaxAdultOccupancy,attr"`
	MaxChildOccupancy string    `json:"MaxChildOccupancy,attr"`
	MaxOccupancy      string    `json:"MaxOccupancy,attr"`
	MinOccupancy      string    `json:"MinOccupancy,attr"`
	RoomTypeName      string    `json:"RoomTypeName,attr"`
	TypeRoom          TypeRoom  `json:"TypeRoom"`
	Amenities         Amenities `json:"Amenities"`
}

type TypeRoom struct {
	BedTypeCode  string `json:"BedTypeCode,attr"`
	RoomTypeCode string `json:"RoomTypeCode,attr"`
	NonSmoking   string `json:"NonSmoking,attr"`
}

type Amenities struct {
	Amenity []Amenity `json:"Amenity"`
}

type Amenity struct {
	RoomAmenityCode string `json:"RoomAmenityCode,attr"`
	Quantity        string `json:"Quantity,attr"`
}

type Restaurants struct {
	Restaurant []Restaurant `json:"Restaurant"`
}

type Restaurant struct {
	RestaurantName     string             `json:"RestaurantName,attr"`
	OfferBreakfast     string             `json:"OfferBreakfast,attr"`
	OfferDinner        string             `json:"OfferDinner,attr"`
	CuisineCodes       CuisineCodes       `json:"CuisineCodes"`
	OperationSchedules OperationSchedules `json:"OperationSchedules"`
	Features           Features           `json:"Features"`
}

type CuisineCodes struct {
	CuisineCode []CuisineCode `json:"CuisineCode"`
}

type CuisineCode struct {
	Code string `xml:"Code,attr"`
}

type OperationSchedules struct {
	OperationSchedule []OperationSchedule `json:"OperationSchedule"`
}

type OperationSchedule struct {
	OperationTimes OperationTimes `json:"OperationTimes"`
}

type AreaInfo struct {
	Attractions Attractions `json:"Attractions"`
}

type Attractions struct {
	Attraction []Attraction `json:"Attraction"`
}

type Attraction struct {
	AttractionName         string `json:"AttractionName,attr"`
	AttractionCategoryCode string `json:"AttractionCategoryCode,attr"`
	Distance               string `json:"Distance,attr"`
	DistanceUnit           string `json:"DistanceUnit,attr,omitempty"`
	LanguageCode           string `json:"LanguageCode,attr,omitempty"`
}

type Policies struct {
	Policy []Policy `json:"Policy"`
}

type Policy struct {
	GuaranteePaymentPolicy string       `json:"GuaranteePaymentPolicy"`
	PolicyInfo             PolicyInfo   `json:"PolicyInfo"`
	PetsPolicies           PetsPolicies `json:"PetsPolicies"`
	TaxPolicies            TaxPolicies  `json:"TaxPolicies"`
	FeePolicies            FeePolicies  `json:"FeePolicies"`
}

type PolicyInfo struct {
	CheckInTime  string `json:"CheckInTime,attr"`
	CheckOutTime string `json:"CheckOutTime,attr"`
}

type PetsPolicies struct {
	PetsAllowedCode string       `json:"PetsAllowedCode,attr"`
	PetsPolicy      []PetsPolicy `json:"PetsPolicy"`
}

type PetsPolicy struct {
	NonRefundableFee string `json:"NonRefundableFee,attr"`
}

type TaxPolicies struct {
	TaxPolicy []TaxPolicy `json:"TaxPolicy"`
}

type TaxPolicy struct {
	Code            string `json:"Code,attr"`
	Percent         string `json:"Percent,attr"`
	Type            string `json:"Type,attr"`
	Amount          string `json:"Amount,attr"`
	ChargeFrequency string `json:"ChargeFrequency,attr"`
}

type FeePolicies struct {
	FeePolicy FeePolicy `json:"FeePolicy"`
}

type FeePolicy struct {
	Code            string `json:"Code,attr"`
	Amount          string `json:"Amount,attr"`
	Type            string `json:"Type,attr"`
	ChargeFrequency string `json:"ChargeFrequency,attr"`
}

type AffiliationInfo struct {
	Awards Awards `json:"Awards"`
}

type Awards struct {
	Award Award `json:"Award"`
}

type Award struct {
	Provider string `json:"Provider,attr"`
	Rating   string `json:"Rating,attr"`
}

type MultimediaDescriptions struct {
	MultimediaDescription []MultimediaDescription `json:"MultimediaDescription"`
}

type MultimediaDescription struct {
	ImageItems ImageItems `json:"ImageItems"`
}

type ImageItems struct {
	ImageItem []ImageItem `json:"ImageItem"`
}

type ImageItem struct {
	ImageFormat ImageFormat `json:"ImageFormat"`
}

type ImageFormat struct {
	Main string `json:"Main,attr,omitempty"`
	Sort string `json:"Sort,attr,omitempty"`
	URL  string `json:"URL"`
}

type ContactInfo struct {
	ContactProfileType string    `json:"ContactProfileType,attr"`
	Addresses          Addresses `json:"Addresses"`
	Names              Names     `json:"Names,omitempty"`
	Emails             Emails    `json:"Emails,omitempty"`
	Phones             Phones    `json:"Phones,omitempty"`
}

type Addresses struct {
	Address []Address `json:"Address"`
}

type Address struct {
	Language    string    `json:"Language,attr,omitempty"`
	AddressLine string    `json:"AddressLine"`
	CityName    string    `json:"CityName"`
	PostalCode  string    `json:"PostalCode"`
	CountryName string    `json:"CountryName"`
	HotelName   string    `json:"HotelName,omitempty"`
	StateProv   StateProv `json:"StateProv,omitempty"`
}

type StateProv struct {
	StateCode string `json:"StateCode,attr"`
}

type Names struct {
	Name []Name `json:"Name"`
}

type Name struct {
	Gender    string `json:"Gender,attr"`
	Language  string `json:"Language,attr"`
	GivenName string `json:"GivenName"`
	Surname   string `json:"Surname"`
	JobTitle  string `json:"JobTitle"`
}

type Emails struct {
	Email []string `json:"Email"`
}

type Phones struct {
	Phone []Phone `json:"Phone"`
}

type Phone struct {
	Extension     string `json:"Extension,attr,omitempty"`
	PhoneNumber   string `json:"PhoneNumber,attr"`
	PhoneTechType string `json:"PhoneTechType,attr"`
}

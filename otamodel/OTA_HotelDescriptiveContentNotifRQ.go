package otamodel

type OTAHotelDescriptiveContentNotifRQ struct {
	PrimaryLangID            string                   `json:"PrimaryLangID"`
	EchoToken                string                   `json:"EchoToken"`
	TimeStamp                string                   `json:"TimeStamp"`
	CorrelationID            string                   `json:"CorrelationID"`
	Version                  string                   `json:"Version"`
	Target                   string                   `json:"Target"`
	HotelDescriptiveContents HotelDescriptiveContents `json:"HotelDescriptiveContents"`
}

type HotelDescriptiveContents struct {
	HotelDescriptiveContent []HotelDescriptiveContent `json:"HotelDescriptiveContent"`
}

type HotelDescriptiveContent struct {
	HotelName              string                 `json:"HotelName"`
	HotelCode              string                 `json:"HotelCode"`
	LanguageCode           string                 `json:"LanguageCode"`
	CurrencyCode           string                 `json:"CurrencyCode"`
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
	Descriptions             HotelDescriptions        `json:"Descriptions,omitempty"`
}

type CategoryCodes struct {
	GuestRoomInfo GuestRoomInfo `json:"GuestRoomInfo"`
	HotelCategory HotelCategory `json:"HotelCategory,omitempty"`
}

type GuestRoomInfo struct {
	Quantity string `json:"Quantity"`
}

type HotelCategory struct {
	Code       string `json:"Code"`
	ExistsCode string `json:"ExistsCode"`
}

type Languages struct {
	Language []Language `json:"Language"`
}

type Language struct {
	LanguageCode string `json:"LanguageCode"`
}

type Position struct {
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

type Services struct {
	Service []Service `json:"Service"`
}

type HotelDescriptions struct {
	Description []HotelDescription `json:"Description"`
}

type HotelDescription struct {
	DescriptiveText string `json:"DescriptiveText,omitempty"`
}

type Service struct {
	Code           string         `json:"Code"`
	Included       string         `json:"Included,omitempty"`
	Price          string         `json:"Price,omitempty"`
	CurrencyCode   string         `json:"CurrencyCode,omitempty"`
	ExistsCode     string         `json:"ExistsCode,omitempty"`
	Quantity       string         `json:"Quantity,omitempty"`
	OperationTimes OperationTimes `json:"OperationTimes,omitempty"`
	Features       Features       `json:"Features,omitempty"`
}

type OperationTimes struct {
	OperationTime []DaysOfWeek `json:"OperationTime"`
}

type Features struct {
	Feature []Feature `json:"Feature"`
}

type Feature struct {
	ID              string `json:"ID,omitempty"`
	DescriptiveText string `json:"DescriptiveText,omitempty"`
	Charge          Charge `json:"Charge"`
}

type Charge struct {
	Amount string `json:"Amount"`
}

type OwnershipManagementInfos struct {
	OwnershipManagementInfo []OwnershipManagementInfo `json:"OwnershipManagementInfo"`
}

type OwnershipManagementInfo struct {
	CompanyName CompanyName `json:"CompanyName"`
}

type CompanyName struct {
	Code string `json:"Code"`
}

type FacilityInfo struct {
	GuestRooms GuestRooms `json:"GuestRooms"`
}

type GuestRooms struct {
	GuestRoom []GuestRoom `json:"GuestRoom"`
}

type GuestRoom struct {
	MaxAdultOccupancy string    `json:"MaxAdultOccupancy"`
	MaxChildOccupancy string    `json:"MaxChildOccupancy"`
	MaxOccupancy      string    `json:"MaxOccupancy"`
	MinOccupancy      string    `json:"MinOccupancy"`
	RoomTypeName      string    `json:"RoomTypeName"`
	TypeRoom          TypeRoom  `json:"TypeRoom"`
	Amenities         Amenities `json:"Amenities"`
}

type TypeRoom struct {
	BedTypeCode  string `json:"BedTypeCode"`
	RoomTypeCode string `json:"RoomTypeCode"`
	NonSmoking   string `json:"NonSmoking"`
}

type Amenities struct {
	Amenity []Amenity `json:"Amenity"`
}

type Amenity struct {
	RoomAmenityCode string `json:"RoomAmenityCode"`
	Quantity        string `json:"Quantity"`
}

type Restaurants struct {
	Restaurant []Restaurant `json:"Restaurant"`
}

type Restaurant struct {
	RestaurantName     string             `json:"RestaurantName"`
	OfferBreakfast     string             `json:"OfferBreakfast"`
	OfferDinner        string             `json:"OfferDinner"`
	CuisineCodes       CuisineCodes       `json:"CuisineCodes"`
	OperationSchedules OperationSchedules `json:"OperationSchedules"`
	Features           Features           `json:"Features"`
}

type CuisineCodes struct {
	CuisineCode []CuisineCode `json:"CuisineCode"`
}

type CuisineCode struct {
	Code string `xml:"Code"`
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
	AttractionName         string `json:"AttractionName"`
	AttractionCategoryCode string `json:"AttractionCategoryCode"`
	Distance               string `json:"Distance"`
	DistanceUnit           string `json:"DistanceUnit,omitempty"`
	LanguageCode           string `json:"LanguageCode,omitempty"`
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
	CheckInTime  string `json:"CheckInTime"`
	CheckOutTime string `json:"CheckOutTime"`
}

type PetsPolicies struct {
	PetsAllowedCode string       `json:"PetsAllowedCode"`
	PetsPolicy      []PetsPolicy `json:"PetsPolicy"`
}

type PetsPolicy struct {
	NonRefundableFee string `json:"NonRefundableFee"`
}

type TaxPolicies struct {
	TaxPolicy []TaxPolicy `json:"TaxPolicy"`
}

type TaxPolicy struct {
	Code            string `json:"Code"`
	Percent         string `json:"Percent"`
	Type            string `json:"Type"`
	Amount          string `json:"Amount"`
	ChargeFrequency string `json:"ChargeFrequency"`
}

type FeePolicies struct {
	FeePolicy FeePolicy `json:"FeePolicy"`
}

type FeePolicy struct {
	Code            string `json:"Code"`
	Amount          string `json:"Amount"`
	Type            string `json:"Type"`
	ChargeFrequency string `json:"ChargeFrequency"`
}

type AffiliationInfo struct {
	Awards Awards `json:"Awards"`
}

type Awards struct {
	Award Award `json:"Award"`
}

type Award struct {
	Provider string `json:"Provider"`
	Rating   string `json:"Rating"`
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
	Main string `json:"Main,omitempty"`
	Sort string `json:"Sort,omitempty"`
	URL  string `json:"URL"`
}

type ContactInfo struct {
	ContactProfileType string    `json:"ContactProfileType"`
	Addresses          Addresses `json:"Addresses"`
	Names              Names     `json:"Names,omitempty"`
	Emails             Emails    `json:"Emails,omitempty"`
	Phones             Phones    `json:"Phones,omitempty"`
	URLs               []URLs    `json:"URLs,omitempty"`
}

type Addresses struct {
	Address []Address `json:"Address"`
}

type URLs struct {
	URL string `json:"URL"`
}

type Address struct {
	Language    string    `json:"Language,omitempty"`
	AddressLine string    `json:"AddressLine"`
	CityName    string    `json:"CityName"`
	PostalCode  string    `json:"PostalCode"`
	CountryName string    `json:"CountryName"`
	HotelName   string    `json:"HotelName,omitempty"`
	StateProv   StateProv `json:"StateProv,omitempty"`
}

type StateProv struct {
	StateCode string `json:"StateCode"`
}

type Names struct {
	Name []Name `json:"Name"`
}

type Name struct {
	Gender    string `json:"Gender"`
	Language  string `json:"Language"`
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
	Extension     string `json:"Extension,omitempty"`
	PhoneNumber   string `json:"PhoneNumber"`
	PhoneTechType string `json:"PhoneTechType"`
}

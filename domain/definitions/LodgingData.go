// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// LodgingData represents class LodgingData
type LodgingData struct {
	Charges                            *[]LodgingCharge `json:"charges,omitempty"`
	CheckInDate                        *string          `json:"checkInDate,omitempty"`
	CheckOutDate                       *string          `json:"checkOutDate,omitempty"`
	FolioNumber                        *string          `json:"folioNumber,omitempty"`
	IsConfirmedReservation             *bool            `json:"isConfirmedReservation,omitempty"`
	IsFacilityFireSafetyConform        *bool            `json:"isFacilityFireSafetyConform,omitempty"`
	IsNoShow                           *bool            `json:"isNoShow,omitempty"`
	IsPreferenceSmokingRoom            *bool            `json:"isPreferenceSmokingRoom,omitempty"`
	NumberOfAdults                     *int32           `json:"numberOfAdults,omitempty"`
	NumberOfNights                     *int32           `json:"numberOfNights,omitempty"`
	NumberOfRooms                      *int32           `json:"numberOfRooms,omitempty"`
	ProgramCode                        *string          `json:"programCode,omitempty"`
	PropertyCustomerServicePhoneNumber *string          `json:"propertyCustomerServicePhoneNumber,omitempty"`
	PropertyPhoneNumber                *string          `json:"propertyPhoneNumber,omitempty"`
	RenterName                         *string          `json:"renterName,omitempty"`
	Rooms                              *[]LodgingRoom   `json:"rooms,omitempty"`
}

// NewLodgingData constructs a new LodgingData
func NewLodgingData() *LodgingData {
	return &LodgingData{}
}

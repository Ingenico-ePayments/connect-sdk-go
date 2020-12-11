// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// LodgingRoom represents class LodgingRoom
type LodgingRoom struct {
	DailyRoomRate                  *string `json:"dailyRoomRate,omitempty"`
	DailyRoomRateCurrencyCode      *string `json:"dailyRoomRateCurrencyCode,omitempty"`
	DailyRoomTaxAmount             *string `json:"dailyRoomTaxAmount,omitempty"`
	DailyRoomTaxAmountCurrencyCode *string `json:"dailyRoomTaxAmountCurrencyCode,omitempty"`
	NumberOfNightsAtRoomRate       *int32  `json:"numberOfNightsAtRoomRate,omitempty"`
	RoomLocation                   *string `json:"roomLocation,omitempty"`
	RoomNumber                     *string `json:"roomNumber,omitempty"`
	TypeOfBed                      *string `json:"typeOfBed,omitempty"`
	TypeOfRoom                     *string `json:"typeOfRoom,omitempty"`
}

// NewLodgingRoom constructs a new LodgingRoom
func NewLodgingRoom() *LodgingRoom {
	return &LodgingRoom{}
}

package reservations

import (
	"be-library/app/model"
	"time"
)

type ResponseReservation struct{
	model.BaseModel
	MemberId string `json:"member_id" binding:"required"`
	BookHubId string `json:"book_hub_id" binding:"required"`
	PickupDeadline time.Time `json:"pickup_deadline" binding:"required"`
}

func (ResponseReservation) TableName()string{
	return "reservations"
}
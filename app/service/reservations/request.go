package reservations

import "time"

type RequestReservation struct {
	MemberId       string    `json:"member_id" binding:"required"`
	BookHubId      string    `json:"book_hub_id" binding:"required"`
	PickupDeadline time.Time `json:"pickup_deadline" binding:"required"`
}

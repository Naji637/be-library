package member

type RequestMember struct {
	MembershipNumber string `json:"membership_number" binding:"required"`
	Name			 string `json:"name" binding:"required"`
	Alumni bool `json:"alumni" binding:"required"`
}
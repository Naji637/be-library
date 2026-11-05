package member

type RequestMember struct {
	Name                   string `json:"name" binding:"required"`
	Alumni                 *bool  `json:"alumni" binding:"required"`
	MemberClassificationId string `json:"member_classification_id" binding:"required"`
}

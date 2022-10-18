package request

type MemberCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"adress"`
	Phone   string `json:"phone"`
	Image   string `json:"image"`
}

type MemberUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"adress"`
	Phone   string `json:"phone"`
	Image   string `json:"image"`
}

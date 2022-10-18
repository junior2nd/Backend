package request

type ApartmentCreateRequest struct{
	Name     string         `json:"name"`
	Branch   string         `json:"branch"`
	Phone    string         `json:"phoen"`
	Adress   string         `json:"adress"`
	Image    string         `json:"image"`
}
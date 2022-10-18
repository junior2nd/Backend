package request

type RoomtypeCreateRequest struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Price  int    `json:"price"`
}

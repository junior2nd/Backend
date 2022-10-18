package request

type RoomCreateRequest struct{
	Name        string `json:"name"`
	Status      int    `json:"status"`
}
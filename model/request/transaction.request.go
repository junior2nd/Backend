package request


type TransactionCreateRequest struct {
	Status bool    `json:"status"`
	Evalue float64 `json:"e_value" gorm:"type:decimal(10,2)"`
	Wvalue float64 `json:"w_value" gorm:"type:decimal(10,2)"`
}

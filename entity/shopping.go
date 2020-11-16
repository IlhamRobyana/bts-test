package entity

// Shopping is used to store the shopping data
type Shopping struct {
	ID          uint64 `json:"id" gorm:"primary_key;type:bigserial"`
	Name        string `json:"name"`
	CreatedDate string `json:"createddate"`
	Timestamp
}

type ShoppingCreateRequest struct {
	Shopping Shopping `json:"shopping"`
}

package entity

// User is used to store the user's data
type User struct {
	ID       uint64 `json:"id" gorm:"primary_key;type:bigserial"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Phone    string `json:"phone"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Postcode string `json:"postcode"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Timestamp
}

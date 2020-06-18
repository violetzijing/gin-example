package models

// User is gorm model
type User struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

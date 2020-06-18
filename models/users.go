package models

// User is gorm model
type User struct {
	ID           int    `gorm:"primary_key" json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	City         string `json:"city"`
	PasswordHash string
}

// ReadFromMap reads from map and assign fields to user
func (u *User) ReadFromMap(m map[string]interface{}) {
	u.ID = int(m["id"].(float64))
	u.Name = m["name"].(string)
}

func (u *User) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":   u.ID,
		"name": u.Name,
		"Age":  u.Age,
		"City": u.City,
	}
}

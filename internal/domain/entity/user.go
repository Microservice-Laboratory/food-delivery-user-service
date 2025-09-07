package entity

type UserRole string

const (
	CustomerRole UserRole = "customer"
	DriverRole   UserRole = "driver"
	Restaurant   UserRole = "restaurant"
)

type User struct {
	BaseEntity
	FirstName   string   `gorm:"size:100;not null"`
	LastName    string   `gorm:"size:100;not null"`
	Email       string   `gorm:"size:100;uniqueIndex:user_idx_email;not null"`
	Password    string   `gorm:"not null"`
	Phone       string   `gorm:"uniqueIndex:user_idx_phone;not null"`
	Role        UserRole `gorm:"type:varchar(30);not null"`
	UserAddress []UserAddress
}

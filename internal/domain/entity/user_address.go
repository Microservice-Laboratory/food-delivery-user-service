package entity

type UserAddress struct {
	BaseEntity
	User      User
	ZipCode   string `gorm:"not null"`
	Country   string
	State     string
	City      string
	Street    string
	IsDefault bool
	UserId    uint
}

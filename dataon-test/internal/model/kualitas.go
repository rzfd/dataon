package model

type Kualitas struct {
	ID      uint   `gorm:"primaryKey"`
	Tingkat int    `gorm:"not null"`
	Nama    string `gorm:"not null"`
	Harga   int    `gorm:"not null"`
}

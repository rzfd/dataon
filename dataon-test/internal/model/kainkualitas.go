package model

type KainKualitas struct {
	ID         uint     `gorm:"primaryKey"`
	KainID     uint     `gorm:"not null"`
	KualitasID uint     `gorm:"not null"`
	Kain       Kain     `gorm:"foreignKey:KainID"`
	Kualitas   Kualitas `gorm:"foreignKey:KualitasID"`
}

package model

type JenisKain struct {
	ID   uint   `gorm:"primaryKey"`
	Kode string `gorm:"unique;not null"`
}

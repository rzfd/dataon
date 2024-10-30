package model

type Kain struct {
	ID          uint      `gorm:"primaryKey"`
	Nama        string    `gorm:"not null"`
	JenisKainID uint      `gorm:"not null"` // Foreign key to JenisKain
	JenisKain   JenisKain `gorm:"foreignKey:JenisKainID"`
}

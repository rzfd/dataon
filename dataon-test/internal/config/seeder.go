package config

import (
	"log"

	"github.com/rzfd/dataon-test/internal/model"
)

func SeedDB() {
	jenisKains := []model.JenisKain{
		{Kode: "STB"},
		{Kode: "NTB"},
	}
	if err := DB.Create(&jenisKains).Error; err != nil {
		log.Fatalf("failed to seed jenis_kain: %v", err)
	}

	kains := []model.Kain{
		{Nama: "Sutra", JenisKainID: jenisKains[0].ID},
		{Nama: "Katun", JenisKainID: jenisKains[1].ID},
	}
	if err := DB.Create(&kains).Error; err != nil {
		log.Fatalf("failed to seed kain: %v", err)
	}

	kualitasSTB := []model.Kualitas{
		{Tingkat: 1, Nama: "Sangat Bagus", Harga: 10000000},
		{Tingkat: 2, Nama: "Bagus", Harga: 9000000},
		{Tingkat: 3, Nama: "Cukup Bagus", Harga: 8000000},
	}

	kualitasNTB := []model.Kualitas{
		{Tingkat: 1, Nama: "Sangat Bagus", Harga: 10000000},
		{Tingkat: 2, Nama: "Bagus", Harga: 10000000},
		{Tingkat: 3, Nama: "Cukup Bagus", Harga: 10000000},
	}

	if err := DB.Create(&kualitasSTB).Error; err != nil {
		log.Fatalf("failed to seed kualitas STB: %v", err)
	}
	if err := DB.Create(&kualitasNTB).Error; err != nil {
		log.Fatalf("failed to seed kualitas NTB: %v", err)
	}

	kainKualitas := []model.KainKualitas{
		{KainID: kains[0].ID, KualitasID: kualitasSTB[0].ID},
		{KainID: kains[0].ID, KualitasID: kualitasSTB[1].ID},
		{KainID: kains[0].ID, KualitasID: kualitasSTB[2].ID},
		{KainID: kains[1].ID, KualitasID: kualitasNTB[0].ID},
		{KainID: kains[1].ID, KualitasID: kualitasNTB[1].ID},
		{KainID: kains[1].ID, KualitasID: kualitasNTB[2].ID},
	}
	if err := DB.Create(&kainKualitas).Error; err != nil {
		log.Fatalf("failed to seed kain_kualitas: %v", err)
	}
}

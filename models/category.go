package models

import "github.com/jinzhu/gorm"

// Model untuk kategori produk
type Category struct {
	gorm.Model
	Name        string    `json:"name"`                  // Nama kategori
	Description string    `json:"description"`           // Deskripsi kategori
	Products    []Product `gorm:"foreignkey:CategoryID"` // Relasi satu ke banyak dengan produk
}

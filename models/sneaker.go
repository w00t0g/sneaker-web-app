package models

import "time"

// Sneaker struct represents a sneaker
type Sneaker struct {
	ID            int        `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     time.Time  `json:"-"`
	DeletedAt     *time.Time `gorm:"index" json:"-"`
	Brand         string     `json:"brand"`
	Model         string     `json:"model"`
	Color         string     `json:"color"`
	Platform      string     `json:"platform"`
	PurchaseDate  string     `json:"purchaseDate"`
	PurchasePrice int        `json:"purchasePrice"`
	Quantity      int        `json:"quantity"`
}

// SoldSneaker struct represents a sold sneaker
type SoldSneaker struct {
	ID        int        `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	SneakerID int        `json:"sneakerID"` // Foreign Key
	Price     int        `json:"price"`
	Quantity  int        `json:"quantity"`

	Sneaker Sneaker `gorm:"foreignKey:SneakerID" json:"sneaker"` // One-to-One relationship with Sneaker
}

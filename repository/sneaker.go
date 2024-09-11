package repository

import (
	"Sneaker_Inventory/models"

	"gorm.io/gorm"
)

// SneakerRepository is an interface for interacting with the database
type SneakerRepository interface {
	AddSneaker(models.Sneaker) (*models.Sneaker, error)
	GetSneaker(int) (*models.Sneaker, error)
	GetSneakers() ([]models.Sneaker, error)
	UpdateSneaker(models.Sneaker) (*models.Sneaker, error)
	DeleteSneaker(int) error
	SellSneaker(models.SoldSneaker) (*models.SoldSneaker, error)
	SoldSneakers() ([]models.SoldSneaker, error)
	DeleteSoldSneakerBySneakerID(int) error

	SearchSneakerByModel(string) ([]models.Sneaker, error)
}

type sneakerRepository struct {
	db *gorm.DB
}

// NewSneakerRepository creates a new instance of SneakerRepository
func NewSneakerRepository(db *gorm.DB) SneakerRepository {
	return &sneakerRepository{db}
}

// AddSneaker adds a new sneaker to the database
func (s *sneakerRepository) AddSneaker(sneaker models.Sneaker) (*models.Sneaker, error) {
	result := s.db.Create(&sneaker)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sneaker, nil
}

// GetSneaker retrieves a sneaker from the database
func (s *sneakerRepository) GetSneaker(id int) (*models.Sneaker, error) {
	var sneaker models.Sneaker
	result := s.db.First(&sneaker, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sneaker, nil
}

// GetSneakers retrieves all sneakers from the database
func (s *sneakerRepository) GetSneakers() ([]models.Sneaker, error) {
	var sneakers []models.Sneaker
	result := s.db.Find(&sneakers)
	if result.Error != nil {
		return nil, result.Error
	}
	return sneakers, nil
}

// UpdateSneaker updates a sneaker in the database
func (s *sneakerRepository) UpdateSneaker(sneaker models.Sneaker) (*models.Sneaker, error) {
	result := s.db.Save(&sneaker)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sneaker, nil
}

// SellSneaker adds a sold sneaker to the database
func (s *sneakerRepository) SellSneaker(soldSneaker models.SoldSneaker) (*models.SoldSneaker, error) {
	result := s.db.Create(&soldSneaker)
	if result.Error != nil {
		return nil, result.Error
	}
	return &soldSneaker, nil
}

// SoldSneakers retrieves all sold sneakers from the database
func (s *sneakerRepository) SoldSneakers() ([]models.SoldSneaker, error) {
	var soldSneakers []models.SoldSneaker
	result := s.db.Preload("Sneaker").Find(&soldSneakers)
	if result.Error != nil {
		return nil, result.Error
	}
	return soldSneakers, nil
}

// DeleteSoldSneakerBySneakerID deletes a sold sneaker from the database
func (s *sneakerRepository) DeleteSoldSneakerBySneakerID(id int) error {
	result := s.db.Where("sneaker_id = ?", id).Delete(&models.SoldSneaker{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteSneaker deletes a sneaker from the database
func (s *sneakerRepository) DeleteSneaker(id int) error {
	result := s.db.Delete(&models.Sneaker{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SearchSneakerByModel searches for sneakers by model
func (s *sneakerRepository) SearchSneakerByModel(model string) ([]models.Sneaker, error) {
	var sneakers []models.Sneaker

	result := s.db.Where("model LIKE ?", "%"+model+"%").Find(&sneakers)
	if result.Error != nil {
		return nil, result.Error
	}
	return sneakers, nil
}

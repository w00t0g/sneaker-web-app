package repository

// unit test for
// // SneakerRepository is an interface for interacting with the database
// type SneakerRepository interface {
// 	AddSneaker(models.Sneaker) (*models.Sneaker, error)
// 	GetSneaker(int) (*models.Sneaker, error)
// 	GetSneakers() ([]models.Sneaker, error)
// 	UpdateSneaker(models.Sneaker) (*models.Sneaker, error)
// 	DeleteSneaker(int) error
// 	SellSneaker(models.SoldSneaker) (*models.SoldSneaker, error)
// 	SoldSneakers() ([]models.SoldSneaker, error)
// 	DeleteSoldSneakerBySneakerID(int) error

// 	SearchSneakerByModel(string) ([]models.Sneaker, error)
// }

// type sneakerRepository struct {
// 	db *gorm.DB
// }

// // NewSneakerRepository creates a new instance of SneakerRepository
// func NewSneakerRepository(db *gorm.DB) SneakerRepository {
// 	return &sneakerRepository{db}
// }

import (
	"Sneaker_Inventory/models"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var repo SneakerRepository

func getRepo() SneakerRepository {
	if repo != nil {
		return repo
	}

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.Sneaker{}, &models.SoldSneaker{})
	repo = NewSneakerRepository(db)

	return repo
}

// TestAddSneaker tests the AddSneaker method
func TestAddSneaker(t *testing.T) {
	repo := getRepo()

	sneaker := models.Sneaker{
		Brand:         "Nike",
		Model:         "Air Max 90",
		Color:         "White",
		Platform:      "Air",
		PurchaseDate:  "2021-01-01",
		PurchasePrice: 100,
		Quantity:      10,
	}

	_sneaker, err := repo.AddSneaker(sneaker)
	if err != nil {
		t.Errorf("AddSneaker failed: %v", err)
	}

	if _sneaker.ID == 0 {
		t.Errorf("AddSneaker failed: expected ID to be set")
	}
}

// TestGetSneaker tests the GetSneaker method
func TestGetSneaker(t *testing.T) {
	repo := getRepo()

	sneaker := models.Sneaker{
		Brand:         "Nike",
		Model:         "Air Max 90",
		Color:         "White",
		Platform:      "Air",
		PurchaseDate:  "2021-01-01",
		PurchasePrice: 100,
		Quantity:      10,
	}

	_sneaker, err := repo.AddSneaker(sneaker)
	if err != nil {
		t.Errorf("AddSneaker failed: %v", err)
	}

	_sneaker, err = repo.GetSneaker(_sneaker.ID)
	if err != nil {
		t.Errorf("GetSneaker failed: %v", err)
	}

	if _sneaker.ID == 0 {
		t.Errorf("GetSneaker failed: expected ID to be set")
	}
}

// TestGetSneakers tests the GetSneakers method
func TestGetSneakers(t *testing.T) {
	repo := getRepo()

	sneaker := models.Sneaker{
		Brand:         "Nike",
		Model:         "Air Max 90",
		Color:         "White",
		Platform:      "Air",
		PurchaseDate:  "2021-01-01",
		PurchasePrice: 100,
		Quantity:      10,
	}

	_, err := repo.AddSneaker(sneaker)
	if err != nil {
		t.Errorf("AddSneaker failed: %v", err)
	}

	sneakers, err := repo.GetSneakers()
	if err != nil {
		t.Errorf("GetSneakers failed: %v", err)
	}

	if len(sneakers) == 0 {
		t.Errorf("GetSneakers failed: expected at least one sneaker")
	}
}

// TestUpdateSneaker tests the UpdateSneaker method
func TestUpdateSneaker(t *testing.T) {
	repo := getRepo()

	sneaker := models.Sneaker{
		Brand:         "Nike",
		Model:         "Air Max 90",
		Color:         "White",
		Platform:      "Air",
		PurchaseDate:  "2021-01-01",
		PurchasePrice: 100,
		Quantity:      10,
	}

	_sneaker, err := repo.AddSneaker(sneaker)
	if err != nil {
		t.Errorf("AddSneaker failed: %v", err)
	}

	_sneaker.Color = "Black"
	_sneaker, err = repo.UpdateSneaker(*_sneaker)
	if err != nil {
		t.Errorf("UpdateSneaker failed: %v", err)
	}

	if _sneaker.Color != "Black" {
		t.Errorf("UpdateSneaker failed: expected Color to be updated")
	}
}

// TestDeleteSneaker tests the DeleteSneaker method
func TestDeleteSneaker(t *testing.T) {
	repo := getRepo()

	sneaker := models.Sneaker{
		Brand:         "Nike",
		Model:         "Air Max 90",
		Color:         "White",
		Platform:      "Air",
		PurchaseDate:  "2021-01-01",
		PurchasePrice: 100,
		Quantity:      10,
	}

	_sneaker, err := repo.AddSneaker(sneaker)
	if err != nil {
		t.Errorf("AddSneaker failed: %v", err)
	}

	err = repo.DeleteSneaker(_sneaker.ID)
	if err != nil {
		t.Errorf("DeleteSneaker failed: %v", err)
	}

	// Check if the sneaker was deleted
	_, err = repo.GetSneaker(_sneaker.ID)
	if err == nil {
		t.Errorf("DeleteSneaker failed: expected error")
	}
}

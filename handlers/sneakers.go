package handlers

import (
	"Sneaker_Inventory/models"
	"Sneaker_Inventory/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type SneakerHandler struct {
	sp repository.SneakerRepository
}

// NewHandler creates a new instance of SneakerHandler
func NewHandler(sneakerRepository repository.SneakerRepository) *SneakerHandler {
	return &SneakerHandler{sneakerRepository}
}

// Add Sneaker
func (sh *SneakerHandler) AddSneaker(w http.ResponseWriter, r *http.Request) {
	var sneaker models.Sneaker
	err := json.NewDecoder(r.Body).Decode(&sneaker)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	storedSneaker, err := sh.sp.AddSneaker(sneaker)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// set json header
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storedSneaker)
}

// Get Sneaker
func (sh *SneakerHandler) GetSneaker(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	sneaker, err := sh.sp.GetSneaker(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(sneaker)
}

// Get Sneakers
func (sh *SneakerHandler) GetSneakers(w http.ResponseWriter, r *http.Request) {
	var sneakers []models.Sneaker

	query := r.FormValue("search")
	if query != "" {
		var err error
		sneakers, err = sh.sp.SearchSneakerByModel(query)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
	} else {
		var err error
		sneakers, err = sh.sp.GetSneakers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
	}

	json.NewEncoder(w).Encode(sneakers)
}

// Update Sneaker
func (h *SneakerHandler) UpdateSneaker(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	sneakerID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	var sneaker models.Sneaker
	err = json.NewDecoder(r.Body).Decode(&sneaker)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	storedSneaker, err := h.sp.GetSneaker(sneakerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	storedSneaker.Brand = sneaker.Brand
	storedSneaker.Model = sneaker.Model
	storedSneaker.Color = sneaker.Color
	storedSneaker.Platform = sneaker.Platform
	storedSneaker.PurchaseDate = sneaker.PurchaseDate
	storedSneaker.PurchasePrice = sneaker.PurchasePrice
	storedSneaker.Quantity = sneaker.Quantity

	_, err = h.sp.UpdateSneaker(*storedSneaker)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(storedSneaker)
}

// Sell Sneaker
func (h *SneakerHandler) SellSneaker(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	sneakerID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	var soldSneaker models.SoldSneaker
	err = json.NewDecoder(r.Body).Decode(&soldSneaker)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	storedSneaker, err := h.sp.GetSneaker(sneakerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	storedSneaker.Quantity = storedSneaker.Quantity - soldSneaker.Quantity
	_, err = h.sp.UpdateSneaker(*storedSneaker)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	soldSneaker.SneakerID = storedSneaker.ID

	storedSoldSneaker, err := h.sp.SellSneaker(soldSneaker)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(storedSoldSneaker)
}

// Sold Sneakers
func (h *SneakerHandler) SoldSneakers(w http.ResponseWriter, r *http.Request) {
	soldSneakers, err := h.sp.SoldSneakers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(soldSneakers)
}

// Delete Sneaker
func (h *SneakerHandler) DeleteSneaker(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	sneakerID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	err = h.sp.DeleteSoldSneakerBySneakerID(sneakerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	err = h.sp.DeleteSneaker(sneakerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Sneaker deleted successfully"})
}

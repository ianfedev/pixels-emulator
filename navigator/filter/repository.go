package repository

import (
	model "pixels-emulator/core/model"
	"sync"

	"gorm.io/gorm"
)

// RoomFilterHandler handles room filtering logic.
type RoomFilterHandler struct {
	ID         string
	FilterFunc func(*gorm.DB) *gorm.DB
}

// filters store the needed filters.
var (
	filters   = make(map[string]RoomFilterHandler)
	filtersMu sync.Mutex
)

// RegisterFilter adds a new filter dynamically.
func RegisterFilter(id string, filterFunc func(*gorm.DB) *gorm.DB) {
	filtersMu.Lock()
	defer filtersMu.Unlock()

	filters[id] = RoomFilterHandler{
		ID:         id,
		FilterFunc: filterFunc,
	}
}

// FilterExists checks if a filter is registered.
func FilterExists(filterID string) bool {
	filtersMu.Lock()
	defer filtersMu.Unlock()

	_, exists := filters[filterID]
	return exists
}

// GetRoomsByFilter applies the selected filter.
func GetRoomsByFilter(db *gorm.DB, filterID string) ([]model.Room, error) {
	if !FilterExists(filterID) {
		return nil, nil
	}

	filtersMu.Lock()
	filter := filters[filterID]
	filtersMu.Unlock()

	var rooms []model.Room
	db = filter.FilterFunc(db)
	err := db.Find(&rooms).Error
	return rooms, err
}

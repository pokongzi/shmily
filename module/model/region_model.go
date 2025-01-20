package model

import (
    "gorm.io/gorm"
    "module/entity"
)

// RegionModel handles database operations for Region
type RegionModel struct {
    DB *gorm.DB
}

// CreateRegion creates a new region record
func (m *RegionModel) CreateRegion(region *entity.Region) error {
    return m.DB.Create(region).Error
}

// GetRegion retrieves a region by ID
func (m *RegionModel) GetRegion(id int64) (*entity.Region, error) {
    var region entity.Region
    err := m.DB.First(&region, id).Error
    return &region, err
}

// UpdateRegion updates an existing region record
func (m *RegionModel) UpdateRegion(region *entity.Region) error {
    return m.DB.Save(region).Error
}

// DeleteRegion deletes a region by ID
func (m *RegionModel) DeleteRegion(id int64) error {
    return m.DB.Delete(&entity.Region{}, id).Error
}

// NewRegionModel creates a new instance of RegionModel
func NewRegionModel(db *gorm.DB) *RegionModel {
    return &RegionModel{
        DB: db,
    }
} 
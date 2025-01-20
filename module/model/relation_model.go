package model

import (
    "gorm.io/gorm"
    "module/entity"
)

// RelationModel handles database operations for Relation
type RelationModel struct {
    DB *gorm.DB
}

// CreateRelation creates a new relation record
func (m *RelationModel) CreateRelation(relation *entity.Relation) error {
    return m.DB.Create(relation).Error
}

// GetRelation retrieves a relation by ID
func (m *RelationModel) GetRelation(id int64) (*entity.Relation, error) {
    var relation entity.Relation
    err := m.DB.First(&relation, id).Error
    return &relation, err
}

// UpdateRelation updates an existing relation record
func (m *RelationModel) UpdateRelation(relation *entity.Relation) error {
    return m.DB.Save(relation).Error
}

// DeleteRelation deletes a relation by ID
func (m *RelationModel) DeleteRelation(id int64) error {
    return m.DB.Delete(&entity.Relation{}, id).Error
}

// NewRelationModel creates a new instance of RelationModel
func NewRelationModel(db *gorm.DB) *RelationModel {
    return &RelationModel{
        DB: db,
    }
} 
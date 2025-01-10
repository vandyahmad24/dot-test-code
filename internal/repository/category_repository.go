package repository

import (
	"context"
	"dot-test-vandy/internal/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(ctx context.Context, cateogry *model.Category) error
	GetAll(ctx context.Context) ([]*model.Category, error)
	GetByID(ctx context.Context, id uint) (*model.Category, error)
	Update(ctx context.Context, cateogry *model.Category) error
	Delete(ctx context.Context, id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Create(ctx context.Context, cateogry *model.Category) error {
	return r.db.Create(cateogry).Error
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]*model.Category, error) {
	var cateogrys []*model.Category
	err := r.db.Preload("Books").Find(&cateogrys).Error
	if err != nil {
		return nil, err
	}
	return cateogrys, nil
}
func (r *categoryRepository) GetByID(ctx context.Context, id uint) (*model.Category, error) {
	var cateogry model.Category
	return &cateogry, r.db.Preload("Books").First(&cateogry, id).Error
}
func (r *categoryRepository) Update(ctx context.Context, cateogry *model.Category) error {
	return r.db.Save(cateogry).Error
}
func (r *categoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("category_id = ?", id).Delete(&model.Book{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.Category{}, id).Error
	})
}

package repository

import (
	"context"
	"dot-test-vandy/internal/model"
	"errors"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(ctx context.Context, book *model.Book) error
	GetAll(ctx context.Context) ([]*model.Book, error)
	GetByID(ctx context.Context, id uint) (*model.Book, error)
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, id uint) error
	GetByCateogryID(ctx context.Context, id uint) ([]*model.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) Create(ctx context.Context, book *model.Book) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if book.CategoryID == 0 {
			var category model.Category
			if err := tx.FirstOrCreate(&category, model.Category{Name: "Uncategorized"}).Error; err != nil {
				return err
			}
			book.CategoryID = category.ID
		}
		return tx.Create(&book).Error
	})
	if err != nil {
		return errors.New("failed to create book")
	}
	return nil
}

func (r *bookRepository) GetAll(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	err := r.db.Preload("Category").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (r *bookRepository) GetByID(ctx context.Context, id uint) (*model.Book, error) {
	var book model.Book
	return &book, r.db.Preload("Category").First(&book, id).Error
}
func (r *bookRepository) Update(ctx context.Context, book *model.Book) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&model.Book{}, book.ID).Error; err != nil {
			return errors.New("book not found")
		}

		var category model.Category
		if err := tx.First(&category, book.CategoryID).Error; err != nil {
			return errors.New("category not found")
		}

		// Update the book
		return tx.Save(book).Error
	})

	if err != nil {
		return err
	}
	return nil

}
func (r *bookRepository) Delete(ctx context.Context, id uint) error {
	return r.db.Delete(&model.Book{}, id).Error
}
func (r *bookRepository) GetByCateogryID(ctx context.Context, id uint) ([]*model.Book, error) {
	var books []*model.Book
	return books, r.db.Where("category_id = ?", id).Find(&books).Error
}

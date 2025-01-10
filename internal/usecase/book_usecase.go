package usecase

import (
	"context"
	"dot-test-vandy/internal/model"
	"dot-test-vandy/internal/repository"
	"encoding/json"
	"fmt"
)

type BookUsecase interface {
	GetAll(ctx context.Context) ([]*model.Book, error)
	GetByID(ctx context.Context, id uint) (*model.Book, error)
	Create(ctx context.Context, book *model.Book) error
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, id uint) error
}

type bookUsecase struct {
	bookRepository repository.BookRepository
	redis          repository.RedisRepository
}

func NewBookUsecase(bookRepository repository.BookRepository, redis repository.RedisRepository) BookUsecase {
	return &bookUsecase{
		bookRepository: bookRepository,
		redis:          redis,
	}
}

func (u *bookUsecase) GetAll(ctx context.Context) ([]*model.Book, error) {
	key := "books"
	booksString, err := u.redis.Get(ctx, key)
	if err == nil && booksString != "" {
		var books []*model.Book
		err = json.Unmarshal([]byte(booksString), &books)
		if err != nil {
			return nil, err
		}

		return books, nil
	}

	books, err := u.bookRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	u.setAllBooksToRedis(ctx, books)

	return books, nil
}

func (u *bookUsecase) GetByID(ctx context.Context, id uint) (*model.Book, error) {
	key := fmt.Sprintf("book:%d", id)
	bookString, err := u.redis.Get(ctx, key)
	if err == nil && bookString != "" {
		var book model.Book
		err = json.Unmarshal([]byte(bookString), &book)
		if err != nil {
			return nil, err
		}

		return &book, nil
	}

	book, err := u.bookRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = u.setBookToRedis(ctx, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (u *bookUsecase) Create(ctx context.Context, book *model.Book) error {
	err := u.bookRepository.Create(ctx, book)
	if err != nil {
		return err
	}

	books, _ := u.bookRepository.GetAll(ctx)
	u.setAllBooksToRedis(ctx, books)

	return nil

}

func (u *bookUsecase) Update(ctx context.Context, book *model.Book) error {
	err := u.bookRepository.Update(ctx, book)
	if err != nil {
		return err
	}

	u.setBookToRedis(ctx, book)
	books, _ := u.bookRepository.GetAll(ctx)
	u.setAllBooksToRedis(ctx, books)

	return nil

}

func (u *bookUsecase) Delete(ctx context.Context, id uint) error {
	err := u.bookRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	u.redis.Delete(ctx, fmt.Sprintf("book:%d", id))
	books, _ := u.bookRepository.GetAll(ctx)
	u.setAllBooksToRedis(ctx, books)

	return nil
}

func (u *bookUsecase) setBookToRedis(ctx context.Context, book *model.Book) error {
	key := fmt.Sprintf("book:%d", book.ID)
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return err
	}

	return u.redis.Set(ctx, key, bookJSON, 0)
}

func (u *bookUsecase) setAllBooksToRedis(ctx context.Context, books []*model.Book) ([]*model.Book, error) {
	// set all in redis
	booksJSON, err := json.Marshal(books)
	if err != nil {
		return nil, err
	}
	err = u.redis.Set(ctx, "books", booksJSON, 0)
	if err != nil {
		return nil, err
	}
	return books, nil
}

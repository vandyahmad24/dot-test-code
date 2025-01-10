package setup

import (
	"dot-test-vandy/config"
	"dot-test-vandy/internal/handler"
	infrastructure "dot-test-vandy/internal/infrastructure/database"
	"dot-test-vandy/internal/repository"
	"dot-test-vandy/internal/usecase"
)

type Services struct {
	CategoryHandler *handler.CategoryHandler
	BookHandler     *handler.BookHandler
}

func NewServices() *Services {
	db := infrastructure.ConnectMysql()
	cfg := config.NewConfig()
	redis := repository.NewRedisRepository(cfg.REDIS_HOST, cfg.REDIS_PASS, 0)
	categoryRepo := repository.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo, redis)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	bookRepo := repository.NewBookRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepo, redis)
	bookHandler := handler.NewBookHandler(bookUsecase)

	return &Services{
		CategoryHandler: categoryHandler,
		BookHandler:     bookHandler,
	}

}

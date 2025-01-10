package usecase

import (
	"context"
	"dot-test-vandy/internal/model"
	"dot-test-vandy/internal/repository"
	"encoding/json"
	"fmt"
)

type CategoryUsecase interface {
	GetAll(ctx context.Context) ([]*model.Category, error)
	GetByID(ctx context.Context, id uint) (*model.Category, error)
	Create(ctx context.Context, category *model.Category) error
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id uint) error
}

type categoryUsecase struct {
	categoryRepository repository.CategoryRepository
	redis              repository.RedisRepository
}

func NewCategoryUsecase(categoryRepository repository.CategoryRepository, redis repository.RedisRepository) CategoryUsecase {
	return &categoryUsecase{
		categoryRepository: categoryRepository,
		redis:              redis,
	}
}

func (c *categoryUsecase) GetAll(ctx context.Context) ([]*model.Category, error) {
	// get all in redis
	var categori []*model.Category
	categoriString, err := c.redis.Get(ctx, "categories")
	if err != nil || categoriString == "" {
		categories, err := c.categoryRepository.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		categori = categories
		return c.setAllCategoriesToRedis(ctx, categori)
	}

	err = json.Unmarshal([]byte(categoriString), &categori)
	if err != nil {
		return nil, err
	}

	return categori, nil

}

func (c *categoryUsecase) GetByID(ctx context.Context, id uint) (*model.Category, error) {
	key := fmt.Sprintf("category:%d", id)
	categoriString, err := c.redis.Get(ctx, key)
	if err != nil || categoriString == "" {
		category, err := c.categoryRepository.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}
		return c.setCategoryToRedis(ctx, category)
	}

	var category model.Category
	err = json.Unmarshal([]byte(categoriString), &category)
	if err != nil {
		return nil, err
	}

	return &category, nil

}

func (c *categoryUsecase) setAllCategoriesToRedis(ctx context.Context, categories []*model.Category) ([]*model.Category, error) {
	// set all in redis
	crg, err := json.Marshal(categories)
	if err != nil {
		return nil, err
	}
	err = c.redis.Set(ctx, "categories", crg, 0)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *categoryUsecase) setCategoryToRedis(ctx context.Context, category *model.Category) (*model.Category, error) {
	// set to redis by id
	key := fmt.Sprintf("category:%d", category.ID)
	crg, err := json.Marshal(category)
	if err != nil {
		return nil, err
	}
	err = c.redis.Set(ctx, key, crg, 0)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *categoryUsecase) Create(ctx context.Context, category *model.Category) error {
	err := c.categoryRepository.Create(ctx, category)
	if err != nil {
		return err
	}

	_, err = c.setCategoryToRedis(ctx, category)
	if err != nil {
		return err
	}

	ctg, err := c.categoryRepository.GetAll(ctx)
	if err != nil {
		return err
	}

	_, err = c.setAllCategoriesToRedis(ctx, ctg)
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryUsecase) Update(ctx context.Context, category *model.Category) error {
	err := c.categoryRepository.Update(ctx, category)
	if err != nil {
		return err
	}

	_, err = c.setCategoryToRedis(ctx, category)
	if err != nil {
		return err
	}

	ctg, err := c.categoryRepository.GetAll(ctx)
	if err != nil {
		return err
	}

	_, err = c.setAllCategoriesToRedis(ctx, ctg)
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryUsecase) Delete(ctx context.Context, id uint) error {
	err := c.categoryRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	ctg, err := c.categoryRepository.GetAll(ctx)
	if err != nil {
		return err
	}

	_, err = c.setAllCategoriesToRedis(ctx, ctg)
	if err != nil {
		return err
	}

	return nil
}

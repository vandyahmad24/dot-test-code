package handler

import (
	"dot-test-vandy/internal/model"
	"dot-test-vandy/internal/request"
	"dot-test-vandy/internal/usecase"
	"dot-test-vandy/lib/wrapper"
	"net/http"

	"github.com/gin-gonic/gin"
	stringManupulation "github.com/vandyahmad24/alat-bantu/strings"
)

type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase}
}

func (h *CategoryHandler) GetAll(ctx *gin.Context) {
	resp, err := h.categoryUsecase.GetAll(ctx)
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	wrapper.WrapSuccessResponse(ctx, resp, "Success")
	return
}

func (h *CategoryHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := stringManupulation.ToUint32(id)
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	resp, err := h.categoryUsecase.GetByID(ctx, uint(idUint))
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	wrapper.WrapSuccessResponse(ctx, resp, "Success")
	return
}

func (h *CategoryHandler) Create(ctx *gin.Context) {
	var category request.CategoryRequest
	if err := ctx.ShouldBindJSON(&category); err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err := h.categoryUsecase.Create(ctx, &model.Category{Name: category.Name})
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	wrapper.WrapSuccessResponse(ctx, nil, "Success")
	return
}

func (h *CategoryHandler) Update(ctx *gin.Context) {
	var category request.CategoryRequest
	if err := ctx.ShouldBindJSON(&category); err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	id := ctx.Param("id")
	idUint, err := stringManupulation.ToUint32(id)
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = h.categoryUsecase.Update(ctx, &model.Category{ID: uint(idUint), Name: category.Name})
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	wrapper.WrapSuccessResponse(ctx, nil, "Success")
	return
}

func (h *CategoryHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := stringManupulation.ToUint32(id)
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = h.categoryUsecase.Delete(ctx, uint(idUint))
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	wrapper.WrapSuccessResponse(ctx, nil, "Success")
	return
}

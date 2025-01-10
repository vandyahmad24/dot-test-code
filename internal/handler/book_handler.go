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

type BookHandler struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(bookUsecase usecase.BookUsecase) *BookHandler {
	return &BookHandler{bookUsecase}
}

func (h *BookHandler) GetAll(ctx *gin.Context) {
	resp, err := h.bookUsecase.GetAll(ctx)
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	wrapper.WrapSuccessResponse(ctx, resp, "Success")
	return
}

func (h *BookHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := stringManupulation.ToUint32(id)
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	resp, err := h.bookUsecase.GetByID(ctx, uint(idUint))
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	wrapper.WrapSuccessResponse(ctx, resp, "Success")
	return
}

func (h *BookHandler) Create(ctx *gin.Context) {
	var book request.BookRequest
	if err := ctx.ShouldBindJSON(&book); err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err := h.bookUsecase.Create(ctx, &model.Book{
		Title:      book.Title,
		Author:     book.Author,
		CategoryID: book.CategoryID,
	})
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	wrapper.WrapSuccessResponse(ctx, nil, "Success")
	return
}

func (h *BookHandler) Update(ctx *gin.Context) {
	var book request.BookRequest
	if err := ctx.ShouldBindJSON(&book); err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	id := ctx.Param("id")
	idUint, err := stringManupulation.ToUint32(id)
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = h.bookUsecase.Update(ctx, &model.Book{
		ID:         uint(idUint),
		Title:      book.Title,
		Author:     book.Author,
		CategoryID: book.CategoryID,
	})
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	wrapper.WrapSuccessResponse(ctx, nil, "Success")
	return
}

func (h *BookHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := stringManupulation.ToUint32(id)
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = h.bookUsecase.Delete(ctx, uint(idUint))
	if err != nil {
		wrapper.WrapErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	wrapper.WrapSuccessResponse(ctx, nil, "Success")
	return
}

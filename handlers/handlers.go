package handlers

import (
	ogenspec "codogenerator/spec"
	"codogenerator/storage"
	"context"
	"log"
)

type Handler struct {
	repo storage.Repository
}

// NewHandler creates new struct Server
func NewHandler(r storage.Repository) *Handler {
	return &Handler{repo: r}
}

func (h *Handler) GetAllExpenses(ctx context.Context, params ogenspec.GetAllExpensesParams) (ogenspec.GetAllExpensesRes, error) {

	allExpenses, err := h.repo.GetAllExpenses(ctx, int(params.UserId))
	if err != nil {
		log.Printf("error get expenses:%s", err.Error())
		h.NewError(ctx, err)
		return nil, err
	}
	return &ogenspec.AllExpenses{Data: allExpenses}, nil
}

func (h *Handler) NewError(ctx context.Context, err error) *ogenspec.ErrorResponseStatusCode {
	// TODO
	// errorResponse := ogenspec.ErrorResponse{Message: err.Error()}
	return &ogenspec.ErrorResponseStatusCode{}
}

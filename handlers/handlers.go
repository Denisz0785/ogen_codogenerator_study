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

	allExpenses, err := h.repo.GetAllExpenses(ctx, params.UserId)
	if err != nil {
		log.Printf("error get expenses:%s", err.Error())
		h.NewError(ctx, err)
		return nil, err
	}
	return &ogenspec.AllExpenses{Data: allExpenses}, nil
}

func (h *Handler) DeleteExpense(ctx context.Context, params ogenspec.DeleteExpenseParams) error {

	err := h.repo.DeleteExpense(ctx, params.UserId, params.ExpenseID)
	if err != nil {
		log.Printf("error delete expense:%s", err.Error())
		h.NewError(ctx, err)
		return err
	}
	return nil
}

func (h *Handler) NewError(ctx context.Context, err error) *ogenspec.ErrorResponseStatusCode {
	// TODO statusCode
	message := ogenspec.OptString{}
	message.SetTo(err.Error())
	errorResponse := ogenspec.ErrorResponse{}
	errorResponse.SetMessage(message)
	errorResponseStatusCode := ogenspec.ErrorResponseStatusCode{}
	errorResponseStatusCode.SetResponse(errorResponse)

	return &errorResponseStatusCode
}

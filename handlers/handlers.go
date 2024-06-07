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

type SecurityServer struct {
	id int
}
type key string

func (s *SecurityServer) HandleApiKeyAuth(ctx context.Context, operationName string, t ogenspec.ApiKeyAuth) (context.Context, error) {
	id, err := parseToken(t.Token)
	s.id = id

	var userID key = "userID"
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, userID, s.id)
	return ctx, nil
}

func (h *Handler) GetAllExpenses(ctx context.Context) (ogenspec.GetAllExpensesRes, error) {
	var userID key = "userID"
	id := ctx.Value(userID).(int)

	allExpenses, err := h.repo.GetAllExpenses(ctx, id)
	if err != nil {
		log.Printf("error get expenses:%s", err.Error())
		h.NewError(ctx, err)
		return nil, err
	}
	return &ogenspec.AllExpenses{Data: allExpenses}, nil
}

func (h *Handler) DeleteExpense(ctx context.Context, params ogenspec.DeleteExpenseParams) error {
	var userID key = "userID"
	id := ctx.Value(userID).(int)
	err := h.repo.DeleteExpense(ctx, id, params.ExpenseID)
	if err != nil {
		log.Printf("error delete expense:%s", err.Error())
		h.NewError(ctx, err)
		return err
	}
	return nil
}

func (h *Handler) SignUp(ctx context.Context, req *ogenspec.CreateUserRequest) (ogenspec.SignUpRes, error) {

	req.Pass = hashPassword(req.Pass)
	id, err := h.repo.CreateUser(ctx, req)
	if err != nil {
		log.Printf("error crate user:%s", err.Error())
		h.NewError(ctx, err)
		return nil, err
	}
	ID := ogenspec.CreateUserResponse{
		ID: id,
	}

	return &ID, nil
}

func (h *Handler) SignIn(ctx context.Context, req *ogenspec.AuthRequest) (ogenspec.SignInRes, error) {

	token, err := h.generateToken(req.Name, req.Pass)
	if err != nil {
		log.Println(err)
		h.NewError(ctx, err)
		return nil, err
	}
	t := ogenspec.NewOptString(token)
	response := ogenspec.AuthResponse{}
	response.SetToken(t)

	return &response, nil

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

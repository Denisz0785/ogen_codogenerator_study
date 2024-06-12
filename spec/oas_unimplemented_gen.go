// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// DeleteExpense implements DeleteExpense operation.
//
// Deletes an expense.
//
// DELETE /api/expenses
func (UnimplementedHandler) DeleteExpense(ctx context.Context, params DeleteExpenseParams) error {
	return ht.ErrNotImplemented
}

// GetAllExpenses implements GetAllExpenses operation.
//
// Returns all expenses.
//
// GET /api/expenses
func (UnimplementedHandler) GetAllExpenses(ctx context.Context) (r GetAllExpensesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// SignIn implements signIn operation.
//
// Validate user.
//
// POST /auth/sign-in
func (UnimplementedHandler) SignIn(ctx context.Context, req *AuthRequest) (r SignInRes, _ error) {
	return r, ht.ErrNotImplemented
}

// SignUp implements signUp operation.
//
// Creates a new user.
//
// POST /auth/sign-up
func (UnimplementedHandler) SignUp(ctx context.Context, req *CreateUserRequest) (r SignUpRes, _ error) {
	return r, ht.ErrNotImplemented
}

package storage

import ogenspec "codogenerator/spec"

// convertExpenseToOgenspecExpense converts a storage.Expense object to
// an ogenspec.Expense object.

func convertExpenseToOgenspecExpense(expense Expense) ogenspec.Expense {
	// Create an ogenspec.Expense object and initialize its fields with
	// the corresponding fields from the input Expense object.
	return ogenspec.Expense{
		ID:            expense.Id,                            // Copy the ID field
		ExpenseTypeID: expense.ExpenseTypeId,                 // Copy the ExpenseTypeID field
		SpentMoney:    expense.SpentMoney,                    // Copy the SpentMoney field
		ReatedAt:      ogenspec.NewOptDateTime(expense.Time), // Copy the ReatedAt field
	}
}

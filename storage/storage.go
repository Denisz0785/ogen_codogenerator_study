package storage

import (
	ogenspec "codogenerator/spec"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"codogenerator/model"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	// GetTypesExpenseUser(ctx context.Context, userId int) ([]dto.ExpensesType, error)
	// GetUserId(ctx context.Context, expenseID int) (int, error)
	// IsExpenseTypeExists(ctx context.Context, expType string) (bool, error)
	// IsExpenseExists(ctx context.Context, expenseID int) (bool, error)
	// CreateExpenseType(ctx context.Context, tx pgx.Tx, expType string, userId int) (int, error)
	// GetExpenseTypeID(ctx context.Context, tx pgx.Tx, expType string) (int, error)
	// SetExpenseTimeAndSpent(ctx context.Context, tx pgx.Tx, expTypeId int, timeSpent string, spent float64) (int, error)
	// AddFileExpense(ctx context.Context, filepath string, expId int, typeFile string) error
	// CreateUserExpense(ctx context.Context, expenseData *dto.CreateExpense, userId int) (int, error)
	GetAllExpenses(ctx context.Context, userId int) ([]ogenspec.Expense, error)
	DeleteExpense(ctx context.Context, userID, expenseID int) error
	// DeleteExpense(ctx context.Context, expenseId, userId int) (int, error)
	// DeleteFile(ctx context.Context, pathFile string, expenseId int) error
	// GetExpense(ctx context.Context, userID, expenseID int) (*dto.Expense, error)
	// UpdateExpense(ctx context.Context, expenseID int, newExpense *dto.Expense) error
	CreateUser(ctx context.Context, user *ogenspec.CreateUserRequest) (int, error)
	GetUser(userName, hashPassword string) (*model.User, error)
}

// Expense represents an expense in the database.
type Expense struct {
	Id            int       `json:"id"`
	ExpenseTypeId string    `json:"expense_type_id" db:"expense_type_id"`
	Time          time.Time `json:"time" db:"reated_at"`
	SpentMoney    float64   `json:"spent_money" db:"spent_money"`
}

// ExpenseRepo create custom struct which contains descriptor of connection to database
type ExpenseRepo struct {
	conn *pgx.Conn
}

// NewExpenseRepo create ExpenseRepo
func NewExpenseRepo(conn *pgx.Conn) *ExpenseRepo {
	return &ExpenseRepo{conn: conn}
}

// ConnectToDB connects to DB
func ConnectToDB(ctx context.Context, myurl string) (*pgx.Conn, error) {
	log.Println(os.Getenv(myurl))
	conn, err := pgx.Connect(ctx, os.Getenv(myurl))
	if err != nil {
		err = fmt.Errorf("unable to connect to database: %v", err)
		return nil, err
	}
	return conn, nil
}

// GetAllExpenses retrieves all expenses for a given user from the database.

func (r *ExpenseRepo) GetAllExpenses(ctx context.Context, userId int) ([]ogenspec.Expense, error) {
	// SQL query to select expenses from the database.
	query := `
		 SELECT
			e.id,
			e.expense_type_id,
			e.reated_at,
			e.spent_money
		FROM
			users u
		JOIN expense_type et ON u.id = et.users_id
		JOIN expense e ON e.expense_type_id = et.id
		WHERE
			u.id = $1
	`

	// Execute the query and retrieve the expenses.
	log.Println("storage", userId)
	rows, err := r.conn.Query(ctx, query, userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Convert the rows to a slice of Expense structs.
	expenses, err := pgx.CollectRows(rows, pgx.RowToStructByName[Expense])
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Convert the Expense structs to ogenspec.Expense structs.
	ogenTypeExpense := make([]ogenspec.Expense, len(expenses))

	for i, v := range expenses {
		ogenTypeExpense[i] = convertExpenseToOgenspecExpense(v)
	}

	return ogenTypeExpense, nil
}

func (r *ExpenseRepo) DeleteExpense(ctx context.Context, userID, expenseID int) error {
	var idDeleteExpense int
	query := `DELETE FROM expense WHERE id IN (select e.id from users u join  expense_type et
		ON u.id=et.users_id join expense e on e.expense_type_id=et.id where u.id=$1
		and e.id=$2) returning id`
	//_, err := r.conn.Exec(ctx, query, userId, expenseId)
	err := r.conn.QueryRow(ctx, query, userID, expenseID).Scan(&idDeleteExpense)
	if err != nil {
		if idDeleteExpense == 0 {
			log.Println("id expense does not exist")
			return errors.New("expense does not exist")
		}
		log.Println(err)
		return err
	}
	return nil
}

func (r *ExpenseRepo) CreateUser(ctx context.Context, user *ogenspec.CreateUserRequest) (int, error) {
	var userId int

	query := fmt.Sprintf("INSERT INTO %s (name,surname,login,pass,email) values ($1,$2,$3,$4,$5) RETURNING id", "users")
	err := r.conn.QueryRow(ctx, query, user.Name, user.Surname, user.Login, user.Pass, user.Email).Scan(&userId)
	if err != nil {
		log.Printf("error create user:%s", err.Error())
		return -1, err
	}
	return userId, nil
}

func (r *ExpenseRepo) GetUser(userName, hashPassword string) (*model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1 and pass=$2", "users")
	err := r.conn.QueryRow(context.Background(), query, userName, hashPassword).Scan(&user.Id)
	if err != nil {
		log.Println("error get user", err.Error())
		return nil, err
	}
	return &user, nil
}

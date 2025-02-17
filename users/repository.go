package users

import (
	dbq "asyncgo/database/sqlc"
	"asyncgo/utils"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type UserRepositoryImpl struct {
	queries *dbq.Queries
}

func NewUserRepository(queries *dbq.Queries) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		queries: queries,
	}
}

func (r *UserRepositoryImpl) CreateUserRepo(ctx context.Context, arg dbq.CreateUserParams) (*dbq.User, error) {

	NanoID, NanoIDErr := utils.NanoIDS()
	if NanoIDErr != nil {
		log.Println("nano_id geenration error")
		return nil, fmt.Errorf("failed to generate NanoID: %w", NanoIDErr)

	}

	password, err := utils.HashPassword(arg.HashedPassword)
	if err != nil {
		log.Println("❌ Password hashing failed:", err)
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := dbq.CreateUserParams{
		ID:             NanoID,
		Email:          arg.Email,
		HashedPassword: password,
		CreatedAt:      arg.CreatedAt,
	}

	createdUser, err := r.queries.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return &createdUser, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*dbq.User, error) {

	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found: %w", err) // Handle missing user explicitly
		}
		return nil, fmt.Errorf("database error retrieving user: %w", err)
	}
	return &user, nil

}

func (r *UserRepositoryImpl) GetUserById(ctx context.Context, id string) (*dbq.User, error) {

	user, err := r.queries.GetUserById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found: %w", err) // Handle missing user explicitly
		}
		return nil, fmt.Errorf("database error retrieving user: %w", err)
	}
	return &user, nil

}

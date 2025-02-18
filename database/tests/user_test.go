package dbq

import (
	dbq "asyncgo/database/sqlc"
	"asyncgo/utils"
	"context"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

const testPassword = "Test@1234" // Generic test password

func CreateTestUser(t *testing.T) dbq.CreateUserParams {
	t.Helper() // Marks function as a helper for better debugging

	email := GenerateRandomEmail()
	//password := GenerateRandomPassword(8)

	hashedPassword, err := utils.HashPassword(testPassword)
	require.NoError(t, err)

	id, err := utils.NanoIDS()
	require.NoError(t, err)

	userParams := dbq.CreateUserParams{
		ID:             id,
		Email:          email,
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now(),
	}

	// 🟢 **Insert test user into the database**
	_, err = testQueries.CreateUser(context.Background(), userParams)
	require.NoError(t, err)

	return userParams
}

func TestCreateUser(t *testing.T) {

	// 🟢 **Ensure fresh test database before running test**
	err := TearDownTestDb()
	require.NoError(t, err)

	// 🟢 **Create test user**
	userParams := CreateTestUser(t)

	// 🟢 **Retrieve created user**
	createdUser, err := testQueries.GetUserById(context.Background(), userParams.ID)
	require.NoError(t, err)
	require.NotNil(t, createdUser)

	// 🟢 **Validate user data**
	require.Equal(t, userParams.Email, createdUser.Email)
	require.NotEqual(t, userParams.HashedPassword, testPassword, "Password should be hashed")

	require.Equal(t, userParams.ID, createdUser.ID)

	// 🟢 **Clean up after test**
	err = TearDownTestDb()
	require.NoError(t, err)
}

func TestGetUserByEmail(t *testing.T) {
	// 🟢 **Ensure fresh test database before running test**
	err := TearDownTestDb()
	require.NoError(t, err)

	// 🟢 **Create test user**
	userParams := CreateTestUser(t)

	// 🟢 **Retrieve user by email**
	user, err := testQueries.GetUserByEmail(context.Background(), userParams.Email)
	require.NoError(t, err)
	require.NotNil(t, user)

	// 🟢 **Validate user data**
	require.Equal(t, userParams.Email, user.Email)
	require.Equal(t, userParams.ID, user.ID)

	// 🟢 **Clean up after test**
	err = TearDownTestDb()
	require.NoError(t, err)
}

func TestGetUserById(t *testing.T) {
	// 🟢 **Ensure fresh test database before running test**
	err := TearDownTestDb()
	require.NoError(t, err)

	// 🟢 **Create test user**
	userParams := CreateTestUser(t)

	// 🟢 **Retrieve user by ID**
	user, err := testQueries.GetUserById(context.Background(), userParams.ID)
	require.NoError(t, err)
	require.NotNil(t, user)

	// 🟢 **Validate user data**
	require.Equal(t, userParams.ID, user.ID)
	require.Equal(t, userParams.Email, user.Email)

	// 🟢 **Clean up after test**
	err = TearDownTestDb()
	require.NoError(t, err)
}

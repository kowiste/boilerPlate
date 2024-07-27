package mysql

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kowiste/boilerplate/src/model/user"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock := NewMockDB()
	defer db.DB()

	repo := MySQL{db: db}
	ctx := context.Background()

	user := &user.User{
		ID:       "1",
		Name:     "John",
		LastName: "Doe",
		Age:      30,
	}

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `users` \\(`id`,`name`,`last_name`,`age`\\) VALUES \\(\\?,\\?,\\?,\\?\\)$").
		WithArgs(user.ID, user.Name, user.LastName, user.Age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	id, err := repo.CreateUser(ctx, user)

	assert.NoError(t, err)
	assert.Equal(t, "1", id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUsers(t *testing.T) {
	db, mock := NewMockDB()
	defer db.DB()

	repo := MySQL{db: db}
	ctx := context.Background()

	rows := sqlmock.NewRows([]string{"id"}).AddRow("1")

	mock.ExpectQuery("^SELECT \\* FROM `users`").WillReturnRows(rows)

	users, err := repo.Users(ctx, &user.FindUsersInput{})

	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "1", users[0].ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestUserByID(t *testing.T) {
	db, mock := NewMockDB()
	defer db.DB()

	repo := MySQL{db: db}
	ctx := context.Background()

	// Create a mock row for the user with ID "1"
	rows := sqlmock.NewRows([]string{"id"}).AddRow("1")

	// Update the ExpectQuery to match the actual query executed
	mock.ExpectQuery(`(?i)^SELECT \* FROM `+"`users`"+` WHERE id=\? ORDER BY `+"`users`.`id`"+` LIMIT \?`).
		WithArgs("1", 1). // Ensure this matches the arguments being passed
		WillReturnRows(rows)

	user, err := repo.UserByID(ctx, "1")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "1", user.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestUpdateUser(t *testing.T) {
        db, mock := NewMockDB()
        defer db.DB()

        repo := MySQL{db: db}
        ctx := context.Background()

        user := &user.User{
                ID:   "1",
                Name: "New Name",
        }

        mock.ExpectBegin()

        // Corrected regex to match the actual SQL statement
        mock.ExpectExec("^UPDATE `users` SET `name` = \\? WHERE id = \\?$").
                WithArgs(user.Name, user.ID).
                WillReturnResult(sqlmock.NewResult(1, 1))

        mock.ExpectCommit()

        err := repo.UpdateUser(ctx, user)

        assert.NoError(t, err)
        assert.NoError(t, mock.ExpectationsWereMet())
}





func TestDeleteUser(t *testing.T) {
	db, mock := NewMockDB()
	defer db.DB()

	repo := MySQL{db: db}
	ctx := context.Background()

	mock.ExpectBegin()
	mock.ExpectExec("^DELETE FROM `users` WHERE id = \\?").WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.DeleteUser(ctx, "1")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

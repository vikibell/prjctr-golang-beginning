package user

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.23"))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return gormDB, mock, nil
}

func TestFindOneById(t *testing.T) {
	db, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to set up mock database: %v", err)
	}
	repo := &Repository{db: db}

	expectedUser := User{
		ID:         1,
		Name:       "Alice",
		Surname:    "Johnson",
		Email:      "alice@example.com",
		Age:        30,
		Sex:        "Female",
		City:       "New York",
		TaxiCount:  5,
		Profession: "Engineer",
	}
	rows := sqlmock.NewRows([]string{
		"id", "name", "surname", "email", "age", "sex", "city", "taxi_count", "profession",
	}).AddRow(
		expectedUser.ID, expectedUser.Name, expectedUser.Surname, expectedUser.Email,
		expectedUser.Age, expectedUser.Sex, expectedUser.City, expectedUser.TaxiCount, expectedUser.Profession,
	)

	mock.ExpectQuery("SELECT \\* FROM `users` WHERE `users`.`id` = \\? ORDER BY `users`.`id` LIMIT \\?").
		WithArgs(1, 1).
		WillReturnRows(rows)

	user, err := repo.FindOneById(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestFindAll(t *testing.T) {
	db, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to set up mock database: %v", err)
	}
	repo := &Repository{db: db}

	page := 2
	limit := 2
	offset := (page - 1) * limit

	expectedUsers := []User{
		{ID: 2, Name: "Vika", Surname: "Tkach", Email: "tk@example.com", Age: 31},
		{ID: 3, Name: "Vika", Surname: "Suslova", Email: "sus@example.com", Age: 30},
	}
	expectedTotal := int64(5)

	mock.ExpectQuery("(?i)SELECT count\\(\\*\\) FROM `users`").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expectedTotal))

	rows := sqlmock.NewRows([]string{"id", "name", "surname", "email", "age"}).
		AddRow(2, "Vika", "Tkach", "tk@example.com", 31).
		AddRow(3, "Vika", "Suslova", "sus@example.com", 30)

	mock.ExpectQuery("SELECT \\* FROM `users` LIMIT \\? OFFSET \\?").
		WithArgs(limit, offset).
		WillReturnRows(rows)

	users, total, err := repo.FindAll(page, limit)
	assert.NoError(t, err)
	assert.Equal(t, expectedTotal, total)
	assert.Equal(t, expectedUsers, users)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

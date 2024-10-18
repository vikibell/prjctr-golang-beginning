package statistic

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	t.Helper()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.23"))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	return gormDB, mock
}

func TestFindOneByID(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := &Repository{db: db}

	expectedStatistic := Statistic{
		ID:           1,
		City:         "Kyiv",
		AverageTrips: 10,
		AgeRange:     Elderly,
	}

	rows := sqlmock.NewRows([]string{
		"id", "city", "average_trips", "age_range",
	}).AddRow(
		expectedStatistic.ID, expectedStatistic.City, expectedStatistic.AverageTrips, expectedStatistic.AgeRange,
	)

	mock.ExpectQuery("SELECT \\* FROM `statistics` WHERE `statistics`.`id` = \\? ORDER BY `statistics`.`id` LIMIT \\?").
		WithArgs(1, 1).
		WillReturnRows(rows)

	statistic, err := repo.FindOneByID(1)
	require.NoError(t, err)
	assert.Equal(t, expectedStatistic, statistic)
}

func TestFindAll(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := &Repository{db: db}

	page := 2
	limit := 2
	offset := (page - 1) * limit

	expectedUsers := []Statistic{
		{ID: 2, City: "Chernihiv", AverageTrips: 35, AgeRange: Adult},
		{ID: 3, City: "Kyiv", AverageTrips: 55, AgeRange: Adult},
	}
	expectedTotal := int64(5)

	mock.ExpectQuery("(?i)SELECT count\\(\\*\\) FROM `statistics`").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expectedTotal))

	rows := sqlmock.NewRows([]string{"id", "city", "average_trips", "age_range"}).
		AddRow(2, "Chernihiv", 35, Adult).
		AddRow(3, "Kyiv", 55, Adult)

	mock.ExpectQuery("SELECT \\* FROM `statistics` LIMIT \\? OFFSET \\?").
		WithArgs(limit, offset).
		WillReturnRows(rows)

	statistics, total, err := repo.FindAll(page, limit)
	require.NoError(t, err)
	assert.Equal(t, expectedTotal, total)
	assert.Equal(t, expectedUsers, statistics)
}

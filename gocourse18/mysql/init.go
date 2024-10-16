package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/statistic"
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
)

const dsn = "vika:123@tcp(127.0.0.1:3320)/db?parseTime=true"

var usersSchema = `
CREATE TABLE IF NOT EXISTS users
(
    id         bigint unsigned auto_increment primary key,
    name       longtext         not null,
    surname    longtext         not null,
    email      longtext         not null,
    age        bigint default 0 not null,
    sex        longtext         not null,
    city       longtext         not null,
    taxi_count bigint default 0 null,
    profession longtext         null
);`

var statisticsSchema = `
CREATE TABLE IF NOT EXISTS statistics
(   
    id            bigint unsigned auto_increment primary key,
    city          longtext         not null,
    average_trips bigint default 0 null,
    age_range     int    default 0 null
);`

func init() {
	populateDB()
}

func GetConnection() *gorm.DB {
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return gormdb
}

func populateDB() {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("Failed to connect database")
	}

	db.MustExec(usersSchema)
	populateUsers(db)
	db.MustExec(statisticsSchema)
	populateStatistics(db)
	db.Close()
}

func populateUsers(db *sqlx.DB) {
	query := `TRUNCATE TABLE users`
	_, trancErr := db.Exec(query)
	if trancErr != nil {
		slog.Error("Failed to truncate table", "error", trancErr)
		return
	}

	users := service.GetUsersFromFile()

	tx := db.MustBegin()
	_, insertErr := tx.NamedExec(`
		INSERT INTO users (name, surname, email, age, sex, city, taxi_count, profession) 
		VALUES (:name, :surname, :email, :age, :sex, :city, :taxi_count, :profession)`,
		users)
	if insertErr != nil {
		slog.Error("Failed to insert into table", "error", insertErr)
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		slog.Error("Failed to commit changes", "error", commitErr)
		return
	}
}

func populateStatistics(db *sqlx.DB) {
	query := `TRUNCATE TABLE statistics`
	_, trancErr := db.Exec(query)
	if trancErr != nil {
		slog.Error("Failed to truncate table", "error", trancErr)
		return
	}

	var statisticsResult []statistic.Statistic
	err := db.Select(&statisticsResult, `
		SELECT city,
		  CASE
           WHEN age BETWEEN 1 AND 8 THEN 1
           WHEN age BETWEEN 9 AND 21 THEN 2
           WHEN age BETWEEN 22 AND 40 THEN 3
           WHEN age BETWEEN 41 AND 60 THEN 4
           WHEN age BETWEEN 61 AND 100 THEN 5
          END AS age_range,
		  round(avg(taxi_count), 0) as average_trips
		FROM users
		GROUP BY city, age_range
		ORDER BY city;
       `)
	if err != nil {
		slog.Error("Failed to select", "error", err)
		return
	}

	tx := db.MustBegin()

	_, insertErr := tx.NamedExec(`
		INSERT INTO statistics (city, average_trips, age_range) 
		VALUES (:city, :average_trips, :age_range)`,
		statisticsResult)
	if insertErr != nil {
		slog.Error("Failed to insert into table", "error", insertErr)
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		slog.Error("Failed to commit changes", "error", commitErr)
		return
	}
}

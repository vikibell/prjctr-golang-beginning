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

const DSN = "vika:123@tcp(127.0.0.1:3320)/db?parseTime=true"

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

type DB struct {
	Connection *gorm.DB
}

func NewDB(dsn string) *DB {
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &DB{
		Connection: gormdb,
	}
}

func init() {
	populateDB()
}

func populateDB() {
	db, err := sqlx.Connect("mysql", DSN)
	if err != nil {
		panic("Failed to connect database")
	}

	db.MustExec(usersSchema)
	err = populateUsers(db)
	if err != nil {
		slog.Error("Failed to populate users table", "error", err)
		return
	}

	db.MustExec(statisticsSchema)
	err = populateStatistics(db)
	if err != nil {
		slog.Error("Failed to populate statistics table", "error", err)
		return
	}

	db.Close()
}

func populateUsers(db *sqlx.DB) error {
	query := `TRUNCATE TABLE users`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	users := service.GetUsersFromFile()

	tx := db.MustBegin()
	_, err = tx.NamedExec(`
		INSERT INTO users (name, surname, email, age, sex, city, taxi_count, profession) 
		VALUES (:name, :surname, :email, :age, :sex, :city, :taxi_count, :profession)`,
		users)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func populateStatistics(db *sqlx.DB) error {
	query := `TRUNCATE TABLE statistics`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	var statisticsResult []statistic.Statistic
	err = db.Select(&statisticsResult, `
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
		return err
	}

	tx := db.MustBegin()

	_, err = tx.NamedExec(`
		INSERT INTO statistics (city, average_trips, age_range) 
		VALUES (:city, :average_trips, :age_range)`,
		statisticsResult)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

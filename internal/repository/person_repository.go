package repository

import (
	"database/sql"
	"fmt"
	"junior/internal/config"
	"junior/internal/model"
	"junior/pkg/logger"

	_ "github.com/lib/pq"
)

type PersonRepository struct {
	db *sql.DB
}

func NewPersonRepository() *PersonRepository {
	// Получаем значения из конфигурации
	host := config.GetEnv("DB_HOST")
	port := config.GetEnv("DB_PORT")
	user := config.GetEnv("DB_USER")
	password := config.GetEnv("DB_PASSWORD")
	dbname := config.GetEnv("DB_NAME")



	// Проверяем, что все необходимые значения установлены
	if host == "" || port == "" || dbname == "" {
		logger.Log.Error("Missing required database configuration")
		panic("Missing required database configuration")
	}

	// Формируем строку подключения
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	logger.Log.Debug("Connecting to database with DSN: ", dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Log.Error("Failed to connect to database: ", err)
		panic(err)
	}

	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		logger.Log.Error("Failed to ping database: ", err)
		panic(err)
	}

	logger.Log.Info("Successfully connected to database")
	return &PersonRepository{db: db}
}

func (r *PersonRepository) Create(p model.Person) (model.Person, error) {
	logger.Log.Debug("Creating new person in database")

	err := r.db.QueryRow(
		"INSERT INTO persons (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, p.Nationality).Scan(&p.ID)
	if err != nil {
		logger.Log.Error("Failed to create person in database: ", err)
		return p, err
	}

	logger.Log.Debug("Successfully created person with ID: ", p.ID)
	return p, nil
}

func (r *PersonRepository) GetAll() ([]model.Person, error) {
	logger.Log.Debug("Getting all persons from database")

	rows, err := r.db.Query("SELECT id, name, surname, patronymic, age, gender, nationality FROM persons")
	if err != nil {
		logger.Log.Error("Failed to get persons from database: ", err)
		return nil, err
	}
	defer rows.Close()

	var people []model.Person
	for rows.Next() {
		var p model.Person
		if err := rows.Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &p.Nationality); err != nil {
			logger.Log.Error("Failed to scan person row: ", err)
			return nil, err
		}
		people = append(people, p)
	}

	logger.Log.Debug("Successfully retrieved ", len(people), " persons from database")
	return people, nil
}

func (r *PersonRepository) GetByID(id int) (model.Person, error) {
	logger.Log.Debug("Getting person by ID from database: ", id)

	var p model.Person
	err := r.db.QueryRow("SELECT id, name, surname, patronymic, age, gender, nationality FROM persons WHERE id=$1", id).
		Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &p.Nationality)
	if err != nil {
		logger.Log.Error("Failed to get person by ID from database: ", err)
		return p, err
	}

	logger.Log.Debug("Successfully retrieved person with ID: ", id)
	return p, nil
}

func (r *PersonRepository) Update(id int, p model.Person) (model.Person, error) {
	logger.Log.Debug("Updating person in database with ID: ", id)

	_, err := r.db.Exec("UPDATE persons SET name=$1, surname=$2, patronymic=$3, age=$4, gender=$5, nationality=$6 WHERE id=$7",
		p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, p.Nationality, id)
	if err != nil {
		logger.Log.Error("Failed to update person in database: ", err)
		return p, err
	}

	p.ID = id
	logger.Log.Debug("Successfully updated person with ID: ", id)
	return p, nil
}

func (r *PersonRepository) Delete(id int) error {
	logger.Log.Debug("Deleting person from database with ID: ", id)

	_, err := r.db.Exec("DELETE FROM persons WHERE id=$1", id)
	if err != nil {
		logger.Log.Error("Failed to delete person from database: ", err)
		return err
	}

	logger.Log.Debug("Successfully deleted person with ID: ", id)
	return nil
}

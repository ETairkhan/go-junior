package repository

import (
    "database/sql"
    "junior/internal/model"
    "junior/internal/config"
    _ "github.com/lib/pq"
)

type PersonRepository struct {
    db *sql.DB
}

func NewPersonRepository() *PersonRepository {
    dsn := "host=" + config.GetEnv("DB_HOST") +
        " port=" + config.GetEnv("DB_PORT") +
        " user=" + config.GetEnv("DB_USER") +
        " password=" + config.GetEnv("DB_PASSWORD") +
        " dbname=" + config.GetEnv("DB_NAME") + " sslmode=disable"

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        panic(err)
    }

    return &PersonRepository{db: db}
}

func (r *PersonRepository) Create(p model.Person) (model.Person, error) {
    err := r.db.QueryRow(
        "INSERT INTO persons (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
        p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, p.Nationality).Scan(&p.ID)
    return p, err
}

func (r *PersonRepository) GetAll() ([]model.Person, error) {
    rows, err := r.db.Query("SELECT id, name, surname, patronymic, age, gender, nationality FROM persons")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var people []model.Person
    for rows.Next() {
        var p model.Person
        rows.Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &p.Nationality)
        people = append(people, p)
    }
    return people, nil
}

func (r *PersonRepository) GetByID(id int) (model.Person, error) {
    var p model.Person
    err := r.db.QueryRow("SELECT id, name, surname, patronymic, age, gender, nationality FROM persons WHERE id=$1", id).
        Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &p.Nationality)
    return p, err
}

func (r *PersonRepository) Update(id int, p model.Person) (model.Person, error) {
    _, err := r.db.Exec("UPDATE persons SET name=$1, surname=$2, patronymic=$3, age=$4, gender=$5, nationality=$6 WHERE id=$7",
        p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, p.Nationality, id)
    p.ID = id
    return p, err
}

func (r *PersonRepository) Delete(id int) error {
    _, err := r.db.Exec("DELETE FROM persons WHERE id=$1", id)
    return err
}

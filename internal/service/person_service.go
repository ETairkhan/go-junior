package service

import (
    "junior/internal/client"
    "junior/internal/model"
    "junior/internal/repository"
)

type PersonService struct {
    repo *repository.PersonRepository
}

func NewPersonService(repo *repository.PersonRepository) *PersonService {
    return &PersonService{
        repo: repo,
    }
}

func (s *PersonService) CreatePerson(p model.Person) (model.Person, error) {
    age, _ := client.GetAge(p.Name)
    gender, _ := client.GetGender(p.Name)
    nationality, _ := client.GetNationality(p.Name)

    p.Age = age
    p.Gender = gender
    p.Nationality = nationality

    return s.repo.Create(p)
}

func (s *PersonService) GetPeople() ([]model.Person, error) {
    return s.repo.GetAll()
}

func (s *PersonService) GetPersonByID(id int) (model.Person, error) {
    return s.repo.GetByID(id)
}

func (s *PersonService) UpdatePerson(id int, p model.Person) (model.Person, error) {
    return s.repo.Update(id, p)
}

func (s *PersonService) DeletePerson(id int) error {
    return s.repo.Delete(id)
}

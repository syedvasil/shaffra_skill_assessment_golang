package user

import (
	repoModels "github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/repositories/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockery --name=Repository --case underscore
type Repository interface {
	CreateUser(user repoModels.User) error
	GetUsers(filter interface{}, offset, limit int) ([]repoModels.User, error)
	GetUserByID(id primitive.ObjectID) (repoModels.User, error)
	UpdateUser(id primitive.ObjectID, user repoModels.User) error
	DeleteUser(id primitive.ObjectID) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) CreateUser(user repoModels.User) error {
	return s.repo.CreateUser(user)
}

func (s *Service) GetUsers(page, limit int) ([]repoModels.User, error) {
	offset := (page - 1) * limit
	filter := bson.M{"deleted_at": bson.M{"$exists": false}}
	return s.repo.GetUsers(filter, offset, limit)
}

func (s *Service) GetUserByID(id primitive.ObjectID) (repoModels.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *Service) UpdateUser(id primitive.ObjectID, user repoModels.UserUpdateReq) error {

	oldUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	oldUser.ForUpdate(user)

	return s.repo.UpdateUser(id, oldUser)
}

func (s *Service) DeleteUser(id primitive.ObjectID) error {
	return s.repo.DeleteUser(id)
}

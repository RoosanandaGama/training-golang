package slice

import "training-golang/session-4-sample-separate-layer/step-4/entity"

type IUserRepository interface {
	GetAllUsers() []entity.User
}

type userRepository struct {
	db     []entity.User
	nextID int
}

func NewUserRepository(db []entity.User) IUserRepository {
	return &userRepository{
		db:     db,
		nextID: 1,
	}
}

func (r *userRepository) GetAllUsers() []entity.User {
	return r.db
}

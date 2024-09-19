package postgresgorm

import (
	"context"
	"errors"
	"log"
	"session-16-crud-user-docker-compose/entity"

	"gorm.io/gorm"
)

type GormDBIface interface {
	WithContext(ctx context.Context) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

// IUserRepository mendefinisikan interface untuk repositry pengguna
type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

// userRepository adalah implementasi dari IUserRepository yang menggunakan slice untuk menyimpan data pengguna
type userRepository struct {
	db GormDBIface
}

// newUserRepository membuat instance baru dari UserRepository
func NewUserRepository(db GormDBIface) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		log.Printf("Error creating user: %v\n", err)
		return entity.User{}, err
	}

	return *user, nil

}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Select("id", "name", "email", "password", "created_at", "updated_at").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, err
		}

		log.Printf("Error getting user by id: %v\n", err)
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	var existingUser entity.User
	if err := r.db.WithContext(ctx).Select("id", "name", "email", "password", "created_at", "updated_at").First(&existingUser, id).Error; err != nil {
		log.Printf("Error finding user by id: %v\n", err)
		return entity.User{}, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	if err := r.db.WithContext(ctx).Save(&existingUser).Error; err != nil {
		log.Printf("Error updating user: %v\n", err)
		return entity.User{}, err
	}

	return existingUser, nil

}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		log.Printf("Error deleting user: %\n", err)
		return err
	}

	return nil
}

// GetAllUsers mengembalikan semua pengguna
func (r *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	if err := r.db.WithContext(ctx).Find(&users).Select("id", "name", "email", "password", "created_at", "updated_at").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return users, nil
}

package repositories

import (
	"errors"

	"github.com/Danangoffic/go-merce/app/models"
)

func (r *RepositoryApp) GetUserByID(user models.User) (models.User, error) {
	if err := r.db.Find(&user).Order("ID DESC").Error; err != nil {
		return models.User{}, err
	}

	if user.ID == 0 {
		return user, errors.New("user data not found")
	}
	return user, nil
}

func (r *RepositoryApp) CreateUser(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

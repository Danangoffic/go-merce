package services

import "github.com/Danangoffic/go-merce/app/models"

func (s *service) GetUserByID(ID uint) (models.User, error) {
	user, err := s.repositories.GetUserByID(models.User{ID: ID})
	if err != nil {
		return user, err
	}
	return user, nil
}

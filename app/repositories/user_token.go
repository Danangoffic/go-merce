package repositories

import (
	"time"

	"github.com/Danangoffic/go-merce/app/models"
	"github.com/Danangoffic/go-merce/app/utils"
)

func (p *RepositoryApp) DetailData(token string) (models.UserToken, error) {
	var userToken models.UserToken
	err := p.db.Where(&models.UserToken{Token: token}).Find(userToken).Error
	if err != nil {
		return models.UserToken{}, err
	}
	return userToken, nil
}

func (p *RepositoryApp) GetUserTokenData(token string) (models.User, error) {
	var userToken models.UserToken
	err := p.db.Where(&models.UserToken{Token: token}).Find(userToken).Error
	if err != nil {
		return models.User{}, err
	}
	return userToken.User, nil
}

func (p *RepositoryApp) GenerateUserToken(userToken models.UserToken) (string, error) {
	token := utils.GenerateToken(32)
	hashedToken := utils.HashTokenString(token)
	// Simpan data token, expiredAt, dan UpdatedAt ke database
	// expired at dalam 1 jam
	expired := time.Minute * 60
	currentTime := time.Now()
	userToken.Token = hashedToken
	userToken.ExpiredAt = currentTime.Add(expired)
	userToken.UpdatedAt = currentTime

	err := p.db.Save(userToken).Error
	if err != nil {
		return "", err
	}

	return token, nil
}

func (p *RepositoryApp) UpdateTokenActive(userToken models.UserToken) (string, error) {
	token := utils.GenerateToken(32)
	hashedToken := utils.HashTokenString(token)
	// Simpan data token, expiredAt, dan UpdatedAt ke database
	// expired at dalam 1 jam
	expired := time.Minute * 60
	currentTime := time.Now()
	userToken.Token = hashedToken
	userToken.ExpiredAt = currentTime.Add(expired)
	userToken.UpdatedAt = currentTime

	err := p.db.Save(userToken).Error
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *RepositoryApp) DeleteExpiredToken(userToken models.UserToken) {
	r.db.Delete(&userToken)
}

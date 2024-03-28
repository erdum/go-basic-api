package auth

import (
	"gorm.io/gorm"
)

type FirebaseAuthService struct {
	db *gorm.DB
}

func NewFirebaseAuth(db *gorm.DB) AuthService {
	return &FirebaseAuthService{db: db}
}

func (auth *FirebaseAuthService) AuthenticateWithThirdParty(
	idToken string,
) (bool, error) {
	return true, nil
}

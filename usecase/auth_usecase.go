package usecase

import (
	"errors"

	"github.com/josevitorrodriguess/productsAPI/repository"
	"github.com/josevitorrodriguess/productsAPI/services"
)

type AuthUSeCase struct {
	merchantRepo repository.MerchantRepository
}

func NewAuthUsecase(merchantRepo repository.MerchantRepository) AuthUSeCase {
	return AuthUSeCase{merchantRepo}
}

func (au *AuthUSeCase) Login(email, password string) (string, error) {
    merchant, err := au.merchantRepo.FindByEmail(email)
    if err != nil {
        return "", err
    }

    if merchant == nil {
        return "", errors.New("cannot find merchant")
    }

    hashedPassword := services.SHA256Encoder(password)

    if merchant.Password != hashedPassword {
        return "", errors.New("invalid credentials")
    }

    token, err := services.NewJWTService().GenerateToken(merchant.ID)
    if err != nil {
        return "", err
    }

    return token, nil
}

package usecase

import (
	"errors"

	"github.com/josevitorrodriguess/productsAPI/repository"
	"github.com/josevitorrodriguess/productsAPI/services"
)

type AuthMerchantUseCase struct {
	merchantRepo repository.MerchantRepository
}

func NewAuthMerchantUsecase(merchantRepo repository.MerchantRepository) AuthMerchantUseCase {
	return AuthMerchantUseCase{merchantRepo}
}

func (au *AuthMerchantUseCase) Login(email, password string) (string, error) {
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

type AuthClientUsecase struct {
	clientRepo repository.ClientRepository
}

func NewAuthClientUsecase(clientRepo repository.ClientRepository) AuthClientUsecase {
	return AuthClientUsecase{clientRepo}
}

func (au *AuthClientUsecase) LoginClient(email, password string) (string, error) {
	client, err := au.clientRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if client == nil {
		return "", errors.New("cannot find merchant")
	}

	hashedPassword := services.SHA256Encoder(password)

	if client.Password != hashedPassword {
		return "", errors.New("invalid credentials")
	}

	token, err := services.NewJWTService().GenerateToken(client.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

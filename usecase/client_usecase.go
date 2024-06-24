package usecase

import (
	"fmt"

	"github.com/josevitorrodriguess/productsAPI/model"
	"github.com/josevitorrodriguess/productsAPI/repository"
	"github.com/josevitorrodriguess/productsAPI/services"
)

type ClientUsecase struct {
	repository repository.ClientRepository
}

func NewClientUsecase(repo repository.ClientRepository) ClientUsecase {
	return ClientUsecase{
		repository: repo,
	}
}

func (cu *ClientUsecase) GetClients() ([]model.Client, error) {
	return cu.repository.GetClients()
}

func (cu *ClientUsecase) GetClientByID(id_client int) (*model.Client, error) {
	merchant, err := cu.repository.GetClientByID(id_client)
	if err != nil {
		return nil, err
	}

	return merchant, nil
}

func (cu *ClientUsecase) CreateClient(client model.Client) (model.Client, error) {
	
	isValidEmail, err := services.ValidateEmail(client.Email)
	if err != nil {
		return model.Client{}, fmt.Errorf("failed to validate email: %v", err)
	}
	if !isValidEmail {
		return model.Client{}, fmt.Errorf("invalid email format: %s", client.Email)
	}

	
	isValidCPF, err := services.ValidateCPF(client.CPF)
	if err != nil {
		return model.Client{}, fmt.Errorf("failed to validate CPF: %v", err)
	}
	if !isValidCPF {
		return model.Client{}, fmt.Errorf("invalid CPF: %s", client.CPF)
	}

	
	client.Password = services.SHA256Encoder(client.Password)

	
	clientID, err := cu.repository.CreateClient(client)
	if err != nil {
		return model.Client{}, fmt.Errorf("failed to create client in repository: %v", err)
	}

	
	client.ID = clientID

	return client, nil
}



func (cu *ClientUsecase) DeleteClient(id_client int) error {
	return cu.repository.DeleteClient(id_client)
}

package usecase

import (
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

	clientId, err := cu.repository.CreateClient(client)
	if err != nil {
		return model.Client{}, err
	}

	client.Password = services.SHA256Encoder(client.Password)
	client.ID = clientId

	return client, nil
}

func (cu *ClientUsecase) DeleteClient(id_client int) error {
	return  cu.repository.DeleteClient(id_client)
}
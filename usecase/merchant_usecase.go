package usecase

import (
	"github.com/josevitorrodriguess/productsAPI/model"
	"github.com/josevitorrodriguess/productsAPI/repository"
)

type MerchantUsecase struct {
	repository repository.MerchantRepository
}

func NewMerchantUsecase(repo repository.MerchantRepository) MerchantUsecase {
	return MerchantUsecase{
		repository: repo,
	}
}

func (mu *MerchantUsecase) GetMerchants() ([]model.Merchant, error) {
	return mu.repository.GetMerchants()
}

func (mu *MerchantUsecase) GetMerchantById(id_merchant int) (*model.Merchant, error) {
	merchant, err := mu.repository.GetMerchantByID(id_merchant)
	if err != nil {
		return  nil, err
	}

	return merchant, nil
}

func (mu *MerchantUsecase) CreateMerchant(merchant model.Merchant) (model.Merchant, error) {

	merchantId, err := mu.repository.CreateMerchant(merchant)
	if err != nil {
		return model.Merchant{}, err
	}

	merchant.ID = merchantId

	return merchant, nil
}

func (mu *MerchantUsecase) DeleteMerchant(id_merchant int) error {
	return  mu.repository.DeleteMerchant(id_merchant)
}
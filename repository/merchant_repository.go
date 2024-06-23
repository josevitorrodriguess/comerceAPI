package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/josevitorrodriguess/productsAPI/model"
	"github.com/josevitorrodriguess/productsAPI/services"
)

type MerchantRepository struct {
	connection *sql.DB
}

func NewMerchantRepository(conn *sql.DB) MerchantRepository {
	return MerchantRepository{
		connection: conn,
	}
}

func ( mr *MerchantRepository) GetMerchants() ([]model.Merchant, error) {
	query := "SELECT id_merchant, name, product_type, email FROM merchant"
	rows, err := mr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var merchantList []model.Merchant

	for rows.Next() {
		var merchantObj model.Merchant
		err := rows.Scan(
			&merchantObj.ID,
			&merchantObj.Name,
			&merchantObj.TypeProduct,
			&merchantObj.Email,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		
		merchantList = append(merchantList, merchantObj)
	}
	return merchantList, nil
}

func (mr *MerchantRepository) GetMerchantByID(merchant_id int) (*model.Merchant, error) {
	
	query := "SELECT id_merchant, name, product_type, email FROM merchant WHERE id_merchant=?"
	row := mr.connection.QueryRow(query, merchant_id)
	
	var merchant model.Merchant
	err := row.Scan(
			&merchant.ID,
			&merchant.Name,
			&merchant.TypeProduct,
			&merchant.Email,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}
	return &merchant, nil
}

func (mr *MerchantRepository) CreateMerchant(merchant model.Merchant) (int, error) {
	var id int64
	
	query, err := mr.connection.Prepare("INSERT INTO merchant (name, product_type, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer query.Close()

	result, err := query.Exec(merchant.Name, merchant.TypeProduct, merchant.Email, services.SHA256Encoder(merchant.Password))
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return int(id), nil
}

func (mr *MerchantRepository) DeleteMerchant(id_merchant int) error {

	query := "DELETE FROM merchant WHERE id_merchant = ?"

	result, err := mr.connection.Exec(query, id_merchant)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no product found with id %d", id_merchant)
	}

	return nil
}

func (mr *MerchantRepository) FindByEmail(email string) (*model.Merchant, error) {
    var merchant model.Merchant

    query := "SELECT id_merchant, name, product_type, email, password FROM merchant WHERE email = ?"
    err := mr.connection.QueryRow(query, email).Scan(
        &merchant.ID,
        &merchant.Name,
        &merchant.TypeProduct,
        &merchant.Email,
        &merchant.Password,
    )
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil // Nenhum comerciante encontrado com o e-mail fornecido
        }
        fmt.Println(err)
        return nil, err
    }

    return &merchant, nil
}



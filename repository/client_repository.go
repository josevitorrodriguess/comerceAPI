package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/josevitorrodriguess/productsAPI/model"
	"github.com/josevitorrodriguess/productsAPI/services"
)

type ClientRepository struct {
	connection *sql.DB
}

func NewClientRepository(conn *sql.DB) ClientRepository {
	return ClientRepository{
		connection: conn,
	}
}

func (cr *ClientRepository) GetClients() ([]model.Client, error) {
	query := "SELECT id, name, cpf, email FROM client"
	rows, err := cr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var clientList []model.Client

	for rows.Next() {
		var clientObj model.Client
		err := rows.Scan(
			&clientObj.ID,
			&clientObj.Name,
			&clientObj.CPF,
			&clientObj.Email,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		clientList = append(clientList, clientObj)
	}
	return clientList, nil
}

func (cr *ClientRepository) GetClientByID(client_id int) (*model.Client, error) {

	query := "SELECT id, name, cpf, email FROM client WHERE id=?"
	row := cr.connection.QueryRow(query, client_id)

	var client model.Client
	err := row.Scan(
		&client.ID,
		&client.Name,
		&client.CPF,
		&client.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}
	return &client, nil
}

func (cr *ClientRepository) CreateClient(client model.Client) (int, error) {
	var id int64

	query, err := cr.connection.Prepare("INSERT INTO client (name, cpf, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare query: %v", err)
	}
	defer query.Close()

	result, err := query.Exec(client.Name, client.CPF, client.Email, services.SHA256Encoder(client.Password))
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %v", err)
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %v", err)
	}

	return int(id), nil
}

func (cr *ClientRepository) DeleteClient(id_client int) error {

	query := "DELETE FROM client WHERE id = ?"

	result, err := cr.connection.Exec(query, id_client)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no product found with id %d", id_client)
	}

	return nil
}

func (cr *ClientRepository) FindByEmail(email string) (*model.Client, error) {
	var client model.Client

	query := "SELECT id, name,cpf, email, password FROM client WHERE email = ?"
	err := cr.connection.QueryRow(query, email).Scan(
		&client.ID,
		&client.Name,
		&client.CPF,
		&client.Email,
		&client.Password,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Nenhum comerciante encontrado com o e-mail fornecido
		}
		fmt.Println(err)
		return nil, err
	}

	return &client, nil
}

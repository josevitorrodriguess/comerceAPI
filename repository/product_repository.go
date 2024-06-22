package repository

import (
	"database/sql"
	"fmt"


	"github.com/josevitorrodriguess/productsAPI/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(conn *sql.DB) ProductRepository {
	return ProductRepository{
		connection: conn,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var productList []model.Product

	for rows.Next() {
		var productObj model.Product
		err := rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		productList = append(productList, productObj)
	}

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int64

	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer query.Close()

	result, err := query.Exec(product.Name, product.Price)
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

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {
	query := "SELECT id, product_name, price FROM product WHERE id = ?"
	row := pr.connection.QueryRow(query, id_product)

	var product model.Product
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	return &product, nil
}

func (pr *ProductRepository) DeleteProduct(id_product int) error {

	query := "DELETE FROM product WHERE id = ?"

	result, err := pr.connection.Exec(query, id_product)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no product found with id %d", id_product)
	}

	return nil
}

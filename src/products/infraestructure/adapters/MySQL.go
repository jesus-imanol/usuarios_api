package adapters

import (
	"fmt"
	"log"
	"productos-api/src/core"
	"productos-api/src/products/domain/entities"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() (*MySQL, error) {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}, nil
}

func (mysql *MySQL) CreateProduct(product *entities.Product) error {
	query := `
		INSERT INTO products (
			name, 
			price, 
			quantity
		) VALUES (?, ?, ?)
	`
	result, err := mysql.conn.ExecutePreparedQuery(
		query,
		product.Name,
		product.Price,
		product.Quantity,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if result != nil {
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 1 {
			log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
			lastInsertID, err := result.LastInsertId()
			if err != nil {
				fmt.Println(err)
				return err
			}
			product.Id = int32(lastInsertID)
		} else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
	} else {
		log.Printf("[MySQL] - Resultado de la consulta es nil.")
	}
	return nil
}

func (mysql *MySQL) UpdateProduct(product *entities.Product) error {
	query := `
		UPDATE products 
		SET name = ?, price = ?, quantity = ? 
		WHERE id = ?
	`
	result, err := mysql.conn.ExecutePreparedQuery(
		query,
		product.Name,
		product.Price,
		product.Quantity,
		product.Id,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if result != nil {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) GetProductById(id int32) (*entities.Product, error) {
	query := `SELECT id, name, price, quantity FROM products WHERE id = ?`
	rows, err := mysql.conn.ExecuteQuery(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var product entities.Product
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			return nil, err
		}
		return &product, nil
	}
	return nil, fmt.Errorf("producto no encontrado")
}

func (mysql *MySQL) GetAllProducts() ([]*entities.Product, error) {
	query := `SELECT id, name, price, quantity FROM products`
	rows, err := mysql.conn.ExecuteQuery(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (mysql *MySQL) DeleteProduct(id int32) error {
	query := `DELETE FROM products WHERE id = ?`
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return err
	}

	if result != nil {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

package adapters

import (
	"fmt"
	"log"
	"productos-api/src/core"
	"productos-api/src/users/domain/entities"
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

func (mysql *MySQL) Register(user *entities.User) error {
	query := `
		INSERT INTO users (
			full_name, 
			email, 
			password_hash
		) VALUES (?, ?, ?)
	`
	result, err := mysql.conn.ExecutePreparedQuery(
		query,
		user.FullName,
		user.Email,
		user.PasswordHash,
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
			user.Id = int32(lastInsertID)
		} else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
	} else {
		log.Printf("[MySQL] - Resultado de la consulta es nil.")
	}
	return nil
}

func (mysql *MySQL) Update(id int32, fullname string, email string, passwordHash string) error {
	query := "UPDATE users SET full_name = ?, email = ?, password_hash = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, fullname, email, passwordHash, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if result != nil {
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 1 {
			log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
		} else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
	} else {
		log.Printf("[MySQL] - Resultado de la consulta es nil.")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]*entities.User, error) {
	query := "SELECT * FROM users WHERE deleted = 0"
	rows, err := mysql.conn.ExecuteQuery(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	var users []*entities.User
	var deleted bool
	var createdAt string
	var updatedAt string
	for rows.Next() {
		user := entities.User{}
		err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.PasswordHash, &createdAt, &updatedAt, &deleted)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (mysql *MySQL) Delete(id int32) error {
	query := "UPDATE users SET deleted = 1 WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("[MySQL] - Error al ejecutar la consulta: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[MySQL] - Error al obtener las filas afectadas: %v", err)
		return err
	}
	if rowsAffected == 0 {
		log.Printf("[MySQL] - Ninguna fila fue afectada. Usuario con ID %d no encontrado.", id)
		return fmt.Errorf("usuario con ID %d no encontrado", id)
	}

	log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	return nil
}

func (mysql *MySQL) GetById(id int32) (*entities.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	rows, err := mysql.conn.ExecuteQuery(query, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	var user entities.User
	var createdAt string
	var updatedAt string
	var deleted bool
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.PasswordHash, &createdAt, &updatedAt, &deleted)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return &user, nil
}

func (mysql *MySQL) UploadPicture(id int32, urlPicture string) error {
	query := "UPDATE users SET profile_picture = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, urlPicture, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if result != nil {
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 1 {
			log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
		} else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
	} else {
		log.Printf("[MySQL] - Resultado de la consulta es nil.")
	}
	return nil
}

func (mysql *MySQL) Login(email string) (*entities.User, error) {
	query := "SELECT * FROM users WHERE email = ? AND deleted = 0"
	rows, err := mysql.conn.ExecuteQuery(query, email)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	var user entities.User
	var createdAt string
	var updatedAt string
	var deleted bool
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.PasswordHash, &createdAt, &updatedAt, &deleted)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	if user.Id == 0 {
		return nil, fmt.Errorf("usuario no encontrado")
	}
	return &user, nil
}

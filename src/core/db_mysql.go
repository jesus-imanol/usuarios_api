package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Conn_MySQL struct {
	DB  *sql.DB
	Err string
}

func GetDBPool() *Conn_MySQL {
	errorMsg := ""

	// Cambiado de Fatalf a Println para que no mate la app en Railway
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: No se encontró archivo .env, usando variables de entorno del sistema")
	}

	// Obtener las variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSchema := os.Getenv("DB_SCHEMA")
	if dbHost == "" || dbUser == "" || dbPass == "" || dbSchema == "" {
		log.Fatalf("Una o más variables de entorno están vacías.")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbSchema)
	log.Printf("DSN: %s", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		errorMsg = fmt.Sprintf("error al abrir la base de datos: %v", err)
		return &Conn_MySQL{Err: errorMsg}
	}

	// Configuración del pool de conexiones
	db.SetMaxOpenConns(10)

	// Probar la conexión
	if err := db.Ping(); err != nil {
		db.Close()
		errorMsg = fmt.Sprintf("error al verificar la conexión a la base de datos: %v", err)
		return &Conn_MySQL{Err: errorMsg}
	}

	fmt.Println("Connected database")
	return &Conn_MySQL{DB: db, Err: errorMsg}
}

func (conn *Conn_MySQL) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	return result, nil
}

func (conn *Conn_MySQL) ExecuteQuery(query string, values ...interface{}) (*sql.Rows, error) {
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	return rows, nil
}

func (conn *Conn_MySQL) Close() {
	if conn.DB != nil {
		conn.DB.Close()
	}
}
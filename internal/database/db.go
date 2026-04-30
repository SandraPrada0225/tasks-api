package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// se usa _ para registrar el driver automaticamente dentro de database
// variable global de conexion
var DB *sql.DB

func Connect() {
	//String de conexión: usuario, clave, base, host (socket en Mac)
	connStr := "user=milena password=1234 dbname=heladeria sslmode=disable host=/tmp"

	var err error
	//abre la conexion
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//verificamos la conexion
	err = DB.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la DB")
	}

	log.Println("Conectado a PostgreSQL 🚀")
}

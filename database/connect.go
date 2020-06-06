package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db  *sqlx.DB
	err error
)

// ConnectDB opens a connection to the mariadb server (with the mysql driver github.com/go-sql-driver/mysql)
func ConnectDB(dbcs string) {

	db, err = sqlx.Open("mysql", dbcs)
	if err != nil {
		log.Panicln("Erro ao tentar abrir conexão com o banco de dados - ", err)
		return
	}
	log.Println("Conexão com BD estabelecida com sucesso!")
}

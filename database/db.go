package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
    host := os.Getenv("HOST")
    user := os.Getenv("USER")
    password := os.Getenv("PASSWORD")
    dbname := os.Getenv("DBNAME")
    port := os.Getenv("DBPORT")
    stringDeConexao := "host="+ host +" user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
    DB, err = gorm.Open(postgres.Open(stringDeConexao))
    if err != nil {
        log.Panic("Erro ao conectar com banco de dados")
    }

    DB.AutoMigrate(&models.Aluno{})
}
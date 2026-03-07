package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/ropehapi/kaizen-secretary/internal/routines"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	c := cron.New(cron.WithSeconds()) // habilita campo de segundos

	_, err := c.AddFunc("0 55 13 7 3 *", routines.RememberScoutMonthlyFees)
	if err != nil {
	    panic(err)
	}

	c.Start()

	// mantém a aplicação rodando
	select {}
}

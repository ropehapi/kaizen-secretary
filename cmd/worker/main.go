package main

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/ropehapi/kaizen-secretary/internal/logger"
	"github.com/ropehapi/kaizen-secretary/internal/routines"
)

func main() {
	logger.Init()

	if err := godotenv.Load(); err != nil {
		slog.Warn("arquivo .env não encontrado, usando variáveis de ambiente do sistema", "error", err)
	}

	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc("0 */5 * * * *", routines.RememberScoutMonthlyFees)
	if err != nil {
		slog.Error("falha ao registrar cron job", "error", err)
		panic(err)
	}

	slog.Info("kaizen-secretary iniciado, aguardando execução dos jobs...")

	c.Start()

	select {}
}

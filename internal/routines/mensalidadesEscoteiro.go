package routines

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func RememberScoutMonthlyFees() {
	taxpayers := map[string]string{
		"Pedrinho": "5543936180709",
	}

	month := getMonthInPortuguese()

	for name, phone := range taxpayers {
		message := fmt.Sprintf("Olá, %s, passando para lembrar sobre Contribuição mensal do Grupo Escoteiro Guarani, referente ao mês de %s. Enviar comprovante no whatsApp *PIX GRUPO GUARANI*.\nObs: Essa é uma mensagem automática. Caso já tenha feito o pagamento, por favor desconsidere.", name, month)
		
		payload := map[string]interface{}{
			"number":  phone,
			"message": message,
		}

		body, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", os.Getenv("MESSAGING_OFFICER_HOST")+":"+os.Getenv("MESSAGING_OFFICER_PORT")+"/api/send-message", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", "corinthians-gigante")
		req.Header.Set("x-Session-Id", "pedro")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Erro ao enviar requisição:", err)
			return
		}
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)

		var prettyJSON interface{} // pode ser []interface{} ou map[string]interface{}
		if err := json.Unmarshal(respBody, &prettyJSON); err == nil {
			formatted, _ := json.MarshalIndent(prettyJSON, "", "  ")
			fmt.Println("HTTP Status:", resp.Status)
			fmt.Println(string(formatted))
		} else {
			fmt.Println("HTTP Status:", resp.Status)
			fmt.Println(string(respBody))
		}
	}
}

func getMonthInPortuguese() string {
	portugueseMonths := [...]string{
		"Janeiro", 
		"Fevereiro", 
		"Março", 
		"Abril", 
		"Maio", 
		"Junho",
		"Julho", 
		"Agosto", 
		"Setembro", 
		"Outubro", 
		"Novembro", 
		"Dezembro",
	}

	actualMonth := time.Now().Month()
	monthNamePortuguese := portugueseMonths[actualMonth-1]

	return monthNamePortuguese
}

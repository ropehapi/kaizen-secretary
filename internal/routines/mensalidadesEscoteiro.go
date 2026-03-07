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
		"":                                       "",
		"CRISLAINE FERREIRA  DE PAULA":           "43996244613",
		"FAGNER JUNIOR MASSEI":                   "43999556660",
		"MAICON FIER":                            "",
		"CAROLINA DE OLIVEIRA TENORIO":           "43991514942",
		"ROSANGELA CRISTINA ALVES HAAGSMA":       "",
		"EGLAIA DE CARVALHO CHERON":              "",
		"ALINE FERREIRA MARCHI":                  "43988706262",
		"DELAIR APARECIDA ALVES DOS SANTOS":      "",
		"DIRCE ELY MAIHACK":                      "43996888151",
		"Maria Gloria dos Santos Miyasaki":       "",
		"RUBIA SIMONI PRIMO":                     "",
		"MARIA DE LOURDES ASSIZ VIEIRA":          "",
		"RICARDO FORSTER":                        "",
		"Rosiane Andréia Ribeiro Teixeira":       "",
		"EDGAR JOSE SCHUSTER":                    "43999571116",
		"EVERTON HENRIQUE FORTI":                 "43999010525",
		"JAIR DONIZETE STEFANI":                  "4399720306",
		"JAQUELINE AMADEU BORASCHI":              "43996403151",
		"Gabriella Pitoli Schauff":               "",
		"NATALIA CRISTINA DO CARMO":              "",
		"Lucineia Antonia de Oliveira Pereira":   "",
		"ELEUTERIO DA SILVA FERNANDES":           "",
		"Gisele Mazer Hofmam":                    "",
		"ALESSANDRO PEREIRA JAQUES":              "",
		"MARCIA TEIXEIRA MARCOS":                 "",
		"LEILA CRISTINA RODRIGUES":               "43991130240",
		"Anderson Melo da Silva":                 "43996661413",
		"Josiane Salmazo Devara":                 "",
		"Eduardo Augusto Matiuzzi":               "",
		"LUCI ANE FERNANDES GARCIA DA SILVA":     "43991181378",
		"CLAUDIA MOREIRA MARQUEZINI":             "43999518077",
		"DAIANE GONÇALVES DE SOUZA VALÉRIO":      "43996370162",
		"VANIA MARIA FERREIRA":                   "",
		"ROSELI A. MONTEIRO YOSHIMURA":           "",
		"JULIANA BARCELLOS DE OLIVEIRA":          "43999282866",
		"CARLA BEAZI":                            "",
		"CLEONI ADEMIR PEREIRA":                  "",
		"João Fernando da Cunha":                 "",
		"Débora Garcia Prescendo de Godoy Bueno": "",
		"Aline Fernandes Rodrigues Nandi":        "43999296846",
		"FABIANY LOPES":                          "",
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

		time.Sleep(10 * time.Second)
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

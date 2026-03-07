package routines

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func RememberScoutMonthlyFees() {
	taxpayers := map[string]string{
		// "":                                       "43988706262",
		"CRISLAINE FERREIRA  DE PAULA":           "5543988590192",
		// "FAGNER JUNIOR MASSEI":                   55"",
		"MAICON FIER":                            "5543996520477",
		"CAROLINA DE OLIVEIRA TENORIO":           "5543981514942",
		"ROSANGELA CRISTINA ALVES HAAGSMA":       "5543996674698",
		"EGLAIA DE CARVALHO CHERON":              "5543999772677",
		"ALINE FERREIRA MARCHI":                  "5543988706262",
		"DELAIR APARECIDA ALVES DOS SANTOS":      "5543991199827",
		"DIRCE ELY MAIHACK":                      "5543988540078",
		"Maria Gloria dos Santos Miyasaki":       "5543988323660",
		"RUBIA SIMONI PRIMO":                     "5543998608622",
		"MARIA DE LOURDES ASSIZ VIEIRA":          "5543991553713",
		"RICARDO FORSTER":                        "5543999556662",
		"Rosiane Andréia Ribeiro Teixeira":       "5543991731318",
		"EDGAR JOSE SCHUSTER":                    "5543998458266",
		// "EVERTON HENRIQUE FORTI":                 55"",
		"JAQUELINE AMADEU BORASCHI":              "5543996403151",
		"Gabriella Pitoli Schauff":               "5543991530112",
		"NATALIA CRISTINA DO CARMO":              "5543998682678",
		"Lucineia Antonia de Oliveira Pereira":   "5543998448075",
		"ELEUTERIO DA SILVA FERNANDES":           "5543991181378",
		"Gisele Mazer Hofmam":                    "5543999660219",
		"ALESSANDRO PEREIRA JAQUES":              "5543991730168",
		"MARCIA TEIXEIRA MARCOS":                 "5543999808467",
		"LEILA CRISTINA RODRIGUES":               "5543991130240",
		"Anderson Melo da Silva":                 "5543996661413",
		"Eduardo Augusto Matiuzzi":               "5543996461302",
		// "LUCI ANE FERNANDES GARCIA DA SILVA":     55"",
		"CLAUDIA MOREIRA MARQUEZINI":             "5543999518077",
		"DAIANE GONÇALVES DE SOUZA VALÉRIO":      "5543996370162",
		"VANIA MARIA FERREIRA":                   "5543984187433",
		"JULIANA BARCELLOS DE OLIVEIRA":          "5543999307945",
		"CARLA BEAZI":                            "5543999714313",
		"CLEONI ADEMIR PEREIRA":                  "5545991176628",
		"João Fernando da Cunha":                 "5543999187335",
		"Débora Garcia Prescendo de Godoy Bueno": "5543999158494",
		"Aline Fernandes Rodrigues Nandi":        "5543999296846",
		"FABIANY LOPES":                          "5561983129340",
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
		req.Header.Set("x-api-key", os.Getenv("MESSAGING_OFFICER_API_KEY"))
		req.Header.Set("x-Session-Id", os.Getenv("MESSAGING_OFFICER_SESSION_ID"))

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

		time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)
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

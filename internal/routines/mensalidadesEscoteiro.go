package routines

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func RememberScoutMonthlyFees() {
	taxpayers := map[string]string{
		"Pedrinho": "5543936180709",
	}

	for name, phone := range taxpayers {
		payload := map[string]interface{}{
			"number":  phone,
			"message": "Olá " + name + " paga eu.",
		}

		body, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", os.Getenv("MESSAGING_OFFICER_HOST")+":"+os.Getenv("MESSAGING_OFFICER_PORT")+"/send-message", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")

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

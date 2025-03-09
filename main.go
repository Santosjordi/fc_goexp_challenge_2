package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ViaCEP struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Estado     string `json:"estado"`
	Cidade     string `json:"localidade"`
	Bairro     string `json:"bairro"`
}

type BrasilApiCep struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"street"`
	Estado     string `json:"state"`
	Cidade     string `json:"city"`
	Bairro     string `json:"neighborhood"`
}

// Isso é uma gambiarra?
type CEP struct {
	Api        string
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Estado     string `json:"estado"`
	Cidade     string `json:"cidade"`
	Bairro     string `json:"bairro"`
}

func (c CEP) PrintAsMap() {
	properties := map[string]string{
		"Api":        c.Api,
		"Cep":        c.Cep,
		"Logradouro": c.Logradouro,
		"Estado":     c.Estado,
		"Cidade":     c.Cidade,
		"Bairro":     c.Bairro,
	}
	for key, value := range properties {
		log.Printf("%s: %s\n", key, value)
	}
}

func convertViaCEP(v ViaCEP) CEP {
	return CEP{
		Api:        "ViaCEP",
		Cep:        v.Cep,
		Logradouro: v.Logradouro,
		Estado:     v.Estado,
		Cidade:     v.Cidade,
		Bairro:     v.Bairro,
	}
}

func convertBrasilApiCep(b BrasilApiCep) CEP {
	return CEP{
		Api:        "BrasilAPI",
		Cep:        b.Cep,
		Logradouro: b.Logradouro,
		Estado:     b.Estado,
		Cidade:     b.Cidade,
		Bairro:     b.Bairro,
	}
}

// [X] Query the same CEP from two different APIs
// [X] Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
// [X] O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.
// [X] Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.
func main() {
	for _, cep := range os.Args[1:] {
		fasterApi := make(chan CEP)
		go func() {
			response, err := BuscaViaCep(cep)
			if err != nil {
				log.Fatalf("Error getting cep: %v", err)
			}
			msg := convertViaCEP(*response)
			fasterApi <- msg
		}()

		go func() {
			responseBRAPI, err := BuscaBrasilApiCep(cep)
			if err != nil {
				log.Fatalf("Error getting cep: %v", err)
			}
			msg := convertBrasilApiCep(*responseBRAPI)
			fasterApi <- msg
		}()

		select {
		case msg := <-fasterApi:
			msg.PrintAsMap()
		case msg := <-fasterApi:
			msg.PrintAsMap()
		case <-time.After(1 * time.Second):
			log.Fatalf("Timeout")
		}
	}
}

func BuscaViaCep(cep string) (*ViaCEP, error) {
	// time.Sleep(2 * time.Second)
	resp, error := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var c ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil
}

func BuscaBrasilApiCep(cep string) (*BrasilApiCep, error) {
	// time.Sleep(2 * time.Second)
	resp, error := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var c BrasilApiCep
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil
}

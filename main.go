package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func main() {
	for _, cep := range os.Args[1:] {
		response, err := BuscaViaCep(cep)
		if err != nil {
			log.Fatalf("Error getting cep: %v", err)
		}
		fmt.Println("ViaCEP")
		fmt.Printf("Cep: %s\nLogradouro: %s\nBairro: %s\nCidade: %s\nEstado: %s\n",
			response.Cep, response.Logradouro, response.Bairro, response.Cidade, response.Estado)

		responseBRAPI, err := BuscaBrasilApiCep(cep)
		if err != nil {
			log.Fatalf("Error getting cep: %v", err)
		}
		fmt.Println("BrasilAPI")
		fmt.Printf("Cep: %s\nLogradouro: %s\nBairro: %s\nCidade: %s\nEstado: %s\n",
			responseBRAPI.Cep, responseBRAPI.Logradouro, responseBRAPI.Bairro, responseBRAPI.Cidade, responseBRAPI.Estado)
	}
}

func BuscaViaCep(cep string) (*ViaCEP, error) {
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

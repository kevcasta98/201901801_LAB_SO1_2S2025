package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Mensaje string `json:"mensaje"`
}

const (
	studentName = "Kevin Castañeda"
	studentID   = "201901801"
	vmName      = "VM2"
	apiName     = "API3"
	port        = ":8083"
)

func baseMessage() string {
	return fmt.Sprintf("Hola, responde la API: %s en la %s, desarrollada por el estudiante %s con carnet: %s",
		apiName, vmName, studentName, studentID)
}

func callAPI(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("Error llamando a %s: %v", url, err)
	}
	defer resp.Body.Close()

	var res Response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "Error al decodificar respuesta"
	}
	return res.Mensaje
}

func main() {
	http.HandleFunc("/api3/"+studentID+"/llamar-api1", func(w http.ResponseWriter, r *http.Request) {
		target := callAPI("http://api1:8081/api1/" + studentID + "/llamar-api2")
		json.NewEncoder(w).Encode(Response{Mensaje: baseMessage() + " → " + target})
	})

	http.HandleFunc("/api3/"+studentID+"/llamar-api2", func(w http.ResponseWriter, r *http.Request) {
		target := callAPI("http://api2:8082/api2/" + studentID + "/llamar-api1")
		json.NewEncoder(w).Encode(Response{Mensaje: baseMessage() + " → " + target})
	})

	fmt.Println(apiName + " escuchando en puerto " + port)
	http.ListenAndServe(port, nil)
}

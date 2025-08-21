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
	vmName      = "VM1"
	apiName     = "API1"
	port        = ":8081"
)

func baseMessage() string {
	return fmt.Sprintf("Hola, responde la API: %s en la %s, desarrollada por el estudiante %s con carnet: %s",
		apiName, vmName, studentName, studentID)
}

func callAPI(url string, depth int) string {
	if depth <= 0 {
		return "Fin de la cadena de llamadas."
	}

	resp, err := http.Get(url + "?depth=" + fmt.Sprint(depth-1))
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
	http.HandleFunc("/api1/"+studentID+"/llamar-api2", func(w http.ResponseWriter, r *http.Request) {
		depth := 2 // máximo número de saltos permitidos
		target := callAPI("http://api2:8082/api2/"+studentID+"/llamar-api3", depth)
		json.NewEncoder(w).Encode(Response{Mensaje: baseMessage() + " → " + target})
	})

	http.HandleFunc("/api1/"+studentID+"/llamar-api3", func(w http.ResponseWriter, r *http.Request) {
		depth := 2
		target := callAPI("http://api3:8083/api3/"+studentID+"/llamar-api2", depth)
		json.NewEncoder(w).Encode(Response{Mensaje: baseMessage() + " → " + target})
	})

	fmt.Println(apiName + " escuchando en puerto " + port)
	http.ListenAndServe(port, nil)
}

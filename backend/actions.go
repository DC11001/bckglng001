package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Response(w http.ResponseWriter, status int, results Model) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func Index(w http.ResponseWriter, r *http.Request) {
	info := IndexModel{Name: "Analisis y Diseño 1", Website: "Josué Daniel Caal Torres - Practica 1 - Grupo 4"}
	fmt.Println(info)
	u, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u)) //
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(info)
}
func Login(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	var response LoginResponse

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Realiza la lógica de inicio de sesión y establece valores en la respuesta...
	response.User = request.User
	response.Valido = isValidCredentials(request.User, request.Pass)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest
	var response CreateUserResponse

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realiza la lógica de creación de usuario y establece valores en la respuesta...

	json.NewEncoder(w).Encode(response)
}

func ModifyUserHandler(w http.ResponseWriter, r *http.Request) {
	var request ModifyUserRequest
	var response ModifyUserResponse

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realiza la lógica de modificación de usuario y establece valores en la respuesta...

	json.NewEncoder(w).Encode(response)
}

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	var request CreateRoomRequest
	var response CreateRoomResponse

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realiza la lógica de creación de habitación y establece valores en la respuesta...

	json.NewEncoder(w).Encode(response)
}

func ModifyRoomHandler(w http.ResponseWriter, r *http.Request) {
	var request ModifyRoomRequest
	var response ModifyRoomResponse

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realiza la lógica de modificación de habitación y establece valores en la respuesta...

	json.NewEncoder(w).Encode(response)
}

func DeleteRoomHandler(w http.ResponseWriter, r *http.Request) {
	var request DeleteRoomRequest
	var response DeleteRoomResponse

	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realiza la lógica de eliminación de habitación y establece valores en la respuesta...

	json.NewEncoder(w).Encode(response)
}

func SearchRoomHandler(w http.ResponseWriter, r *http.Request) {
	var request SearchRoomRequest
	var response SearchRoomResponse

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realiza la lógica de búsqueda de habitación y establece valores en la respuesta...

	json.NewEncoder(w).Encode(response)
}

func SearchRoomByStateHandler(w http.ResponseWriter, r *http.Request) {
	var request SearchRoomByStateRequest
	var response SearchRoomByStateResponse

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realiza la lógica de búsqueda de habitación por estado y establece valores en la respuesta...

	json.NewEncoder(w).Encode(response)
}

func CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	var request CreateReservationRequest
	var response CreateReservationResponse

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realiza la lógica de creación de reserva y establece valores en la respuesta...

	json.NewEncoder(w).Encode(response)
}

func ModifyReservationHandler(w http.ResponseWriter, r *http.Request) {
	var request ModifyReservationRequest
	var response ModifyReservationResponse

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realiza la lógica de modificación de reserva y establece valores en la respuesta...

	json.NewEncoder(w).Encode(response)
}

func DeleteReservationHandler(w http.ResponseWriter, r *http.Request) {
	var request DeleteReservationRequest
	var response DeleteReservationResponse

	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realiza la lógica de eliminación de reserva y establece valores en la respuesta...

	json.NewEncoder(w).Encode(response)
}

func ClientReportHandler(w http.ResponseWriter, r *http.Request) {
	var request ClientReportRequest
	var response ClientReportResponse

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realiza la lógica de generación de informe para el cliente y establece valores en la respuesta...

	json.NewEncoder(w).Encode(response)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data Model
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	log.Println("Removed:\n", data)
	fmt.Fprintf(w, "Item removed")
}

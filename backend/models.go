package main

import "time"

type Model struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type IndexModel struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

type LoginRequest struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type LoginResponse struct {
	User   string `json:"user"`
	Valido bool   `json:"valido"`
}

type CreateUserRequest struct {
	User     string `json:"user"`
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     string `json:"Edad"`
	Genero   string `json:"Genero"`
	Pass     string `json:"pass"`
}

type CreateUserResponse struct {
	User   string `json:"user"`
	Creado bool   `json:"creado"`
}

type ModifyUserRequest struct {
	User      string `json:"user"`
	Modificar string `json:"modificar"`
	Valor     string `json:"valor"`
}

type ModifyUserResponse struct {
	User       string `json:"user"`
	Modificado bool   `json:"modificado"`
}

type CreateRoomRequest struct {
	Codigo int    `json:"codigo"`
	Estado string `json:"estado"`
}

type CreateRoomResponse struct {
	Codigo int  `json:"codigo"`
	Creado bool `json:"creado"`
}

type ModifyRoomRequest struct {
	Codigo int    `json:"codigo"`
	Estado string `json:"estado"`
}

type ModifyRoomResponse struct {
	Codigo     int  `json:"codigo"`
	Modificado bool `json:"modificado"`
}

type DeleteRoomRequest struct {
	Codigo int `json:"codigo"`
}

type DeleteRoomResponse struct {
	Codigo    int  `json:"codigo"`
	Eliminado bool `json:"eliminado"`
}

type SearchRoomRequest struct {
	Codigo int `json:"codigo"`
}

type SearchRoomResponse struct {
	Codigo int    `json:"codigo"`
	Estado string `json:"estado"`
}

type SearchRoomByStateRequest struct {
	Estado string `json:"estado"`
}

type SearchRoomByStateResponse struct {
	Respuesta []int `json:"respuesta"`
}

type CreateReservationRequest struct {
	User         string    `json:"user"`
	Habitacion   int       `json:"habitacion"`
	FechaIngreso time.Time `json:"Fecha_ingreso"`
	FechaSalida  time.Time `json:"Fecha_salida"`
}

type CreateReservationResponse struct {
	User   string `json:"user"`
	Creado bool   `json:"creado"`
}

type ModifyReservationRequest struct {
	User         string    `json:"user"`
	Habitacion   int       `json:"habitacion"`
	FechaIngreso time.Time `json:"Fecha_ingreso"`
	FechaSalida  time.Time `json:"Fecha_salida"`
}

type ModifyReservationResponse struct {
	User       string `json:"user"`
	Modificado bool   `json:"modificado"`
}

type DeleteReservationRequest struct {
	User       string `json:"user"`
	Habitacion int    `json:"habitacion"`
}

type DeleteReservationResponse struct {
	User      string `json:"user"`
	Eliminado bool   `json:"eliminado"`
}

type ClientReportRequest struct {
	User string `json:"user"`
}

type ClientReportResponse struct {
	Respuesta []int `json:"respuesta"`
}

type GeneralReportResponse struct {
	Respuesta []int `json:"respuesta"`
}

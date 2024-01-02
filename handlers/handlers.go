package handlers

import (
	"apirestgo/db"
	"apirestgo/models"
	"encoding/json"
	"fmt"
	"net/http"
)


func GetUsers(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json") //defino el tipo de contenido que voy a devolver

	db.Connect()
	users := models.ListUsers()
	db.Close()
	
	//transformo el objeto a JSON
	output, _ := json.Marshal(users) //devuelve dos valores, un valor tipo byte y un error
	fmt.Fprintln(rw, string(output))


}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Get User")
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Create User")
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Update User")
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Delete User")
}


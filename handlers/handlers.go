package handlers

import (
	"apirestgo/db"
	"apirestgo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	rw.Header().Set("Content-Type", "application/json") 

	//obtengo el id del usuario
	vars := mux.Vars(r) //variables del url

	//convierto el id a entero
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user := models.GetUser(int64(userId))
	db.Close()


	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json") 

	//obtengo TODO el registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body) //decodifico el json a un objeto
	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
		
	} else {
		db.Connect()
		user.Save()
		db.Close()
	}

	output, _ := json.Marshal(user) //convierto el objeto a json
	fmt.Fprintln(rw, string(output))
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json") 

	//obtengo TODO el registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body) //decodifico el json a un objeto

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
		
	} else {
		db.Connect()
		user.Save()
		db.Close()
	}

	output, _ := json.Marshal(user) //convierto el objeto a json
	fmt.Fprintln(rw, string(output))
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json") 

	//obtengo el id del usuario
	vars := mux.Vars(r) //variables del url

	//convierto el id a entero
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user := models.GetUser(int64(userId))
	user.Delete()
	db.Close()


	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}


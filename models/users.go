package models

import (
	"apirestgo/db"
)


type User struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type Users []User


const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`


//constructor de un nuevo usuario
func NewUser(username string, password string, email string) *User{
	user := &User{
		Username: username,
		Password: password,
		Email: email,
	}
	return user
}

//crear usuario e insertar en la base de datos
func CreateUser(username, password, email string) *User{
	user := NewUser(username, password, email)
	user.Save()
	return user
}

//insertar registro
func(user *User) insert(){
	sql := "INSERT users SET username=?, password=?, email=?"

	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)

	user.Id, _ = result.LastInsertId() //devuelve el id del ultimo registro insertado

}


//listar todos los usuarios
func ListUsers() Users{
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, _ := db.Query(sql) 

	//recorro el rows para obtener cada registro
	for rows.Next(){
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users
}


//obtener un registro por id
func GetUser(id int64) *User{
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := db.Query(sql, id)

	for rows.Next(){
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}

	return user

}

//actualizar un registro
func(user *User) update(){
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)


}

//guardar o editar un registro
func(user *User) Save(){
	if user.Id == 0 {
		user.insert()
	}else{
		user.update()
	}
}

//eliminar un registro
func(user *User) Delete(){
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)
}
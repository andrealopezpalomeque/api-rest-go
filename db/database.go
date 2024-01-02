package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3307)/database
const url = "root@tcp(localhost:3307)/goweb_db"
//guarda la conexion
var db *sql.DB

//realizar la conexion
func Connect(){
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexion exitosa")
	db = connection

}
//cerrar la conexion
func Close(){
	db.Close()
}

//verificar la conexion
func Ping(){
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Pong")
}

//verificar si la TABLA existe
func ExistsTable(tableName string) bool{
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return rows.Next() //devuelve un booleano
}

//------CREAR TABLAS-----
func CreateTable(schema string, name string){
	if !ExistsTable(name){	
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println("Error:", err)
		}}

}

//REINICIAR EL REGISTRO DE UNA TABLA
func TruncateTable(tableName string){
	sql := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println("Error:", err)
	}
}


//polimosfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error){
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return result, err
}

//polimosfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error){
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return rows, err
}





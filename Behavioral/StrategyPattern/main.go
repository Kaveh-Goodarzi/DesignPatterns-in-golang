package main

import (
	"fmt"
)

type IDBConnection interface {
	connect()
}

type DBConnection struct {
	Db IDBConnection
}

func (con DBConnection) DBConnect() {
	con.Db.connect()
}

type MySQLConnection struct {
	ConnectionString string
}

func (con MySQLConnection) connect() {
	fmt.Println("MySQL" + con.ConnectionString)
}

type PostgreSQLConnection struct {
	ConnectionString string
}

func (con PostgreSQLConnection) connect() {
	fmt.Println("PostgreSQL" + con.ConnectionString)
}

type MongoDBConnection struct {
	ConnectionString string
}

func (con MongoDBConnection) connect() {
	fmt.Println("MongoDB" + con.ConnectionString)
}

func main() {

	mySQL := MySQLConnection{ConnectionString: "my SQL DB is connected"}
	con := DBConnection{Db: mySQL}
	con.DBConnect()

	fmt.Println("======================")


	postgreSQL := PostgreSQLConnection{ConnectionString: "pg SQL DB is connected"}
	con1 := DBConnection{Db: postgreSQL}
	con1.DBConnect()

	fmt.Println("======================")
	
	mongoDB := MongoDBConnection{ConnectionString: "mongo DB is connected"}
	con2 := DBConnection{Db: mongoDB}
	con2.DBConnect()
}